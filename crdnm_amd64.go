//+build !noasm !appengine
package main

import (
	"unsafe"
//	"fmt"
)

//go:nosplit
//go:noescape
func _crmdnAsUint64() (v unsafe.Pointer)

//go:nosplit
//go:noescape
func _crmdnAsUint32() (v unsafe.Pointer)

//go:nosplit
//go:noescape
func _crmdnAsUint16() (v unsafe.Pointer)


func CrmdnAsUint64() uint64 {
	x := (uint64)(uintptr(_crmdnAsUint64()))
	return x
}

func CrmdnAsUint32() uint32 {
	x := (uint32)(uintptr(_crmdnAsUint32()))
	return x
}

func CrmdnAsUint16() uint16 {
	x := (uint16)(uintptr(_crmdnAsUint64()))
	return x
}

func main() {
	x := CrmdnAsUint64()
	println(x)
}
