package clipboard

//go:generate mkwinsyscall -output zsyscall_windows.go syscall_windows.go

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

//sys	openClipboard(h syscall.Handle) (err error) = User32.OpenClipboard
//sys	closeClipboard() (err error) = User32.CloseClipboard
//sys	emptyClipboard() (err error) = User32.EmptyClipboard
//sys	registerClipboardFormat(name string) (id uint32, err error) = User32.RegisterClipboardFormatW
//sys	enumClipboardFormats(format uint32) (id uint32, err error) = User32.EnumClipboardFormats
//sys	getClipboardFormatName(format uint32, lpszFormatName *uint16, cchMaxCount int32) (len int32, err error) = User32.GetClipboardFormatNameW
//sys	getClipboardData(uFormat uint32) (h syscall.Handle, err error) = User32.GetClipboardData
//sys	setClipboardData(uFormat uint32, hMem syscall.Handle) (h syscall.Handle, err error) = User32.SetClipboardData
//sys	isClipboardFormatAvailable(uFormat uint32) (err error) = User32.IsClipboardFormatAvailable
//sys	AddClipboardFormatListener(hWnd syscall.Handle) (err error) = User32.AddClipboardFormatListener
//sys	RemoveClipboardFormatListener(hWnd syscall.Handle) (err error) = User32.RemoveClipboardFormatListener

//sys	getProcessHeap() (hHeap syscall.Handle, err error) = Kernel32.GetProcessHeap
//sys	heapAlloc(hHeap syscall.Handle, dwFlags uint32, dwSize uintptr) (lpMem uintptr, err error) = Kernel32.HeapAlloc
//sys	heapFree(hHeap syscall.Handle, dwFlags uint32, lpMem uintptr) (err error) = Kernel32.HeapFree
//sys	heapSize(hHeap syscall.Handle, dwFlags uint32, lpMem uintptr) (size uintptr, err error) [failretval==^uintptr(r0)] = Kernel32.HeapSize

//sys	dragQueryFile(hDrop syscall.Handle, iFile int, buf *uint16, len uint32) (n int, err error) = Shell32.DragQueryFileW
