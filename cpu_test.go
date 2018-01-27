package m68k_test

import (
	"io"
	"testing"

	"github.com/cavaliercoder/go-m68k/m68ktest"
)

func TestReset(t *testing.T) {
	// init dirty processor
	p := m68ktest.NewProcessor(t)
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
	p := m68ktest.LoadFile(t, "testdata/test-op-addi.h68")
	m68ktest.AssertRun(t, p)
	m68ktest.AssertDataRegister(t, p, 0, 0xFE)
	m68ktest.AssertDataRegister(t, p, 1, 0xFFFE)
	m68ktest.AssertDataRegister(t, p, 2, 0xFFFFFFFE)
	m68ktest.AssertDataRegister(t, p, 3, 0xFFFFFFFF)
}

func TestMoveL(t *testing.T) {
	p := m68ktest.LoadFile(t, "testdata/test-op-move-l.h68")
	m68ktest.AssertRun(t, p)

	// source: data register
	m68ktest.AssertDataRegister(t, p, 1, 0x44304430)
	m68ktest.AssertAddressRegister(t, p, 1, 0x44304430)
	m68ktest.AssertLong(t, p, 0x2000, 0x44304430)
	m68ktest.AssertLong(t, p, 0x2004, 0x44314431)
	m68ktest.AssertLong(t, p, 0x2008, 0x44324432)
	// m68ktest.AssertLong(t, p, 0x200C, 0x44334433)
	m68ktest.AssertLong(t, p, 0x2010, 0x44344434)
	m68ktest.AssertLong(t, p, 0x12000, 0x44344434)

	// source: address register
	m68ktest.AssertDataRegister(t, p, 2, 0x41304130)
	m68ktest.AssertAddressRegister(t, p, 2, 0x41304130)
	m68ktest.AssertLong(t, p, 0x3000, 0x41304130)
	m68ktest.AssertLong(t, p, 0x3004, 0x41314131)
	m68ktest.AssertLong(t, p, 0x3008, 0x41324132)
	m68ktest.AssertLong(t, p, 0x300C, 0x41334133)
	m68ktest.AssertLong(t, p, 0x13000, 0x41334133)

	// source: immediate
	m68ktest.AssertDataRegister(t, p, 3, 0x49304930)
	m68ktest.AssertAddressRegister(t, p, 3, 0x49304930)
	m68ktest.AssertLong(t, p, 0xF000, 0x49304930)
	m68ktest.AssertLong(t, p, 0xF004, 0x49314931)
	m68ktest.AssertLong(t, p, 0xF008, 0x49324932)
	m68ktest.AssertLong(t, p, 0xF00C, 0x49334933)
	m68ktest.AssertLong(t, p, 0x1F000, 0x49334933)
}

func TestOri(t *testing.T) {
	p := m68ktest.LoadFile(t, "testdata/test-op-ori.h68")
	p.D[2] = 0x22
	p.D[3] = 0x2222
	p.D[4] = 0x22222222
	m68ktest.AssertRun(t, p)
	m68ktest.AssertStatusRegister(t, p, 0x1F)
	m68ktest.AssertDataRegister(t, p, 1, 0xFFFFFFFF)
	m68ktest.AssertDataRegister(t, p, 2, 0x33)
	m68ktest.AssertDataRegister(t, p, 3, 0x3333)
	m68ktest.AssertDataRegister(t, p, 4, 0x33333333)
}

func TestAndi(t *testing.T) {
	b := []byte{
		0x02, 0x01, 0x00, 0x0F, // andi.b #$0F,D1
		0x02, 0x41, 0x0F, 0xFF, // andi.w #$0FFF,D1
		0x02, 0x81, 0x0F, 0x0F, 0xFF, 0xFF, // andi.l #$0F0FFFFF,D1
	}
	// p := m68ktest.LoadFile(t, "testdata/test-op-andi.h68")
	p := m68ktest.LoadBytes(t, b)
	p.D[1] = 0x55555555
	m68ktest.AssertRun(t, p)
	// m68ktest.AssertStatusRegister(t, p, 0x1F)
	m68ktest.AssertDataRegister(t, p, 1, 0x05050505)
}

var progAndiOri = m68ktest.LoadBytes(nil, []byte{
	0x00, 0x80, 0xFF, 0xFF, 0xFF, 0xFF, // ori.l	#$FFFFFFFF,D0
	0x02, 0x80, 0x55, 0x55, 0x55, 0x55, // andi.l	#$55555555,D0
	0x00, 0x80, 0xAA, 0xAA, 0xAA, 0xAA, // ori.l	#$AAAAAAAA,D0
	0x02, 0x80, 0x0F, 0x0F, 0x0F, 0x0F, // andi.l	#$0F0F0F0F,D0
	0x00, 0x80, 0xF0, 0xF0, 0xF0, 0xF0, // ori.l	#$F0F0F0F0,D0
	0x02, 0x80, 0x0F, 0x0F, 0x0F, 0x0F, // andi.l	#$0F0F0F0F,D0
	0x00, 0x80, 0x00, 0xFF, 0xFF, 0x00, // ori.l	#$00FFFF00,D0
	0x02, 0x80, 0xF0, 0xF0, 0xF0, 0xF0, // andi.l	#$F0F0F0F0,D0
})

func BenchmarkAndiOri(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if err := progAndiOri.Run(); err != io.EOF {
			panic(err)
		}
	}
}
