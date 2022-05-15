//go:build !noasm || !appengine
// +build !noasm !appengine

package crndm

import (
	"testing"
)

// func TestRdseedAsuint16(t *testing.T) {
// 	const rangeMax = 65100
// 	var v uint16
// 	for i := 1000; i < rangeMax; i++ {
// 		v = uint16(i)
// 		if !RdseedAsUint16(&v) {
// 			t.Fatalf("RdseedAsuint16 error: %d", i)
// 		}
// 	}
// }

// func TestRdseedAsUint32(t *testing.T) {
// 	const rangeMax = 4000000
// 	var v uint32
// 	for i := rangeMax - 1000; i < rangeMax; i++ {
// 		v = uint32(i)
// 		if !RdseedAsUint32(&v) {
// 			t.Fatalf("RdseedAsuint32 error: %d", i)
// 		}
// 	}
// }

// func TestRdseedAsuint64(t *testing.T) {
// 	const rangeMax = 30303030303
// 	var v uint64
// 	for i := rangeMax - 1000; i < rangeMax; i++ {
// 		v = uint64(i)
// 		if !RdseedAsUint64(&v) {
// 			t.Fatalf("RdseedAsuint64 error: %d", i)
// 		}
// 	}
// }

// func TestRdseedAsUint64_2(t *testing.T) {
// 	t.Log("working...")
// 	v := uint64(94821932)
// 	if !RdseedAsUint64(&v) {
// 		t.Fatal("TestRdseedAsuint64_2")
// 	}
// }

func TestCrmdnAsUint64(t *testing.T) {
	for i := 0; i < 150; i++ {
		if CrmdnAsUint64() == 0 {
			t.Fatal("CrmdnAsUint64 returns zero")
		}
	}
}

func BenchmarkCrmdnAsUint64(b *testing.B) {
	var v uint64
	for i := 0; i < b.N; i++ {
		v = CrmdnAsUint64()
	}
	_ = v
}

func TestCrmdnAsUint32(t *testing.T) {
	for i := 0; i < 150; i++ {
		if CrmdnAsUint32() == 0 {
			t.Fatal("CrmdnAsUint32 returns zero")
		}
	}
}

func BenchmarkCrmdnAsUint32(b *testing.B) {
	var v uint32
	for i := 0; i < b.N; i++ {
		v = CrmdnAsUint32()
	}
	_ = v
}

func TestCrmdnAsUint16(t *testing.T) {
	for i := 0; i < 150; i++ {
		if CrmdnAsUint16() == 0 {
			t.Fatal("CrmdnAsUint16 returns zero")
		}
	}
}

func BenchmarkCrmdnAsUint16(b *testing.B) {
	var v uint16
	for i := 0; i < b.N; i++ {
		v = CrmdnAsUint16()
	}
	_ = v
}
