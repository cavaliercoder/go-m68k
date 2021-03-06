package m68kmem

type rom struct {
	b []byte
}

// NewROM returns Read-Only Memory initialized with the given byte data. This
// memory can be used to load programs and data for a 68000 processor. Memory
// access is protected by an AddressBus.
func NewROM(b []byte) Memory {
	// TODO: check that size is addressable
	return AttachBus(&rom{b})
}

func (m *rom) Read(addr int, p []byte) (n int, err error) {
	if addr < 0 || addr >= len(m.b) {
		return 0, AccessViolationError(uint32(addr))
	}
	n = copy(p, m.b[addr:])
	return
}

// Write always returns an AccessViolationError.
func (m *rom) Write(addr int, p []byte) (n int, err error) {
	err = accessViolationError(addr)
	return
}

// Reset does nothing, but is required to implement the Memory interface.
func (m *rom) Reset() (err error) {
	// no-op
	return
}
