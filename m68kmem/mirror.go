package m68kmem

// Mirror expands the address range of a Memory device by mirroring it.
type Mirror struct {
	M           Memory
	Size, Range uint32
}

// NewMirror returns a new Mirror for the given Memory interface. If n is the
// size of the mirrored device, then r is the size of the range across which it
// is mirrored.
func NewMirror(m Memory, n, r uint32) *Mirror {
	return &Mirror{M: m, Size: n, Range: r}
}

// Write writes to the underlying Memory interface.
func (m *Mirror) Write(addr int, p []byte) (n int, err error) {
	if addr < 0 || uint32(addr) >= m.Range {
		return 0, AccessViolationError(uint32(addr))
	}
	addr = int(uint32(addr) % m.Size)
	return m.M.Write(addr, p)
}

// Read reads from the underlying Memory interface.
func (m *Mirror) Read(addr int, p []byte) (n int, err error) {
	if addr < 0 || uint32(addr) >= m.Range {
		return 0, AccessViolationError(uint32(addr))
	}
	addr = int(uint32(addr) % m.Size)
	return m.M.Read(addr, p)
}
