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
	p.SR = 0xFFFFFFFF
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
	if p.SR != 0 {
		t.Error("Condition code register was not reset")
	}
	if p.PC != 0 {
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

var progHelloWorld = m68ktest.LoadBytes(nil, []byte{
	0x41, 0xFA, 0x00, 0x12, 0x12, 0x18, 0x67, 0x08, 0x10, 0x3C, 0x00, 0x06, 0x4E,
	0x4F, 0x60, 0xF4, 0x4E, 0x72, 0x27, 0x00, 0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20,
	0x77, 0x6F, 0x72, 0x6C, 0x64, 0x21, 0x0D, 0x0A, 0x00,
})

func BenchmarkHelloWorld(b *testing.B) {
	for n := 0; n < b.N; n++ {
		if err := progHelloWorld.Run(); err != io.EOF {
			panic(err)
		}
	}
}
