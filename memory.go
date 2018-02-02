package m68k

import (
	"io"
	"unsafe"
)

// Memory is an interface for any IO device that is addressable via memory
// mapping. This interface will be satisfied by random access memory, a virtual
// memory manager, peripheral devices, etc.
type Memory interface {
	Read(addr int, p []byte) (n int, err error)
	Write(addr int, p []byte) (n int, err error)
}

type ram struct {
	b []byte
}

// NewRAM returns Random Access Memory initialized to the given size. This
// memory can used to load programs and data for a 68000 processor.
func NewRAM(size uint32) Memory {
	return &ram{
		b: make([]byte, size),
	}
}

func (m *ram) Read(addr int, p []byte) (n int, err error) {
	if addr < 0 || addr >= len(m.b) {
		return 0, io.EOF
	}
	n = copy(p, m.b[addr:])
	return
}

func (m *ram) Write(addr int, p []byte) (n int, err error) {
	if addr < 0 || addr >= len(m.b) {
		return 0, ErrAddressOutOfBounds
	}
	n = copy(m.b[addr:], p)
	// TODO: raise exception vector instead of returning errors?
	if n < len(p) {
		return n, io.ErrShortWrite
	}
	return
}

// A MemoryDecoder provides native type decoding for the underlying Memory
// interface.
type MemoryDecoder struct {
	M Memory

	buf [4]byte
}

// Write writes to the underlying Memory interface.
func (m *MemoryDecoder) Write(addr int, p []byte) (n int, err error) {
	return m.M.Write(addr, p)
}

// Read reads from the underlying Memory interface.
func (m *MemoryDecoder) Read(addr int, p []byte) (n int, err error) {
	return m.M.Read(addr, p)
}

// Byte reads a single byte from the underlying Memory interface.
func (m *MemoryDecoder) Byte(addr int) (b byte, err error) {
	_, err = m.M.Read(addr, m.buf[:1])
	b = m.buf[0]
	return
}

// Word reads a single 16-bit unsigned word from the underlying Memory
// interface.
func (m *MemoryDecoder) Word(addr int) (n uint16, err error) {
	_, err = m.M.Read(addr, m.buf[:2])
	n = uint16(m.buf[0])<<8 | uint16(m.buf[1])
	return
}

// Sword reads a single 16-bit signed word from the underlying Memory interface.
func (m *MemoryDecoder) Sword(addr int) (n int16, err error) {
	_, err = m.M.Read(addr, m.buf[:2])
	u := uint16(m.buf[0])<<8 | uint16(m.buf[1])
	n = *(*int16)(unsafe.Pointer(&u))
	return
}

// Long reads a single 32-bit unsigned long from the underlying Memory
// interface.
func (m *MemoryDecoder) Long(addr int) (n uint32, err error) {
	_, err = m.M.Read(addr, m.buf[:4])
	n = uint32(m.buf[0])<<24 | uint32(m.buf[1])<<16 | uint32(m.buf[2])<<8 | uint32(m.buf[3])
	return
}

// Slong reads a single 32-bit signed long from the underlying Memory interface.
func (m *MemoryDecoder) Slong(addr int) (n int32, err error) {
	_, err = m.M.Read(addr, m.buf[:4])
	u := uint32(m.buf[0])<<24 | uint32(m.buf[1])<<16 | uint32(m.buf[2])<<8 | uint32(m.buf[3])
	n = *(*int32)(unsafe.Pointer(&u))
	return
}

// clear wipes the given Memory device, resetting all reachable addresses to
// zero.
func clear(m Memory) {
	if m == nil {
		return
	}
	addr := 0
	zero := [32]byte{}
	for {
		n, err := m.Write(addr, zero[:])
		if err == io.ErrShortWrite || err == ErrAddressOutOfBounds {
			break
		}
		if err != nil {
			panic(err)
		}
		addr += n
	}
}
