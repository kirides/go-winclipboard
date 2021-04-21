package winsys

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

//go:generate go run golang.org/x/sys/windows/mkwinsyscall -output zsyscall_windows.go syscall_windows.go

type (
	BOOL          uint32
	BOOLEAN       byte
	BYTE          byte
	DWORD         uint32
	DWORD64       uint64
	HANDLE        uintptr
	HLOCAL        uintptr
	LARGE_INTEGER int64
	LONG          int32
	LPVOID        uintptr
	SIZE_T        uintptr
	UINT          uint32
	ULONG_PTR     uintptr
	ULONGLONG     uint64
	WORD          uint16

	WPARAM   uintptr
	LPARAM   uintptr
	LRESULT  uintptr
	HOOKPROC func(int, WPARAM, LPARAM) LRESULT
)

//sys	setWindowsHookExW(idHook int32, lpfn unsafe.Pointer, hmod syscall.Handle, dwThreadId uint32) (h syscall.Handle, err error) = User32.SetWindowsHookExW

//sys	OpenClipboard(h syscall.Handle) (err error) = User32.OpenClipboard
//sys	CloseClipboard() (err error) = User32.CloseClipboard
//sys	EmptyClipboard() (err error) = User32.EmptyClipboard
//sys	RegisterClipboardFormat(name string) (id uint32, err error) = User32.RegisterClipboardFormatW
//sys	EnumClipboardFormats(format uint32) (id uint32, err error) = User32.EnumClipboardFormats
//sys	GetClipboardFormatName(format uint32, lpszFormatName *uint16, cchMaxCount int32) (len int32, err error) = User32.GetClipboardFormatNameW
//sys	GetClipboardData(uFormat uint32) (h syscall.Handle, err error) = User32.GetClipboardData
//sys	SetClipboardData(uFormat uint32, hMem syscall.Handle) (h syscall.Handle, err error) = User32.SetClipboardData
//sys	IsClipboardFormatAvailable(uFormat uint32) (err error) = User32.IsClipboardFormatAvailable
//sys	AddClipboardFormatListener(hWnd syscall.Handle) (err error) = User32.AddClipboardFormatListener
//sys	RemoveClipboardFormatListener(hWnd syscall.Handle) (err error) = User32.RemoveClipboardFormatListener

//sys	GetProcessHeap() (hHeap syscall.Handle, err error) = Kernel32.GetProcessHeap
//sys	HeapAlloc(hHeap syscall.Handle, dwFlags uint32, dwSize uintptr) (lpMem uintptr, err error) = Kernel32.HeapAlloc
//sys	HeapFree(hHeap syscall.Handle, dwFlags uint32, lpMem uintptr) (err error) = Kernel32.HeapFree
//sys	HeapSize(hHeap syscall.Handle, dwFlags uint32, lpMem uintptr) (size uintptr, err error) [failretval==^uintptr(r0)] = Kernel32.HeapSize

//sys	DragQueryFile(hDrop syscall.Handle, iFile int, buf *uint16, len uint32) (n int, err error) = Shell32.DragQueryFileW

var (

	// Clipboard
	procSHGetPathFromIDListEx = modShell32.NewProc("SHGetPathFromIDListEx")
	procSHGetKnownFolderPath  = modShell32.NewProc("SHGetKnownFolderPath")
)

func GetDesktopFolder(token uintptr) (string, error) {
	var desktopId = windows.GUID{Data1: 0xB4BFCC3A, Data2: 0xDB2C, Data3: 0x424C, Data4: [8]byte{0xB0, 0x29, 0x7F, 0xE9, 0x9A, 0x87, 0xC6, 0x41}}
	var retVal unsafe.Pointer
	r1, _, e1 := syscall.Syscall6(procSHGetKnownFolderPath.Addr(), 4, uintptr(unsafe.Pointer(&desktopId)), uintptr(0), uintptr(token), uintptr(unsafe.Pointer(&retVal)), 0, 0)
	if r1 != 0 {
		return "", e1
	}
	defer windows.CoTaskMemFree(retVal)
	v := windows.UTF16PtrToString((*uint16)(retVal))
	return v, nil
}
func ShGetPathFromIDList(pidl uintptr, buf []uint16) (err error) {
	r1, _, e1 := syscall.Syscall6(procSHGetPathFromIDListEx.Addr(), 4, uintptr(pidl), uintptr(unsafe.Pointer(&buf[0])), uintptr(len(buf)), 0, 0, 0)
	if r1 == 0 {
		err = e1
	}
	return
}
