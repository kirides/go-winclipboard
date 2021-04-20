package clipboard

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Contains Win32 API wrappers and structs

var (

	// Clipboard
	procSHGetPathFromIDListEx = modShell32.NewProc("SHGetPathFromIDListEx")
	procSHGetKnownFolderPath  = modShell32.NewProc("SHGetKnownFolderPath")
)

// SECTION : CLIPBOARD

var desktopId = windows.GUID{Data1: 0xB4BFCC3A, Data2: 0xDB2C, Data3: 0x424C, Data4: [8]byte{0xB0, 0x29, 0x7F, 0xE9, 0x9A, 0x87, 0xC6, 0x41}}

func getDesktopFolder(token uintptr) (string, error) {
	var retVal unsafe.Pointer
	r1, _, e1 := syscall.Syscall6(procSHGetKnownFolderPath.Addr(), 4, uintptr(unsafe.Pointer(&desktopId)), uintptr(0), uintptr(token), uintptr(unsafe.Pointer(&retVal)), 0, 0)
	if r1 != 0 {
		return "", e1
	}
	defer windows.CoTaskMemFree(retVal)
	v := windows.UTF16PtrToString((*uint16)(retVal))
	return v, nil
}

func dragQueryFileSlice(hDrop syscall.Handle, iFile int, buf []uint16) (int, error) {
	var b *uint16
	if len(buf) > 0 {
		b = &buf[0]
	}
	return dragQueryFile(hDrop, iFile, b, uint32(len(buf)))
}

func shGetPathFromIDListW(pidl uintptr, buf []uint16) (err error) {
	r1, _, e1 := syscall.Syscall6(procSHGetPathFromIDListEx.Addr(), 4, uintptr(pidl), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)), 0, 0, 0)
	if r1 == 0 {
		err = e1
	}
	return
}

func setClipboardDataSlice(id uint32, value []byte) error {
	hHeap, err := getProcessHeap()
	if hHeap == 0 {
		return err
	}
	hMem, err := heapAlloc(hHeap, 0, uintptr(len(value)))
	if err != nil {
		return err
	}
	copy((*[1 << 31]byte)(unsafe.Pointer(hMem))[:len(value):len(value)], value)
	r, e1 := setClipboardData(id, syscall.Handle(hMem))
	if r != syscall.Handle(hMem) {
		e2 := heapFree(hHeap, 0, hMem)
		if e2 != nil {
			return fmt.Errorf("failed to free heap memory: %v. %w", e2, e1)
		}
		return e1
	}
	return nil
}

type _FILEGROUPDESCRIPTORW struct {
	nItems uint32
	fgd    *_FILEDESCRIPTORW
}

const _MAX_PATH = 260

type _FILEDESCRIPTORW struct {
	dwFlags          uint32
	clsid            windows.GUID
	size             [2]int32
	pointl           [2]int32
	dwFileAttributes uint32
	ftCreationTime   windows.Filetime
	ftLastAccessTime windows.Filetime
	ftLastWriteTime  windows.Filetime
	nFileSizeHigh    uint32
	nFileSizeLow     uint32
	cFileName        [_MAX_PATH]uint16
}
