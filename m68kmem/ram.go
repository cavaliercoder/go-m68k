package m68kmem

import "io"

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
		return 0, accessViolationError(addr)
	}
	n = copy(m.b[addr:], p)
	// TODO: raise exception vector instead of returning errors?
	if n < len(p) {
		return n, io.ErrShortWrite
	}
	return
}
