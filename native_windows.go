package clipboard

import (
	"fmt"
	"syscall"
	"unsafe"

	"github.com/kirides/go-winclipboard/internal/winsys"
	"golang.org/x/sys/windows"
)

// Contains Win32 API wrappers and structs

// SECTION : CLIPBOARD

func dragQueryFileSlice(hDrop syscall.Handle, iFile int, buf []uint16) (int, error) {
	var b *uint16
	if len(buf) > 0 {
		b = &buf[0]
	}
	return winsys.DragQueryFile(hDrop, iFile, b, uint32(len(buf)))
}

func setClipboardDataSlice(id uint32, value []byte) error {
	hHeap, err := winsys.GetProcessHeap()
	if hHeap == 0 {
		return err
	}
	hMem, err := winsys.HeapAlloc(hHeap, 0, uintptr(len(value)))
	if err != nil {
		return err
	}
	copy((*[1 << 31]byte)(unsafe.Pointer(hMem))[:len(value):len(value)], value)
	r, e1 := winsys.SetClipboardData(id, syscall.Handle(hMem))
	if r != syscall.Handle(hMem) {
		e2 := winsys.HeapFree(hHeap, 0, hMem)
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
