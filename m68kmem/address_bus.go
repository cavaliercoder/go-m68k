package m68kmem

type addressBus struct {
	M Memory
}

// AttachBus protects all accesses to the given Memory device with a 68000
// A23-A0 Address Bus emulator. The bus ensures all 32-bit addresses are reduced
// to a 24-bit address by masking the 8 most significant bits.
func AttachBus(m Memory) Memory {
	return &addressBus{M: m}
}

func (m *addressBus) Read(addr int, p []byte) (n int, err error) {
	addr, err = maskAddr(addr)
	if err != nil {
		return
	}
	return m.M.Read(addr, p)
}

func (m *addressBus) Write(addr int, p []byte) (n int, err error) {
	addr, err = maskAddr(addr)
	if err != nil {
		return
	}
	return m.M.Write(addr, p)
}

func maskAddr(addr int) (n int, err error) {
	if addr < 0 {
		err = AccessViolationError(uint32(addr))
		return
	}
	n = int(uint32(addr) & MaxAddress)
	return
}
