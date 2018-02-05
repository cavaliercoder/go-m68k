package m68kmem

import "os"

func ExampleNewTracer() {
	m := NewTracer("ram", os.Stdout, NewRAM(64))
	m.Write(0x10, []byte("Hello world!\n"))
	m.Read(0x10, make([]byte, 13))

	// Output:
	// 00000010 [ram] wrote 13 bytes: 48656c6c6f20776f726c64210a
	// 00000010 [ram] read 13 bytes: 48656c6c6f20776f726c64210a
}
