package m68kmem

import (
	"errors"
)

// A Mapper allows a Motorola 68000 processor to communicate with external
// devices using a memory address, mapped into its addressable range.
type Mapper struct {
	mappings *mapping
}

type mapping struct {
	m          Memory
	start, end int
	next       *mapping
}

// NewMapper returns a new MemoryMapper.
func NewMapper() *Mapper {
	return &Mapper{}
}

// Map maps the given Memory addressable device into the given address range.
func (mm *Mapper) Map(m Memory, start, end int) (err error) {
	if uint32(start) < MinAddress || uint32(start) > MaxAddress {
		return errors.New("start address out of range")
	}
	if end < start || uint32(end) < MinAddress || uint32(end) > MaxAddress {
		return errors.New("end address out of range")
	}
	n := &mapping{
		m:     m,
		start: start,
		end:   end,
	}
	mm.mappings = insertMapping(mm.mappings, n)
	return
}

func insertMapping(root, m *mapping) *mapping {
	for n := root; n != nil; n = n.next {
		if n.next == nil || n.next.start >= m.start {
			m.next = n.next
			n.next = m
			if m.next != nil && m.start == m.next.start {
				// deref any mappings with duplicate start address
				m.next = m.next.next
			}
			return root
		}
	}
	return m
}

func (mm *Mapper) Read(addr int, p []byte) (n int, err error) {
	addr, err = maskAddr(addr)
	if err != nil {
		return
	}
	m := mm.mappings
	for i := addr; i < addr+len(p); {
		if m == nil {
			return 0, accessViolationError(addr)
		}
		for m != nil {
			if m.start <= i && m.end >= i {
				var nn int
				nn, err = m.m.Read(i-m.start, p[i-addr:])
				if err != nil {
					return
				}
				n += nn
				i += nn
				m = m.next
				break
			}
			m = m.next
		}
	}
	return
}

func (mm *Mapper) Write(addr int, p []byte) (n int, err error) {
	addr, err = maskAddr(addr)
	if err != nil {
		return
	}
	for m := mm.mappings; m != nil; m = m.next {
		if m.start <= addr && m.end >= addr {
			addr -= m.start
			n, err = m.m.Write(addr, p)
			return
		}
	}
	return 0, accessViolationError(addr)
}

// Reset calls Reset on all mapped Memory devices. Mappings are preserved.
func (mm *Mapper) Reset() (err error) {
	for m := mm.mappings; m != nil; m = m.next {
		err = m.m.Reset()
		if err != nil {
			return
		}
	}
	return
}
