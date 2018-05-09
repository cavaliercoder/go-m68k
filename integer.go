package m68k

import (
	"unsafe"
)

var (
	signMask = []uint32{
		0x00000080, // SizeByte
		0x00008000, // SizeWord
		0x80000000, // SizeLong
	}

	sizeFilter = []uint32{
		0x000000FF, // SizeByte
		0x0000FFFF, // SizeWord
		0xFFFFFFFF, // SizeLong
	}

	sizeMask = []uint32{
		0xFFFFFF00, // SizeByte
		0xFFFF0000, // SizeWord
		0x00000000, // SizeLong
	}
)

// Int represents M68000 internal encoding of an integer and provides methods
// for emulating the its arithmetic behaviors.
type Int struct {
	value, flags uint32
}

// NewInt returns a new Int initialized with bits v, of size sz.
func NewInt(v uint32, sz int) Int {
	if sz < 0 || sz > SizeLong {
		panic("invalid size")
	}
	i := Int{
		value: v,
		flags: uint32(sz) << 6,
	}
	return i
}

func (i Int) Size() int {
	return int(i.flags >> 6)
}

func (i Int) Status() uint32 {
	return uint32(i.flags & 0x0F)
}

// SignBit returns only the sign bit of the integer, or zero if it is not set.
func (i Int) SignBit() uint32 {
	return i.value & signMask[i.Size()]
}

// IsNegative returns true if the sign it is set for the integer.
func (i Int) IsNegative() bool {
	return signBit(i.value, i.Size()) != 0
}

// Byte returns the 8 least significant bits of the integer as a byte.
func (i Int) Byte() byte {
	return byte(i.value & 0xFF)
}

// Word returns the 16 least significant bits of the integer as a uint16.
func (i Int) Word() uint16 {
	return uint16(i.value & 0xFFFF)
}

// Long returns all 32 bits of the integer as a uint32.
func (i Int) Long() uint32 {
	return i.value
}

// SignedLong returns all 32 bits of the integer as an int32.
func (i Int) SignedLong() int32 {
	i = i.SignExtend(SizeLong)
	return *(*int32)(unsafe.Pointer(&i.value))
}

// SignExtend returns an Int extended to the given size using signed arithmetic.
func (i Int) SignExtend(sz int) Int {
	if i.IsNegative() {
		i.value |= sizeMask[i.Size()] & sizeFilter[sz]
	}
	return i
}

// Add returns the sum of two integers and sets its status code to enable
// overflow and carry detection, etc. The number of bits mutated is limited by
// the width of the operand. For example, if a Byte is added to a Word, only the
// eight least significant bits will be modified by the function. The first 24
// bit of the result will equal the first 24 of the Int.
func (i Int) Add(v Int) Int {
	sz := v.Size()
	a := i.value & sizeFilter[sz]
	b := v.value & sizeFilter[sz]
	n := a + b

	// keep uneffected bytes
	i.value = i.value&sizeMask[sz] | n&sizeFilter[sz]

	// set status flags
	as, bs, ns := signBit(a, sz), signBit(b, sz), signBit(n, sz)
	if n == 0 {
		i.flags |= StatusZero
	}
	if ns != 0 {
		i.flags |= StatusNegative
	}
	if as == bs && as != ns {
		// see Hacker's Delight pg. 2-13
		i.flags |= StatusOverflow
	}
	if xor(xor(n, a), b) != 0 {
		// given: c = (x + y) ⊕ x ⊕ y
		i.flags |= StatusCarry
	}

	return i
}

// Sub returns the difference between two integers, similarly to Add.
func (i Int) Sub(v Int) Int {
	// given: x - y = x + ¬y + 1
	i = i.Add(NewInt(1, i.Size()))
	v.value = ^v.value
	return i.Add(v)
}

// signBit returns only the sign bit for the given value of size sz.
func signBit(v uint32, sz int) uint32 {
	return v & signMask[sz]
}

// xor returns a ⊕ b
func xor(a, b uint32) uint32 {
	// given: x ⊕ y = (x | y) - (x & y)
	return (a | b) - (a & b)
}
