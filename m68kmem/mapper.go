package m68kmem

import "errors"

type Mapper struct {
	mappings *mapping
}

type mapping struct {
	m          Memory
	start, end int
	next       *mapping
}

func NewMapper() *Mapper {
	return &Mapper{}
}

func (mm *Mapper) Map(m Memory, start, end int) error {
	for p := mm.mappings; p != nil; p = p.next {
		if p.end < start && (p.next == nil || p.next.start > end) {
			p.next = &mapping{
				m:     m,
				start: start,
				end:   end,
				next:  p.next,
			}
			return nil
		}
	}
	return errors.New("mapping collision")
}

func (mm *Mapper) Read(addr int, p []byte) (n int, err error) {
	for m := mm.mappings; m != nil; m = m.next {
		if m.start <= addr && m.end >= addr {
			addr -= m.start
			return m.m.Read(addr, p)
		}
	}
	err = ErrAccessViolation
	return
}

func (mm *Mapper) Write(addr int, p []byte) (n int, err error) {
	for m := mm.mappings; m != nil; m = m.next {
		if m.start <= addr && m.end >= addr {
			addr -= m.start
			return m.m.Write(addr, p)
		}
	}
	err = ErrAccessViolation
	return
}
