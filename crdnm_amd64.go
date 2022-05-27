// Copyrigth 2022 - mailto:seantolstoyevski@protonmail.com
//
// The source codes of this project are licensed under the MIT license.
//
// It is provided without any warranty.

//go:build !noasm || !appengine
// +build !noasm !appengine

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

// CrmdnAsUint64 generates random number as uint64.
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

//go:noescape
func _RdseedAsUint16(val unsafe.Pointer) (r unsafe.Pointer)

//go:noescape
func _RdseedAsUint32(val unsafe.Pointer) (r unsafe.Pointer)

//go:noescape
func _RdseedAsUint64(val unsafe.Pointer) (r unsafe.Pointer)

func RdseedAsUint16(seed *uint16) bool {
	r := _RdseedAsUint16(unsafe.Pointer(seed))
	i := (uint8)(uintptr(r))
	return i == 1
}

func RdseedAsUint32(seed *uint32) bool {
	r := _RdseedAsUint32(unsafe.Pointer(seed))
	i := (uint8)(uintptr(r))
	return i == 1
}

func RdseedAsUint64(seed *uint64) bool {
	r := _RdseedAsUint64(unsafe.Pointer(seed))
	i := (uint8)(uintptr(r))
	return i == 1
}
