package clipboard

import (
	"encoding/binary"
	"errors"
	"fmt"
	"path/filepath"
	"reflect"
	"syscall"
	"time"
	"unicode/utf16"
	"unsafe"

	"golang.org/x/sys/windows"
)

func getClipboardDataHGlobalSlice(id uint32) (s []byte, err error) {
	r1, _, e1 := syscall.Syscall(procGetClipboardData.Addr(), 1, uintptr(id), 0, 0)
	if r1 == 0 {
		err = e1
	}
	hHeap, _ := getProcessHeap()
	size, e2 := heapSize(hHeap, 0, r1)

	if size == 0 {
		return nil, e2
	}
	s = byteSliceFromUintptr(r1, int(size))
	return s, nil
}

// GetFileGroupDescriptor returns a slice containing file metadata (filename + filesize) in the FileGroupDescriptorW slot
func GetFileGroupDescriptor() ([]FileInfo, error) {
	id, err := registerClipboardFormat("FileGroupDescriptorW")
	if err != nil {
		return nil, err
	}
	if err := openClipboard(0); err != nil {
		return nil, err
	}
	defer closeClipboard()

	ok, err := isClipboardFormatAvailable(id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("Format not available.")
	}
	h, err := getClipboardDataHGlobalSlice(id)
	if err != nil {
		return nil, err
	}

	items := []_FILEDESCRIPTORW{}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&items))
	hdr.Data = uintptr(unsafe.Pointer(&h[4]))
	hdr.Len = int(binary.LittleEndian.Uint32(h))
	result := []FileInfo{}

	for i := 0; i < len(items); i++ {
		fd := items[i]
		result = append(result, FileInfo{
			Name: syscall.UTF16ToString(fd.cFileName[:]),
			Size: int64(fd.nFileSizeHigh)<<32 | int64(fd.nFileSizeLow),
		})
	}
	return result, nil
}

// returns a slice containing the filepaths in the H_DROP(15) slot
func GetHDROP() ([]string, error) {
	const id = 15
	if err := openClipboard(0); err != nil {
		return nil, err
	}
	defer closeClipboard()

	ok, err := isClipboardFormatAvailable(id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("Format not available.")
	}
	h, err := getClipboardData(id)
	if err != nil {
		return nil, err
	}
	buf := make([]uint16, 72)
	const NUM_ENTRIES = 0xFFFFFFFF
	n, err := dragQueryFile(syscall.Handle(h), NUM_ENTRIES, nil)
	if n > 0 {
		var result []string
		for i := 0; i < n; i++ {
			reqBufSize, err := dragQueryFile(syscall.Handle(h), i, nil)
			if reqBufSize > len(buf) {
				buf = make([]uint16, reqBufSize+1)
			}
			n, err := dragQueryFile(syscall.Handle(h), i, buf)
			if err != nil {
				return nil, err
			}
			result = append(result, string(utf16.Decode(buf[:n])))
		}
		return result, nil
	}
	return []string{}, nil
}

// GetShellIDListArray DEPRECATED
func GetShellIDListArray() ([]string, error) {
	id, err := registerClipboardFormat("Shell IDList Array")
	if err := openClipboard(0); err != nil {
		return nil, err
	}
	defer closeClipboard()

	ok, err := isClipboardFormatAvailable(id)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("Format not available.")
	}
	h, err := getClipboardData(id)
	if err != nil {
		return nil, err
	}

	nItems := int(binary.LittleEndian.Uint32(byteSliceFromUintptr(h, 4)))
	result := []string{}
	if nItems > 0 {
		buf := [1024]uint16{}
		offset := shGetPIDLValue(h, 0)
		if err := shGetPathFromIDListW(uintptr(h)+offset, buf[:]); err != nil {
			return nil, err
		}
		rootDir := syscall.UTF16ToString(buf[:])
		desktopDir, err := getDesktopFolder(0)
		if err != nil {
			return nil, err
		}
		for i := 0; i < nItems; i++ {
			offset := shGetPIDLValue(h, i+1)
			if err := shGetPathFromIDListW(uintptr(h)+offset, buf[:]); err != nil {
				return nil, err
			}
			relPath, err := filepath.Rel(desktopDir, syscall.UTF16ToString(buf[:]))
			if err != nil {
				return nil, err
			}
			result = append(result, filepath.Join(rootDir, relPath))
		}
	}
	return result, nil
}

func shGetPIDLValue(cidl uintptr, offset int) uintptr {
	return uintptr(binary.LittleEndian.Uint32(byteSliceFromUintptr(cidl+4+(4*uintptr(offset)), 4)))
}

// Formats returns a slice that contains all formats currently avaiable in the clipboard
func Formats() ([]int, error) {
	if err := openClipboard(0); err != nil {
		return nil, err
	}
	defer closeClipboard()

	var f uint = 0
	var err error
	var result []int
	retries := 0
	for {
		f, err = enumClipboardFormats(f)
		if err != nil {
			if errors.Is(err, windows.ERROR_SUCCESS) {
				break
			}
			if errors.Is(err, windows.ERROR_CLIPBOARD_NOT_OPEN) ||
				errors.Is(err, windows.ERROR_ACCESS_DENIED) {
				if retries < 3 {
					if err := openClipboard(0); err == nil {
						retries++
						time.Sleep(time.Millisecond * time.Duration(retries))
						continue
					}
				}
			}
			return nil, err
		}
		result = append(result, int(f))
	}
	return result, nil
}

var ErrUnknownClipboardFormat = errors.New("Unkown clipboard format")

var predefinedFormatNames = map[uint]string{
	1:  "CF_TEXT",
	2:  "CF_BITMAP",
	3:  "CF_METAFILEPICT",
	4:  "CF_SYLK",
	5:  "CF_DIF",
	6:  "CF_TIFF",
	7:  "CF_OEMTEXT",
	8:  "CF_DIB",
	9:  "CF_PALETTE",
	10: "CF_PENDATA",
	11: "CF_RIFF",
	12: "CF_WAVE",
	13: "CF_UNICODETEXT",
	14: "CF_ENHMETAFILE",
	15: "CF_HDROP",
	16: "CF_LOCALE",
	17: "CF_DIBV5",

	0x0080: "CF_OWNERDISPLAY",
	0x0081: "CF_DSPTEXT",
	0x0082: "CF_DSPBITMAP",
	0x0083: "CF_DSPMETAFILEPICT",
	0x008E: "CF_DSPENHMETAFILE",
	0x0200: "CF_PRIVATEFIRST",
	0x02FF: "CF_PRIVATELAST",
	0x0300: "CF_GDIOBJFIRST",
	0x03FF: "CF_GDIOBJLAST",
}

func predefinedFormatName(id uint) (string, error) {
	if v, ok := predefinedFormatNames[id]; ok {
		return v, nil
	}
	if id > 0x0200 && id < 0x02FF {
		return "PRIVATE", nil
	}
	if id > 0x0300 && id < 0x03FF {
		return "GDIOBJ", nil
	}
	return "", fmt.Errorf("Unsupported format %d. %w", id, ErrUnknownClipboardFormat)
}

func isRegisteredClipboardFormat(id uint) bool {
	return id >= 0xC000 && id <= 0xFFFF
}

// FormatName returns a readable name for the passed id.
//
// Being either a pre-defined name, or through a call to GetClipboardFormatNameW)
func FormatName(id int) (string, error) {
	if isRegisteredClipboardFormat(uint(id)) {
		buf := [256]uint16{}
		n, err := getClipboardFormatName(uint(id), buf[:])
		if err != nil {
			return "", err
		}
		return string(utf16.Decode(buf[:n])), nil
	}

	return predefinedFormatName(uint(id))
}
