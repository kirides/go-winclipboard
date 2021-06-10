package winsys

import (
	"fmt"
	"io"
	"syscall"
	"unsafe"
)

type iUnknownVtbl struct {
	// every COM object starts with these three
	QueryInterface uintptr
	AddRef         uintptr
	Release        uintptr
}

type ISequentialStreamVtbl struct {
	iUnknownVtbl
	Read  uintptr
	Write uintptr
}

type IStreamVtbl struct {
	ISequentialStreamVtbl
	Seek         uintptr
	SetSize      uintptr
	CopyTo       uintptr
	Commit       uintptr
	Revert       uintptr
	LockRegion   uintptr
	UnlockRegion uintptr
	Stat         uintptr
	Clone        uintptr
}

type IStream struct {
	vtbl *IStreamVtbl
}

func (obj *IStream) Read(buffer []byte) (int, error) {
	bufPtr := &buffer[0]
	var read uint32
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.Read,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(bufPtr)),
		uintptr(len(buffer)),
		uintptr(unsafe.Pointer(&read)),
		0,
		0,
	)
	if ret == _S_FALSE {
		return int(read), io.EOF
	}
	if ret != _S_OK {
		return int(read), HRESULT(ret)
	}
	return int(read), nil
}
func (obj *IStream) Close() error {
	return obj.Release()
}
func (obj *IStream) Release() error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	if ret != _S_OK {
		return HRESULT(ret)
	}
	return nil
}

type Tymed uint32

const (
	TymedNULL     Tymed = 0x0
	TymedHGLOBAL  Tymed = 0x1
	TymedFILE     Tymed = 0x2
	TymedISTREAM  Tymed = 0x4
	TymedISTORAGE Tymed = 0x8
	TymedGDI      Tymed = 0x10
	TymedMFPICT   Tymed = 0x20
	TymedENHMF    Tymed = 0x40
)

type IUnknown struct {
	vtbl iUnknownVtbl
}

func (obj *IUnknown) Release() error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	if ret != _S_OK {
		return HRESULT(ret)
	}
	return nil
}

type iDataObjectVtbl struct {
	iUnknownVtbl
	GetData               uintptr
	GetDataHere           uintptr
	QueryGetData          uintptr
	GetCanonicalFormatEtc uintptr
	SetData               uintptr
	EnumFormatEtc         uintptr
	DAdvise               uintptr
	DUnadvise             uintptr
	EnumDAdvise           uintptr
}

func (m STGMEDIUM) Release() error {
	if m.PUnkForRelease == nil {
		return ReleaseStgMedium(&m)
	}
	return nil
}

func (m STGMEDIUM) Stream() (*IStream, error) {
	if m.Tymed != TymedISTREAM {
		return nil, fmt.Errorf("invalid Tymed")
	}
	return (*IStream)(unsafe.Pointer(m.UnionMember)), nil
}

func (m STGMEDIUM) Bytes() ([]byte, error) {
	if m.Tymed != TymedHGLOBAL {
		return nil, fmt.Errorf("invalid Tymed")
	}
	size, err := GlobalSize(m.UnionMember)
	if err != nil {
		return nil, fmt.Errorf("could not determine size")
	}

	lpMem, err := GlobalLock(m.UnionMember)
	if err != nil {
		return nil, fmt.Errorf("could not determine size")
	}
	defer GlobalUnlock(m.UnionMember)

	data := (*[1 << 30]byte)(unsafe.Pointer(lpMem))

	result := make([]byte, size)
	copy(result, data[:size])
	return result, nil
}

type IDataObject struct {
	vtbl *iDataObjectVtbl
}

func (obj *IDataObject) Release() error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	if ret != _S_OK {
		return HRESULT(ret)
	}
	return nil
}

func (obj *IDataObject) GetData(formatEtc *FORMATETC, medium *STGMEDIUM) error {
	s2 := unsafe.Sizeof(*medium)
	_ = s2
	// return obj.getDataUnsafe(formatEtc, medium)
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetData,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(formatEtc)),
		uintptr(unsafe.Pointer(medium)),
	)

	if ret != _S_OK {
		return HRESULT(ret)
	}
	return nil
}

func (obj *IDataObject) GetDataHere(formatEtc *FORMATETC, medium *STGMEDIUM) error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.GetDataHere,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(formatEtc)),
		uintptr(unsafe.Pointer(medium)),
	)

	if ret != _S_OK {
		return HRESULT(ret)
	}
	return nil
}
func (obj *IDataObject) QueryGetData(formatEtc *FORMATETC) error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.QueryGetData,
		2,
		uintptr(unsafe.Pointer(obj)),
		uintptr(unsafe.Pointer(formatEtc)),
		0,
	)

	if ret != _S_OK {
		return HRESULT(ret)
	}
	return nil
}
func (obj *IDataObject) EnumFormatEtc(direction uint32, pIEnumFORMATETC **IEnumFORMATETC) error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.EnumFormatEtc,
		3,
		uintptr(unsafe.Pointer(obj)),
		uintptr(direction),
		uintptr(unsafe.Pointer(pIEnumFORMATETC)),
	)

	if ret != _S_OK {
		return HRESULT(ret)
	}
	return nil
}

type iEnumFORMATETCVtbl struct {
	iUnknownVtbl
	Next  uintptr
	Skip  uintptr
	Reset uintptr
	Clone uintptr
}
type IEnumFORMATETC struct {
	vtbl *iEnumFORMATETCVtbl
}

func (obj *IEnumFORMATETC) Release() error {
	ret, _, _ := syscall.Syscall(
		obj.vtbl.Release,
		1,
		uintptr(unsafe.Pointer(obj)),
		0,
		0,
	)
	if ret != _S_OK {
		return HRESULT(ret)
	}
	return nil
}

func (obj *IEnumFORMATETC) Next(formatEtc []FORMATETC, celtFetched *uint32) error {
	ret, _, _ := syscall.Syscall6(
		obj.vtbl.Next,
		4,
		uintptr(unsafe.Pointer(obj)),
		uintptr(len(formatEtc)),
		uintptr(unsafe.Pointer(&formatEtc[0])),
		uintptr(unsafe.Pointer(celtFetched)),
		0,
		0,
	)
	if ret != 1 {
		return HRESULT(ret)
	}
	return nil
}
