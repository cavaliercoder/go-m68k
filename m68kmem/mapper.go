package m68kmem

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

func NewMapper() *Mapper {
	return &Mapper{}
}

func (mm *Mapper) Map(m Memory, start, end int) error {
	mm.mappings = &mapping{
		m:     m,
		start: start,
		end:   end,
		next:  mm.mappings,
	}
	return nil
}

func (mm *Mapper) Read(addr int, p []byte) (n int, err error) {
	for m := mm.mappings; m != nil; m = m.next {
		if m.start <= addr && m.end >= addr {
			addr -= m.start
			return m.m.Read(addr, p)
		}
	}
	return 0, accessViolationError(addr)
}

func (mm *Mapper) Write(addr int, p []byte) (n int, err error) {
	for m := mm.mappings; m != nil; m = m.next {
		if m.start <= addr && m.end >= addr {
			addr -= m.start
			return m.m.Write(addr, p)
		}
	}
	return 0, accessViolationError(addr)
}
