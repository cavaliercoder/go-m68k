package m68k_test

import (
	"io"
	"os"
	"testing"

	"github.com/cavaliercoder/go-m68k"
	"github.com/cavaliercoder/go-m68k/srec"
)

type testLogWriter struct {
	t *testing.T
}

func (t *testLogWriter) Write(p []byte) (n int, err error) {
	t.t.Logf("%s", p)
	n = len(p)
	return
}

func testProc(t *testing.T) *m68k.Processor {
	p := &m68k.Processor{
		TraceWriter: &testLogWriter{t},
	}
	p.Reset()
	return p
}

func loadFile(t *testing.T, path string) *m68k.Processor {
	p := testProc(t)
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

func loadBytes(t *testing.T, b []byte) *m68k.Processor {
	p := testProc(t)
	if _, err := p.M.Write(0x1000, b); err != nil {
		panic(err)
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
		m68k.Dump(os.Stdout, p.M)
	}
}

func assertDataRegister(t *testing.T, p *m68k.Processor, r int, v uint32) {
	if p.D[r] != v {
		t.Errorf("expected data value 0x%08X in D%d, got 0x%08X", v, r, p.D[r])
	}
}

func assertAddressRegister(t *testing.T, p *m68k.Processor, r int, v uint32) {
	if p.A[r] != v {
		t.Errorf("expected address value 0x%08X in A%d, got 0x%08X", v, r, p.A[r])
	}
}

func assertCCRRegister(t *testing.T, p *m68k.Processor, v uint32) {
	if p.CCR != v {
		t.Errorf("expected value 0x%X in CCR, got 0x%X", v, p.CCR)
	}
}

func assertLong(t *testing.T, p *m68k.Processor, addr uint32, v uint32) {
	b := make([]byte, 4)
	_, err := p.M.Read(int(addr), b)
	if err != nil {
		t.Fatal(err)
	}
	actual := uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
	if actual != v {
		t.Errorf("expected value 0x%08X at 0x%08X, got 0x%08X", v, addr, actual)
		m68k.DumpState(os.Stdout, p)
		m68k.Dump(os.Stdout, p.M)
	}
}

func TestReset(t *testing.T) {
	// init dirty processor
	p := testProc(t)
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

// func TestAddi(t *testing.T) {
// 	p := load("test-op-addi.h68")
// 	assertRun(t, p)
// 	assertDataRegister(t, p, 0, 0xFE)
// 	assertDataRegister(t, p, 1, 0xFFFE)
// 	assertDataRegister(t, p, 2, 0xFFFFFFFE)
// 	assertDataRegister(t, p, 3, 0xFFFFFFFF)
// }

func TestMoveL(t *testing.T) {
	p := loadFile(t, "test-op-move-l.h68")
	assertRun(t, p)

	// source: data register
	assertDataRegister(t, p, 1, 0x44304430)
	assertAddressRegister(t, p, 1, 0x44304430)
	assertLong(t, p, 0x2000, 0x44304430)
	assertLong(t, p, 0x2004, 0x44314431)
	assertLong(t, p, 0x2008, 0x44324432)
	// assertLong(t, p, 0x200C, 0x44334433)
	assertLong(t, p, 0x2010, 0x44344434)
	assertLong(t, p, 0x12000, 0x44344434)

	// source: address register
	assertDataRegister(t, p, 2, 0x41304130)
	assertAddressRegister(t, p, 2, 0x41304130)
	assertLong(t, p, 0x3000, 0x41304130)
	assertLong(t, p, 0x3004, 0x41314131)
	assertLong(t, p, 0x3008, 0x41324132)
	assertLong(t, p, 0x300C, 0x41334133)
	assertLong(t, p, 0x13000, 0x41334133)

	// source: immediate
	assertDataRegister(t, p, 3, 0x49304930)
	assertAddressRegister(t, p, 3, 0x49304930)
	assertLong(t, p, 0xF000, 0x49304930)
	assertLong(t, p, 0xF004, 0x49314931)
	assertLong(t, p, 0xF008, 0x49324932)
	assertLong(t, p, 0xF00C, 0x49334933)
	assertLong(t, p, 0x1F000, 0x49334933)
}

func TestOri(t *testing.T) {
	p := loadFile(t, "test-op-ori.h68")
	assertRun(t, p)
	assertCCRRegister(t, p, 0x1F)
	assertDataRegister(t, p, 1, 0xFFFFFFFF)
	assertDataRegister(t, p, 2, 0x33)
	assertDataRegister(t, p, 3, 0x3333)
	assertDataRegister(t, p, 4, 0x33333333)
}
