package m68kmem

import (
	"bytes"
	"testing"
)

func TestMemory(t *testing.T) {
	m := NewRAM(16)
	data := []byte("0123456789ABCDEF")

	// TODO: test init state is 0x00000000...

	t.Run("Write", func(t *testing.T) {
		n, err := m.Write(0x00, data)
		if err != nil {
			t.Errorf("error writing to memory: %v", err)
		}
		if n != 16 {
			t.Errorf("expected to write 16 bytes to memory, got %d", n)
		}
	})

	t.Run("Read", func(t *testing.T) {
		b := make([]byte, 16)
		n, err := m.Read(0, b)
		if err != nil {
			t.Errorf("error reading from memory: %v", err)
		}
		if n != 16 {
			t.Errorf("expected to read 16 bytes from memory, got %d", n)
		}
		if !bytes.Equal(b, data) {
			t.Errorf("expected read to return '%s', got '%s'", data, b)
		}
	})

	t.Run("WriteWithOffset", func(t *testing.T) {
		n, err := m.Write(8, data[:8])
		if err != nil {
			t.Errorf("error writing to memory: %v", err)
		}
		if n != 8 {
			t.Errorf("expected to write 8 bytes to memory, got %d", n)
		}

		expect := []byte("0123456701234567")
		b := make([]byte, 16)
		m.Read(0, b)
		if !bytes.Equal(b, expect) {
			t.Errorf("expected memory state to be '%s' after write, got '%s'", expect, b)
		}
	})

	t.Run("ReadWithOffset", func(t *testing.T) {
		expect := []byte("45670123")
		b := make([]byte, 8)
		n, err := m.Read(4, b)
		if err != nil {
			t.Errorf("error reading from memory: %v", err)
		}
		if n != 8 {
			t.Errorf("expected to read 8 bytes from memory, got %d", n)
		}
		if !bytes.Equal(b, expect) {
			t.Errorf("expected read to return '%s', got '%s'", expect, b)
		}
	})
}
