package m68k_test

import (
	"io"
	"os"
	"testing"

	"github.com/cavaliercoder/go-m68k"
	"github.com/cavaliercoder/go-m68k/srec"
)

func testProc() *m68k.Processor {
	p := &m68k.Processor{
		TraceWriter: os.Stderr,
	}
	p.Reset()
	return p
}

func load(path string) *m68k.Processor {
	p := testProc()
	f, err := os.Open("./testdata/" + path)
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

func assertRun(t *testing.T, p *m68k.Processor) {
	err := p.Jump(0x1000)
	if err != nil {
		t.Errorf("error resetting program counter: %v", err)
	}
	err = p.Run()
	if err == io.EOF {
		return
	}
	if err != nil {
		t.Errorf("error running program: %v", err)
	}
}

func assertDataRegister(t *testing.T, p *m68k.Processor, r int, v uint32) {
	if p.D[r] != v {
		t.Errorf("expected value 0x%08X in D%d, got 0x%08X", v, r, p.D[r])
	}
}

func TestReset(t *testing.T) {
	// init dirty processor
	p := testProc()
	for i := 0; i < 8; i++ {
		p.A[i], p.D[i] = 0xFFFFFFFF, 0xFFFFFFFF
	}
	p.CCR = 0xFFFFFFFF
	p.PC = 0xFFFFFFFF
	p.M.Write(0x1000, []byte{0xFF, 0xFF, 0xFF, 0xFF})

	// reset and test
	p.Reset()
	for i := 0; i < 8; i++ {
		if p.A[i] != 0 {
			t.Errorf("Address register %d was not reset", i)
		}
		if p.D[i] != 0 {
			t.Errorf("Data register %d was not reset", i)
		}
	}
	if p.CCR != 0 {
		t.Error("Condition code register was not reset")
	}
	if p.PC != 0x1000 {
		t.Error("Program counter not reset")
	}
	b := make([]byte, 4)
	_, err := p.M.Read(0x1000, b)
	if err != nil {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		if b[i] != 0 {
			t.Error("Memory was not zeroed")
			break
		}
	}
}

func TestAddi(t *testing.T) {
	p := load("test-op-addi.h68")
	assertRun(t, p)
	assertDataRegister(t, p, 0, 0xFE)
	assertDataRegister(t, p, 1, 0xFFFE)
	assertDataRegister(t, p, 2, 0xFFFFFFFE)
	assertDataRegister(t, p, 3, 0xFFFFFFFF)
}

func TestMovel(t *testing.T) {
	p := load("test-op-move-l.h68")
	assertRun(t, p)
	assertDataRegister(t, p, 7, 0xD7)
}
