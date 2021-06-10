package winsys

type FORMATETC struct {
	ClipFormat     uint16
	_padding0      [2]byte
	DvTargetDevice uintptr
	Aspect         uint32
	Index          int32
	Tymed          uint32
}

type STGMEDIUM struct {
	Tymed          Tymed
	UnionMember    uintptr
	PUnkForRelease uintptr
}
