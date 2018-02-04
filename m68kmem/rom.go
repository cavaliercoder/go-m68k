package m68kmem

import (
	"io"
)

type rom struct {
	b []byte
}

// NewROM returns Read-Only Memory initialized with the given byte data. This
// memory can be used to load programs and data for a 68000 processor.
func NewROM(b []byte) Memory {
	return &rom{b}
}

func (m *rom) Read(addr int, p []byte) (n int, err error) {
	if addr < 0 {
		return 0, AccessViolationError(uint32(addr))
	}
	if addr >= len(m.b) {
		return 0, io.EOF
	}
	n = copy(p, m.b[addr:])
	return
}

func (m *rom) Write(addr int, p []byte) (n int, err error) {
	return 0, accessViolationError(addr)
}
