package main

import "unsafe"

func IntToByteArray(num uint32) []byte {
	const size = uint32(unsafe.Sizeof(num))
	arr := (*(*[size]byte)(unsafe.Pointer(&num)))[:]
	return arr
}

func ByteArrayToInt(arr []byte) uint32 {
	val := uint32(0)
	size := len(arr)
	for i := 0; i < size; i++ {
		*(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(&val)) + uintptr(i))) = arr[i]
	}
	return val
}
