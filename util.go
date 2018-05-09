package m68k

// byteToInt32 sign extends the given byte to an Int32.
func byteToInt32(b byte) int32 {
	return NewInt(uint32(b), SizeByte).SignedLong()
}

// wordToInt32 sign extends the given word to an Int32.
func wordToInt32(n uint16) int32 {
	return NewInt(uint32(n), SizeWord).SignedLong()
}

// longToInt32 decodes the given long as a signed Int32.
func longToInt32(n uint32) int32 {
	return int32(n)
}

// decodeEA returns the effective address segment (bits 0 - 5) of an opcode.
func decodeEA(op uint16) uint16 {
	return op & 0x003F
}

// decodeExEA returns the effective destination address segment of an opcode,
// for instructions that have an effective address in bit 6 - 11. The result is
// flipped so that the mode and register portions of the ea match the encoding
// used by typical ea fields.
func decodeExEA(op uint16) uint16 {
	return ((op & 0x01C0) >> 3) | ((op & 0x0E00) >> 9)
}
