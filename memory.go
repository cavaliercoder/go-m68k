package m68k

import (
	"fmt"
	"io"
)

type Memory interface {
	Read(addr int, p []byte) (n int, err error)
	Write(addr int, p []byte) (n int, err error)
}

type memory struct {
	b []byte
}

func NewMemory(size uint32) Memory {
	return &memory{
		b: make([]byte, size),
	}
}

func (m *memory) Read(addr int, p []byte) (n int, err error) {
	if addr >= len(m.b) {
		return 0, io.EOF
	}
	n = copy(p, m.b[addr:])
	return
}

func (m *memory) Write(addr int, p []byte) (n int, err error) {
	n = copy(m.b[addr:], p)
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

// Dump prints the content of a Memory bank to the given writer in hexidecimal
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
				fmt.Fprintf(w, "%08X  ...\n", i)
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