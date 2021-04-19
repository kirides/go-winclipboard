package clipboard

import (
	"fmt"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// Contains Win32 API wrappers and structs

var (
	modUser32   = windows.NewLazySystemDLL("User32.dll")
	modShell32  = windows.NewLazySystemDLL("Shell32.dll")
	modKernel32 = windows.NewLazySystemDLL("Kernel32.dll")

	// Clipboard

	procOpenClipboard                 = modUser32.NewProc("OpenClipboard")
	procCloseClipboard                = modUser32.NewProc("CloseClipboard")
	procEmptyClipboard                = modUser32.NewProc("EmptyClipboard")
	procRegisterClipboardFormat       = modUser32.NewProc("RegisterClipboardFormatW")
	procEnumClipboardFormats          = modUser32.NewProc("EnumClipboardFormats")
	procGetClipboardFormatName        = modUser32.NewProc("GetClipboardFormatNameW")
	procAddClipboardFormatListener    = modUser32.NewProc("AddClipboardFormatListener")
	procRemoveClipboardFormatListener = modUser32.NewProc("RemoveClipboardFormatListener")
	procGetClipboardData              = modUser32.NewProc("GetClipboardData")
	procSetClipboardData              = modUser32.NewProc("SetClipboardData")
	procIsClipboardFormatAvailable    = modUser32.NewProc("IsClipboardFormatAvailable")

	procDragQueryFileW        = modShell32.NewProc("DragQueryFileW")
	procSHGetPathFromIDListEx = modShell32.NewProc("SHGetPathFromIDListEx")
	procSHGetKnownFolderPath  = modShell32.NewProc("SHGetKnownFolderPath")

	procGetProcessHeap = modKernel32.NewProc("GetProcessHeap")
	procHeapSize       = modKernel32.NewProc("HeapSize")
	procHeapAlloc      = modKernel32.NewProc("HeapAlloc")
	procHeapFree       = modKernel32.NewProc("HeapFree")
)

func getProcessHeap() (syscall.Handle, error) {
	r1, _, e1 := syscall.Syscall(procGetProcessHeap.Addr(), 0, 0, 0, 0)
	if r1 == 0 {
		return 0, e1
	}
	return syscall.Handle(r1), e1
}
func heapSize(hHeap syscall.Handle, dwFlags uint32, lpMem uintptr) (uintptr, error) {
	r1, _, e1 := syscall.Syscall(procHeapSize.Addr(), 3, uintptr(hHeap), uintptr(dwFlags), uintptr(lpMem))
	if r1 == unsignedMinusOne {
		return 0, e1
	}
	return r1, e1
}
func heapAlloc(hHeap syscall.Handle, dwFlags uint32, dwSize uint32) (uintptr, error) {
	r1, _, _ := syscall.Syscall(procHeapAlloc.Addr(), 3, uintptr(hHeap), uintptr(dwFlags), uintptr(dwSize))
	if r1 == 0 {
		return 0, fmt.Errorf("Could not allocate")
	}
	return r1, nil
}
func heapFree(hHeap syscall.Handle, dwFlags uint32, lpMem uintptr) error {
	r1, _, e1 := syscall.Syscall(procHeapFree.Addr(), 3, uintptr(hHeap), uintptr(dwFlags), uintptr(lpMem))
	if r1 == 0 {
		return e1
	}
	return nil
}

// SECTION : CLIPBOARD

func getDesktopFolder(token uintptr) (string, error) {
	desktopId := windows.GUID{Data1: 0xB4BFCC3A, Data2: 0xDB2C, Data3: 0x424C, Data4: [8]byte{0xB0, 0x29, 0x7F, 0xE9, 0x9A, 0x87, 0xC6, 0x41}}
	var retVal unsafe.Pointer
	r1, _, e1 := syscall.Syscall6(procSHGetKnownFolderPath.Addr(), 4, uintptr(unsafe.Pointer(&desktopId)), uintptr(0), uintptr(token), uintptr(unsafe.Pointer(&retVal)), 0, 0)
	if r1 != 0 {
		return "", e1
	}
	defer windows.CoTaskMemFree(retVal)
	v := windows.UTF16PtrToString((*uint16)(retVal))
	return v, nil
}

func dragQueryFile(hDrop syscall.Handle, iFile int, buf []uint16) (int, error) {
	var r1 uintptr
	var e1 syscall.Errno
	if buf == nil {
		r1, _, e1 = syscall.Syscall6(procDragQueryFileW.Addr(), 4, uintptr(hDrop), uintptr(iFile), uintptr(0), uintptr(0), 0, 0)
	} else {
		r1, _, e1 = syscall.Syscall6(procDragQueryFileW.Addr(), 4, uintptr(hDrop), uintptr(iFile), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)), 0, 0)
	}
	if r1 == 0 {
		return 0, e1
	}
	return int(r1), nil
}

func shGetPathFromIDListW(pidl uintptr, buf []uint16) (err error) {
	r1, _, e1 := syscall.Syscall6(procSHGetPathFromIDListEx.Addr(), 4, uintptr(pidl), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)), 0, 0, 0)
	if r1 == 0 {
		err = e1
	}
	return
}

func openClipboard(hWndNewOwner syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procOpenClipboard.Addr(), 1, uintptr(hWndNewOwner), 0, 0)
	if r1 == 0 {
		err = e1
	}
	return
}
func emptyClipboard() (err error) {
	r1, _, e1 := syscall.Syscall(procEmptyClipboard.Addr(), 0, 0, 0, 0)
	if r1 == 0 {
		err = e1
	}
	return
}
func getClipboardData(id uint32) (h uintptr, err error) {
	r1, _, e1 := syscall.Syscall(procGetClipboardData.Addr(), 1, uintptr(id), 0, 0)
	if r1 == 0 {
		err = e1
	}
	h = r1
	return
}
func setClipboardData(id uint32, value []byte) error {
	hHeap, err := getProcessHeap()
	if hHeap == 0 {
		return err
	}
	hMem, err := heapAlloc(hHeap, 0, uint32(len(value)))
	if err != nil {
		return err
	}
	b := byteSliceFromUintptr(uintptr(hMem), len(value))
	copy(b, value)

	r1, _, e1 := syscall.Syscall(procSetClipboardData.Addr(), 2, uintptr(id), uintptr(hMem), 0)
	if r1 == 0 {
		e2 := heapFree(hHeap, 0, hMem)
		if e2 != nil {
			return fmt.Errorf("Failed to free heap memory: %v. %w", e2, e1)
		}
		return e1
	}
	return nil
}

func isClipboardFormatAvailable(id uint32) (ok bool, err error) {
	r1, _, e1 := syscall.Syscall(procIsClipboardFormatAvailable.Addr(), 1, uintptr(id), 0, 0)
	if r1 == 0 {
		err = e1
	}
	ok = r1 != 0
	return
}

func closeClipboard() (err error) {
	r1, _, e1 := syscall.Syscall(procCloseClipboard.Addr(), 0, 0, 0, 0)
	if r1 == 0 {
		err = e1
	}
	return
}

func registerClipboardFormat(format string) (id uint32, err error) {
	ptr, e1 := syscall.UTF16PtrFromString(format)
	if e1 != nil {
		err = e1
		return
	}
	r1, _, e1 := syscall.Syscall(procRegisterClipboardFormat.Addr(), 1, uintptr(unsafe.Pointer(ptr)), 0, 0)
	if r1 == 0 {
		err = e1
	}
	id = uint32(r1)
	return
}

func enumClipboardFormats(format uint) (uint, error) {
	r1, _, e1 := syscall.Syscall(procEnumClipboardFormats.Addr(), 1, uintptr(format), 0, 0)
	if r1 == 0 {
		return 0, e1
	}
	return uint(r1), nil
}

func getClipboardFormatName(format uint, buf []uint16) (n int, err error) {
	r1, _, e1 := syscall.Syscall(procGetClipboardFormatName.Addr(), 3, uintptr(format), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)))
	if r1 == 0 {
		err = e1
	}
	n = int(r1)
	return
}

func AddClipboardFormatListener(hwnd syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procAddClipboardFormatListener.Addr(), 1, uintptr(hwnd), 0, 0)
	if r1 == 0 {
		err = e1
	}
	return
}
func RemoveClipboardFormatListener(hwnd syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procRemoveClipboardFormatListener.Addr(), 1, uintptr(hwnd), 0, 0)
	if r1 == 0 {
		err = e1
	}
	return
}

type _FILEGROUPDESCRIPTORW struct {
	nItems uint32
	fgd    [1]_FILEDESCRIPTORW
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
