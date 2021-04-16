package clipboard

import (
	"reflect"
	"unsafe"
)

type FileInfo struct {
	Name string
	Size int64
}

func byteSliceFromUintptr(v uintptr, len int) []byte {
	res := []byte{}
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&res))
	hdr.Data = v
	hdr.Len = len
	return res
}
