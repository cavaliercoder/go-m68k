/*
Package dump contains functions for printing the internal state of a Motorola
68000 processor emulator for debug purposes.
*/
package dump

import (
	"fmt"
	"io"

	m68k "github.com/cavaliercoder/go-m68k"
)

// Processor prints the internal state of of a 68000 processor to the given
// writer in hexidecimal format.
func Processor(w io.Writer, p *m68k.Processor) {
	fmt.Fprintf(w, "SR: %08X PC: %08X\n", p.SR, p.PC)
	for i := 0; i < len(p.D); i++ {
		fmt.Fprintf(w, "D%d: %08X A%d: %08X\n", i, p.D[i], i, p.A[i])
	}
}

func bytesZero(b []byte) bool {
	for i := 0; i < len(b); i++ {
		if b[i] != 0 {
			return false
		}
	}
	return true
}

// Memory prints the contents of a Memory device to the given writer in
// hexidecimal format.
func Memory(w io.Writer, m m68k.Memory) {
	b := make([]byte, 16)
	elip := false
	for i := 0; ; i += 16 {
		_, err := m.Read(i, b)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Fprintf(w, "%v\n", err)
			return
		}
		if bytesZero(b) {
			if !elip {
				fmt.Fprint(w, "*\n")
				elip = true
			}
			continue
		}
		elip = false
		fmt.Fprintf(w, "%08X ", i)
		for j := 0; j < 8; j++ {
			fmt.Fprintf(w, " %02X", b[j])
		}
		w.Write([]byte{' '})
		for j := 8; j < 16; j++ {
			fmt.Fprintf(w, " %02X", b[j])
		}
		w.Write([]byte("  |"))
		for j := 0; j < 16; j++ {
			if b[j] > 31 && b[j] < 126 {
				w.Write([]byte{b[j]})
			} else {
				w.Write([]byte{'.'})
			}
		}
		w.Write([]byte("|\n"))
	}
}
