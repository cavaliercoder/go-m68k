package m68k

import (
	"errors"
	"io"
)

var (
	ErrAddressOutOfBounds = errors.New("memory address out of bounds")
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
