package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrInvalidOperandSize              = errors.New("invalid operand size")
	ErrInvalidEffectiveAddressMode     = errors.New("invalid effective address mode")
	ErrInvalidEffectiveAddressRegister = errors.New("invalid effective address register")
)

func main() {
	w := os.Stdout
	fmt.Fprintf(w, "package m68k\n\n")
	funcs := []uint16{}
	for op := uint16(0x0000); op < 0xFFFF; op++ {
		fn, err := genMove(op)
		if err != nil {
			// fmt.Fprintf(os.Stderr, "Opcode %04X: %v\n", op, err)
		} else {
			w.WriteString(fn)
			funcs = append(funcs, op)
		}
	}

	fmt.Fprintf(w, "func (c *Processor) mapFn(op uint16) func() {\n")
	fmt.Fprintf(w, "	table := map[uint16]func(){\n")
	for i := 0; i < len(funcs); i++ {
		fmt.Fprintf(w, "		0x%04X: c.op%04X,\n", funcs[i], funcs[i])
	}
	fmt.Fprintf(w, "	}\n")
	fmt.Fprintf(w, "	return table[op]\n")
	fmt.Fprintf(w, "}\n")
}

func tab(w io.Writer, n int) io.Writer {
	return NewIndentWriter(w, "\t", n)
}

func printOpFuncHdr(w io.Writer, op uint16) {
	fmt.Fprintf(w, "func (c *Processor) op%04X() {\n", op)
	fmt.Fprintln(w, "	pc := c.PC")
	fmt.Fprintln(w, "	c.PC += 2")
}

func printOpFuncFtr(w io.Writer, op, sz, src, dst string, args []string) {
	b := &bytes.Buffer{}
	for i := 0; i < len(args); i++ {
		fmt.Fprintf(b, ", %s", args[i])
	}
	fmt.Fprintf(w, "	c.tracef(\"%%04X %s.%s %s,%s\\n\", pc%s)\n", op, sz, src, dst, b)
	fmt.Fprint(w, "}\n\n")
}

func printReadMem(w io.Writer, addr string, sz uint16) {
	buflen := []int{1, 2, 4}[sz]
	fmt.Fprintf(w, "_, c.err = c.M.Read(int(%s), c.buf[:%d])\n", addr, buflen)
	fmt.Fprintln(w, "if c.err != nil {")
	fmt.Fprintln(w, "	return")
	fmt.Fprintln(w, "}")
	switch sz {
	case 0:
		fmt.Fprintln(w, "v := c.buf[0]")
	case 1:
		fmt.Fprintln(w, "v := uint16(c.buf[0])<<8 | uint16(c.buf[1])")
	case 0x02:
		fmt.Fprintln(w, "v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24")
	}
}

func genMove(op uint16) (string, error) {
	if op&0xC000 != 0 {
		return "", errors.New("invalid move operation")
	}

	w := &bytes.Buffer{}
	ops, szs, ss, ds := "move", "?", "?", "?"
	sz := (op & 0x3000) >> 12 // operand size
	sm := (op & 0x0038) >> 3  // source mode
	sr := op & 0x0007         // source register
	dm := (op & 0x01C0) >> 6  // dest mode
	dr := (op & 0x0E00) >> 9  // dest register

	if sz == 0x03 {
		return "", ErrInvalidOperandSize
	}
	szs = []string{"b", "w", "l"}[sz]
	// optype := []string{"byte", "uint16", "uint32"}[sz]
	args := []string{}

	printOpFuncHdr(w, op)

	// read source value
	switch sm {
	default:
		return "", errors.New("invalid source mode")

	case 0x00: // data register
		fmt.Fprintf(w, "	v := c.D[%d]\n", sr)
		ss = fmt.Sprintf("D%d", sr)

	case 0x01: // address register
		fmt.Fprintf(w, "	v := c.A[%d]\n", sr)
		ss = fmt.Sprintf("A%d", sr)

	case 0x02: // memory address
		printReadMem(tab(w, 1), fmt.Sprintf("c.A[%d]", sr), sz)
		ss = fmt.Sprintf("(A%d)", sr)

	case 0x07:
		switch sr {
		default:
			return "", errors.New("invalid source register")

		case 0x00, 0x01: // absolute short/long
			fmt.Fprintf(w, "	v := c.readImm(0x02)\n")
			ss = "$%X"
			args = append(args, "v")

		case 0x04: // immediate
			fmt.Fprintf(w, "	v := c.readImm(0x02)\n")
			ss = "#$%X"
			args = append(args, "v")
		}
	}

	// write to destination
	switch dm {
	default:
		return "", errors.New("invalid destination mode")

	case 0x00: // data register
		fmt.Fprintf(w, "	c.D[%d] = uint32(v)\n", dr)
		ds = fmt.Sprintf("D%d", dr)

	case 0x01: // address register
		// TODO: move.l to An is illegal but supported by assemblers
		fmt.Fprintf(w, "	c.A[%d] = uint32(v)\n", dr)
		ds = fmt.Sprintf("A%d", dr)
		ops = "movea"

	case 0x02: // memory address
		fmt.Fprintf(w, "	c.writeLong(c.A[%d], uint32(v))\n", dr)
		ds = fmt.Sprintf("(A%d)", dr)

	case 0x03: // memory address with post-increment
		fmt.Fprintf(w, "	c.writeLong(c.A[%d], uint32(v))\n", dr)
		fmt.Fprintf(w, "	c.A[%d] += 4\n", dr)
		ds = fmt.Sprintf("(A%d)+", dr)

	case 0x04: // memory address with pre-decrement
		fmt.Fprintf(w, "	c.A[%d] -= 4\n", dr)
		fmt.Fprintf(w, "	c.writeLong(c.A[%d], uint32(v))\n", dr)
		ds = fmt.Sprintf("-(A%d)", dr)

	case 0x05: // memory address with displacement
		fmt.Fprint(w, "	disp := c.readImm(0x01)\n")
		fmt.Fprintf(w, "	addr := disp + c.A[%d]\n", dr)
		fmt.Fprint(w, "	c.writeLong(addr, uint32(v))\n")
		ds = fmt.Sprintf("(%%d,A%d)", dr)
		args = append(args, "disp")

	case 0x07: // other
		switch dr {
		default:
			return "", errors.New("invalid destination register")

		case 0x00: // absolute word
			fmt.Fprint(w, "	addr := c.readImm(0x01)\n")
			fmt.Fprint(w, "	c.writeLong(addr, uint32(v))\n")
			ds = "$%X"
			args = append(args, "addr")

		case 0x01: // absolute long
			fmt.Fprint(w, "	addr := c.readImm(0x02)\n")
			fmt.Fprint(w, "	c.writeLong(addr, uint32(v))\n")
			ds = "$%X"
			args = append(args, "addr")
		}
	}
	printOpFuncFtr(w, ops, szs, ss, ds, args)
	return w.String(), nil
}

func genORI(op uint16) (string, error) {
	w := &bytes.Buffer{}

	ins := "ori"
	opl := "#$%X"
	opr := "?"
	szid := (op & 0xC0) >> 6
	mod := (op & 0x38) >> 3
	reg := (op & 0x07)

	fmt.Fprintf(w, "func (c *Processor) op%04X() {\n", op)

	// read immediate value
	switch szid {
	case 0x00: // byte
		fmt.Fprintf(w, "	v := c.readImmByte()\n")
		switch mod {
		case 0x00: // data register
			opr = fmt.Sprintf("D%d", reg)
			fmt.Fprintf(w, "	c.D[%d] = uint32(byte(c.D[%d]) ^ v)\n", reg, reg)

		case 0x07: // other
			switch reg {
			case 0x04: // CCR
				opr = "CCR"
				fmt.Fprintf(w, "	c.CCR = (c.CCR ^ uint32(v)) & 0xFF\n")

			default:
				return "", ErrInvalidEffectiveAddressRegister
			}

		default:
			return "", ErrInvalidEffectiveAddressMode
		}

	case 0x01: // word
		fmt.Fprintf(w, "	v := c.readImmWord()\n")
		switch mod {
		case 0x00: // data register
			opr = fmt.Sprintf("D%d", reg)
			fmt.Fprintf(w, "	c.D[%d] = uint32(uint16(c.D[%d]) ^ v)\n", reg, reg)

		case 0x07: // other
			switch reg {
			case 0x04: // SP
				opr = "SP"
				fmt.Fprintf(w, "	c.A[7] = (c.A[7] & 0xFF) ^ uint32(v)\n")

			default:
				return "", ErrInvalidEffectiveAddressRegister
			}

		default:
			return "", ErrInvalidEffectiveAddressMode
		}

	case 0x02: // long
		fmt.Fprintf(w, "	v := c.readImmLong()\n")
		switch mod {
		case 0x00: // data register
			opr = fmt.Sprintf("D%d", reg)
			fmt.Fprintf(w, "	c.D[%d] = c.D[%d] ^ v\n", reg, reg)

		default:
			return "", ErrInvalidEffectiveAddressMode
		}

	default:
		return "", ErrInvalidOperandSize
	}

	szs := []byte{'b', 'w', 'l', '?'}[szid]
	fmt.Fprintf(w, "	c.tracef(\"%%04X %s.%c %s,%s\\n\", addr, v)\n", ins, szs, opl, opr)
	fmt.Fprintf(w, "}\n\n")
	return w.String(), nil
}
