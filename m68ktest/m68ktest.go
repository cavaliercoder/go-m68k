/*
Package m68ktest provides functions for testing Go code against the Motorola
68000 chipset emulator.
*/
package m68ktest

import (
	"io"
	"os"
	"testing"

	"github.com/cavaliercoder/go-m68k"
	"github.com/cavaliercoder/go-m68k/dump"
	"github.com/cavaliercoder/go-m68k/srec"
)

type testLogWriter struct {
	t *testing.T
}

func (t *testLogWriter) Write(p []byte) (n int, err error) {
	if t.t == nil || p == nil {
		return
	}
	t.t.Logf("%s", p)
	n = len(p)
	return
}

// NewProcessor returns a 68000 processor that will log trace messages to the
// given test context.
func NewProcessor(t *testing.T) *m68k.Processor {
	p := &m68k.Processor{
		TraceWriter: &testLogWriter{t},
	}
	p.Reset()
	return p
}

// LoadFile returns a new 680000 processor with the given S-Record format
// program loaded into memory.
func LoadFile(t *testing.T, path string) *m68k.Processor {
	p := NewProcessor(t)
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	recs, err := srec.Read(f)
	if err != nil {
		panic(err)
	}
	for _, rec := range recs {
		if rec.IsData() {
			if _, err := p.M.Write(rec.Address(), rec.Data()); err != nil {
				panic(err)
			}
		}
	}
	return p
}

// LoadBytes returns a new 68000 processor with the given raw program data
// loaded into memory.
func LoadBytes(t *testing.T, b []byte) *m68k.Processor {
	p := NewProcessor(t)
	if _, err := p.M.Write(0x1000, b); err != nil {
		panic(err)
	}
	return p
}

// AssertRun assert that the processor can run the program loaded into memory.
func AssertRun(t *testing.T, p *m68k.Processor) {
	p.PC = 0x1000
	err := p.Run()
	if err == io.EOF {
		return
	}
	if err != nil {
		t.Errorf("error running program: %v", err)
		dump.Memory(p.TraceWriter, p.M)
	}
}

func AssertDataRegister(t *testing.T, p *m68k.Processor, r int, v uint32) {
	if p.D[r] != v {
		t.Errorf("expected data value 0x%08X in D%d, got 0x%08X", v, r, p.D[r])
	}
}

func AssertAddressRegister(t *testing.T, p *m68k.Processor, r int, v uint32) {
	if p.A[r] != v {
		t.Errorf("expected address value 0x%08X in A%d, got 0x%08X", v, r, p.A[r])
	}
}

func AssertStatusRegister(t *testing.T, p *m68k.Processor, v uint32) {
	if p.SR != v {
		t.Errorf("expected value 0x%X in SR, got 0x%X", v, p.SR)
	}
}

func AssertLong(t *testing.T, p *m68k.Processor, addr uint32, v uint32) {
	b := make([]byte, 4)
	_, err := p.M.Read(int(addr), b)
	if err != nil {
		t.Fatal(err)
	}
	actual := uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
	if actual != v {
		t.Errorf("expected value 0x%08X at 0x%08X, got 0x%08X", v, addr, actual)
		dump.Processor(os.Stdout, p)
		dump.Memory(os.Stdout, p.M)
	}
}
