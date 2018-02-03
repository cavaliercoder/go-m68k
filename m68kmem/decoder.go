package m68kmem

import "unsafe"

// A Decoder provides native type decoding for a Memory interface.
type Decoder struct {
	M Memory

	buf [4]byte
}

func NewDecoder(m Memory) *Decoder {
	return &Decoder{M: m}
}

// Write writes to the underlying Memory interface.
func (m *Decoder) Write(addr int, p []byte) (n int, err error) {
	return m.M.Write(addr, p)
}

// Read reads from the underlying Memory interface.
func (m *Decoder) Read(addr int, p []byte) (n int, err error) {
	return m.M.Read(addr, p)
}

// Byte reads a single byte from the underlying Memory interface.
func (m *Decoder) Byte(addr int) (b byte, err error) {
	_, err = m.M.Read(addr, m.buf[:1])
	b = m.buf[0]
	return
}

// Word reads a single 16-bit unsigned word from the underlying Memory
// interface.
func (m *Decoder) Word(addr int) (n uint16, err error) {
	_, err = m.M.Read(addr, m.buf[:2])
	n = uint16(m.buf[0])<<8 | uint16(m.buf[1])
	return
}

// Sword reads a single 16-bit signed word from the underlying Memory interface.
func (m *Decoder) Sword(addr int) (n int16, err error) {
	_, err = m.M.Read(addr, m.buf[:2])
	u := uint16(m.buf[0])<<8 | uint16(m.buf[1])
	n = *(*int16)(unsafe.Pointer(&u))
	return
}

// Long reads a single 32-bit unsigned long from the underlying Memory
// interface.
func (m *Decoder) Long(addr int) (n uint32, err error) {
	_, err = m.M.Read(addr, m.buf[:4])
	n = uint32(m.buf[0])<<24 | uint32(m.buf[1])<<16 | uint32(m.buf[2])<<8 | uint32(m.buf[3])
	return
}

// Slong reads a single 32-bit signed long from the underlying Memory interface.
func (m *Decoder) Slong(addr int) (n int32, err error) {
	_, err = m.M.Read(addr, m.buf[:4])
	u := uint32(m.buf[0])<<24 | uint32(m.buf[1])<<16 | uint32(m.buf[2])<<8 | uint32(m.buf[3])
	n = *(*int32)(unsafe.Pointer(&u))
	return
}
