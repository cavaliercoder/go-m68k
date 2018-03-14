package m68kmem

import (
	"bytes"
	"testing"
)

func TestMirror(t *testing.T) {
	m := NewMirror(NewRAM(0x10), 0x10, 0x40)
	data := []byte("0123456789ABCDEF")

	t.Run("WriteBefore", func(t *testing.T) {
		_, err := m.Write(-16, data)
		if !isAccessViolationError(err) {
			t.Errorf("expected access violation, got: %v", err)
		}
	})

	t.Run("WriteStart", func(t *testing.T) {
		n, err := m.Write(0x00, data[:8])
		if err != nil {
			t.Errorf("error writing to memory: %v", err)
		}
		if n != 8 {
			t.Errorf("expected to write 16 bytes to memory, got %d", n)
		}
	})

	t.Run("WriteEnd", func(t *testing.T) {
		n, err := m.Write(0x38, data[8:])
		if err != nil {
			t.Errorf("error writing to memory: %v", err)
		}
		if n != 8 {
			t.Errorf("expected to write 16 bytes to memory, got %d", n)
		}
	})

	t.Run("WriteAfter", func(t *testing.T) {
		_, err := m.Write(0x40, data)
		if !isAccessViolationError(err) {
			t.Errorf("expected access violation, got: %v", err)
		}
	})

	t.Run("ReadBefore", func(t *testing.T) {
		b := make([]byte, 16)
		_, err := m.Read(-16, b)
		if !isAccessViolationError(err) {
			t.Errorf("expected access violation, got: %v", err)
		}
	})

	t.Run("ReadStart", func(t *testing.T) {
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

	t.Run("ReadEnd", func(t *testing.T) {
		b := make([]byte, 16)
		n, err := m.Read(0x30, b)
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

	t.Run("ReadAfter", func(t *testing.T) {
		b := make([]byte, 16)
		_, err := m.Read(0x40, b)
		if !isAccessViolationError(err) {
			t.Errorf("expected access violation, got: %v", err)
		}
	})
}
