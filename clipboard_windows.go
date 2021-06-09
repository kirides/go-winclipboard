package clipboard

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"reflect"
	"syscall"
	"time"
	"unicode/utf16"
	"unsafe"

	"github.com/kirides/go-winclipboard/internal/winsys"

	"golang.org/x/sys/windows"
	"golang.org/x/text/encoding/unicode"
)

func AddClipboardFormatListener(h syscall.Handle) error {
	return winsys.AddClipboardFormatListener(h)
}
func RemoveClipboardFormatListener(h syscall.Handle) error {
	return winsys.RemoveClipboardFormatListener(h)
}

func getClipboardDataHGlobalSlice(id uint32) (s []byte, err error) {
	r1, e1 := winsys.GetClipboardData(id)
	if e1 != nil {
		return nil, e1
	}
	hHeap, _ := winsys.GetProcessHeap()
	size, e2 := winsys.HeapSize(hHeap, 0, uintptr(r1))

	if size == 0 {
		return nil, e2
	}
	s = byteSliceFromUintptr(uintptr(r1), int(size))
	return s, nil
}

func SetData(id uint, data []byte) error {
	return setClipboardDataSlice(uint32(id), data)
}

func getUnicodeBytes(text string) ([]byte, error) {
	enc := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	buf := bytes.NewBuffer(nil)
	_, err := enc.Writer(buf).Write([]byte(text))
	if err != nil {
		return nil, err
	}
	buf.WriteByte(0)
	buf.WriteByte(0)
	return buf.Bytes(), nil
}

func Empty() error {
	if err := winsys.OpenClipboard(0); err != nil {
		return err
	}
	defer winsys.CloseClipboard()
	return winsys.EmptyClipboard()
}

func SetUnicodeText(text string) error {
	data, err := getUnicodeBytes(text)
	if err != nil {
		return err
	}
	const CF_UNICODETEXT = 13
	if err := winsys.OpenClipboard(0); err != nil {
		return err
	}

	defer winsys.CloseClipboard()
	return setClipboardDataSlice(CF_UNICODETEXT, data)
}

// GetFileGroupDescriptor returns a slice containing file metadata (filename + filesize) in the FileGroupDescriptorW slot
func GetFileGroupDescriptor() ([]FileInfo, error) {
	id, err := winsys.RegisterClipboardFormat("FileGroupDescriptorW")
	if err != nil {
		return nil, err
	}
	if err := winsys.OpenClipboard(0); err != nil {
		return nil, err
	}
	defer winsys.CloseClipboard()

	if err := winsys.IsClipboardFormatAvailable(id); err != nil {
		return nil, err
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

type comStreamWrapper struct {
	iStream *winsys.IStream
	medium  winsys.STGMEDIUM
}

func (s *comStreamWrapper) Read(buf []byte) (int, error) {
	return s.iStream.Read(buf)
}
func (s *comStreamWrapper) Close() error {
	return s.medium.Release()
}
func GetFileContents() ([]io.ReadCloser, error) {
	fds, err := GetFileGroupDescriptor()
	if err != nil {
		return nil, err
	}

	var dataObject *winsys.IDataObject
	if err := winsys.OleGetClipboard(&dataObject); err != nil {
		return nil, err
	}
	defer dataObject.Release()
	id1, _ := winsys.RegisterClipboardFormat("FileContents")

	const DVASPECT_CONTENT = 0x1
	format := winsys.FORMATETC{
		ClipFormat:     uint16(id1),
		DvTargetDevice: 0,
		Aspect:         DVASPECT_CONTENT,
		Index:          -1,
		Tymed:          uint32(winsys.TymedHGLOBAL | winsys.TymedISTREAM),
	}

	var result []io.ReadCloser
	for i := 0; i < len(fds); i++ {
		format.Index = int32(i)
		var medium winsys.STGMEDIUM
		if err := dataObject.GetData(&format, &medium); err != nil {
			return nil, err
		}

		stream, err := medium.Stream()
		if err != nil {
			medium.Release()
			return nil, err
		}
		result = append(result, &comStreamWrapper{iStream: stream, medium: medium})
	}

	return result, nil
}
func GetFileContent(index int) (io.ReadCloser, error) {
	var dataObject *winsys.IDataObject
	if err := winsys.OleGetClipboard(&dataObject); err != nil {
		return nil, err
	}
	defer dataObject.Release()
	id1, _ := winsys.RegisterClipboardFormat("FileContents")

	const DVASPECT_CONTENT = 0x1
	format := winsys.FORMATETC{
		ClipFormat:     uint16(id1),
		DvTargetDevice: 0,
		Aspect:         DVASPECT_CONTENT,
		Index:          int32(index),
		Tymed:          uint32(winsys.TymedHGLOBAL | winsys.TymedISTREAM),
	}

	var medium winsys.STGMEDIUM
	if err := dataObject.GetData(&format, &medium); err != nil {
		return nil, err
	}

	stream, err := medium.Stream()
	if err != nil {
		medium.Release()
		return nil, err
	}
	return &comStreamWrapper{iStream: stream, medium: medium}, nil
}

// returns a slice containing the filepaths in the H_DROP(15) slot
func GetHDROP() ([]string, error) {

	const id = 15
	// --------------

	// ----------

	if err := winsys.OpenClipboard(0); err != nil {
		return nil, err
	}
	defer winsys.CloseClipboard()

	if err := winsys.IsClipboardFormatAvailable(id); err != nil {
		return nil, err
	}

	h, err := winsys.GetClipboardData(id)
	if err != nil {
		return nil, err
	}
	buf := make([]uint16, 72)
	const NUM_ENTRIES = 0xFFFFFFFF
	n, err := dragQueryFileSlice(syscall.Handle(h), NUM_ENTRIES, nil)
	if err != nil {
		return nil, err
	}
	if n > 0 {
		var result []string
		for i := uint(0); i < n; i++ {
			reqBufSize, err := dragQueryFileSlice(syscall.Handle(h), i, nil)
			if err != nil {
				return nil, err
			}
			if int(reqBufSize) > len(buf) {
				buf = make([]uint16, reqBufSize+1)
			}
			n, err := dragQueryFileSlice(syscall.Handle(h), i, buf)
			if err != nil {
				return nil, err
			}
			result = append(result, string(utf16.Decode(buf[:n])))
		}
		return result, nil
	}
	return []string{}, nil
}

// Formats returns a slice that contains all formats currently avaiable in the clipboard
func Formats() ([]int, error) {
	if err := winsys.OpenClipboard(0); err != nil {
		return nil, err
	}
	defer winsys.CloseClipboard()

	var f uint32 = 0
	var err error
	var result []int
	retries := 0
	for {
		f, err = winsys.EnumClipboardFormats(f)
		if err != nil {
			if errors.Is(err, syscall.EINVAL) {
				break
			}
			if errors.Is(err, windows.ERROR_CLIPBOARD_NOT_OPEN) ||
				errors.Is(err, windows.ERROR_ACCESS_DENIED) {
				if retries < 3 {
					if err := winsys.OpenClipboard(0); err == nil {
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

var ErrUnknownClipboardFormat = errors.New("unkown clipboard format")

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
	return "", fmt.Errorf("unsupported format %d. %w", id, ErrUnknownClipboardFormat)
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
		n, err := winsys.GetClipboardFormatName(uint32(id), &buf[0], int32(len(buf)))
		if err != nil {
			return "", err
		}
		return string(utf16.Decode(buf[:n])), nil
	}

	return predefinedFormatName(uint(id))
}
