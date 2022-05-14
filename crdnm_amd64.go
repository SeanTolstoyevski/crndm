//go:build !noasm || !appengine
// +build !noasm !appengine


// package crndm (create randoms) is a rdrand instruction package.
//
// for more information, see: https://en.wikipedia.org/wiki/RDRAND
package crndm

import (
	"unsafe"
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

// CrmdnAsUint16 generates random number as uint64.
func CrmdnAsUint64() uint64 {
	x := (uint64)(uintptr(_crmdnAsUint64()))
	return x
}

// CrmdnAsUint32 generates random number as uint32.
func CrmdnAsUint32() uint32 {
	x := (uint32)(uintptr(_crmdnAsUint32()))
	return x
}

// CrmdnAsUint16 generates random number as uint16.
func CrmdnAsUint16() uint16 {
	x := (uint16)(uintptr(_crmdnAsUint16()))
	return x
}
