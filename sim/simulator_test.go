package sim

import (
	"bytes"
	"io"
	"testing"

	"github.com/cavaliercoder/go-m68k/m68ktest"
)

func TestSim(t *testing.T) {
	p := m68ktest.LoadBytes(t, []byte{0x4E, 0x4F})
	_, err := New(p, p.TraceWriter, nil)
	if err != nil {
		t.Fatal(err)
	}
	m68ktest.AssertRun(t, p)
}

func TestStdout(t *testing.T) {
	b := &bytes.Buffer{}
	p := m68ktest.LoadBytes(t, []byte{
		0x41, 0xFA, 0x00, 0x12,
		0x12, 0x18,
		0x67, 0x08,
		0x10, 0x3C, 0x00, 0x06,
		0x4E, 0x4F,
		0x60, 0xF4,
		0x4E, 0x72, 0x27, 0x00,
		0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x20, 0x77, 0x6F, 0x72, 0x6C, 0x64, 0x21, 0x0D, 0x0A, 0x00,
	})
	s, err := New(p, b, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.Run(); err != io.EOF {
		t.Fatal(err)
	}
	expect := "Hello world!\r\n"
	actual := b.String()
	if actual != expect {
		t.Errorf("expected '%q', got '%q'", expect, actual)
	}
}
