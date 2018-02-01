package m68k

import (
	"errors"
	"fmt"
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

func bytesZero(b []byte) bool {
	for i := 0; i < len(b); i++ {
		if b[i] != 0 {
			return false
		}
	}
	return true
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

// Dump prints the content of a Memory device to the given writer in hexidecimal
// format.
func Dump(w io.Writer, m Memory) {
	b := make([]byte, 16)
	elip := false
	for i := 0; ; i += 16 {
		_, err := m.Read(i, b)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Fprintf(w, "%v\n", err)
			return
		}
		if bytesZero(b) {
			if !elip {
				fmt.Fprint(w, "*\n")
				elip = true
			}
			continue
		}
		elip = false
		fmt.Fprintf(w, "%08X ", i)
		for j := 0; j < 8; j++ {
			fmt.Fprintf(w, " %02X", b[j])
		}
		w.Write([]byte{' '})
		for j := 8; j < 16; j++ {
			fmt.Fprintf(w, " %02X", b[j])
		}
		w.Write([]byte("  |"))
		for j := 0; j < 16; j++ {
			if b[j] > 31 && b[j] < 126 {
				w.Write([]byte{b[j]})
			} else {
				w.Write([]byte{'.'})
			}
		}
		w.Write([]byte("|\n"))
	}
}
