package m68k

func (c *Processor) op0000() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b D0,D0\n", pc)
}

func (c *Processor) op0001() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b D1,D0\n", pc)
}

func (c *Processor) op0002() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b D2,D0\n", pc)
}

func (c *Processor) op0003() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b D3,D0\n", pc)
}

func (c *Processor) op0004() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b D4,D0\n", pc)
}

func (c *Processor) op0005() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b D5,D0\n", pc)
}

func (c *Processor) op0006() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b D6,D0\n", pc)
}

func (c *Processor) op0007() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b D7,D0\n", pc)
}

func (c *Processor) op0008() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b A0,D0\n", pc)
}

func (c *Processor) op0009() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b A1,D0\n", pc)
}

func (c *Processor) op000A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b A2,D0\n", pc)
}

func (c *Processor) op000B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b A3,D0\n", pc)
}

func (c *Processor) op000C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b A4,D0\n", pc)
}

func (c *Processor) op000D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b A5,D0\n", pc)
}

func (c *Processor) op000E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b A6,D0\n", pc)
}

func (c *Processor) op000F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b A7,D0\n", pc)
}

func (c *Processor) op0010() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b (A0),D0\n", pc)
}

func (c *Processor) op0011() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b (A1),D0\n", pc)
}

func (c *Processor) op0012() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b (A2),D0\n", pc)
}

func (c *Processor) op0013() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b (A3),D0\n", pc)
}

func (c *Processor) op0014() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b (A4),D0\n", pc)
}

func (c *Processor) op0015() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b (A5),D0\n", pc)
}

func (c *Processor) op0016() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b (A6),D0\n", pc)
}

func (c *Processor) op0017() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.b (A7),D0\n", pc)
}

func (c *Processor) op0038() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[0] = uint32(v)
	c.tracef("%04X move.b $%X,D0\n", pc, v)
}

func (c *Processor) op0039() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[0] = uint32(v)
	c.tracef("%04X move.b $%X,D0\n", pc, v)
}

func (c *Processor) op003C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[0] = uint32(v)
	c.tracef("%04X move.b #$%X,D0\n", pc, v)
}

func (c *Processor) op0040() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b D0,A0\n", pc)
}

func (c *Processor) op0041() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b D1,A0\n", pc)
}

func (c *Processor) op0042() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b D2,A0\n", pc)
}

func (c *Processor) op0043() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b D3,A0\n", pc)
}

func (c *Processor) op0044() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b D4,A0\n", pc)
}

func (c *Processor) op0045() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b D5,A0\n", pc)
}

func (c *Processor) op0046() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b D6,A0\n", pc)
}

func (c *Processor) op0047() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b D7,A0\n", pc)
}

func (c *Processor) op0048() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b A0,A0\n", pc)
}

func (c *Processor) op0049() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b A1,A0\n", pc)
}

func (c *Processor) op004A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b A2,A0\n", pc)
}

func (c *Processor) op004B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b A3,A0\n", pc)
}

func (c *Processor) op004C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b A4,A0\n", pc)
}

func (c *Processor) op004D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b A5,A0\n", pc)
}

func (c *Processor) op004E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b A6,A0\n", pc)
}

func (c *Processor) op004F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b A7,A0\n", pc)
}

func (c *Processor) op0050() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b (A0),A0\n", pc)
}

func (c *Processor) op0051() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b (A1),A0\n", pc)
}

func (c *Processor) op0052() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b (A2),A0\n", pc)
}

func (c *Processor) op0053() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b (A3),A0\n", pc)
}

func (c *Processor) op0054() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b (A4),A0\n", pc)
}

func (c *Processor) op0055() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b (A5),A0\n", pc)
}

func (c *Processor) op0056() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b (A6),A0\n", pc)
}

func (c *Processor) op0057() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b (A7),A0\n", pc)
}

func (c *Processor) op0078() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b $%X,A0\n", pc, v)
}

func (c *Processor) op0079() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b $%X,A0\n", pc, v)
}

func (c *Processor) op007C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] = uint32(v)
	c.tracef("%04X movea.b #$%X,A0\n", pc, v)
}

func (c *Processor) op0080() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D0,(A0)\n", pc)
}

func (c *Processor) op0081() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D1,(A0)\n", pc)
}

func (c *Processor) op0082() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D2,(A0)\n", pc)
}

func (c *Processor) op0083() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D3,(A0)\n", pc)
}

func (c *Processor) op0084() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D4,(A0)\n", pc)
}

func (c *Processor) op0085() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D5,(A0)\n", pc)
}

func (c *Processor) op0086() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D6,(A0)\n", pc)
}

func (c *Processor) op0087() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D7,(A0)\n", pc)
}

func (c *Processor) op0088() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A0,(A0)\n", pc)
}

func (c *Processor) op0089() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A1,(A0)\n", pc)
}

func (c *Processor) op008A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A2,(A0)\n", pc)
}

func (c *Processor) op008B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A3,(A0)\n", pc)
}

func (c *Processor) op008C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A4,(A0)\n", pc)
}

func (c *Processor) op008D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A5,(A0)\n", pc)
}

func (c *Processor) op008E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A6,(A0)\n", pc)
}

func (c *Processor) op008F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A7,(A0)\n", pc)
}

func (c *Processor) op0090() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A0),(A0)\n", pc)
}

func (c *Processor) op0091() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A1),(A0)\n", pc)
}

func (c *Processor) op0092() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A2),(A0)\n", pc)
}

func (c *Processor) op0093() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A3),(A0)\n", pc)
}

func (c *Processor) op0094() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A4),(A0)\n", pc)
}

func (c *Processor) op0095() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A5),(A0)\n", pc)
}

func (c *Processor) op0096() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A6),(A0)\n", pc)
}

func (c *Processor) op0097() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A7),(A0)\n", pc)
}

func (c *Processor) op00B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b $%X,(A0)\n", pc, v)
}

func (c *Processor) op00B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b $%X,(A0)\n", pc, v)
}

func (c *Processor) op00BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b #$%X,(A0)\n", pc, v)
}

func (c *Processor) op00C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b D0,(A0)+\n", pc)
}

func (c *Processor) op00C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b D1,(A0)+\n", pc)
}

func (c *Processor) op00C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b D2,(A0)+\n", pc)
}

func (c *Processor) op00C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b D3,(A0)+\n", pc)
}

func (c *Processor) op00C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b D4,(A0)+\n", pc)
}

func (c *Processor) op00C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b D5,(A0)+\n", pc)
}

func (c *Processor) op00C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b D6,(A0)+\n", pc)
}

func (c *Processor) op00C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b D7,(A0)+\n", pc)
}

func (c *Processor) op00C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b A0,(A0)+\n", pc)
}

func (c *Processor) op00C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b A1,(A0)+\n", pc)
}

func (c *Processor) op00CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b A2,(A0)+\n", pc)
}

func (c *Processor) op00CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b A3,(A0)+\n", pc)
}

func (c *Processor) op00CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b A4,(A0)+\n", pc)
}

func (c *Processor) op00CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b A5,(A0)+\n", pc)
}

func (c *Processor) op00CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b A6,(A0)+\n", pc)
}

func (c *Processor) op00CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b A7,(A0)+\n", pc)
}

func (c *Processor) op00D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b (A0),(A0)+\n", pc)
}

func (c *Processor) op00D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b (A1),(A0)+\n", pc)
}

func (c *Processor) op00D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b (A2),(A0)+\n", pc)
}

func (c *Processor) op00D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b (A3),(A0)+\n", pc)
}

func (c *Processor) op00D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b (A4),(A0)+\n", pc)
}

func (c *Processor) op00D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b (A5),(A0)+\n", pc)
}

func (c *Processor) op00D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b (A6),(A0)+\n", pc)
}

func (c *Processor) op00D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b (A7),(A0)+\n", pc)
}

func (c *Processor) op00F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b $%X,(A0)+\n", pc, v)
}

func (c *Processor) op00F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b $%X,(A0)+\n", pc, v)
}

func (c *Processor) op00FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.b #$%X,(A0)+\n", pc, v)
}

func (c *Processor) op0100() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D0,-(A0)\n", pc)
}

func (c *Processor) op0101() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D1,-(A0)\n", pc)
}

func (c *Processor) op0102() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D2,-(A0)\n", pc)
}

func (c *Processor) op0103() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D3,-(A0)\n", pc)
}

func (c *Processor) op0104() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D4,-(A0)\n", pc)
}

func (c *Processor) op0105() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D5,-(A0)\n", pc)
}

func (c *Processor) op0106() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D6,-(A0)\n", pc)
}

func (c *Processor) op0107() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b D7,-(A0)\n", pc)
}

func (c *Processor) op0108() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A0,-(A0)\n", pc)
}

func (c *Processor) op0109() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A1,-(A0)\n", pc)
}

func (c *Processor) op010A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A2,-(A0)\n", pc)
}

func (c *Processor) op010B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A3,-(A0)\n", pc)
}

func (c *Processor) op010C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A4,-(A0)\n", pc)
}

func (c *Processor) op010D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A5,-(A0)\n", pc)
}

func (c *Processor) op010E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A6,-(A0)\n", pc)
}

func (c *Processor) op010F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b A7,-(A0)\n", pc)
}

func (c *Processor) op0110() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A0),-(A0)\n", pc)
}

func (c *Processor) op0111() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A1),-(A0)\n", pc)
}

func (c *Processor) op0112() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A2),-(A0)\n", pc)
}

func (c *Processor) op0113() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A3),-(A0)\n", pc)
}

func (c *Processor) op0114() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A4),-(A0)\n", pc)
}

func (c *Processor) op0115() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A5),-(A0)\n", pc)
}

func (c *Processor) op0116() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A6),-(A0)\n", pc)
}

func (c *Processor) op0117() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b (A7),-(A0)\n", pc)
}

func (c *Processor) op0138() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b $%X,-(A0)\n", pc, v)
}

func (c *Processor) op0139() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b $%X,-(A0)\n", pc, v)
}

func (c *Processor) op013C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.b #$%X,-(A0)\n", pc, v)
}

func (c *Processor) op0140() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D0,(%d,A0)\n", pc, disp)
}

func (c *Processor) op0141() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D1,(%d,A0)\n", pc, disp)
}

func (c *Processor) op0142() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D2,(%d,A0)\n", pc, disp)
}

func (c *Processor) op0143() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D3,(%d,A0)\n", pc, disp)
}

func (c *Processor) op0144() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D4,(%d,A0)\n", pc, disp)
}

func (c *Processor) op0145() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D5,(%d,A0)\n", pc, disp)
}

func (c *Processor) op0146() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D6,(%d,A0)\n", pc, disp)
}

func (c *Processor) op0147() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D7,(%d,A0)\n", pc, disp)
}

func (c *Processor) op0148() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A0,(%d,A0)\n", pc, disp)
}

func (c *Processor) op0149() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A1,(%d,A0)\n", pc, disp)
}

func (c *Processor) op014A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A2,(%d,A0)\n", pc, disp)
}

func (c *Processor) op014B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A3,(%d,A0)\n", pc, disp)
}

func (c *Processor) op014C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A4,(%d,A0)\n", pc, disp)
}

func (c *Processor) op014D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A5,(%d,A0)\n", pc, disp)
}

func (c *Processor) op014E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A6,(%d,A0)\n", pc, disp)
}

func (c *Processor) op014F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A7,(%d,A0)\n", pc, disp)
}

func (c *Processor) op0150() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A0),(%d,A0)\n", pc, disp)
}

func (c *Processor) op0151() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A1),(%d,A0)\n", pc, disp)
}

func (c *Processor) op0152() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A2),(%d,A0)\n", pc, disp)
}

func (c *Processor) op0153() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A3),(%d,A0)\n", pc, disp)
}

func (c *Processor) op0154() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A4),(%d,A0)\n", pc, disp)
}

func (c *Processor) op0155() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A5),(%d,A0)\n", pc, disp)
}

func (c *Processor) op0156() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A6),(%d,A0)\n", pc, disp)
}

func (c *Processor) op0157() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A7),(%d,A0)\n", pc, disp)
}

func (c *Processor) op0178() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A0)\n", pc, v, disp)
}

func (c *Processor) op0179() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A0)\n", pc, v, disp)
}

func (c *Processor) op017C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b #$%X,(%d,A0)\n", pc, v, disp)
}

func (c *Processor) op01C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D0,$%X\n", pc, addr)
}

func (c *Processor) op01C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D1,$%X\n", pc, addr)
}

func (c *Processor) op01C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D2,$%X\n", pc, addr)
}

func (c *Processor) op01C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D3,$%X\n", pc, addr)
}

func (c *Processor) op01C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D4,$%X\n", pc, addr)
}

func (c *Processor) op01C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D5,$%X\n", pc, addr)
}

func (c *Processor) op01C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D6,$%X\n", pc, addr)
}

func (c *Processor) op01C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D7,$%X\n", pc, addr)
}

func (c *Processor) op01C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A0,$%X\n", pc, addr)
}

func (c *Processor) op01C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A1,$%X\n", pc, addr)
}

func (c *Processor) op01CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A2,$%X\n", pc, addr)
}

func (c *Processor) op01CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A3,$%X\n", pc, addr)
}

func (c *Processor) op01CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A4,$%X\n", pc, addr)
}

func (c *Processor) op01CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A5,$%X\n", pc, addr)
}

func (c *Processor) op01CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A6,$%X\n", pc, addr)
}

func (c *Processor) op01CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A7,$%X\n", pc, addr)
}

func (c *Processor) op01D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A0),$%X\n", pc, addr)
}

func (c *Processor) op01D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A1),$%X\n", pc, addr)
}

func (c *Processor) op01D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A2),$%X\n", pc, addr)
}

func (c *Processor) op01D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A3),$%X\n", pc, addr)
}

func (c *Processor) op01D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A4),$%X\n", pc, addr)
}

func (c *Processor) op01D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A5),$%X\n", pc, addr)
}

func (c *Processor) op01D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A6),$%X\n", pc, addr)
}

func (c *Processor) op01D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A7),$%X\n", pc, addr)
}

func (c *Processor) op01F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,$%X\n", pc, v, addr)
}

func (c *Processor) op01F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,$%X\n", pc, v, addr)
}

func (c *Processor) op01FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b #$%X,$%X\n", pc, v, addr)
}

func (c *Processor) op0200() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b D0,D1\n", pc)
}

func (c *Processor) op0201() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b D1,D1\n", pc)
}

func (c *Processor) op0202() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b D2,D1\n", pc)
}

func (c *Processor) op0203() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b D3,D1\n", pc)
}

func (c *Processor) op0204() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b D4,D1\n", pc)
}

func (c *Processor) op0205() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b D5,D1\n", pc)
}

func (c *Processor) op0206() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b D6,D1\n", pc)
}

func (c *Processor) op0207() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b D7,D1\n", pc)
}

func (c *Processor) op0208() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b A0,D1\n", pc)
}

func (c *Processor) op0209() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b A1,D1\n", pc)
}

func (c *Processor) op020A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b A2,D1\n", pc)
}

func (c *Processor) op020B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b A3,D1\n", pc)
}

func (c *Processor) op020C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b A4,D1\n", pc)
}

func (c *Processor) op020D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b A5,D1\n", pc)
}

func (c *Processor) op020E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b A6,D1\n", pc)
}

func (c *Processor) op020F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b A7,D1\n", pc)
}

func (c *Processor) op0210() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b (A0),D1\n", pc)
}

func (c *Processor) op0211() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b (A1),D1\n", pc)
}

func (c *Processor) op0212() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b (A2),D1\n", pc)
}

func (c *Processor) op0213() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b (A3),D1\n", pc)
}

func (c *Processor) op0214() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b (A4),D1\n", pc)
}

func (c *Processor) op0215() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b (A5),D1\n", pc)
}

func (c *Processor) op0216() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b (A6),D1\n", pc)
}

func (c *Processor) op0217() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.b (A7),D1\n", pc)
}

func (c *Processor) op0238() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[1] = uint32(v)
	c.tracef("%04X move.b $%X,D1\n", pc, v)
}

func (c *Processor) op0239() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[1] = uint32(v)
	c.tracef("%04X move.b $%X,D1\n", pc, v)
}

func (c *Processor) op023C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[1] = uint32(v)
	c.tracef("%04X move.b #$%X,D1\n", pc, v)
}

func (c *Processor) op0240() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b D0,A1\n", pc)
}

func (c *Processor) op0241() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b D1,A1\n", pc)
}

func (c *Processor) op0242() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b D2,A1\n", pc)
}

func (c *Processor) op0243() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b D3,A1\n", pc)
}

func (c *Processor) op0244() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b D4,A1\n", pc)
}

func (c *Processor) op0245() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b D5,A1\n", pc)
}

func (c *Processor) op0246() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b D6,A1\n", pc)
}

func (c *Processor) op0247() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b D7,A1\n", pc)
}

func (c *Processor) op0248() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b A0,A1\n", pc)
}

func (c *Processor) op0249() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b A1,A1\n", pc)
}

func (c *Processor) op024A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b A2,A1\n", pc)
}

func (c *Processor) op024B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b A3,A1\n", pc)
}

func (c *Processor) op024C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b A4,A1\n", pc)
}

func (c *Processor) op024D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b A5,A1\n", pc)
}

func (c *Processor) op024E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b A6,A1\n", pc)
}

func (c *Processor) op024F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b A7,A1\n", pc)
}

func (c *Processor) op0250() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b (A0),A1\n", pc)
}

func (c *Processor) op0251() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b (A1),A1\n", pc)
}

func (c *Processor) op0252() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b (A2),A1\n", pc)
}

func (c *Processor) op0253() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b (A3),A1\n", pc)
}

func (c *Processor) op0254() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b (A4),A1\n", pc)
}

func (c *Processor) op0255() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b (A5),A1\n", pc)
}

func (c *Processor) op0256() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b (A6),A1\n", pc)
}

func (c *Processor) op0257() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b (A7),A1\n", pc)
}

func (c *Processor) op0278() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b $%X,A1\n", pc, v)
}

func (c *Processor) op0279() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b $%X,A1\n", pc, v)
}

func (c *Processor) op027C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] = uint32(v)
	c.tracef("%04X movea.b #$%X,A1\n", pc, v)
}

func (c *Processor) op0280() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D0,(A1)\n", pc)
}

func (c *Processor) op0281() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D1,(A1)\n", pc)
}

func (c *Processor) op0282() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D2,(A1)\n", pc)
}

func (c *Processor) op0283() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D3,(A1)\n", pc)
}

func (c *Processor) op0284() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D4,(A1)\n", pc)
}

func (c *Processor) op0285() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D5,(A1)\n", pc)
}

func (c *Processor) op0286() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D6,(A1)\n", pc)
}

func (c *Processor) op0287() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D7,(A1)\n", pc)
}

func (c *Processor) op0288() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A0,(A1)\n", pc)
}

func (c *Processor) op0289() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A1,(A1)\n", pc)
}

func (c *Processor) op028A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A2,(A1)\n", pc)
}

func (c *Processor) op028B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A3,(A1)\n", pc)
}

func (c *Processor) op028C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A4,(A1)\n", pc)
}

func (c *Processor) op028D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A5,(A1)\n", pc)
}

func (c *Processor) op028E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A6,(A1)\n", pc)
}

func (c *Processor) op028F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A7,(A1)\n", pc)
}

func (c *Processor) op0290() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A0),(A1)\n", pc)
}

func (c *Processor) op0291() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A1),(A1)\n", pc)
}

func (c *Processor) op0292() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A2),(A1)\n", pc)
}

func (c *Processor) op0293() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A3),(A1)\n", pc)
}

func (c *Processor) op0294() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A4),(A1)\n", pc)
}

func (c *Processor) op0295() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A5),(A1)\n", pc)
}

func (c *Processor) op0296() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A6),(A1)\n", pc)
}

func (c *Processor) op0297() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A7),(A1)\n", pc)
}

func (c *Processor) op02B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b $%X,(A1)\n", pc, v)
}

func (c *Processor) op02B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b $%X,(A1)\n", pc, v)
}

func (c *Processor) op02BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b #$%X,(A1)\n", pc, v)
}

func (c *Processor) op02C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b D0,(A1)+\n", pc)
}

func (c *Processor) op02C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b D1,(A1)+\n", pc)
}

func (c *Processor) op02C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b D2,(A1)+\n", pc)
}

func (c *Processor) op02C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b D3,(A1)+\n", pc)
}

func (c *Processor) op02C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b D4,(A1)+\n", pc)
}

func (c *Processor) op02C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b D5,(A1)+\n", pc)
}

func (c *Processor) op02C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b D6,(A1)+\n", pc)
}

func (c *Processor) op02C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b D7,(A1)+\n", pc)
}

func (c *Processor) op02C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b A0,(A1)+\n", pc)
}

func (c *Processor) op02C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b A1,(A1)+\n", pc)
}

func (c *Processor) op02CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b A2,(A1)+\n", pc)
}

func (c *Processor) op02CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b A3,(A1)+\n", pc)
}

func (c *Processor) op02CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b A4,(A1)+\n", pc)
}

func (c *Processor) op02CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b A5,(A1)+\n", pc)
}

func (c *Processor) op02CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b A6,(A1)+\n", pc)
}

func (c *Processor) op02CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b A7,(A1)+\n", pc)
}

func (c *Processor) op02D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b (A0),(A1)+\n", pc)
}

func (c *Processor) op02D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b (A1),(A1)+\n", pc)
}

func (c *Processor) op02D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b (A2),(A1)+\n", pc)
}

func (c *Processor) op02D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b (A3),(A1)+\n", pc)
}

func (c *Processor) op02D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b (A4),(A1)+\n", pc)
}

func (c *Processor) op02D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b (A5),(A1)+\n", pc)
}

func (c *Processor) op02D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b (A6),(A1)+\n", pc)
}

func (c *Processor) op02D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b (A7),(A1)+\n", pc)
}

func (c *Processor) op02F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b $%X,(A1)+\n", pc, v)
}

func (c *Processor) op02F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b $%X,(A1)+\n", pc, v)
}

func (c *Processor) op02FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.b #$%X,(A1)+\n", pc, v)
}

func (c *Processor) op0300() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D0,-(A1)\n", pc)
}

func (c *Processor) op0301() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D1,-(A1)\n", pc)
}

func (c *Processor) op0302() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D2,-(A1)\n", pc)
}

func (c *Processor) op0303() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D3,-(A1)\n", pc)
}

func (c *Processor) op0304() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D4,-(A1)\n", pc)
}

func (c *Processor) op0305() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D5,-(A1)\n", pc)
}

func (c *Processor) op0306() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D6,-(A1)\n", pc)
}

func (c *Processor) op0307() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b D7,-(A1)\n", pc)
}

func (c *Processor) op0308() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A0,-(A1)\n", pc)
}

func (c *Processor) op0309() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A1,-(A1)\n", pc)
}

func (c *Processor) op030A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A2,-(A1)\n", pc)
}

func (c *Processor) op030B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A3,-(A1)\n", pc)
}

func (c *Processor) op030C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A4,-(A1)\n", pc)
}

func (c *Processor) op030D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A5,-(A1)\n", pc)
}

func (c *Processor) op030E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A6,-(A1)\n", pc)
}

func (c *Processor) op030F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b A7,-(A1)\n", pc)
}

func (c *Processor) op0310() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A0),-(A1)\n", pc)
}

func (c *Processor) op0311() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A1),-(A1)\n", pc)
}

func (c *Processor) op0312() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A2),-(A1)\n", pc)
}

func (c *Processor) op0313() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A3),-(A1)\n", pc)
}

func (c *Processor) op0314() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A4),-(A1)\n", pc)
}

func (c *Processor) op0315() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A5),-(A1)\n", pc)
}

func (c *Processor) op0316() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A6),-(A1)\n", pc)
}

func (c *Processor) op0317() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b (A7),-(A1)\n", pc)
}

func (c *Processor) op0338() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b $%X,-(A1)\n", pc, v)
}

func (c *Processor) op0339() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b $%X,-(A1)\n", pc, v)
}

func (c *Processor) op033C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.b #$%X,-(A1)\n", pc, v)
}

func (c *Processor) op0340() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D0,(%d,A1)\n", pc, disp)
}

func (c *Processor) op0341() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D1,(%d,A1)\n", pc, disp)
}

func (c *Processor) op0342() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D2,(%d,A1)\n", pc, disp)
}

func (c *Processor) op0343() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D3,(%d,A1)\n", pc, disp)
}

func (c *Processor) op0344() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D4,(%d,A1)\n", pc, disp)
}

func (c *Processor) op0345() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D5,(%d,A1)\n", pc, disp)
}

func (c *Processor) op0346() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D6,(%d,A1)\n", pc, disp)
}

func (c *Processor) op0347() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D7,(%d,A1)\n", pc, disp)
}

func (c *Processor) op0348() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A0,(%d,A1)\n", pc, disp)
}

func (c *Processor) op0349() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A1,(%d,A1)\n", pc, disp)
}

func (c *Processor) op034A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A2,(%d,A1)\n", pc, disp)
}

func (c *Processor) op034B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A3,(%d,A1)\n", pc, disp)
}

func (c *Processor) op034C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A4,(%d,A1)\n", pc, disp)
}

func (c *Processor) op034D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A5,(%d,A1)\n", pc, disp)
}

func (c *Processor) op034E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A6,(%d,A1)\n", pc, disp)
}

func (c *Processor) op034F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A7,(%d,A1)\n", pc, disp)
}

func (c *Processor) op0350() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A0),(%d,A1)\n", pc, disp)
}

func (c *Processor) op0351() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A1),(%d,A1)\n", pc, disp)
}

func (c *Processor) op0352() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A2),(%d,A1)\n", pc, disp)
}

func (c *Processor) op0353() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A3),(%d,A1)\n", pc, disp)
}

func (c *Processor) op0354() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A4),(%d,A1)\n", pc, disp)
}

func (c *Processor) op0355() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A5),(%d,A1)\n", pc, disp)
}

func (c *Processor) op0356() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A6),(%d,A1)\n", pc, disp)
}

func (c *Processor) op0357() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A7),(%d,A1)\n", pc, disp)
}

func (c *Processor) op0378() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A1)\n", pc, v, disp)
}

func (c *Processor) op0379() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A1)\n", pc, v, disp)
}

func (c *Processor) op037C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b #$%X,(%d,A1)\n", pc, v, disp)
}

func (c *Processor) op03C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D0,$%X\n", pc, addr)
}

func (c *Processor) op03C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D1,$%X\n", pc, addr)
}

func (c *Processor) op03C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D2,$%X\n", pc, addr)
}

func (c *Processor) op03C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D3,$%X\n", pc, addr)
}

func (c *Processor) op03C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D4,$%X\n", pc, addr)
}

func (c *Processor) op03C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D5,$%X\n", pc, addr)
}

func (c *Processor) op03C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D6,$%X\n", pc, addr)
}

func (c *Processor) op03C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D7,$%X\n", pc, addr)
}

func (c *Processor) op03C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A0,$%X\n", pc, addr)
}

func (c *Processor) op03C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A1,$%X\n", pc, addr)
}

func (c *Processor) op03CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A2,$%X\n", pc, addr)
}

func (c *Processor) op03CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A3,$%X\n", pc, addr)
}

func (c *Processor) op03CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A4,$%X\n", pc, addr)
}

func (c *Processor) op03CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A5,$%X\n", pc, addr)
}

func (c *Processor) op03CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A6,$%X\n", pc, addr)
}

func (c *Processor) op03CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A7,$%X\n", pc, addr)
}

func (c *Processor) op03D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A0),$%X\n", pc, addr)
}

func (c *Processor) op03D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A1),$%X\n", pc, addr)
}

func (c *Processor) op03D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A2),$%X\n", pc, addr)
}

func (c *Processor) op03D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A3),$%X\n", pc, addr)
}

func (c *Processor) op03D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A4),$%X\n", pc, addr)
}

func (c *Processor) op03D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A5),$%X\n", pc, addr)
}

func (c *Processor) op03D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A6),$%X\n", pc, addr)
}

func (c *Processor) op03D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A7),$%X\n", pc, addr)
}

func (c *Processor) op03F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,$%X\n", pc, v, addr)
}

func (c *Processor) op03F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,$%X\n", pc, v, addr)
}

func (c *Processor) op03FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b #$%X,$%X\n", pc, v, addr)
}

func (c *Processor) op0400() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b D0,D2\n", pc)
}

func (c *Processor) op0401() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b D1,D2\n", pc)
}

func (c *Processor) op0402() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b D2,D2\n", pc)
}

func (c *Processor) op0403() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b D3,D2\n", pc)
}

func (c *Processor) op0404() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b D4,D2\n", pc)
}

func (c *Processor) op0405() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b D5,D2\n", pc)
}

func (c *Processor) op0406() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b D6,D2\n", pc)
}

func (c *Processor) op0407() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b D7,D2\n", pc)
}

func (c *Processor) op0408() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b A0,D2\n", pc)
}

func (c *Processor) op0409() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b A1,D2\n", pc)
}

func (c *Processor) op040A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b A2,D2\n", pc)
}

func (c *Processor) op040B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b A3,D2\n", pc)
}

func (c *Processor) op040C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b A4,D2\n", pc)
}

func (c *Processor) op040D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b A5,D2\n", pc)
}

func (c *Processor) op040E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b A6,D2\n", pc)
}

func (c *Processor) op040F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b A7,D2\n", pc)
}

func (c *Processor) op0410() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b (A0),D2\n", pc)
}

func (c *Processor) op0411() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b (A1),D2\n", pc)
}

func (c *Processor) op0412() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b (A2),D2\n", pc)
}

func (c *Processor) op0413() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b (A3),D2\n", pc)
}

func (c *Processor) op0414() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b (A4),D2\n", pc)
}

func (c *Processor) op0415() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b (A5),D2\n", pc)
}

func (c *Processor) op0416() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b (A6),D2\n", pc)
}

func (c *Processor) op0417() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.b (A7),D2\n", pc)
}

func (c *Processor) op0438() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[2] = uint32(v)
	c.tracef("%04X move.b $%X,D2\n", pc, v)
}

func (c *Processor) op0439() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[2] = uint32(v)
	c.tracef("%04X move.b $%X,D2\n", pc, v)
}

func (c *Processor) op043C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[2] = uint32(v)
	c.tracef("%04X move.b #$%X,D2\n", pc, v)
}

func (c *Processor) op0440() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b D0,A2\n", pc)
}

func (c *Processor) op0441() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b D1,A2\n", pc)
}

func (c *Processor) op0442() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b D2,A2\n", pc)
}

func (c *Processor) op0443() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b D3,A2\n", pc)
}

func (c *Processor) op0444() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b D4,A2\n", pc)
}

func (c *Processor) op0445() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b D5,A2\n", pc)
}

func (c *Processor) op0446() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b D6,A2\n", pc)
}

func (c *Processor) op0447() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b D7,A2\n", pc)
}

func (c *Processor) op0448() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b A0,A2\n", pc)
}

func (c *Processor) op0449() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b A1,A2\n", pc)
}

func (c *Processor) op044A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b A2,A2\n", pc)
}

func (c *Processor) op044B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b A3,A2\n", pc)
}

func (c *Processor) op044C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b A4,A2\n", pc)
}

func (c *Processor) op044D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b A5,A2\n", pc)
}

func (c *Processor) op044E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b A6,A2\n", pc)
}

func (c *Processor) op044F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b A7,A2\n", pc)
}

func (c *Processor) op0450() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b (A0),A2\n", pc)
}

func (c *Processor) op0451() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b (A1),A2\n", pc)
}

func (c *Processor) op0452() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b (A2),A2\n", pc)
}

func (c *Processor) op0453() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b (A3),A2\n", pc)
}

func (c *Processor) op0454() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b (A4),A2\n", pc)
}

func (c *Processor) op0455() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b (A5),A2\n", pc)
}

func (c *Processor) op0456() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b (A6),A2\n", pc)
}

func (c *Processor) op0457() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b (A7),A2\n", pc)
}

func (c *Processor) op0478() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b $%X,A2\n", pc, v)
}

func (c *Processor) op0479() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b $%X,A2\n", pc, v)
}

func (c *Processor) op047C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] = uint32(v)
	c.tracef("%04X movea.b #$%X,A2\n", pc, v)
}

func (c *Processor) op0480() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D0,(A2)\n", pc)
}

func (c *Processor) op0481() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D1,(A2)\n", pc)
}

func (c *Processor) op0482() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D2,(A2)\n", pc)
}

func (c *Processor) op0483() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D3,(A2)\n", pc)
}

func (c *Processor) op0484() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D4,(A2)\n", pc)
}

func (c *Processor) op0485() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D5,(A2)\n", pc)
}

func (c *Processor) op0486() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D6,(A2)\n", pc)
}

func (c *Processor) op0487() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D7,(A2)\n", pc)
}

func (c *Processor) op0488() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A0,(A2)\n", pc)
}

func (c *Processor) op0489() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A1,(A2)\n", pc)
}

func (c *Processor) op048A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A2,(A2)\n", pc)
}

func (c *Processor) op048B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A3,(A2)\n", pc)
}

func (c *Processor) op048C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A4,(A2)\n", pc)
}

func (c *Processor) op048D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A5,(A2)\n", pc)
}

func (c *Processor) op048E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A6,(A2)\n", pc)
}

func (c *Processor) op048F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A7,(A2)\n", pc)
}

func (c *Processor) op0490() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A0),(A2)\n", pc)
}

func (c *Processor) op0491() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A1),(A2)\n", pc)
}

func (c *Processor) op0492() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A2),(A2)\n", pc)
}

func (c *Processor) op0493() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A3),(A2)\n", pc)
}

func (c *Processor) op0494() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A4),(A2)\n", pc)
}

func (c *Processor) op0495() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A5),(A2)\n", pc)
}

func (c *Processor) op0496() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A6),(A2)\n", pc)
}

func (c *Processor) op0497() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A7),(A2)\n", pc)
}

func (c *Processor) op04B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b $%X,(A2)\n", pc, v)
}

func (c *Processor) op04B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b $%X,(A2)\n", pc, v)
}

func (c *Processor) op04BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b #$%X,(A2)\n", pc, v)
}

func (c *Processor) op04C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b D0,(A2)+\n", pc)
}

func (c *Processor) op04C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b D1,(A2)+\n", pc)
}

func (c *Processor) op04C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b D2,(A2)+\n", pc)
}

func (c *Processor) op04C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b D3,(A2)+\n", pc)
}

func (c *Processor) op04C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b D4,(A2)+\n", pc)
}

func (c *Processor) op04C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b D5,(A2)+\n", pc)
}

func (c *Processor) op04C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b D6,(A2)+\n", pc)
}

func (c *Processor) op04C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b D7,(A2)+\n", pc)
}

func (c *Processor) op04C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b A0,(A2)+\n", pc)
}

func (c *Processor) op04C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b A1,(A2)+\n", pc)
}

func (c *Processor) op04CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b A2,(A2)+\n", pc)
}

func (c *Processor) op04CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b A3,(A2)+\n", pc)
}

func (c *Processor) op04CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b A4,(A2)+\n", pc)
}

func (c *Processor) op04CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b A5,(A2)+\n", pc)
}

func (c *Processor) op04CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b A6,(A2)+\n", pc)
}

func (c *Processor) op04CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b A7,(A2)+\n", pc)
}

func (c *Processor) op04D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b (A0),(A2)+\n", pc)
}

func (c *Processor) op04D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b (A1),(A2)+\n", pc)
}

func (c *Processor) op04D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b (A2),(A2)+\n", pc)
}

func (c *Processor) op04D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b (A3),(A2)+\n", pc)
}

func (c *Processor) op04D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b (A4),(A2)+\n", pc)
}

func (c *Processor) op04D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b (A5),(A2)+\n", pc)
}

func (c *Processor) op04D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b (A6),(A2)+\n", pc)
}

func (c *Processor) op04D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b (A7),(A2)+\n", pc)
}

func (c *Processor) op04F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b $%X,(A2)+\n", pc, v)
}

func (c *Processor) op04F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b $%X,(A2)+\n", pc, v)
}

func (c *Processor) op04FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.b #$%X,(A2)+\n", pc, v)
}

func (c *Processor) op0500() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D0,-(A2)\n", pc)
}

func (c *Processor) op0501() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D1,-(A2)\n", pc)
}

func (c *Processor) op0502() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D2,-(A2)\n", pc)
}

func (c *Processor) op0503() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D3,-(A2)\n", pc)
}

func (c *Processor) op0504() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D4,-(A2)\n", pc)
}

func (c *Processor) op0505() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D5,-(A2)\n", pc)
}

func (c *Processor) op0506() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D6,-(A2)\n", pc)
}

func (c *Processor) op0507() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b D7,-(A2)\n", pc)
}

func (c *Processor) op0508() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A0,-(A2)\n", pc)
}

func (c *Processor) op0509() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A1,-(A2)\n", pc)
}

func (c *Processor) op050A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A2,-(A2)\n", pc)
}

func (c *Processor) op050B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A3,-(A2)\n", pc)
}

func (c *Processor) op050C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A4,-(A2)\n", pc)
}

func (c *Processor) op050D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A5,-(A2)\n", pc)
}

func (c *Processor) op050E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A6,-(A2)\n", pc)
}

func (c *Processor) op050F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b A7,-(A2)\n", pc)
}

func (c *Processor) op0510() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A0),-(A2)\n", pc)
}

func (c *Processor) op0511() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A1),-(A2)\n", pc)
}

func (c *Processor) op0512() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A2),-(A2)\n", pc)
}

func (c *Processor) op0513() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A3),-(A2)\n", pc)
}

func (c *Processor) op0514() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A4),-(A2)\n", pc)
}

func (c *Processor) op0515() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A5),-(A2)\n", pc)
}

func (c *Processor) op0516() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A6),-(A2)\n", pc)
}

func (c *Processor) op0517() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b (A7),-(A2)\n", pc)
}

func (c *Processor) op0538() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b $%X,-(A2)\n", pc, v)
}

func (c *Processor) op0539() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b $%X,-(A2)\n", pc, v)
}

func (c *Processor) op053C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.b #$%X,-(A2)\n", pc, v)
}

func (c *Processor) op0540() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D0,(%d,A2)\n", pc, disp)
}

func (c *Processor) op0541() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D1,(%d,A2)\n", pc, disp)
}

func (c *Processor) op0542() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D2,(%d,A2)\n", pc, disp)
}

func (c *Processor) op0543() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D3,(%d,A2)\n", pc, disp)
}

func (c *Processor) op0544() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D4,(%d,A2)\n", pc, disp)
}

func (c *Processor) op0545() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D5,(%d,A2)\n", pc, disp)
}

func (c *Processor) op0546() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D6,(%d,A2)\n", pc, disp)
}

func (c *Processor) op0547() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D7,(%d,A2)\n", pc, disp)
}

func (c *Processor) op0548() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A0,(%d,A2)\n", pc, disp)
}

func (c *Processor) op0549() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A1,(%d,A2)\n", pc, disp)
}

func (c *Processor) op054A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A2,(%d,A2)\n", pc, disp)
}

func (c *Processor) op054B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A3,(%d,A2)\n", pc, disp)
}

func (c *Processor) op054C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A4,(%d,A2)\n", pc, disp)
}

func (c *Processor) op054D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A5,(%d,A2)\n", pc, disp)
}

func (c *Processor) op054E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A6,(%d,A2)\n", pc, disp)
}

func (c *Processor) op054F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A7,(%d,A2)\n", pc, disp)
}

func (c *Processor) op0550() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A0),(%d,A2)\n", pc, disp)
}

func (c *Processor) op0551() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A1),(%d,A2)\n", pc, disp)
}

func (c *Processor) op0552() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A2),(%d,A2)\n", pc, disp)
}

func (c *Processor) op0553() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A3),(%d,A2)\n", pc, disp)
}

func (c *Processor) op0554() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A4),(%d,A2)\n", pc, disp)
}

func (c *Processor) op0555() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A5),(%d,A2)\n", pc, disp)
}

func (c *Processor) op0556() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A6),(%d,A2)\n", pc, disp)
}

func (c *Processor) op0557() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A7),(%d,A2)\n", pc, disp)
}

func (c *Processor) op0578() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A2)\n", pc, v, disp)
}

func (c *Processor) op0579() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A2)\n", pc, v, disp)
}

func (c *Processor) op057C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b #$%X,(%d,A2)\n", pc, v, disp)
}

func (c *Processor) op0600() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b D0,D3\n", pc)
}

func (c *Processor) op0601() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b D1,D3\n", pc)
}

func (c *Processor) op0602() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b D2,D3\n", pc)
}

func (c *Processor) op0603() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b D3,D3\n", pc)
}

func (c *Processor) op0604() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b D4,D3\n", pc)
}

func (c *Processor) op0605() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b D5,D3\n", pc)
}

func (c *Processor) op0606() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b D6,D3\n", pc)
}

func (c *Processor) op0607() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b D7,D3\n", pc)
}

func (c *Processor) op0608() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b A0,D3\n", pc)
}

func (c *Processor) op0609() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b A1,D3\n", pc)
}

func (c *Processor) op060A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b A2,D3\n", pc)
}

func (c *Processor) op060B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b A3,D3\n", pc)
}

func (c *Processor) op060C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b A4,D3\n", pc)
}

func (c *Processor) op060D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b A5,D3\n", pc)
}

func (c *Processor) op060E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b A6,D3\n", pc)
}

func (c *Processor) op060F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b A7,D3\n", pc)
}

func (c *Processor) op0610() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b (A0),D3\n", pc)
}

func (c *Processor) op0611() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b (A1),D3\n", pc)
}

func (c *Processor) op0612() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b (A2),D3\n", pc)
}

func (c *Processor) op0613() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b (A3),D3\n", pc)
}

func (c *Processor) op0614() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b (A4),D3\n", pc)
}

func (c *Processor) op0615() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b (A5),D3\n", pc)
}

func (c *Processor) op0616() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b (A6),D3\n", pc)
}

func (c *Processor) op0617() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.b (A7),D3\n", pc)
}

func (c *Processor) op0638() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[3] = uint32(v)
	c.tracef("%04X move.b $%X,D3\n", pc, v)
}

func (c *Processor) op0639() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[3] = uint32(v)
	c.tracef("%04X move.b $%X,D3\n", pc, v)
}

func (c *Processor) op063C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[3] = uint32(v)
	c.tracef("%04X move.b #$%X,D3\n", pc, v)
}

func (c *Processor) op0640() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b D0,A3\n", pc)
}

func (c *Processor) op0641() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b D1,A3\n", pc)
}

func (c *Processor) op0642() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b D2,A3\n", pc)
}

func (c *Processor) op0643() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b D3,A3\n", pc)
}

func (c *Processor) op0644() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b D4,A3\n", pc)
}

func (c *Processor) op0645() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b D5,A3\n", pc)
}

func (c *Processor) op0646() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b D6,A3\n", pc)
}

func (c *Processor) op0647() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b D7,A3\n", pc)
}

func (c *Processor) op0648() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b A0,A3\n", pc)
}

func (c *Processor) op0649() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b A1,A3\n", pc)
}

func (c *Processor) op064A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b A2,A3\n", pc)
}

func (c *Processor) op064B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b A3,A3\n", pc)
}

func (c *Processor) op064C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b A4,A3\n", pc)
}

func (c *Processor) op064D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b A5,A3\n", pc)
}

func (c *Processor) op064E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b A6,A3\n", pc)
}

func (c *Processor) op064F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b A7,A3\n", pc)
}

func (c *Processor) op0650() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b (A0),A3\n", pc)
}

func (c *Processor) op0651() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b (A1),A3\n", pc)
}

func (c *Processor) op0652() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b (A2),A3\n", pc)
}

func (c *Processor) op0653() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b (A3),A3\n", pc)
}

func (c *Processor) op0654() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b (A4),A3\n", pc)
}

func (c *Processor) op0655() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b (A5),A3\n", pc)
}

func (c *Processor) op0656() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b (A6),A3\n", pc)
}

func (c *Processor) op0657() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b (A7),A3\n", pc)
}

func (c *Processor) op0678() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b $%X,A3\n", pc, v)
}

func (c *Processor) op0679() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b $%X,A3\n", pc, v)
}

func (c *Processor) op067C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] = uint32(v)
	c.tracef("%04X movea.b #$%X,A3\n", pc, v)
}

func (c *Processor) op0680() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D0,(A3)\n", pc)
}

func (c *Processor) op0681() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D1,(A3)\n", pc)
}

func (c *Processor) op0682() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D2,(A3)\n", pc)
}

func (c *Processor) op0683() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D3,(A3)\n", pc)
}

func (c *Processor) op0684() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D4,(A3)\n", pc)
}

func (c *Processor) op0685() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D5,(A3)\n", pc)
}

func (c *Processor) op0686() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D6,(A3)\n", pc)
}

func (c *Processor) op0687() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D7,(A3)\n", pc)
}

func (c *Processor) op0688() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A0,(A3)\n", pc)
}

func (c *Processor) op0689() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A1,(A3)\n", pc)
}

func (c *Processor) op068A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A2,(A3)\n", pc)
}

func (c *Processor) op068B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A3,(A3)\n", pc)
}

func (c *Processor) op068C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A4,(A3)\n", pc)
}

func (c *Processor) op068D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A5,(A3)\n", pc)
}

func (c *Processor) op068E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A6,(A3)\n", pc)
}

func (c *Processor) op068F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A7,(A3)\n", pc)
}

func (c *Processor) op0690() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A0),(A3)\n", pc)
}

func (c *Processor) op0691() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A1),(A3)\n", pc)
}

func (c *Processor) op0692() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A2),(A3)\n", pc)
}

func (c *Processor) op0693() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A3),(A3)\n", pc)
}

func (c *Processor) op0694() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A4),(A3)\n", pc)
}

func (c *Processor) op0695() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A5),(A3)\n", pc)
}

func (c *Processor) op0696() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A6),(A3)\n", pc)
}

func (c *Processor) op0697() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A7),(A3)\n", pc)
}

func (c *Processor) op06B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b $%X,(A3)\n", pc, v)
}

func (c *Processor) op06B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b $%X,(A3)\n", pc, v)
}

func (c *Processor) op06BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b #$%X,(A3)\n", pc, v)
}

func (c *Processor) op06C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b D0,(A3)+\n", pc)
}

func (c *Processor) op06C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b D1,(A3)+\n", pc)
}

func (c *Processor) op06C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b D2,(A3)+\n", pc)
}

func (c *Processor) op06C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b D3,(A3)+\n", pc)
}

func (c *Processor) op06C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b D4,(A3)+\n", pc)
}

func (c *Processor) op06C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b D5,(A3)+\n", pc)
}

func (c *Processor) op06C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b D6,(A3)+\n", pc)
}

func (c *Processor) op06C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b D7,(A3)+\n", pc)
}

func (c *Processor) op06C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b A0,(A3)+\n", pc)
}

func (c *Processor) op06C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b A1,(A3)+\n", pc)
}

func (c *Processor) op06CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b A2,(A3)+\n", pc)
}

func (c *Processor) op06CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b A3,(A3)+\n", pc)
}

func (c *Processor) op06CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b A4,(A3)+\n", pc)
}

func (c *Processor) op06CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b A5,(A3)+\n", pc)
}

func (c *Processor) op06CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b A6,(A3)+\n", pc)
}

func (c *Processor) op06CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b A7,(A3)+\n", pc)
}

func (c *Processor) op06D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b (A0),(A3)+\n", pc)
}

func (c *Processor) op06D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b (A1),(A3)+\n", pc)
}

func (c *Processor) op06D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b (A2),(A3)+\n", pc)
}

func (c *Processor) op06D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b (A3),(A3)+\n", pc)
}

func (c *Processor) op06D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b (A4),(A3)+\n", pc)
}

func (c *Processor) op06D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b (A5),(A3)+\n", pc)
}

func (c *Processor) op06D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b (A6),(A3)+\n", pc)
}

func (c *Processor) op06D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b (A7),(A3)+\n", pc)
}

func (c *Processor) op06F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b $%X,(A3)+\n", pc, v)
}

func (c *Processor) op06F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b $%X,(A3)+\n", pc, v)
}

func (c *Processor) op06FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.b #$%X,(A3)+\n", pc, v)
}

func (c *Processor) op0700() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D0,-(A3)\n", pc)
}

func (c *Processor) op0701() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D1,-(A3)\n", pc)
}

func (c *Processor) op0702() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D2,-(A3)\n", pc)
}

func (c *Processor) op0703() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D3,-(A3)\n", pc)
}

func (c *Processor) op0704() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D4,-(A3)\n", pc)
}

func (c *Processor) op0705() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D5,-(A3)\n", pc)
}

func (c *Processor) op0706() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D6,-(A3)\n", pc)
}

func (c *Processor) op0707() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b D7,-(A3)\n", pc)
}

func (c *Processor) op0708() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A0,-(A3)\n", pc)
}

func (c *Processor) op0709() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A1,-(A3)\n", pc)
}

func (c *Processor) op070A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A2,-(A3)\n", pc)
}

func (c *Processor) op070B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A3,-(A3)\n", pc)
}

func (c *Processor) op070C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A4,-(A3)\n", pc)
}

func (c *Processor) op070D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A5,-(A3)\n", pc)
}

func (c *Processor) op070E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A6,-(A3)\n", pc)
}

func (c *Processor) op070F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b A7,-(A3)\n", pc)
}

func (c *Processor) op0710() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A0),-(A3)\n", pc)
}

func (c *Processor) op0711() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A1),-(A3)\n", pc)
}

func (c *Processor) op0712() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A2),-(A3)\n", pc)
}

func (c *Processor) op0713() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A3),-(A3)\n", pc)
}

func (c *Processor) op0714() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A4),-(A3)\n", pc)
}

func (c *Processor) op0715() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A5),-(A3)\n", pc)
}

func (c *Processor) op0716() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A6),-(A3)\n", pc)
}

func (c *Processor) op0717() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b (A7),-(A3)\n", pc)
}

func (c *Processor) op0738() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b $%X,-(A3)\n", pc, v)
}

func (c *Processor) op0739() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b $%X,-(A3)\n", pc, v)
}

func (c *Processor) op073C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.b #$%X,-(A3)\n", pc, v)
}

func (c *Processor) op0740() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D0,(%d,A3)\n", pc, disp)
}

func (c *Processor) op0741() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D1,(%d,A3)\n", pc, disp)
}

func (c *Processor) op0742() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D2,(%d,A3)\n", pc, disp)
}

func (c *Processor) op0743() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D3,(%d,A3)\n", pc, disp)
}

func (c *Processor) op0744() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D4,(%d,A3)\n", pc, disp)
}

func (c *Processor) op0745() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D5,(%d,A3)\n", pc, disp)
}

func (c *Processor) op0746() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D6,(%d,A3)\n", pc, disp)
}

func (c *Processor) op0747() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D7,(%d,A3)\n", pc, disp)
}

func (c *Processor) op0748() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A0,(%d,A3)\n", pc, disp)
}

func (c *Processor) op0749() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A1,(%d,A3)\n", pc, disp)
}

func (c *Processor) op074A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A2,(%d,A3)\n", pc, disp)
}

func (c *Processor) op074B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A3,(%d,A3)\n", pc, disp)
}

func (c *Processor) op074C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A4,(%d,A3)\n", pc, disp)
}

func (c *Processor) op074D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A5,(%d,A3)\n", pc, disp)
}

func (c *Processor) op074E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A6,(%d,A3)\n", pc, disp)
}

func (c *Processor) op074F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A7,(%d,A3)\n", pc, disp)
}

func (c *Processor) op0750() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A0),(%d,A3)\n", pc, disp)
}

func (c *Processor) op0751() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A1),(%d,A3)\n", pc, disp)
}

func (c *Processor) op0752() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A2),(%d,A3)\n", pc, disp)
}

func (c *Processor) op0753() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A3),(%d,A3)\n", pc, disp)
}

func (c *Processor) op0754() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A4),(%d,A3)\n", pc, disp)
}

func (c *Processor) op0755() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A5),(%d,A3)\n", pc, disp)
}

func (c *Processor) op0756() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A6),(%d,A3)\n", pc, disp)
}

func (c *Processor) op0757() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A7),(%d,A3)\n", pc, disp)
}

func (c *Processor) op0778() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A3)\n", pc, v, disp)
}

func (c *Processor) op0779() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A3)\n", pc, v, disp)
}

func (c *Processor) op077C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b #$%X,(%d,A3)\n", pc, v, disp)
}

func (c *Processor) op0800() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b D0,D4\n", pc)
}

func (c *Processor) op0801() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b D1,D4\n", pc)
}

func (c *Processor) op0802() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b D2,D4\n", pc)
}

func (c *Processor) op0803() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b D3,D4\n", pc)
}

func (c *Processor) op0804() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b D4,D4\n", pc)
}

func (c *Processor) op0805() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b D5,D4\n", pc)
}

func (c *Processor) op0806() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b D6,D4\n", pc)
}

func (c *Processor) op0807() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b D7,D4\n", pc)
}

func (c *Processor) op0808() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b A0,D4\n", pc)
}

func (c *Processor) op0809() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b A1,D4\n", pc)
}

func (c *Processor) op080A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b A2,D4\n", pc)
}

func (c *Processor) op080B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b A3,D4\n", pc)
}

func (c *Processor) op080C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b A4,D4\n", pc)
}

func (c *Processor) op080D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b A5,D4\n", pc)
}

func (c *Processor) op080E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b A6,D4\n", pc)
}

func (c *Processor) op080F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b A7,D4\n", pc)
}

func (c *Processor) op0810() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b (A0),D4\n", pc)
}

func (c *Processor) op0811() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b (A1),D4\n", pc)
}

func (c *Processor) op0812() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b (A2),D4\n", pc)
}

func (c *Processor) op0813() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b (A3),D4\n", pc)
}

func (c *Processor) op0814() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b (A4),D4\n", pc)
}

func (c *Processor) op0815() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b (A5),D4\n", pc)
}

func (c *Processor) op0816() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b (A6),D4\n", pc)
}

func (c *Processor) op0817() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.b (A7),D4\n", pc)
}

func (c *Processor) op0838() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[4] = uint32(v)
	c.tracef("%04X move.b $%X,D4\n", pc, v)
}

func (c *Processor) op0839() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[4] = uint32(v)
	c.tracef("%04X move.b $%X,D4\n", pc, v)
}

func (c *Processor) op083C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[4] = uint32(v)
	c.tracef("%04X move.b #$%X,D4\n", pc, v)
}

func (c *Processor) op0840() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b D0,A4\n", pc)
}

func (c *Processor) op0841() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b D1,A4\n", pc)
}

func (c *Processor) op0842() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b D2,A4\n", pc)
}

func (c *Processor) op0843() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b D3,A4\n", pc)
}

func (c *Processor) op0844() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b D4,A4\n", pc)
}

func (c *Processor) op0845() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b D5,A4\n", pc)
}

func (c *Processor) op0846() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b D6,A4\n", pc)
}

func (c *Processor) op0847() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b D7,A4\n", pc)
}

func (c *Processor) op0848() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b A0,A4\n", pc)
}

func (c *Processor) op0849() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b A1,A4\n", pc)
}

func (c *Processor) op084A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b A2,A4\n", pc)
}

func (c *Processor) op084B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b A3,A4\n", pc)
}

func (c *Processor) op084C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b A4,A4\n", pc)
}

func (c *Processor) op084D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b A5,A4\n", pc)
}

func (c *Processor) op084E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b A6,A4\n", pc)
}

func (c *Processor) op084F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b A7,A4\n", pc)
}

func (c *Processor) op0850() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b (A0),A4\n", pc)
}

func (c *Processor) op0851() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b (A1),A4\n", pc)
}

func (c *Processor) op0852() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b (A2),A4\n", pc)
}

func (c *Processor) op0853() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b (A3),A4\n", pc)
}

func (c *Processor) op0854() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b (A4),A4\n", pc)
}

func (c *Processor) op0855() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b (A5),A4\n", pc)
}

func (c *Processor) op0856() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b (A6),A4\n", pc)
}

func (c *Processor) op0857() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b (A7),A4\n", pc)
}

func (c *Processor) op0878() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b $%X,A4\n", pc, v)
}

func (c *Processor) op0879() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b $%X,A4\n", pc, v)
}

func (c *Processor) op087C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] = uint32(v)
	c.tracef("%04X movea.b #$%X,A4\n", pc, v)
}

func (c *Processor) op0880() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D0,(A4)\n", pc)
}

func (c *Processor) op0881() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D1,(A4)\n", pc)
}

func (c *Processor) op0882() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D2,(A4)\n", pc)
}

func (c *Processor) op0883() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D3,(A4)\n", pc)
}

func (c *Processor) op0884() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D4,(A4)\n", pc)
}

func (c *Processor) op0885() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D5,(A4)\n", pc)
}

func (c *Processor) op0886() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D6,(A4)\n", pc)
}

func (c *Processor) op0887() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D7,(A4)\n", pc)
}

func (c *Processor) op0888() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A0,(A4)\n", pc)
}

func (c *Processor) op0889() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A1,(A4)\n", pc)
}

func (c *Processor) op088A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A2,(A4)\n", pc)
}

func (c *Processor) op088B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A3,(A4)\n", pc)
}

func (c *Processor) op088C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A4,(A4)\n", pc)
}

func (c *Processor) op088D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A5,(A4)\n", pc)
}

func (c *Processor) op088E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A6,(A4)\n", pc)
}

func (c *Processor) op088F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A7,(A4)\n", pc)
}

func (c *Processor) op0890() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A0),(A4)\n", pc)
}

func (c *Processor) op0891() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A1),(A4)\n", pc)
}

func (c *Processor) op0892() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A2),(A4)\n", pc)
}

func (c *Processor) op0893() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A3),(A4)\n", pc)
}

func (c *Processor) op0894() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A4),(A4)\n", pc)
}

func (c *Processor) op0895() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A5),(A4)\n", pc)
}

func (c *Processor) op0896() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A6),(A4)\n", pc)
}

func (c *Processor) op0897() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A7),(A4)\n", pc)
}

func (c *Processor) op08B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b $%X,(A4)\n", pc, v)
}

func (c *Processor) op08B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b $%X,(A4)\n", pc, v)
}

func (c *Processor) op08BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b #$%X,(A4)\n", pc, v)
}

func (c *Processor) op08C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b D0,(A4)+\n", pc)
}

func (c *Processor) op08C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b D1,(A4)+\n", pc)
}

func (c *Processor) op08C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b D2,(A4)+\n", pc)
}

func (c *Processor) op08C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b D3,(A4)+\n", pc)
}

func (c *Processor) op08C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b D4,(A4)+\n", pc)
}

func (c *Processor) op08C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b D5,(A4)+\n", pc)
}

func (c *Processor) op08C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b D6,(A4)+\n", pc)
}

func (c *Processor) op08C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b D7,(A4)+\n", pc)
}

func (c *Processor) op08C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b A0,(A4)+\n", pc)
}

func (c *Processor) op08C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b A1,(A4)+\n", pc)
}

func (c *Processor) op08CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b A2,(A4)+\n", pc)
}

func (c *Processor) op08CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b A3,(A4)+\n", pc)
}

func (c *Processor) op08CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b A4,(A4)+\n", pc)
}

func (c *Processor) op08CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b A5,(A4)+\n", pc)
}

func (c *Processor) op08CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b A6,(A4)+\n", pc)
}

func (c *Processor) op08CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b A7,(A4)+\n", pc)
}

func (c *Processor) op08D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b (A0),(A4)+\n", pc)
}

func (c *Processor) op08D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b (A1),(A4)+\n", pc)
}

func (c *Processor) op08D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b (A2),(A4)+\n", pc)
}

func (c *Processor) op08D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b (A3),(A4)+\n", pc)
}

func (c *Processor) op08D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b (A4),(A4)+\n", pc)
}

func (c *Processor) op08D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b (A5),(A4)+\n", pc)
}

func (c *Processor) op08D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b (A6),(A4)+\n", pc)
}

func (c *Processor) op08D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b (A7),(A4)+\n", pc)
}

func (c *Processor) op08F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b $%X,(A4)+\n", pc, v)
}

func (c *Processor) op08F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b $%X,(A4)+\n", pc, v)
}

func (c *Processor) op08FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.b #$%X,(A4)+\n", pc, v)
}

func (c *Processor) op0900() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D0,-(A4)\n", pc)
}

func (c *Processor) op0901() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D1,-(A4)\n", pc)
}

func (c *Processor) op0902() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D2,-(A4)\n", pc)
}

func (c *Processor) op0903() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D3,-(A4)\n", pc)
}

func (c *Processor) op0904() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D4,-(A4)\n", pc)
}

func (c *Processor) op0905() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D5,-(A4)\n", pc)
}

func (c *Processor) op0906() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D6,-(A4)\n", pc)
}

func (c *Processor) op0907() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b D7,-(A4)\n", pc)
}

func (c *Processor) op0908() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A0,-(A4)\n", pc)
}

func (c *Processor) op0909() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A1,-(A4)\n", pc)
}

func (c *Processor) op090A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A2,-(A4)\n", pc)
}

func (c *Processor) op090B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A3,-(A4)\n", pc)
}

func (c *Processor) op090C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A4,-(A4)\n", pc)
}

func (c *Processor) op090D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A5,-(A4)\n", pc)
}

func (c *Processor) op090E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A6,-(A4)\n", pc)
}

func (c *Processor) op090F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b A7,-(A4)\n", pc)
}

func (c *Processor) op0910() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A0),-(A4)\n", pc)
}

func (c *Processor) op0911() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A1),-(A4)\n", pc)
}

func (c *Processor) op0912() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A2),-(A4)\n", pc)
}

func (c *Processor) op0913() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A3),-(A4)\n", pc)
}

func (c *Processor) op0914() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A4),-(A4)\n", pc)
}

func (c *Processor) op0915() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A5),-(A4)\n", pc)
}

func (c *Processor) op0916() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A6),-(A4)\n", pc)
}

func (c *Processor) op0917() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b (A7),-(A4)\n", pc)
}

func (c *Processor) op0938() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b $%X,-(A4)\n", pc, v)
}

func (c *Processor) op0939() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b $%X,-(A4)\n", pc, v)
}

func (c *Processor) op093C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.b #$%X,-(A4)\n", pc, v)
}

func (c *Processor) op0940() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D0,(%d,A4)\n", pc, disp)
}

func (c *Processor) op0941() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D1,(%d,A4)\n", pc, disp)
}

func (c *Processor) op0942() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D2,(%d,A4)\n", pc, disp)
}

func (c *Processor) op0943() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D3,(%d,A4)\n", pc, disp)
}

func (c *Processor) op0944() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D4,(%d,A4)\n", pc, disp)
}

func (c *Processor) op0945() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D5,(%d,A4)\n", pc, disp)
}

func (c *Processor) op0946() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D6,(%d,A4)\n", pc, disp)
}

func (c *Processor) op0947() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D7,(%d,A4)\n", pc, disp)
}

func (c *Processor) op0948() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A0,(%d,A4)\n", pc, disp)
}

func (c *Processor) op0949() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A1,(%d,A4)\n", pc, disp)
}

func (c *Processor) op094A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A2,(%d,A4)\n", pc, disp)
}

func (c *Processor) op094B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A3,(%d,A4)\n", pc, disp)
}

func (c *Processor) op094C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A4,(%d,A4)\n", pc, disp)
}

func (c *Processor) op094D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A5,(%d,A4)\n", pc, disp)
}

func (c *Processor) op094E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A6,(%d,A4)\n", pc, disp)
}

func (c *Processor) op094F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A7,(%d,A4)\n", pc, disp)
}

func (c *Processor) op0950() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A0),(%d,A4)\n", pc, disp)
}

func (c *Processor) op0951() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A1),(%d,A4)\n", pc, disp)
}

func (c *Processor) op0952() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A2),(%d,A4)\n", pc, disp)
}

func (c *Processor) op0953() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A3),(%d,A4)\n", pc, disp)
}

func (c *Processor) op0954() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A4),(%d,A4)\n", pc, disp)
}

func (c *Processor) op0955() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A5),(%d,A4)\n", pc, disp)
}

func (c *Processor) op0956() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A6),(%d,A4)\n", pc, disp)
}

func (c *Processor) op0957() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A7),(%d,A4)\n", pc, disp)
}

func (c *Processor) op0978() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A4)\n", pc, v, disp)
}

func (c *Processor) op0979() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A4)\n", pc, v, disp)
}

func (c *Processor) op097C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b #$%X,(%d,A4)\n", pc, v, disp)
}

func (c *Processor) op0A00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b D0,D5\n", pc)
}

func (c *Processor) op0A01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b D1,D5\n", pc)
}

func (c *Processor) op0A02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b D2,D5\n", pc)
}

func (c *Processor) op0A03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b D3,D5\n", pc)
}

func (c *Processor) op0A04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b D4,D5\n", pc)
}

func (c *Processor) op0A05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b D5,D5\n", pc)
}

func (c *Processor) op0A06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b D6,D5\n", pc)
}

func (c *Processor) op0A07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b D7,D5\n", pc)
}

func (c *Processor) op0A08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b A0,D5\n", pc)
}

func (c *Processor) op0A09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b A1,D5\n", pc)
}

func (c *Processor) op0A0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b A2,D5\n", pc)
}

func (c *Processor) op0A0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b A3,D5\n", pc)
}

func (c *Processor) op0A0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b A4,D5\n", pc)
}

func (c *Processor) op0A0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b A5,D5\n", pc)
}

func (c *Processor) op0A0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b A6,D5\n", pc)
}

func (c *Processor) op0A0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b A7,D5\n", pc)
}

func (c *Processor) op0A10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b (A0),D5\n", pc)
}

func (c *Processor) op0A11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b (A1),D5\n", pc)
}

func (c *Processor) op0A12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b (A2),D5\n", pc)
}

func (c *Processor) op0A13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b (A3),D5\n", pc)
}

func (c *Processor) op0A14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b (A4),D5\n", pc)
}

func (c *Processor) op0A15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b (A5),D5\n", pc)
}

func (c *Processor) op0A16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b (A6),D5\n", pc)
}

func (c *Processor) op0A17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.b (A7),D5\n", pc)
}

func (c *Processor) op0A38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[5] = uint32(v)
	c.tracef("%04X move.b $%X,D5\n", pc, v)
}

func (c *Processor) op0A39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[5] = uint32(v)
	c.tracef("%04X move.b $%X,D5\n", pc, v)
}

func (c *Processor) op0A3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[5] = uint32(v)
	c.tracef("%04X move.b #$%X,D5\n", pc, v)
}

func (c *Processor) op0A40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b D0,A5\n", pc)
}

func (c *Processor) op0A41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b D1,A5\n", pc)
}

func (c *Processor) op0A42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b D2,A5\n", pc)
}

func (c *Processor) op0A43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b D3,A5\n", pc)
}

func (c *Processor) op0A44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b D4,A5\n", pc)
}

func (c *Processor) op0A45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b D5,A5\n", pc)
}

func (c *Processor) op0A46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b D6,A5\n", pc)
}

func (c *Processor) op0A47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b D7,A5\n", pc)
}

func (c *Processor) op0A48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b A0,A5\n", pc)
}

func (c *Processor) op0A49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b A1,A5\n", pc)
}

func (c *Processor) op0A4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b A2,A5\n", pc)
}

func (c *Processor) op0A4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b A3,A5\n", pc)
}

func (c *Processor) op0A4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b A4,A5\n", pc)
}

func (c *Processor) op0A4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b A5,A5\n", pc)
}

func (c *Processor) op0A4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b A6,A5\n", pc)
}

func (c *Processor) op0A4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b A7,A5\n", pc)
}

func (c *Processor) op0A50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b (A0),A5\n", pc)
}

func (c *Processor) op0A51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b (A1),A5\n", pc)
}

func (c *Processor) op0A52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b (A2),A5\n", pc)
}

func (c *Processor) op0A53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b (A3),A5\n", pc)
}

func (c *Processor) op0A54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b (A4),A5\n", pc)
}

func (c *Processor) op0A55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b (A5),A5\n", pc)
}

func (c *Processor) op0A56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b (A6),A5\n", pc)
}

func (c *Processor) op0A57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b (A7),A5\n", pc)
}

func (c *Processor) op0A78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b $%X,A5\n", pc, v)
}

func (c *Processor) op0A79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b $%X,A5\n", pc, v)
}

func (c *Processor) op0A7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] = uint32(v)
	c.tracef("%04X movea.b #$%X,A5\n", pc, v)
}

func (c *Processor) op0A80() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D0,(A5)\n", pc)
}

func (c *Processor) op0A81() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D1,(A5)\n", pc)
}

func (c *Processor) op0A82() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D2,(A5)\n", pc)
}

func (c *Processor) op0A83() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D3,(A5)\n", pc)
}

func (c *Processor) op0A84() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D4,(A5)\n", pc)
}

func (c *Processor) op0A85() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D5,(A5)\n", pc)
}

func (c *Processor) op0A86() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D6,(A5)\n", pc)
}

func (c *Processor) op0A87() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D7,(A5)\n", pc)
}

func (c *Processor) op0A88() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A0,(A5)\n", pc)
}

func (c *Processor) op0A89() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A1,(A5)\n", pc)
}

func (c *Processor) op0A8A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A2,(A5)\n", pc)
}

func (c *Processor) op0A8B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A3,(A5)\n", pc)
}

func (c *Processor) op0A8C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A4,(A5)\n", pc)
}

func (c *Processor) op0A8D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A5,(A5)\n", pc)
}

func (c *Processor) op0A8E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A6,(A5)\n", pc)
}

func (c *Processor) op0A8F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A7,(A5)\n", pc)
}

func (c *Processor) op0A90() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A0),(A5)\n", pc)
}

func (c *Processor) op0A91() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A1),(A5)\n", pc)
}

func (c *Processor) op0A92() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A2),(A5)\n", pc)
}

func (c *Processor) op0A93() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A3),(A5)\n", pc)
}

func (c *Processor) op0A94() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A4),(A5)\n", pc)
}

func (c *Processor) op0A95() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A5),(A5)\n", pc)
}

func (c *Processor) op0A96() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A6),(A5)\n", pc)
}

func (c *Processor) op0A97() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A7),(A5)\n", pc)
}

func (c *Processor) op0AB8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b $%X,(A5)\n", pc, v)
}

func (c *Processor) op0AB9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b $%X,(A5)\n", pc, v)
}

func (c *Processor) op0ABC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b #$%X,(A5)\n", pc, v)
}

func (c *Processor) op0AC0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b D0,(A5)+\n", pc)
}

func (c *Processor) op0AC1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b D1,(A5)+\n", pc)
}

func (c *Processor) op0AC2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b D2,(A5)+\n", pc)
}

func (c *Processor) op0AC3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b D3,(A5)+\n", pc)
}

func (c *Processor) op0AC4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b D4,(A5)+\n", pc)
}

func (c *Processor) op0AC5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b D5,(A5)+\n", pc)
}

func (c *Processor) op0AC6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b D6,(A5)+\n", pc)
}

func (c *Processor) op0AC7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b D7,(A5)+\n", pc)
}

func (c *Processor) op0AC8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b A0,(A5)+\n", pc)
}

func (c *Processor) op0AC9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b A1,(A5)+\n", pc)
}

func (c *Processor) op0ACA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b A2,(A5)+\n", pc)
}

func (c *Processor) op0ACB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b A3,(A5)+\n", pc)
}

func (c *Processor) op0ACC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b A4,(A5)+\n", pc)
}

func (c *Processor) op0ACD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b A5,(A5)+\n", pc)
}

func (c *Processor) op0ACE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b A6,(A5)+\n", pc)
}

func (c *Processor) op0ACF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b A7,(A5)+\n", pc)
}

func (c *Processor) op0AD0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b (A0),(A5)+\n", pc)
}

func (c *Processor) op0AD1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b (A1),(A5)+\n", pc)
}

func (c *Processor) op0AD2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b (A2),(A5)+\n", pc)
}

func (c *Processor) op0AD3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b (A3),(A5)+\n", pc)
}

func (c *Processor) op0AD4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b (A4),(A5)+\n", pc)
}

func (c *Processor) op0AD5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b (A5),(A5)+\n", pc)
}

func (c *Processor) op0AD6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b (A6),(A5)+\n", pc)
}

func (c *Processor) op0AD7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b (A7),(A5)+\n", pc)
}

func (c *Processor) op0AF8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b $%X,(A5)+\n", pc, v)
}

func (c *Processor) op0AF9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b $%X,(A5)+\n", pc, v)
}

func (c *Processor) op0AFC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.b #$%X,(A5)+\n", pc, v)
}

func (c *Processor) op0B00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D0,-(A5)\n", pc)
}

func (c *Processor) op0B01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D1,-(A5)\n", pc)
}

func (c *Processor) op0B02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D2,-(A5)\n", pc)
}

func (c *Processor) op0B03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D3,-(A5)\n", pc)
}

func (c *Processor) op0B04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D4,-(A5)\n", pc)
}

func (c *Processor) op0B05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D5,-(A5)\n", pc)
}

func (c *Processor) op0B06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D6,-(A5)\n", pc)
}

func (c *Processor) op0B07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b D7,-(A5)\n", pc)
}

func (c *Processor) op0B08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A0,-(A5)\n", pc)
}

func (c *Processor) op0B09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A1,-(A5)\n", pc)
}

func (c *Processor) op0B0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A2,-(A5)\n", pc)
}

func (c *Processor) op0B0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A3,-(A5)\n", pc)
}

func (c *Processor) op0B0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A4,-(A5)\n", pc)
}

func (c *Processor) op0B0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A5,-(A5)\n", pc)
}

func (c *Processor) op0B0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A6,-(A5)\n", pc)
}

func (c *Processor) op0B0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b A7,-(A5)\n", pc)
}

func (c *Processor) op0B10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A0),-(A5)\n", pc)
}

func (c *Processor) op0B11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A1),-(A5)\n", pc)
}

func (c *Processor) op0B12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A2),-(A5)\n", pc)
}

func (c *Processor) op0B13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A3),-(A5)\n", pc)
}

func (c *Processor) op0B14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A4),-(A5)\n", pc)
}

func (c *Processor) op0B15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A5),-(A5)\n", pc)
}

func (c *Processor) op0B16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A6),-(A5)\n", pc)
}

func (c *Processor) op0B17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b (A7),-(A5)\n", pc)
}

func (c *Processor) op0B38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b $%X,-(A5)\n", pc, v)
}

func (c *Processor) op0B39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b $%X,-(A5)\n", pc, v)
}

func (c *Processor) op0B3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.b #$%X,-(A5)\n", pc, v)
}

func (c *Processor) op0B40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D0,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D1,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D2,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D3,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D4,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D5,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D6,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D7,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A0,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A1,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A2,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A3,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A4,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A5,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A6,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A7,(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A0),(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A1),(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A2),(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A3),(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A4),(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A5),(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A6),(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A7),(%d,A5)\n", pc, disp)
}

func (c *Processor) op0B78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A5)\n", pc, v, disp)
}

func (c *Processor) op0B79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A5)\n", pc, v, disp)
}

func (c *Processor) op0B7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b #$%X,(%d,A5)\n", pc, v, disp)
}

func (c *Processor) op0C00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b D0,D6\n", pc)
}

func (c *Processor) op0C01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b D1,D6\n", pc)
}

func (c *Processor) op0C02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b D2,D6\n", pc)
}

func (c *Processor) op0C03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b D3,D6\n", pc)
}

func (c *Processor) op0C04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b D4,D6\n", pc)
}

func (c *Processor) op0C05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b D5,D6\n", pc)
}

func (c *Processor) op0C06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b D6,D6\n", pc)
}

func (c *Processor) op0C07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b D7,D6\n", pc)
}

func (c *Processor) op0C08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b A0,D6\n", pc)
}

func (c *Processor) op0C09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b A1,D6\n", pc)
}

func (c *Processor) op0C0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b A2,D6\n", pc)
}

func (c *Processor) op0C0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b A3,D6\n", pc)
}

func (c *Processor) op0C0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b A4,D6\n", pc)
}

func (c *Processor) op0C0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b A5,D6\n", pc)
}

func (c *Processor) op0C0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b A6,D6\n", pc)
}

func (c *Processor) op0C0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b A7,D6\n", pc)
}

func (c *Processor) op0C10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b (A0),D6\n", pc)
}

func (c *Processor) op0C11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b (A1),D6\n", pc)
}

func (c *Processor) op0C12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b (A2),D6\n", pc)
}

func (c *Processor) op0C13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b (A3),D6\n", pc)
}

func (c *Processor) op0C14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b (A4),D6\n", pc)
}

func (c *Processor) op0C15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b (A5),D6\n", pc)
}

func (c *Processor) op0C16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b (A6),D6\n", pc)
}

func (c *Processor) op0C17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.b (A7),D6\n", pc)
}

func (c *Processor) op0C38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[6] = uint32(v)
	c.tracef("%04X move.b $%X,D6\n", pc, v)
}

func (c *Processor) op0C39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[6] = uint32(v)
	c.tracef("%04X move.b $%X,D6\n", pc, v)
}

func (c *Processor) op0C3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[6] = uint32(v)
	c.tracef("%04X move.b #$%X,D6\n", pc, v)
}

func (c *Processor) op0C40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b D0,A6\n", pc)
}

func (c *Processor) op0C41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b D1,A6\n", pc)
}

func (c *Processor) op0C42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b D2,A6\n", pc)
}

func (c *Processor) op0C43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b D3,A6\n", pc)
}

func (c *Processor) op0C44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b D4,A6\n", pc)
}

func (c *Processor) op0C45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b D5,A6\n", pc)
}

func (c *Processor) op0C46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b D6,A6\n", pc)
}

func (c *Processor) op0C47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b D7,A6\n", pc)
}

func (c *Processor) op0C48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b A0,A6\n", pc)
}

func (c *Processor) op0C49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b A1,A6\n", pc)
}

func (c *Processor) op0C4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b A2,A6\n", pc)
}

func (c *Processor) op0C4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b A3,A6\n", pc)
}

func (c *Processor) op0C4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b A4,A6\n", pc)
}

func (c *Processor) op0C4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b A5,A6\n", pc)
}

func (c *Processor) op0C4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b A6,A6\n", pc)
}

func (c *Processor) op0C4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b A7,A6\n", pc)
}

func (c *Processor) op0C50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b (A0),A6\n", pc)
}

func (c *Processor) op0C51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b (A1),A6\n", pc)
}

func (c *Processor) op0C52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b (A2),A6\n", pc)
}

func (c *Processor) op0C53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b (A3),A6\n", pc)
}

func (c *Processor) op0C54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b (A4),A6\n", pc)
}

func (c *Processor) op0C55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b (A5),A6\n", pc)
}

func (c *Processor) op0C56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b (A6),A6\n", pc)
}

func (c *Processor) op0C57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b (A7),A6\n", pc)
}

func (c *Processor) op0C78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b $%X,A6\n", pc, v)
}

func (c *Processor) op0C79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b $%X,A6\n", pc, v)
}

func (c *Processor) op0C7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] = uint32(v)
	c.tracef("%04X movea.b #$%X,A6\n", pc, v)
}

func (c *Processor) op0C80() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D0,(A6)\n", pc)
}

func (c *Processor) op0C81() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D1,(A6)\n", pc)
}

func (c *Processor) op0C82() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D2,(A6)\n", pc)
}

func (c *Processor) op0C83() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D3,(A6)\n", pc)
}

func (c *Processor) op0C84() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D4,(A6)\n", pc)
}

func (c *Processor) op0C85() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D5,(A6)\n", pc)
}

func (c *Processor) op0C86() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D6,(A6)\n", pc)
}

func (c *Processor) op0C87() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D7,(A6)\n", pc)
}

func (c *Processor) op0C88() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A0,(A6)\n", pc)
}

func (c *Processor) op0C89() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A1,(A6)\n", pc)
}

func (c *Processor) op0C8A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A2,(A6)\n", pc)
}

func (c *Processor) op0C8B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A3,(A6)\n", pc)
}

func (c *Processor) op0C8C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A4,(A6)\n", pc)
}

func (c *Processor) op0C8D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A5,(A6)\n", pc)
}

func (c *Processor) op0C8E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A6,(A6)\n", pc)
}

func (c *Processor) op0C8F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A7,(A6)\n", pc)
}

func (c *Processor) op0C90() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A0),(A6)\n", pc)
}

func (c *Processor) op0C91() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A1),(A6)\n", pc)
}

func (c *Processor) op0C92() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A2),(A6)\n", pc)
}

func (c *Processor) op0C93() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A3),(A6)\n", pc)
}

func (c *Processor) op0C94() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A4),(A6)\n", pc)
}

func (c *Processor) op0C95() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A5),(A6)\n", pc)
}

func (c *Processor) op0C96() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A6),(A6)\n", pc)
}

func (c *Processor) op0C97() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A7),(A6)\n", pc)
}

func (c *Processor) op0CB8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b $%X,(A6)\n", pc, v)
}

func (c *Processor) op0CB9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b $%X,(A6)\n", pc, v)
}

func (c *Processor) op0CBC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b #$%X,(A6)\n", pc, v)
}

func (c *Processor) op0CC0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b D0,(A6)+\n", pc)
}

func (c *Processor) op0CC1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b D1,(A6)+\n", pc)
}

func (c *Processor) op0CC2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b D2,(A6)+\n", pc)
}

func (c *Processor) op0CC3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b D3,(A6)+\n", pc)
}

func (c *Processor) op0CC4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b D4,(A6)+\n", pc)
}

func (c *Processor) op0CC5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b D5,(A6)+\n", pc)
}

func (c *Processor) op0CC6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b D6,(A6)+\n", pc)
}

func (c *Processor) op0CC7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b D7,(A6)+\n", pc)
}

func (c *Processor) op0CC8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b A0,(A6)+\n", pc)
}

func (c *Processor) op0CC9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b A1,(A6)+\n", pc)
}

func (c *Processor) op0CCA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b A2,(A6)+\n", pc)
}

func (c *Processor) op0CCB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b A3,(A6)+\n", pc)
}

func (c *Processor) op0CCC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b A4,(A6)+\n", pc)
}

func (c *Processor) op0CCD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b A5,(A6)+\n", pc)
}

func (c *Processor) op0CCE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b A6,(A6)+\n", pc)
}

func (c *Processor) op0CCF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b A7,(A6)+\n", pc)
}

func (c *Processor) op0CD0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b (A0),(A6)+\n", pc)
}

func (c *Processor) op0CD1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b (A1),(A6)+\n", pc)
}

func (c *Processor) op0CD2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b (A2),(A6)+\n", pc)
}

func (c *Processor) op0CD3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b (A3),(A6)+\n", pc)
}

func (c *Processor) op0CD4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b (A4),(A6)+\n", pc)
}

func (c *Processor) op0CD5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b (A5),(A6)+\n", pc)
}

func (c *Processor) op0CD6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b (A6),(A6)+\n", pc)
}

func (c *Processor) op0CD7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b (A7),(A6)+\n", pc)
}

func (c *Processor) op0CF8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b $%X,(A6)+\n", pc, v)
}

func (c *Processor) op0CF9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b $%X,(A6)+\n", pc, v)
}

func (c *Processor) op0CFC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.b #$%X,(A6)+\n", pc, v)
}

func (c *Processor) op0D00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D0,-(A6)\n", pc)
}

func (c *Processor) op0D01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D1,-(A6)\n", pc)
}

func (c *Processor) op0D02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D2,-(A6)\n", pc)
}

func (c *Processor) op0D03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D3,-(A6)\n", pc)
}

func (c *Processor) op0D04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D4,-(A6)\n", pc)
}

func (c *Processor) op0D05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D5,-(A6)\n", pc)
}

func (c *Processor) op0D06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D6,-(A6)\n", pc)
}

func (c *Processor) op0D07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b D7,-(A6)\n", pc)
}

func (c *Processor) op0D08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A0,-(A6)\n", pc)
}

func (c *Processor) op0D09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A1,-(A6)\n", pc)
}

func (c *Processor) op0D0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A2,-(A6)\n", pc)
}

func (c *Processor) op0D0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A3,-(A6)\n", pc)
}

func (c *Processor) op0D0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A4,-(A6)\n", pc)
}

func (c *Processor) op0D0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A5,-(A6)\n", pc)
}

func (c *Processor) op0D0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A6,-(A6)\n", pc)
}

func (c *Processor) op0D0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b A7,-(A6)\n", pc)
}

func (c *Processor) op0D10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A0),-(A6)\n", pc)
}

func (c *Processor) op0D11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A1),-(A6)\n", pc)
}

func (c *Processor) op0D12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A2),-(A6)\n", pc)
}

func (c *Processor) op0D13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A3),-(A6)\n", pc)
}

func (c *Processor) op0D14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A4),-(A6)\n", pc)
}

func (c *Processor) op0D15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A5),-(A6)\n", pc)
}

func (c *Processor) op0D16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A6),-(A6)\n", pc)
}

func (c *Processor) op0D17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b (A7),-(A6)\n", pc)
}

func (c *Processor) op0D38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b $%X,-(A6)\n", pc, v)
}

func (c *Processor) op0D39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b $%X,-(A6)\n", pc, v)
}

func (c *Processor) op0D3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.b #$%X,-(A6)\n", pc, v)
}

func (c *Processor) op0D40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D0,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D1,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D2,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D3,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D4,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D5,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D6,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D7,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A0,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A1,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A2,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A3,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A4,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A5,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A6,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A7,(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A0),(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A1),(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A2),(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A3),(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A4),(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A5),(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A6),(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A7),(%d,A6)\n", pc, disp)
}

func (c *Processor) op0D78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A6)\n", pc, v, disp)
}

func (c *Processor) op0D79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A6)\n", pc, v, disp)
}

func (c *Processor) op0D7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b #$%X,(%d,A6)\n", pc, v, disp)
}

func (c *Processor) op0E00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b D0,D7\n", pc)
}

func (c *Processor) op0E01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b D1,D7\n", pc)
}

func (c *Processor) op0E02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b D2,D7\n", pc)
}

func (c *Processor) op0E03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b D3,D7\n", pc)
}

func (c *Processor) op0E04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b D4,D7\n", pc)
}

func (c *Processor) op0E05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b D5,D7\n", pc)
}

func (c *Processor) op0E06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b D6,D7\n", pc)
}

func (c *Processor) op0E07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b D7,D7\n", pc)
}

func (c *Processor) op0E08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b A0,D7\n", pc)
}

func (c *Processor) op0E09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b A1,D7\n", pc)
}

func (c *Processor) op0E0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b A2,D7\n", pc)
}

func (c *Processor) op0E0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b A3,D7\n", pc)
}

func (c *Processor) op0E0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b A4,D7\n", pc)
}

func (c *Processor) op0E0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b A5,D7\n", pc)
}

func (c *Processor) op0E0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b A6,D7\n", pc)
}

func (c *Processor) op0E0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b A7,D7\n", pc)
}

func (c *Processor) op0E10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b (A0),D7\n", pc)
}

func (c *Processor) op0E11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b (A1),D7\n", pc)
}

func (c *Processor) op0E12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b (A2),D7\n", pc)
}

func (c *Processor) op0E13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b (A3),D7\n", pc)
}

func (c *Processor) op0E14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b (A4),D7\n", pc)
}

func (c *Processor) op0E15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b (A5),D7\n", pc)
}

func (c *Processor) op0E16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b (A6),D7\n", pc)
}

func (c *Processor) op0E17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.b (A7),D7\n", pc)
}

func (c *Processor) op0E38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[7] = uint32(v)
	c.tracef("%04X move.b $%X,D7\n", pc, v)
}

func (c *Processor) op0E39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[7] = uint32(v)
	c.tracef("%04X move.b $%X,D7\n", pc, v)
}

func (c *Processor) op0E3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[7] = uint32(v)
	c.tracef("%04X move.b #$%X,D7\n", pc, v)
}

func (c *Processor) op0E40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b D0,A7\n", pc)
}

func (c *Processor) op0E41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b D1,A7\n", pc)
}

func (c *Processor) op0E42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b D2,A7\n", pc)
}

func (c *Processor) op0E43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b D3,A7\n", pc)
}

func (c *Processor) op0E44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b D4,A7\n", pc)
}

func (c *Processor) op0E45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b D5,A7\n", pc)
}

func (c *Processor) op0E46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b D6,A7\n", pc)
}

func (c *Processor) op0E47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b D7,A7\n", pc)
}

func (c *Processor) op0E48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b A0,A7\n", pc)
}

func (c *Processor) op0E49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b A1,A7\n", pc)
}

func (c *Processor) op0E4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b A2,A7\n", pc)
}

func (c *Processor) op0E4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b A3,A7\n", pc)
}

func (c *Processor) op0E4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b A4,A7\n", pc)
}

func (c *Processor) op0E4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b A5,A7\n", pc)
}

func (c *Processor) op0E4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b A6,A7\n", pc)
}

func (c *Processor) op0E4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b A7,A7\n", pc)
}

func (c *Processor) op0E50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b (A0),A7\n", pc)
}

func (c *Processor) op0E51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b (A1),A7\n", pc)
}

func (c *Processor) op0E52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b (A2),A7\n", pc)
}

func (c *Processor) op0E53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b (A3),A7\n", pc)
}

func (c *Processor) op0E54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b (A4),A7\n", pc)
}

func (c *Processor) op0E55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b (A5),A7\n", pc)
}

func (c *Processor) op0E56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b (A6),A7\n", pc)
}

func (c *Processor) op0E57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b (A7),A7\n", pc)
}

func (c *Processor) op0E78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b $%X,A7\n", pc, v)
}

func (c *Processor) op0E79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b $%X,A7\n", pc, v)
}

func (c *Processor) op0E7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] = uint32(v)
	c.tracef("%04X movea.b #$%X,A7\n", pc, v)
}

func (c *Processor) op0E80() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D0,(A7)\n", pc)
}

func (c *Processor) op0E81() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D1,(A7)\n", pc)
}

func (c *Processor) op0E82() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D2,(A7)\n", pc)
}

func (c *Processor) op0E83() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D3,(A7)\n", pc)
}

func (c *Processor) op0E84() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D4,(A7)\n", pc)
}

func (c *Processor) op0E85() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D5,(A7)\n", pc)
}

func (c *Processor) op0E86() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D6,(A7)\n", pc)
}

func (c *Processor) op0E87() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D7,(A7)\n", pc)
}

func (c *Processor) op0E88() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A0,(A7)\n", pc)
}

func (c *Processor) op0E89() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A1,(A7)\n", pc)
}

func (c *Processor) op0E8A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A2,(A7)\n", pc)
}

func (c *Processor) op0E8B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A3,(A7)\n", pc)
}

func (c *Processor) op0E8C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A4,(A7)\n", pc)
}

func (c *Processor) op0E8D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A5,(A7)\n", pc)
}

func (c *Processor) op0E8E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A6,(A7)\n", pc)
}

func (c *Processor) op0E8F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A7,(A7)\n", pc)
}

func (c *Processor) op0E90() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A0),(A7)\n", pc)
}

func (c *Processor) op0E91() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A1),(A7)\n", pc)
}

func (c *Processor) op0E92() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A2),(A7)\n", pc)
}

func (c *Processor) op0E93() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A3),(A7)\n", pc)
}

func (c *Processor) op0E94() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A4),(A7)\n", pc)
}

func (c *Processor) op0E95() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A5),(A7)\n", pc)
}

func (c *Processor) op0E96() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A6),(A7)\n", pc)
}

func (c *Processor) op0E97() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A7),(A7)\n", pc)
}

func (c *Processor) op0EB8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b $%X,(A7)\n", pc, v)
}

func (c *Processor) op0EB9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b $%X,(A7)\n", pc, v)
}

func (c *Processor) op0EBC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b #$%X,(A7)\n", pc, v)
}

func (c *Processor) op0EC0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b D0,(A7)+\n", pc)
}

func (c *Processor) op0EC1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b D1,(A7)+\n", pc)
}

func (c *Processor) op0EC2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b D2,(A7)+\n", pc)
}

func (c *Processor) op0EC3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b D3,(A7)+\n", pc)
}

func (c *Processor) op0EC4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b D4,(A7)+\n", pc)
}

func (c *Processor) op0EC5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b D5,(A7)+\n", pc)
}

func (c *Processor) op0EC6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b D6,(A7)+\n", pc)
}

func (c *Processor) op0EC7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b D7,(A7)+\n", pc)
}

func (c *Processor) op0EC8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b A0,(A7)+\n", pc)
}

func (c *Processor) op0EC9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b A1,(A7)+\n", pc)
}

func (c *Processor) op0ECA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b A2,(A7)+\n", pc)
}

func (c *Processor) op0ECB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b A3,(A7)+\n", pc)
}

func (c *Processor) op0ECC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b A4,(A7)+\n", pc)
}

func (c *Processor) op0ECD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b A5,(A7)+\n", pc)
}

func (c *Processor) op0ECE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b A6,(A7)+\n", pc)
}

func (c *Processor) op0ECF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b A7,(A7)+\n", pc)
}

func (c *Processor) op0ED0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b (A0),(A7)+\n", pc)
}

func (c *Processor) op0ED1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b (A1),(A7)+\n", pc)
}

func (c *Processor) op0ED2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b (A2),(A7)+\n", pc)
}

func (c *Processor) op0ED3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b (A3),(A7)+\n", pc)
}

func (c *Processor) op0ED4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b (A4),(A7)+\n", pc)
}

func (c *Processor) op0ED5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b (A5),(A7)+\n", pc)
}

func (c *Processor) op0ED6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b (A6),(A7)+\n", pc)
}

func (c *Processor) op0ED7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b (A7),(A7)+\n", pc)
}

func (c *Processor) op0EF8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b $%X,(A7)+\n", pc, v)
}

func (c *Processor) op0EF9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b $%X,(A7)+\n", pc, v)
}

func (c *Processor) op0EFC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.b #$%X,(A7)+\n", pc, v)
}

func (c *Processor) op0F00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D0,-(A7)\n", pc)
}

func (c *Processor) op0F01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D1,-(A7)\n", pc)
}

func (c *Processor) op0F02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D2,-(A7)\n", pc)
}

func (c *Processor) op0F03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D3,-(A7)\n", pc)
}

func (c *Processor) op0F04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D4,-(A7)\n", pc)
}

func (c *Processor) op0F05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D5,-(A7)\n", pc)
}

func (c *Processor) op0F06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D6,-(A7)\n", pc)
}

func (c *Processor) op0F07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b D7,-(A7)\n", pc)
}

func (c *Processor) op0F08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A0,-(A7)\n", pc)
}

func (c *Processor) op0F09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A1,-(A7)\n", pc)
}

func (c *Processor) op0F0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A2,-(A7)\n", pc)
}

func (c *Processor) op0F0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A3,-(A7)\n", pc)
}

func (c *Processor) op0F0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A4,-(A7)\n", pc)
}

func (c *Processor) op0F0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A5,-(A7)\n", pc)
}

func (c *Processor) op0F0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A6,-(A7)\n", pc)
}

func (c *Processor) op0F0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b A7,-(A7)\n", pc)
}

func (c *Processor) op0F10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A0),-(A7)\n", pc)
}

func (c *Processor) op0F11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A1),-(A7)\n", pc)
}

func (c *Processor) op0F12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A2),-(A7)\n", pc)
}

func (c *Processor) op0F13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A3),-(A7)\n", pc)
}

func (c *Processor) op0F14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A4),-(A7)\n", pc)
}

func (c *Processor) op0F15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A5),-(A7)\n", pc)
}

func (c *Processor) op0F16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A6),-(A7)\n", pc)
}

func (c *Processor) op0F17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b (A7),-(A7)\n", pc)
}

func (c *Processor) op0F38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b $%X,-(A7)\n", pc, v)
}

func (c *Processor) op0F39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b $%X,-(A7)\n", pc, v)
}

func (c *Processor) op0F3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.b #$%X,-(A7)\n", pc, v)
}

func (c *Processor) op0F40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D0,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D1,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D2,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D3,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D4,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D5,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D6,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b D7,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A0,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A1,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A2,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A3,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A4,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A5,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A6,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b A7,(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A0),(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A1),(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A2),(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A3),(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A4),(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A5),(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A6),(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:1])
	if c.err != nil {
		return
	}
	v := c.buf[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b (A7),(%d,A7)\n", pc, disp)
}

func (c *Processor) op0F78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A7)\n", pc, v, disp)
}

func (c *Processor) op0F79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b $%X,(%d,A7)\n", pc, v, disp)
}

func (c *Processor) op0F7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.b #$%X,(%d,A7)\n", pc, v, disp)
}

func (c *Processor) op1000() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w D0,D0\n", pc)
}

func (c *Processor) op1001() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w D1,D0\n", pc)
}

func (c *Processor) op1002() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w D2,D0\n", pc)
}

func (c *Processor) op1003() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w D3,D0\n", pc)
}

func (c *Processor) op1004() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w D4,D0\n", pc)
}

func (c *Processor) op1005() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w D5,D0\n", pc)
}

func (c *Processor) op1006() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w D6,D0\n", pc)
}

func (c *Processor) op1007() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w D7,D0\n", pc)
}

func (c *Processor) op1008() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w A0,D0\n", pc)
}

func (c *Processor) op1009() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w A1,D0\n", pc)
}

func (c *Processor) op100A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w A2,D0\n", pc)
}

func (c *Processor) op100B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w A3,D0\n", pc)
}

func (c *Processor) op100C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w A4,D0\n", pc)
}

func (c *Processor) op100D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w A5,D0\n", pc)
}

func (c *Processor) op100E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w A6,D0\n", pc)
}

func (c *Processor) op100F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[0] = uint32(v)
	c.tracef("%04X move.w A7,D0\n", pc)
}

func (c *Processor) op1010() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[0] = uint32(v)
	c.tracef("%04X move.w (A0),D0\n", pc)
}

func (c *Processor) op1011() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[0] = uint32(v)
	c.tracef("%04X move.w (A1),D0\n", pc)
}

func (c *Processor) op1012() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[0] = uint32(v)
	c.tracef("%04X move.w (A2),D0\n", pc)
}

func (c *Processor) op1013() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[0] = uint32(v)
	c.tracef("%04X move.w (A3),D0\n", pc)
}

func (c *Processor) op1014() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[0] = uint32(v)
	c.tracef("%04X move.w (A4),D0\n", pc)
}

func (c *Processor) op1015() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[0] = uint32(v)
	c.tracef("%04X move.w (A5),D0\n", pc)
}

func (c *Processor) op1016() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[0] = uint32(v)
	c.tracef("%04X move.w (A6),D0\n", pc)
}

func (c *Processor) op1017() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[0] = uint32(v)
	c.tracef("%04X move.w (A7),D0\n", pc)
}

func (c *Processor) op1038() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[0] = uint32(v)
	c.tracef("%04X move.w $%X,D0\n", pc, v)
}

func (c *Processor) op1039() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[0] = uint32(v)
	c.tracef("%04X move.w $%X,D0\n", pc, v)
}

func (c *Processor) op103C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[0] = uint32(v)
	c.tracef("%04X move.w #$%X,D0\n", pc, v)
}

func (c *Processor) op1040() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w D0,A0\n", pc)
}

func (c *Processor) op1041() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w D1,A0\n", pc)
}

func (c *Processor) op1042() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w D2,A0\n", pc)
}

func (c *Processor) op1043() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w D3,A0\n", pc)
}

func (c *Processor) op1044() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w D4,A0\n", pc)
}

func (c *Processor) op1045() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w D5,A0\n", pc)
}

func (c *Processor) op1046() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w D6,A0\n", pc)
}

func (c *Processor) op1047() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w D7,A0\n", pc)
}

func (c *Processor) op1048() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w A0,A0\n", pc)
}

func (c *Processor) op1049() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w A1,A0\n", pc)
}

func (c *Processor) op104A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w A2,A0\n", pc)
}

func (c *Processor) op104B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w A3,A0\n", pc)
}

func (c *Processor) op104C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w A4,A0\n", pc)
}

func (c *Processor) op104D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w A5,A0\n", pc)
}

func (c *Processor) op104E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w A6,A0\n", pc)
}

func (c *Processor) op104F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w A7,A0\n", pc)
}

func (c *Processor) op1050() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w (A0),A0\n", pc)
}

func (c *Processor) op1051() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w (A1),A0\n", pc)
}

func (c *Processor) op1052() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w (A2),A0\n", pc)
}

func (c *Processor) op1053() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w (A3),A0\n", pc)
}

func (c *Processor) op1054() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w (A4),A0\n", pc)
}

func (c *Processor) op1055() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w (A5),A0\n", pc)
}

func (c *Processor) op1056() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w (A6),A0\n", pc)
}

func (c *Processor) op1057() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w (A7),A0\n", pc)
}

func (c *Processor) op1078() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w $%X,A0\n", pc, v)
}

func (c *Processor) op1079() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w $%X,A0\n", pc, v)
}

func (c *Processor) op107C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] = uint32(v)
	c.tracef("%04X movea.w #$%X,A0\n", pc, v)
}

func (c *Processor) op1080() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D0,(A0)\n", pc)
}

func (c *Processor) op1081() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D1,(A0)\n", pc)
}

func (c *Processor) op1082() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D2,(A0)\n", pc)
}

func (c *Processor) op1083() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D3,(A0)\n", pc)
}

func (c *Processor) op1084() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D4,(A0)\n", pc)
}

func (c *Processor) op1085() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D5,(A0)\n", pc)
}

func (c *Processor) op1086() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D6,(A0)\n", pc)
}

func (c *Processor) op1087() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D7,(A0)\n", pc)
}

func (c *Processor) op1088() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A0,(A0)\n", pc)
}

func (c *Processor) op1089() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A1,(A0)\n", pc)
}

func (c *Processor) op108A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A2,(A0)\n", pc)
}

func (c *Processor) op108B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A3,(A0)\n", pc)
}

func (c *Processor) op108C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A4,(A0)\n", pc)
}

func (c *Processor) op108D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A5,(A0)\n", pc)
}

func (c *Processor) op108E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A6,(A0)\n", pc)
}

func (c *Processor) op108F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A7,(A0)\n", pc)
}

func (c *Processor) op1090() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A0),(A0)\n", pc)
}

func (c *Processor) op1091() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A1),(A0)\n", pc)
}

func (c *Processor) op1092() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A2),(A0)\n", pc)
}

func (c *Processor) op1093() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A3),(A0)\n", pc)
}

func (c *Processor) op1094() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A4),(A0)\n", pc)
}

func (c *Processor) op1095() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A5),(A0)\n", pc)
}

func (c *Processor) op1096() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A6),(A0)\n", pc)
}

func (c *Processor) op1097() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A7),(A0)\n", pc)
}

func (c *Processor) op10B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w $%X,(A0)\n", pc, v)
}

func (c *Processor) op10B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w $%X,(A0)\n", pc, v)
}

func (c *Processor) op10BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w #$%X,(A0)\n", pc, v)
}

func (c *Processor) op10C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w D0,(A0)+\n", pc)
}

func (c *Processor) op10C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w D1,(A0)+\n", pc)
}

func (c *Processor) op10C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w D2,(A0)+\n", pc)
}

func (c *Processor) op10C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w D3,(A0)+\n", pc)
}

func (c *Processor) op10C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w D4,(A0)+\n", pc)
}

func (c *Processor) op10C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w D5,(A0)+\n", pc)
}

func (c *Processor) op10C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w D6,(A0)+\n", pc)
}

func (c *Processor) op10C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w D7,(A0)+\n", pc)
}

func (c *Processor) op10C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w A0,(A0)+\n", pc)
}

func (c *Processor) op10C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w A1,(A0)+\n", pc)
}

func (c *Processor) op10CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w A2,(A0)+\n", pc)
}

func (c *Processor) op10CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w A3,(A0)+\n", pc)
}

func (c *Processor) op10CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w A4,(A0)+\n", pc)
}

func (c *Processor) op10CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w A5,(A0)+\n", pc)
}

func (c *Processor) op10CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w A6,(A0)+\n", pc)
}

func (c *Processor) op10CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w A7,(A0)+\n", pc)
}

func (c *Processor) op10D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w (A0),(A0)+\n", pc)
}

func (c *Processor) op10D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w (A1),(A0)+\n", pc)
}

func (c *Processor) op10D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w (A2),(A0)+\n", pc)
}

func (c *Processor) op10D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w (A3),(A0)+\n", pc)
}

func (c *Processor) op10D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w (A4),(A0)+\n", pc)
}

func (c *Processor) op10D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w (A5),(A0)+\n", pc)
}

func (c *Processor) op10D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w (A6),(A0)+\n", pc)
}

func (c *Processor) op10D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w (A7),(A0)+\n", pc)
}

func (c *Processor) op10F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w $%X,(A0)+\n", pc, v)
}

func (c *Processor) op10F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w $%X,(A0)+\n", pc, v)
}

func (c *Processor) op10FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.w #$%X,(A0)+\n", pc, v)
}

func (c *Processor) op1100() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D0,-(A0)\n", pc)
}

func (c *Processor) op1101() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D1,-(A0)\n", pc)
}

func (c *Processor) op1102() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D2,-(A0)\n", pc)
}

func (c *Processor) op1103() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D3,-(A0)\n", pc)
}

func (c *Processor) op1104() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D4,-(A0)\n", pc)
}

func (c *Processor) op1105() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D5,-(A0)\n", pc)
}

func (c *Processor) op1106() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D6,-(A0)\n", pc)
}

func (c *Processor) op1107() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w D7,-(A0)\n", pc)
}

func (c *Processor) op1108() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A0,-(A0)\n", pc)
}

func (c *Processor) op1109() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A1,-(A0)\n", pc)
}

func (c *Processor) op110A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A2,-(A0)\n", pc)
}

func (c *Processor) op110B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A3,-(A0)\n", pc)
}

func (c *Processor) op110C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A4,-(A0)\n", pc)
}

func (c *Processor) op110D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A5,-(A0)\n", pc)
}

func (c *Processor) op110E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A6,-(A0)\n", pc)
}

func (c *Processor) op110F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w A7,-(A0)\n", pc)
}

func (c *Processor) op1110() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A0),-(A0)\n", pc)
}

func (c *Processor) op1111() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A1),-(A0)\n", pc)
}

func (c *Processor) op1112() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A2),-(A0)\n", pc)
}

func (c *Processor) op1113() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A3),-(A0)\n", pc)
}

func (c *Processor) op1114() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A4),-(A0)\n", pc)
}

func (c *Processor) op1115() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A5),-(A0)\n", pc)
}

func (c *Processor) op1116() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A6),-(A0)\n", pc)
}

func (c *Processor) op1117() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w (A7),-(A0)\n", pc)
}

func (c *Processor) op1138() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w $%X,-(A0)\n", pc, v)
}

func (c *Processor) op1139() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w $%X,-(A0)\n", pc, v)
}

func (c *Processor) op113C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.w #$%X,-(A0)\n", pc, v)
}

func (c *Processor) op1140() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D0,(%d,A0)\n", pc, disp)
}

func (c *Processor) op1141() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D1,(%d,A0)\n", pc, disp)
}

func (c *Processor) op1142() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D2,(%d,A0)\n", pc, disp)
}

func (c *Processor) op1143() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D3,(%d,A0)\n", pc, disp)
}

func (c *Processor) op1144() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D4,(%d,A0)\n", pc, disp)
}

func (c *Processor) op1145() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D5,(%d,A0)\n", pc, disp)
}

func (c *Processor) op1146() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D6,(%d,A0)\n", pc, disp)
}

func (c *Processor) op1147() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D7,(%d,A0)\n", pc, disp)
}

func (c *Processor) op1148() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A0,(%d,A0)\n", pc, disp)
}

func (c *Processor) op1149() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A1,(%d,A0)\n", pc, disp)
}

func (c *Processor) op114A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A2,(%d,A0)\n", pc, disp)
}

func (c *Processor) op114B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A3,(%d,A0)\n", pc, disp)
}

func (c *Processor) op114C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A4,(%d,A0)\n", pc, disp)
}

func (c *Processor) op114D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A5,(%d,A0)\n", pc, disp)
}

func (c *Processor) op114E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A6,(%d,A0)\n", pc, disp)
}

func (c *Processor) op114F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A7,(%d,A0)\n", pc, disp)
}

func (c *Processor) op1150() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A0),(%d,A0)\n", pc, disp)
}

func (c *Processor) op1151() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A1),(%d,A0)\n", pc, disp)
}

func (c *Processor) op1152() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A2),(%d,A0)\n", pc, disp)
}

func (c *Processor) op1153() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A3),(%d,A0)\n", pc, disp)
}

func (c *Processor) op1154() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A4),(%d,A0)\n", pc, disp)
}

func (c *Processor) op1155() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A5),(%d,A0)\n", pc, disp)
}

func (c *Processor) op1156() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A6),(%d,A0)\n", pc, disp)
}

func (c *Processor) op1157() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A7),(%d,A0)\n", pc, disp)
}

func (c *Processor) op1178() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A0)\n", pc, v, disp)
}

func (c *Processor) op1179() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A0)\n", pc, v, disp)
}

func (c *Processor) op117C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w #$%X,(%d,A0)\n", pc, v, disp)
}

func (c *Processor) op11C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D0,$%X\n", pc, addr)
}

func (c *Processor) op11C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D1,$%X\n", pc, addr)
}

func (c *Processor) op11C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D2,$%X\n", pc, addr)
}

func (c *Processor) op11C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D3,$%X\n", pc, addr)
}

func (c *Processor) op11C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D4,$%X\n", pc, addr)
}

func (c *Processor) op11C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D5,$%X\n", pc, addr)
}

func (c *Processor) op11C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D6,$%X\n", pc, addr)
}

func (c *Processor) op11C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D7,$%X\n", pc, addr)
}

func (c *Processor) op11C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A0,$%X\n", pc, addr)
}

func (c *Processor) op11C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A1,$%X\n", pc, addr)
}

func (c *Processor) op11CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A2,$%X\n", pc, addr)
}

func (c *Processor) op11CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A3,$%X\n", pc, addr)
}

func (c *Processor) op11CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A4,$%X\n", pc, addr)
}

func (c *Processor) op11CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A5,$%X\n", pc, addr)
}

func (c *Processor) op11CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A6,$%X\n", pc, addr)
}

func (c *Processor) op11CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A7,$%X\n", pc, addr)
}

func (c *Processor) op11D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A0),$%X\n", pc, addr)
}

func (c *Processor) op11D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A1),$%X\n", pc, addr)
}

func (c *Processor) op11D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A2),$%X\n", pc, addr)
}

func (c *Processor) op11D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A3),$%X\n", pc, addr)
}

func (c *Processor) op11D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A4),$%X\n", pc, addr)
}

func (c *Processor) op11D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A5),$%X\n", pc, addr)
}

func (c *Processor) op11D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A6),$%X\n", pc, addr)
}

func (c *Processor) op11D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A7),$%X\n", pc, addr)
}

func (c *Processor) op11F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,$%X\n", pc, v, addr)
}

func (c *Processor) op11F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,$%X\n", pc, v, addr)
}

func (c *Processor) op11FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w #$%X,$%X\n", pc, v, addr)
}

func (c *Processor) op1200() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w D0,D1\n", pc)
}

func (c *Processor) op1201() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w D1,D1\n", pc)
}

func (c *Processor) op1202() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w D2,D1\n", pc)
}

func (c *Processor) op1203() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w D3,D1\n", pc)
}

func (c *Processor) op1204() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w D4,D1\n", pc)
}

func (c *Processor) op1205() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w D5,D1\n", pc)
}

func (c *Processor) op1206() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w D6,D1\n", pc)
}

func (c *Processor) op1207() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w D7,D1\n", pc)
}

func (c *Processor) op1208() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w A0,D1\n", pc)
}

func (c *Processor) op1209() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w A1,D1\n", pc)
}

func (c *Processor) op120A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w A2,D1\n", pc)
}

func (c *Processor) op120B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w A3,D1\n", pc)
}

func (c *Processor) op120C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w A4,D1\n", pc)
}

func (c *Processor) op120D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w A5,D1\n", pc)
}

func (c *Processor) op120E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w A6,D1\n", pc)
}

func (c *Processor) op120F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[1] = uint32(v)
	c.tracef("%04X move.w A7,D1\n", pc)
}

func (c *Processor) op1210() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[1] = uint32(v)
	c.tracef("%04X move.w (A0),D1\n", pc)
}

func (c *Processor) op1211() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[1] = uint32(v)
	c.tracef("%04X move.w (A1),D1\n", pc)
}

func (c *Processor) op1212() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[1] = uint32(v)
	c.tracef("%04X move.w (A2),D1\n", pc)
}

func (c *Processor) op1213() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[1] = uint32(v)
	c.tracef("%04X move.w (A3),D1\n", pc)
}

func (c *Processor) op1214() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[1] = uint32(v)
	c.tracef("%04X move.w (A4),D1\n", pc)
}

func (c *Processor) op1215() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[1] = uint32(v)
	c.tracef("%04X move.w (A5),D1\n", pc)
}

func (c *Processor) op1216() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[1] = uint32(v)
	c.tracef("%04X move.w (A6),D1\n", pc)
}

func (c *Processor) op1217() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[1] = uint32(v)
	c.tracef("%04X move.w (A7),D1\n", pc)
}

func (c *Processor) op1238() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[1] = uint32(v)
	c.tracef("%04X move.w $%X,D1\n", pc, v)
}

func (c *Processor) op1239() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[1] = uint32(v)
	c.tracef("%04X move.w $%X,D1\n", pc, v)
}

func (c *Processor) op123C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[1] = uint32(v)
	c.tracef("%04X move.w #$%X,D1\n", pc, v)
}

func (c *Processor) op1240() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w D0,A1\n", pc)
}

func (c *Processor) op1241() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w D1,A1\n", pc)
}

func (c *Processor) op1242() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w D2,A1\n", pc)
}

func (c *Processor) op1243() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w D3,A1\n", pc)
}

func (c *Processor) op1244() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w D4,A1\n", pc)
}

func (c *Processor) op1245() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w D5,A1\n", pc)
}

func (c *Processor) op1246() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w D6,A1\n", pc)
}

func (c *Processor) op1247() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w D7,A1\n", pc)
}

func (c *Processor) op1248() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w A0,A1\n", pc)
}

func (c *Processor) op1249() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w A1,A1\n", pc)
}

func (c *Processor) op124A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w A2,A1\n", pc)
}

func (c *Processor) op124B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w A3,A1\n", pc)
}

func (c *Processor) op124C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w A4,A1\n", pc)
}

func (c *Processor) op124D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w A5,A1\n", pc)
}

func (c *Processor) op124E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w A6,A1\n", pc)
}

func (c *Processor) op124F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w A7,A1\n", pc)
}

func (c *Processor) op1250() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w (A0),A1\n", pc)
}

func (c *Processor) op1251() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w (A1),A1\n", pc)
}

func (c *Processor) op1252() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w (A2),A1\n", pc)
}

func (c *Processor) op1253() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w (A3),A1\n", pc)
}

func (c *Processor) op1254() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w (A4),A1\n", pc)
}

func (c *Processor) op1255() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w (A5),A1\n", pc)
}

func (c *Processor) op1256() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w (A6),A1\n", pc)
}

func (c *Processor) op1257() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w (A7),A1\n", pc)
}

func (c *Processor) op1278() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w $%X,A1\n", pc, v)
}

func (c *Processor) op1279() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w $%X,A1\n", pc, v)
}

func (c *Processor) op127C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] = uint32(v)
	c.tracef("%04X movea.w #$%X,A1\n", pc, v)
}

func (c *Processor) op1280() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D0,(A1)\n", pc)
}

func (c *Processor) op1281() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D1,(A1)\n", pc)
}

func (c *Processor) op1282() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D2,(A1)\n", pc)
}

func (c *Processor) op1283() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D3,(A1)\n", pc)
}

func (c *Processor) op1284() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D4,(A1)\n", pc)
}

func (c *Processor) op1285() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D5,(A1)\n", pc)
}

func (c *Processor) op1286() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D6,(A1)\n", pc)
}

func (c *Processor) op1287() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D7,(A1)\n", pc)
}

func (c *Processor) op1288() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A0,(A1)\n", pc)
}

func (c *Processor) op1289() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A1,(A1)\n", pc)
}

func (c *Processor) op128A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A2,(A1)\n", pc)
}

func (c *Processor) op128B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A3,(A1)\n", pc)
}

func (c *Processor) op128C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A4,(A1)\n", pc)
}

func (c *Processor) op128D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A5,(A1)\n", pc)
}

func (c *Processor) op128E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A6,(A1)\n", pc)
}

func (c *Processor) op128F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A7,(A1)\n", pc)
}

func (c *Processor) op1290() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A0),(A1)\n", pc)
}

func (c *Processor) op1291() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A1),(A1)\n", pc)
}

func (c *Processor) op1292() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A2),(A1)\n", pc)
}

func (c *Processor) op1293() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A3),(A1)\n", pc)
}

func (c *Processor) op1294() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A4),(A1)\n", pc)
}

func (c *Processor) op1295() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A5),(A1)\n", pc)
}

func (c *Processor) op1296() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A6),(A1)\n", pc)
}

func (c *Processor) op1297() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A7),(A1)\n", pc)
}

func (c *Processor) op12B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w $%X,(A1)\n", pc, v)
}

func (c *Processor) op12B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w $%X,(A1)\n", pc, v)
}

func (c *Processor) op12BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w #$%X,(A1)\n", pc, v)
}

func (c *Processor) op12C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w D0,(A1)+\n", pc)
}

func (c *Processor) op12C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w D1,(A1)+\n", pc)
}

func (c *Processor) op12C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w D2,(A1)+\n", pc)
}

func (c *Processor) op12C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w D3,(A1)+\n", pc)
}

func (c *Processor) op12C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w D4,(A1)+\n", pc)
}

func (c *Processor) op12C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w D5,(A1)+\n", pc)
}

func (c *Processor) op12C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w D6,(A1)+\n", pc)
}

func (c *Processor) op12C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w D7,(A1)+\n", pc)
}

func (c *Processor) op12C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w A0,(A1)+\n", pc)
}

func (c *Processor) op12C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w A1,(A1)+\n", pc)
}

func (c *Processor) op12CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w A2,(A1)+\n", pc)
}

func (c *Processor) op12CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w A3,(A1)+\n", pc)
}

func (c *Processor) op12CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w A4,(A1)+\n", pc)
}

func (c *Processor) op12CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w A5,(A1)+\n", pc)
}

func (c *Processor) op12CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w A6,(A1)+\n", pc)
}

func (c *Processor) op12CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w A7,(A1)+\n", pc)
}

func (c *Processor) op12D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w (A0),(A1)+\n", pc)
}

func (c *Processor) op12D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w (A1),(A1)+\n", pc)
}

func (c *Processor) op12D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w (A2),(A1)+\n", pc)
}

func (c *Processor) op12D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w (A3),(A1)+\n", pc)
}

func (c *Processor) op12D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w (A4),(A1)+\n", pc)
}

func (c *Processor) op12D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w (A5),(A1)+\n", pc)
}

func (c *Processor) op12D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w (A6),(A1)+\n", pc)
}

func (c *Processor) op12D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w (A7),(A1)+\n", pc)
}

func (c *Processor) op12F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w $%X,(A1)+\n", pc, v)
}

func (c *Processor) op12F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w $%X,(A1)+\n", pc, v)
}

func (c *Processor) op12FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.w #$%X,(A1)+\n", pc, v)
}

func (c *Processor) op1300() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D0,-(A1)\n", pc)
}

func (c *Processor) op1301() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D1,-(A1)\n", pc)
}

func (c *Processor) op1302() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D2,-(A1)\n", pc)
}

func (c *Processor) op1303() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D3,-(A1)\n", pc)
}

func (c *Processor) op1304() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D4,-(A1)\n", pc)
}

func (c *Processor) op1305() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D5,-(A1)\n", pc)
}

func (c *Processor) op1306() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D6,-(A1)\n", pc)
}

func (c *Processor) op1307() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w D7,-(A1)\n", pc)
}

func (c *Processor) op1308() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A0,-(A1)\n", pc)
}

func (c *Processor) op1309() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A1,-(A1)\n", pc)
}

func (c *Processor) op130A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A2,-(A1)\n", pc)
}

func (c *Processor) op130B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A3,-(A1)\n", pc)
}

func (c *Processor) op130C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A4,-(A1)\n", pc)
}

func (c *Processor) op130D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A5,-(A1)\n", pc)
}

func (c *Processor) op130E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A6,-(A1)\n", pc)
}

func (c *Processor) op130F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w A7,-(A1)\n", pc)
}

func (c *Processor) op1310() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A0),-(A1)\n", pc)
}

func (c *Processor) op1311() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A1),-(A1)\n", pc)
}

func (c *Processor) op1312() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A2),-(A1)\n", pc)
}

func (c *Processor) op1313() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A3),-(A1)\n", pc)
}

func (c *Processor) op1314() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A4),-(A1)\n", pc)
}

func (c *Processor) op1315() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A5),-(A1)\n", pc)
}

func (c *Processor) op1316() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A6),-(A1)\n", pc)
}

func (c *Processor) op1317() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w (A7),-(A1)\n", pc)
}

func (c *Processor) op1338() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w $%X,-(A1)\n", pc, v)
}

func (c *Processor) op1339() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w $%X,-(A1)\n", pc, v)
}

func (c *Processor) op133C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.w #$%X,-(A1)\n", pc, v)
}

func (c *Processor) op1340() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D0,(%d,A1)\n", pc, disp)
}

func (c *Processor) op1341() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D1,(%d,A1)\n", pc, disp)
}

func (c *Processor) op1342() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D2,(%d,A1)\n", pc, disp)
}

func (c *Processor) op1343() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D3,(%d,A1)\n", pc, disp)
}

func (c *Processor) op1344() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D4,(%d,A1)\n", pc, disp)
}

func (c *Processor) op1345() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D5,(%d,A1)\n", pc, disp)
}

func (c *Processor) op1346() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D6,(%d,A1)\n", pc, disp)
}

func (c *Processor) op1347() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D7,(%d,A1)\n", pc, disp)
}

func (c *Processor) op1348() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A0,(%d,A1)\n", pc, disp)
}

func (c *Processor) op1349() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A1,(%d,A1)\n", pc, disp)
}

func (c *Processor) op134A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A2,(%d,A1)\n", pc, disp)
}

func (c *Processor) op134B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A3,(%d,A1)\n", pc, disp)
}

func (c *Processor) op134C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A4,(%d,A1)\n", pc, disp)
}

func (c *Processor) op134D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A5,(%d,A1)\n", pc, disp)
}

func (c *Processor) op134E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A6,(%d,A1)\n", pc, disp)
}

func (c *Processor) op134F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A7,(%d,A1)\n", pc, disp)
}

func (c *Processor) op1350() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A0),(%d,A1)\n", pc, disp)
}

func (c *Processor) op1351() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A1),(%d,A1)\n", pc, disp)
}

func (c *Processor) op1352() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A2),(%d,A1)\n", pc, disp)
}

func (c *Processor) op1353() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A3),(%d,A1)\n", pc, disp)
}

func (c *Processor) op1354() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A4),(%d,A1)\n", pc, disp)
}

func (c *Processor) op1355() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A5),(%d,A1)\n", pc, disp)
}

func (c *Processor) op1356() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A6),(%d,A1)\n", pc, disp)
}

func (c *Processor) op1357() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A7),(%d,A1)\n", pc, disp)
}

func (c *Processor) op1378() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A1)\n", pc, v, disp)
}

func (c *Processor) op1379() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A1)\n", pc, v, disp)
}

func (c *Processor) op137C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w #$%X,(%d,A1)\n", pc, v, disp)
}

func (c *Processor) op13C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D0,$%X\n", pc, addr)
}

func (c *Processor) op13C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D1,$%X\n", pc, addr)
}

func (c *Processor) op13C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D2,$%X\n", pc, addr)
}

func (c *Processor) op13C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D3,$%X\n", pc, addr)
}

func (c *Processor) op13C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D4,$%X\n", pc, addr)
}

func (c *Processor) op13C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D5,$%X\n", pc, addr)
}

func (c *Processor) op13C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D6,$%X\n", pc, addr)
}

func (c *Processor) op13C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D7,$%X\n", pc, addr)
}

func (c *Processor) op13C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A0,$%X\n", pc, addr)
}

func (c *Processor) op13C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A1,$%X\n", pc, addr)
}

func (c *Processor) op13CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A2,$%X\n", pc, addr)
}

func (c *Processor) op13CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A3,$%X\n", pc, addr)
}

func (c *Processor) op13CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A4,$%X\n", pc, addr)
}

func (c *Processor) op13CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A5,$%X\n", pc, addr)
}

func (c *Processor) op13CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A6,$%X\n", pc, addr)
}

func (c *Processor) op13CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A7,$%X\n", pc, addr)
}

func (c *Processor) op13D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A0),$%X\n", pc, addr)
}

func (c *Processor) op13D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A1),$%X\n", pc, addr)
}

func (c *Processor) op13D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A2),$%X\n", pc, addr)
}

func (c *Processor) op13D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A3),$%X\n", pc, addr)
}

func (c *Processor) op13D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A4),$%X\n", pc, addr)
}

func (c *Processor) op13D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A5),$%X\n", pc, addr)
}

func (c *Processor) op13D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A6),$%X\n", pc, addr)
}

func (c *Processor) op13D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A7),$%X\n", pc, addr)
}

func (c *Processor) op13F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,$%X\n", pc, v, addr)
}

func (c *Processor) op13F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,$%X\n", pc, v, addr)
}

func (c *Processor) op13FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w #$%X,$%X\n", pc, v, addr)
}

func (c *Processor) op1400() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w D0,D2\n", pc)
}

func (c *Processor) op1401() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w D1,D2\n", pc)
}

func (c *Processor) op1402() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w D2,D2\n", pc)
}

func (c *Processor) op1403() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w D3,D2\n", pc)
}

func (c *Processor) op1404() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w D4,D2\n", pc)
}

func (c *Processor) op1405() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w D5,D2\n", pc)
}

func (c *Processor) op1406() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w D6,D2\n", pc)
}

func (c *Processor) op1407() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w D7,D2\n", pc)
}

func (c *Processor) op1408() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w A0,D2\n", pc)
}

func (c *Processor) op1409() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w A1,D2\n", pc)
}

func (c *Processor) op140A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w A2,D2\n", pc)
}

func (c *Processor) op140B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w A3,D2\n", pc)
}

func (c *Processor) op140C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w A4,D2\n", pc)
}

func (c *Processor) op140D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w A5,D2\n", pc)
}

func (c *Processor) op140E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w A6,D2\n", pc)
}

func (c *Processor) op140F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[2] = uint32(v)
	c.tracef("%04X move.w A7,D2\n", pc)
}

func (c *Processor) op1410() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[2] = uint32(v)
	c.tracef("%04X move.w (A0),D2\n", pc)
}

func (c *Processor) op1411() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[2] = uint32(v)
	c.tracef("%04X move.w (A1),D2\n", pc)
}

func (c *Processor) op1412() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[2] = uint32(v)
	c.tracef("%04X move.w (A2),D2\n", pc)
}

func (c *Processor) op1413() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[2] = uint32(v)
	c.tracef("%04X move.w (A3),D2\n", pc)
}

func (c *Processor) op1414() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[2] = uint32(v)
	c.tracef("%04X move.w (A4),D2\n", pc)
}

func (c *Processor) op1415() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[2] = uint32(v)
	c.tracef("%04X move.w (A5),D2\n", pc)
}

func (c *Processor) op1416() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[2] = uint32(v)
	c.tracef("%04X move.w (A6),D2\n", pc)
}

func (c *Processor) op1417() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[2] = uint32(v)
	c.tracef("%04X move.w (A7),D2\n", pc)
}

func (c *Processor) op1438() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[2] = uint32(v)
	c.tracef("%04X move.w $%X,D2\n", pc, v)
}

func (c *Processor) op1439() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[2] = uint32(v)
	c.tracef("%04X move.w $%X,D2\n", pc, v)
}

func (c *Processor) op143C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[2] = uint32(v)
	c.tracef("%04X move.w #$%X,D2\n", pc, v)
}

func (c *Processor) op1440() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w D0,A2\n", pc)
}

func (c *Processor) op1441() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w D1,A2\n", pc)
}

func (c *Processor) op1442() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w D2,A2\n", pc)
}

func (c *Processor) op1443() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w D3,A2\n", pc)
}

func (c *Processor) op1444() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w D4,A2\n", pc)
}

func (c *Processor) op1445() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w D5,A2\n", pc)
}

func (c *Processor) op1446() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w D6,A2\n", pc)
}

func (c *Processor) op1447() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w D7,A2\n", pc)
}

func (c *Processor) op1448() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w A0,A2\n", pc)
}

func (c *Processor) op1449() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w A1,A2\n", pc)
}

func (c *Processor) op144A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w A2,A2\n", pc)
}

func (c *Processor) op144B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w A3,A2\n", pc)
}

func (c *Processor) op144C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w A4,A2\n", pc)
}

func (c *Processor) op144D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w A5,A2\n", pc)
}

func (c *Processor) op144E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w A6,A2\n", pc)
}

func (c *Processor) op144F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w A7,A2\n", pc)
}

func (c *Processor) op1450() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w (A0),A2\n", pc)
}

func (c *Processor) op1451() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w (A1),A2\n", pc)
}

func (c *Processor) op1452() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w (A2),A2\n", pc)
}

func (c *Processor) op1453() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w (A3),A2\n", pc)
}

func (c *Processor) op1454() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w (A4),A2\n", pc)
}

func (c *Processor) op1455() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w (A5),A2\n", pc)
}

func (c *Processor) op1456() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w (A6),A2\n", pc)
}

func (c *Processor) op1457() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w (A7),A2\n", pc)
}

func (c *Processor) op1478() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w $%X,A2\n", pc, v)
}

func (c *Processor) op1479() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w $%X,A2\n", pc, v)
}

func (c *Processor) op147C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] = uint32(v)
	c.tracef("%04X movea.w #$%X,A2\n", pc, v)
}

func (c *Processor) op1480() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D0,(A2)\n", pc)
}

func (c *Processor) op1481() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D1,(A2)\n", pc)
}

func (c *Processor) op1482() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D2,(A2)\n", pc)
}

func (c *Processor) op1483() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D3,(A2)\n", pc)
}

func (c *Processor) op1484() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D4,(A2)\n", pc)
}

func (c *Processor) op1485() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D5,(A2)\n", pc)
}

func (c *Processor) op1486() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D6,(A2)\n", pc)
}

func (c *Processor) op1487() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D7,(A2)\n", pc)
}

func (c *Processor) op1488() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A0,(A2)\n", pc)
}

func (c *Processor) op1489() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A1,(A2)\n", pc)
}

func (c *Processor) op148A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A2,(A2)\n", pc)
}

func (c *Processor) op148B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A3,(A2)\n", pc)
}

func (c *Processor) op148C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A4,(A2)\n", pc)
}

func (c *Processor) op148D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A5,(A2)\n", pc)
}

func (c *Processor) op148E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A6,(A2)\n", pc)
}

func (c *Processor) op148F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A7,(A2)\n", pc)
}

func (c *Processor) op1490() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A0),(A2)\n", pc)
}

func (c *Processor) op1491() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A1),(A2)\n", pc)
}

func (c *Processor) op1492() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A2),(A2)\n", pc)
}

func (c *Processor) op1493() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A3),(A2)\n", pc)
}

func (c *Processor) op1494() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A4),(A2)\n", pc)
}

func (c *Processor) op1495() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A5),(A2)\n", pc)
}

func (c *Processor) op1496() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A6),(A2)\n", pc)
}

func (c *Processor) op1497() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A7),(A2)\n", pc)
}

func (c *Processor) op14B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w $%X,(A2)\n", pc, v)
}

func (c *Processor) op14B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w $%X,(A2)\n", pc, v)
}

func (c *Processor) op14BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w #$%X,(A2)\n", pc, v)
}

func (c *Processor) op14C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w D0,(A2)+\n", pc)
}

func (c *Processor) op14C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w D1,(A2)+\n", pc)
}

func (c *Processor) op14C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w D2,(A2)+\n", pc)
}

func (c *Processor) op14C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w D3,(A2)+\n", pc)
}

func (c *Processor) op14C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w D4,(A2)+\n", pc)
}

func (c *Processor) op14C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w D5,(A2)+\n", pc)
}

func (c *Processor) op14C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w D6,(A2)+\n", pc)
}

func (c *Processor) op14C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w D7,(A2)+\n", pc)
}

func (c *Processor) op14C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w A0,(A2)+\n", pc)
}

func (c *Processor) op14C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w A1,(A2)+\n", pc)
}

func (c *Processor) op14CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w A2,(A2)+\n", pc)
}

func (c *Processor) op14CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w A3,(A2)+\n", pc)
}

func (c *Processor) op14CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w A4,(A2)+\n", pc)
}

func (c *Processor) op14CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w A5,(A2)+\n", pc)
}

func (c *Processor) op14CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w A6,(A2)+\n", pc)
}

func (c *Processor) op14CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w A7,(A2)+\n", pc)
}

func (c *Processor) op14D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w (A0),(A2)+\n", pc)
}

func (c *Processor) op14D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w (A1),(A2)+\n", pc)
}

func (c *Processor) op14D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w (A2),(A2)+\n", pc)
}

func (c *Processor) op14D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w (A3),(A2)+\n", pc)
}

func (c *Processor) op14D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w (A4),(A2)+\n", pc)
}

func (c *Processor) op14D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w (A5),(A2)+\n", pc)
}

func (c *Processor) op14D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w (A6),(A2)+\n", pc)
}

func (c *Processor) op14D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w (A7),(A2)+\n", pc)
}

func (c *Processor) op14F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w $%X,(A2)+\n", pc, v)
}

func (c *Processor) op14F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w $%X,(A2)+\n", pc, v)
}

func (c *Processor) op14FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.w #$%X,(A2)+\n", pc, v)
}

func (c *Processor) op1500() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D0,-(A2)\n", pc)
}

func (c *Processor) op1501() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D1,-(A2)\n", pc)
}

func (c *Processor) op1502() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D2,-(A2)\n", pc)
}

func (c *Processor) op1503() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D3,-(A2)\n", pc)
}

func (c *Processor) op1504() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D4,-(A2)\n", pc)
}

func (c *Processor) op1505() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D5,-(A2)\n", pc)
}

func (c *Processor) op1506() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D6,-(A2)\n", pc)
}

func (c *Processor) op1507() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w D7,-(A2)\n", pc)
}

func (c *Processor) op1508() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A0,-(A2)\n", pc)
}

func (c *Processor) op1509() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A1,-(A2)\n", pc)
}

func (c *Processor) op150A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A2,-(A2)\n", pc)
}

func (c *Processor) op150B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A3,-(A2)\n", pc)
}

func (c *Processor) op150C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A4,-(A2)\n", pc)
}

func (c *Processor) op150D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A5,-(A2)\n", pc)
}

func (c *Processor) op150E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A6,-(A2)\n", pc)
}

func (c *Processor) op150F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w A7,-(A2)\n", pc)
}

func (c *Processor) op1510() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A0),-(A2)\n", pc)
}

func (c *Processor) op1511() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A1),-(A2)\n", pc)
}

func (c *Processor) op1512() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A2),-(A2)\n", pc)
}

func (c *Processor) op1513() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A3),-(A2)\n", pc)
}

func (c *Processor) op1514() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A4),-(A2)\n", pc)
}

func (c *Processor) op1515() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A5),-(A2)\n", pc)
}

func (c *Processor) op1516() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A6),-(A2)\n", pc)
}

func (c *Processor) op1517() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w (A7),-(A2)\n", pc)
}

func (c *Processor) op1538() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w $%X,-(A2)\n", pc, v)
}

func (c *Processor) op1539() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w $%X,-(A2)\n", pc, v)
}

func (c *Processor) op153C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.w #$%X,-(A2)\n", pc, v)
}

func (c *Processor) op1540() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D0,(%d,A2)\n", pc, disp)
}

func (c *Processor) op1541() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D1,(%d,A2)\n", pc, disp)
}

func (c *Processor) op1542() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D2,(%d,A2)\n", pc, disp)
}

func (c *Processor) op1543() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D3,(%d,A2)\n", pc, disp)
}

func (c *Processor) op1544() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D4,(%d,A2)\n", pc, disp)
}

func (c *Processor) op1545() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D5,(%d,A2)\n", pc, disp)
}

func (c *Processor) op1546() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D6,(%d,A2)\n", pc, disp)
}

func (c *Processor) op1547() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D7,(%d,A2)\n", pc, disp)
}

func (c *Processor) op1548() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A0,(%d,A2)\n", pc, disp)
}

func (c *Processor) op1549() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A1,(%d,A2)\n", pc, disp)
}

func (c *Processor) op154A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A2,(%d,A2)\n", pc, disp)
}

func (c *Processor) op154B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A3,(%d,A2)\n", pc, disp)
}

func (c *Processor) op154C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A4,(%d,A2)\n", pc, disp)
}

func (c *Processor) op154D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A5,(%d,A2)\n", pc, disp)
}

func (c *Processor) op154E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A6,(%d,A2)\n", pc, disp)
}

func (c *Processor) op154F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A7,(%d,A2)\n", pc, disp)
}

func (c *Processor) op1550() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A0),(%d,A2)\n", pc, disp)
}

func (c *Processor) op1551() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A1),(%d,A2)\n", pc, disp)
}

func (c *Processor) op1552() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A2),(%d,A2)\n", pc, disp)
}

func (c *Processor) op1553() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A3),(%d,A2)\n", pc, disp)
}

func (c *Processor) op1554() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A4),(%d,A2)\n", pc, disp)
}

func (c *Processor) op1555() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A5),(%d,A2)\n", pc, disp)
}

func (c *Processor) op1556() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A6),(%d,A2)\n", pc, disp)
}

func (c *Processor) op1557() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A7),(%d,A2)\n", pc, disp)
}

func (c *Processor) op1578() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A2)\n", pc, v, disp)
}

func (c *Processor) op1579() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A2)\n", pc, v, disp)
}

func (c *Processor) op157C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w #$%X,(%d,A2)\n", pc, v, disp)
}

func (c *Processor) op1600() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w D0,D3\n", pc)
}

func (c *Processor) op1601() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w D1,D3\n", pc)
}

func (c *Processor) op1602() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w D2,D3\n", pc)
}

func (c *Processor) op1603() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w D3,D3\n", pc)
}

func (c *Processor) op1604() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w D4,D3\n", pc)
}

func (c *Processor) op1605() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w D5,D3\n", pc)
}

func (c *Processor) op1606() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w D6,D3\n", pc)
}

func (c *Processor) op1607() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w D7,D3\n", pc)
}

func (c *Processor) op1608() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w A0,D3\n", pc)
}

func (c *Processor) op1609() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w A1,D3\n", pc)
}

func (c *Processor) op160A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w A2,D3\n", pc)
}

func (c *Processor) op160B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w A3,D3\n", pc)
}

func (c *Processor) op160C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w A4,D3\n", pc)
}

func (c *Processor) op160D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w A5,D3\n", pc)
}

func (c *Processor) op160E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w A6,D3\n", pc)
}

func (c *Processor) op160F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[3] = uint32(v)
	c.tracef("%04X move.w A7,D3\n", pc)
}

func (c *Processor) op1610() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[3] = uint32(v)
	c.tracef("%04X move.w (A0),D3\n", pc)
}

func (c *Processor) op1611() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[3] = uint32(v)
	c.tracef("%04X move.w (A1),D3\n", pc)
}

func (c *Processor) op1612() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[3] = uint32(v)
	c.tracef("%04X move.w (A2),D3\n", pc)
}

func (c *Processor) op1613() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[3] = uint32(v)
	c.tracef("%04X move.w (A3),D3\n", pc)
}

func (c *Processor) op1614() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[3] = uint32(v)
	c.tracef("%04X move.w (A4),D3\n", pc)
}

func (c *Processor) op1615() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[3] = uint32(v)
	c.tracef("%04X move.w (A5),D3\n", pc)
}

func (c *Processor) op1616() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[3] = uint32(v)
	c.tracef("%04X move.w (A6),D3\n", pc)
}

func (c *Processor) op1617() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[3] = uint32(v)
	c.tracef("%04X move.w (A7),D3\n", pc)
}

func (c *Processor) op1638() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[3] = uint32(v)
	c.tracef("%04X move.w $%X,D3\n", pc, v)
}

func (c *Processor) op1639() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[3] = uint32(v)
	c.tracef("%04X move.w $%X,D3\n", pc, v)
}

func (c *Processor) op163C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[3] = uint32(v)
	c.tracef("%04X move.w #$%X,D3\n", pc, v)
}

func (c *Processor) op1640() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w D0,A3\n", pc)
}

func (c *Processor) op1641() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w D1,A3\n", pc)
}

func (c *Processor) op1642() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w D2,A3\n", pc)
}

func (c *Processor) op1643() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w D3,A3\n", pc)
}

func (c *Processor) op1644() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w D4,A3\n", pc)
}

func (c *Processor) op1645() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w D5,A3\n", pc)
}

func (c *Processor) op1646() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w D6,A3\n", pc)
}

func (c *Processor) op1647() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w D7,A3\n", pc)
}

func (c *Processor) op1648() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w A0,A3\n", pc)
}

func (c *Processor) op1649() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w A1,A3\n", pc)
}

func (c *Processor) op164A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w A2,A3\n", pc)
}

func (c *Processor) op164B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w A3,A3\n", pc)
}

func (c *Processor) op164C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w A4,A3\n", pc)
}

func (c *Processor) op164D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w A5,A3\n", pc)
}

func (c *Processor) op164E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w A6,A3\n", pc)
}

func (c *Processor) op164F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w A7,A3\n", pc)
}

func (c *Processor) op1650() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w (A0),A3\n", pc)
}

func (c *Processor) op1651() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w (A1),A3\n", pc)
}

func (c *Processor) op1652() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w (A2),A3\n", pc)
}

func (c *Processor) op1653() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w (A3),A3\n", pc)
}

func (c *Processor) op1654() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w (A4),A3\n", pc)
}

func (c *Processor) op1655() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w (A5),A3\n", pc)
}

func (c *Processor) op1656() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w (A6),A3\n", pc)
}

func (c *Processor) op1657() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w (A7),A3\n", pc)
}

func (c *Processor) op1678() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w $%X,A3\n", pc, v)
}

func (c *Processor) op1679() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w $%X,A3\n", pc, v)
}

func (c *Processor) op167C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] = uint32(v)
	c.tracef("%04X movea.w #$%X,A3\n", pc, v)
}

func (c *Processor) op1680() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D0,(A3)\n", pc)
}

func (c *Processor) op1681() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D1,(A3)\n", pc)
}

func (c *Processor) op1682() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D2,(A3)\n", pc)
}

func (c *Processor) op1683() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D3,(A3)\n", pc)
}

func (c *Processor) op1684() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D4,(A3)\n", pc)
}

func (c *Processor) op1685() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D5,(A3)\n", pc)
}

func (c *Processor) op1686() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D6,(A3)\n", pc)
}

func (c *Processor) op1687() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D7,(A3)\n", pc)
}

func (c *Processor) op1688() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A0,(A3)\n", pc)
}

func (c *Processor) op1689() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A1,(A3)\n", pc)
}

func (c *Processor) op168A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A2,(A3)\n", pc)
}

func (c *Processor) op168B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A3,(A3)\n", pc)
}

func (c *Processor) op168C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A4,(A3)\n", pc)
}

func (c *Processor) op168D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A5,(A3)\n", pc)
}

func (c *Processor) op168E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A6,(A3)\n", pc)
}

func (c *Processor) op168F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A7,(A3)\n", pc)
}

func (c *Processor) op1690() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A0),(A3)\n", pc)
}

func (c *Processor) op1691() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A1),(A3)\n", pc)
}

func (c *Processor) op1692() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A2),(A3)\n", pc)
}

func (c *Processor) op1693() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A3),(A3)\n", pc)
}

func (c *Processor) op1694() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A4),(A3)\n", pc)
}

func (c *Processor) op1695() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A5),(A3)\n", pc)
}

func (c *Processor) op1696() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A6),(A3)\n", pc)
}

func (c *Processor) op1697() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A7),(A3)\n", pc)
}

func (c *Processor) op16B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w $%X,(A3)\n", pc, v)
}

func (c *Processor) op16B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w $%X,(A3)\n", pc, v)
}

func (c *Processor) op16BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w #$%X,(A3)\n", pc, v)
}

func (c *Processor) op16C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w D0,(A3)+\n", pc)
}

func (c *Processor) op16C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w D1,(A3)+\n", pc)
}

func (c *Processor) op16C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w D2,(A3)+\n", pc)
}

func (c *Processor) op16C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w D3,(A3)+\n", pc)
}

func (c *Processor) op16C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w D4,(A3)+\n", pc)
}

func (c *Processor) op16C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w D5,(A3)+\n", pc)
}

func (c *Processor) op16C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w D6,(A3)+\n", pc)
}

func (c *Processor) op16C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w D7,(A3)+\n", pc)
}

func (c *Processor) op16C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w A0,(A3)+\n", pc)
}

func (c *Processor) op16C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w A1,(A3)+\n", pc)
}

func (c *Processor) op16CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w A2,(A3)+\n", pc)
}

func (c *Processor) op16CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w A3,(A3)+\n", pc)
}

func (c *Processor) op16CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w A4,(A3)+\n", pc)
}

func (c *Processor) op16CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w A5,(A3)+\n", pc)
}

func (c *Processor) op16CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w A6,(A3)+\n", pc)
}

func (c *Processor) op16CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w A7,(A3)+\n", pc)
}

func (c *Processor) op16D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w (A0),(A3)+\n", pc)
}

func (c *Processor) op16D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w (A1),(A3)+\n", pc)
}

func (c *Processor) op16D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w (A2),(A3)+\n", pc)
}

func (c *Processor) op16D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w (A3),(A3)+\n", pc)
}

func (c *Processor) op16D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w (A4),(A3)+\n", pc)
}

func (c *Processor) op16D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w (A5),(A3)+\n", pc)
}

func (c *Processor) op16D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w (A6),(A3)+\n", pc)
}

func (c *Processor) op16D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w (A7),(A3)+\n", pc)
}

func (c *Processor) op16F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w $%X,(A3)+\n", pc, v)
}

func (c *Processor) op16F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w $%X,(A3)+\n", pc, v)
}

func (c *Processor) op16FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.w #$%X,(A3)+\n", pc, v)
}

func (c *Processor) op1700() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D0,-(A3)\n", pc)
}

func (c *Processor) op1701() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D1,-(A3)\n", pc)
}

func (c *Processor) op1702() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D2,-(A3)\n", pc)
}

func (c *Processor) op1703() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D3,-(A3)\n", pc)
}

func (c *Processor) op1704() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D4,-(A3)\n", pc)
}

func (c *Processor) op1705() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D5,-(A3)\n", pc)
}

func (c *Processor) op1706() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D6,-(A3)\n", pc)
}

func (c *Processor) op1707() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w D7,-(A3)\n", pc)
}

func (c *Processor) op1708() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A0,-(A3)\n", pc)
}

func (c *Processor) op1709() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A1,-(A3)\n", pc)
}

func (c *Processor) op170A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A2,-(A3)\n", pc)
}

func (c *Processor) op170B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A3,-(A3)\n", pc)
}

func (c *Processor) op170C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A4,-(A3)\n", pc)
}

func (c *Processor) op170D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A5,-(A3)\n", pc)
}

func (c *Processor) op170E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A6,-(A3)\n", pc)
}

func (c *Processor) op170F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w A7,-(A3)\n", pc)
}

func (c *Processor) op1710() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A0),-(A3)\n", pc)
}

func (c *Processor) op1711() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A1),-(A3)\n", pc)
}

func (c *Processor) op1712() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A2),-(A3)\n", pc)
}

func (c *Processor) op1713() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A3),-(A3)\n", pc)
}

func (c *Processor) op1714() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A4),-(A3)\n", pc)
}

func (c *Processor) op1715() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A5),-(A3)\n", pc)
}

func (c *Processor) op1716() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A6),-(A3)\n", pc)
}

func (c *Processor) op1717() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w (A7),-(A3)\n", pc)
}

func (c *Processor) op1738() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w $%X,-(A3)\n", pc, v)
}

func (c *Processor) op1739() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w $%X,-(A3)\n", pc, v)
}

func (c *Processor) op173C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.w #$%X,-(A3)\n", pc, v)
}

func (c *Processor) op1740() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D0,(%d,A3)\n", pc, disp)
}

func (c *Processor) op1741() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D1,(%d,A3)\n", pc, disp)
}

func (c *Processor) op1742() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D2,(%d,A3)\n", pc, disp)
}

func (c *Processor) op1743() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D3,(%d,A3)\n", pc, disp)
}

func (c *Processor) op1744() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D4,(%d,A3)\n", pc, disp)
}

func (c *Processor) op1745() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D5,(%d,A3)\n", pc, disp)
}

func (c *Processor) op1746() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D6,(%d,A3)\n", pc, disp)
}

func (c *Processor) op1747() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D7,(%d,A3)\n", pc, disp)
}

func (c *Processor) op1748() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A0,(%d,A3)\n", pc, disp)
}

func (c *Processor) op1749() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A1,(%d,A3)\n", pc, disp)
}

func (c *Processor) op174A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A2,(%d,A3)\n", pc, disp)
}

func (c *Processor) op174B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A3,(%d,A3)\n", pc, disp)
}

func (c *Processor) op174C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A4,(%d,A3)\n", pc, disp)
}

func (c *Processor) op174D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A5,(%d,A3)\n", pc, disp)
}

func (c *Processor) op174E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A6,(%d,A3)\n", pc, disp)
}

func (c *Processor) op174F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A7,(%d,A3)\n", pc, disp)
}

func (c *Processor) op1750() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A0),(%d,A3)\n", pc, disp)
}

func (c *Processor) op1751() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A1),(%d,A3)\n", pc, disp)
}

func (c *Processor) op1752() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A2),(%d,A3)\n", pc, disp)
}

func (c *Processor) op1753() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A3),(%d,A3)\n", pc, disp)
}

func (c *Processor) op1754() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A4),(%d,A3)\n", pc, disp)
}

func (c *Processor) op1755() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A5),(%d,A3)\n", pc, disp)
}

func (c *Processor) op1756() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A6),(%d,A3)\n", pc, disp)
}

func (c *Processor) op1757() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A7),(%d,A3)\n", pc, disp)
}

func (c *Processor) op1778() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A3)\n", pc, v, disp)
}

func (c *Processor) op1779() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A3)\n", pc, v, disp)
}

func (c *Processor) op177C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w #$%X,(%d,A3)\n", pc, v, disp)
}

func (c *Processor) op1800() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w D0,D4\n", pc)
}

func (c *Processor) op1801() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w D1,D4\n", pc)
}

func (c *Processor) op1802() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w D2,D4\n", pc)
}

func (c *Processor) op1803() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w D3,D4\n", pc)
}

func (c *Processor) op1804() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w D4,D4\n", pc)
}

func (c *Processor) op1805() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w D5,D4\n", pc)
}

func (c *Processor) op1806() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w D6,D4\n", pc)
}

func (c *Processor) op1807() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w D7,D4\n", pc)
}

func (c *Processor) op1808() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w A0,D4\n", pc)
}

func (c *Processor) op1809() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w A1,D4\n", pc)
}

func (c *Processor) op180A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w A2,D4\n", pc)
}

func (c *Processor) op180B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w A3,D4\n", pc)
}

func (c *Processor) op180C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w A4,D4\n", pc)
}

func (c *Processor) op180D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w A5,D4\n", pc)
}

func (c *Processor) op180E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w A6,D4\n", pc)
}

func (c *Processor) op180F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[4] = uint32(v)
	c.tracef("%04X move.w A7,D4\n", pc)
}

func (c *Processor) op1810() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[4] = uint32(v)
	c.tracef("%04X move.w (A0),D4\n", pc)
}

func (c *Processor) op1811() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[4] = uint32(v)
	c.tracef("%04X move.w (A1),D4\n", pc)
}

func (c *Processor) op1812() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[4] = uint32(v)
	c.tracef("%04X move.w (A2),D4\n", pc)
}

func (c *Processor) op1813() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[4] = uint32(v)
	c.tracef("%04X move.w (A3),D4\n", pc)
}

func (c *Processor) op1814() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[4] = uint32(v)
	c.tracef("%04X move.w (A4),D4\n", pc)
}

func (c *Processor) op1815() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[4] = uint32(v)
	c.tracef("%04X move.w (A5),D4\n", pc)
}

func (c *Processor) op1816() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[4] = uint32(v)
	c.tracef("%04X move.w (A6),D4\n", pc)
}

func (c *Processor) op1817() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[4] = uint32(v)
	c.tracef("%04X move.w (A7),D4\n", pc)
}

func (c *Processor) op1838() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[4] = uint32(v)
	c.tracef("%04X move.w $%X,D4\n", pc, v)
}

func (c *Processor) op1839() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[4] = uint32(v)
	c.tracef("%04X move.w $%X,D4\n", pc, v)
}

func (c *Processor) op183C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[4] = uint32(v)
	c.tracef("%04X move.w #$%X,D4\n", pc, v)
}

func (c *Processor) op1840() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w D0,A4\n", pc)
}

func (c *Processor) op1841() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w D1,A4\n", pc)
}

func (c *Processor) op1842() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w D2,A4\n", pc)
}

func (c *Processor) op1843() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w D3,A4\n", pc)
}

func (c *Processor) op1844() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w D4,A4\n", pc)
}

func (c *Processor) op1845() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w D5,A4\n", pc)
}

func (c *Processor) op1846() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w D6,A4\n", pc)
}

func (c *Processor) op1847() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w D7,A4\n", pc)
}

func (c *Processor) op1848() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w A0,A4\n", pc)
}

func (c *Processor) op1849() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w A1,A4\n", pc)
}

func (c *Processor) op184A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w A2,A4\n", pc)
}

func (c *Processor) op184B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w A3,A4\n", pc)
}

func (c *Processor) op184C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w A4,A4\n", pc)
}

func (c *Processor) op184D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w A5,A4\n", pc)
}

func (c *Processor) op184E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w A6,A4\n", pc)
}

func (c *Processor) op184F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w A7,A4\n", pc)
}

func (c *Processor) op1850() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w (A0),A4\n", pc)
}

func (c *Processor) op1851() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w (A1),A4\n", pc)
}

func (c *Processor) op1852() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w (A2),A4\n", pc)
}

func (c *Processor) op1853() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w (A3),A4\n", pc)
}

func (c *Processor) op1854() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w (A4),A4\n", pc)
}

func (c *Processor) op1855() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w (A5),A4\n", pc)
}

func (c *Processor) op1856() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w (A6),A4\n", pc)
}

func (c *Processor) op1857() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w (A7),A4\n", pc)
}

func (c *Processor) op1878() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w $%X,A4\n", pc, v)
}

func (c *Processor) op1879() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w $%X,A4\n", pc, v)
}

func (c *Processor) op187C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] = uint32(v)
	c.tracef("%04X movea.w #$%X,A4\n", pc, v)
}

func (c *Processor) op1880() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D0,(A4)\n", pc)
}

func (c *Processor) op1881() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D1,(A4)\n", pc)
}

func (c *Processor) op1882() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D2,(A4)\n", pc)
}

func (c *Processor) op1883() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D3,(A4)\n", pc)
}

func (c *Processor) op1884() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D4,(A4)\n", pc)
}

func (c *Processor) op1885() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D5,(A4)\n", pc)
}

func (c *Processor) op1886() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D6,(A4)\n", pc)
}

func (c *Processor) op1887() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D7,(A4)\n", pc)
}

func (c *Processor) op1888() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A0,(A4)\n", pc)
}

func (c *Processor) op1889() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A1,(A4)\n", pc)
}

func (c *Processor) op188A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A2,(A4)\n", pc)
}

func (c *Processor) op188B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A3,(A4)\n", pc)
}

func (c *Processor) op188C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A4,(A4)\n", pc)
}

func (c *Processor) op188D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A5,(A4)\n", pc)
}

func (c *Processor) op188E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A6,(A4)\n", pc)
}

func (c *Processor) op188F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A7,(A4)\n", pc)
}

func (c *Processor) op1890() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A0),(A4)\n", pc)
}

func (c *Processor) op1891() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A1),(A4)\n", pc)
}

func (c *Processor) op1892() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A2),(A4)\n", pc)
}

func (c *Processor) op1893() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A3),(A4)\n", pc)
}

func (c *Processor) op1894() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A4),(A4)\n", pc)
}

func (c *Processor) op1895() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A5),(A4)\n", pc)
}

func (c *Processor) op1896() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A6),(A4)\n", pc)
}

func (c *Processor) op1897() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A7),(A4)\n", pc)
}

func (c *Processor) op18B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w $%X,(A4)\n", pc, v)
}

func (c *Processor) op18B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w $%X,(A4)\n", pc, v)
}

func (c *Processor) op18BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w #$%X,(A4)\n", pc, v)
}

func (c *Processor) op18C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w D0,(A4)+\n", pc)
}

func (c *Processor) op18C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w D1,(A4)+\n", pc)
}

func (c *Processor) op18C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w D2,(A4)+\n", pc)
}

func (c *Processor) op18C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w D3,(A4)+\n", pc)
}

func (c *Processor) op18C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w D4,(A4)+\n", pc)
}

func (c *Processor) op18C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w D5,(A4)+\n", pc)
}

func (c *Processor) op18C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w D6,(A4)+\n", pc)
}

func (c *Processor) op18C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w D7,(A4)+\n", pc)
}

func (c *Processor) op18C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w A0,(A4)+\n", pc)
}

func (c *Processor) op18C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w A1,(A4)+\n", pc)
}

func (c *Processor) op18CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w A2,(A4)+\n", pc)
}

func (c *Processor) op18CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w A3,(A4)+\n", pc)
}

func (c *Processor) op18CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w A4,(A4)+\n", pc)
}

func (c *Processor) op18CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w A5,(A4)+\n", pc)
}

func (c *Processor) op18CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w A6,(A4)+\n", pc)
}

func (c *Processor) op18CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w A7,(A4)+\n", pc)
}

func (c *Processor) op18D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w (A0),(A4)+\n", pc)
}

func (c *Processor) op18D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w (A1),(A4)+\n", pc)
}

func (c *Processor) op18D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w (A2),(A4)+\n", pc)
}

func (c *Processor) op18D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w (A3),(A4)+\n", pc)
}

func (c *Processor) op18D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w (A4),(A4)+\n", pc)
}

func (c *Processor) op18D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w (A5),(A4)+\n", pc)
}

func (c *Processor) op18D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w (A6),(A4)+\n", pc)
}

func (c *Processor) op18D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w (A7),(A4)+\n", pc)
}

func (c *Processor) op18F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w $%X,(A4)+\n", pc, v)
}

func (c *Processor) op18F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w $%X,(A4)+\n", pc, v)
}

func (c *Processor) op18FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.w #$%X,(A4)+\n", pc, v)
}

func (c *Processor) op1900() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D0,-(A4)\n", pc)
}

func (c *Processor) op1901() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D1,-(A4)\n", pc)
}

func (c *Processor) op1902() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D2,-(A4)\n", pc)
}

func (c *Processor) op1903() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D3,-(A4)\n", pc)
}

func (c *Processor) op1904() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D4,-(A4)\n", pc)
}

func (c *Processor) op1905() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D5,-(A4)\n", pc)
}

func (c *Processor) op1906() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D6,-(A4)\n", pc)
}

func (c *Processor) op1907() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w D7,-(A4)\n", pc)
}

func (c *Processor) op1908() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A0,-(A4)\n", pc)
}

func (c *Processor) op1909() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A1,-(A4)\n", pc)
}

func (c *Processor) op190A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A2,-(A4)\n", pc)
}

func (c *Processor) op190B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A3,-(A4)\n", pc)
}

func (c *Processor) op190C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A4,-(A4)\n", pc)
}

func (c *Processor) op190D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A5,-(A4)\n", pc)
}

func (c *Processor) op190E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A6,-(A4)\n", pc)
}

func (c *Processor) op190F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w A7,-(A4)\n", pc)
}

func (c *Processor) op1910() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A0),-(A4)\n", pc)
}

func (c *Processor) op1911() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A1),-(A4)\n", pc)
}

func (c *Processor) op1912() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A2),-(A4)\n", pc)
}

func (c *Processor) op1913() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A3),-(A4)\n", pc)
}

func (c *Processor) op1914() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A4),-(A4)\n", pc)
}

func (c *Processor) op1915() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A5),-(A4)\n", pc)
}

func (c *Processor) op1916() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A6),-(A4)\n", pc)
}

func (c *Processor) op1917() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w (A7),-(A4)\n", pc)
}

func (c *Processor) op1938() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w $%X,-(A4)\n", pc, v)
}

func (c *Processor) op1939() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w $%X,-(A4)\n", pc, v)
}

func (c *Processor) op193C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.w #$%X,-(A4)\n", pc, v)
}

func (c *Processor) op1940() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D0,(%d,A4)\n", pc, disp)
}

func (c *Processor) op1941() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D1,(%d,A4)\n", pc, disp)
}

func (c *Processor) op1942() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D2,(%d,A4)\n", pc, disp)
}

func (c *Processor) op1943() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D3,(%d,A4)\n", pc, disp)
}

func (c *Processor) op1944() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D4,(%d,A4)\n", pc, disp)
}

func (c *Processor) op1945() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D5,(%d,A4)\n", pc, disp)
}

func (c *Processor) op1946() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D6,(%d,A4)\n", pc, disp)
}

func (c *Processor) op1947() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D7,(%d,A4)\n", pc, disp)
}

func (c *Processor) op1948() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A0,(%d,A4)\n", pc, disp)
}

func (c *Processor) op1949() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A1,(%d,A4)\n", pc, disp)
}

func (c *Processor) op194A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A2,(%d,A4)\n", pc, disp)
}

func (c *Processor) op194B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A3,(%d,A4)\n", pc, disp)
}

func (c *Processor) op194C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A4,(%d,A4)\n", pc, disp)
}

func (c *Processor) op194D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A5,(%d,A4)\n", pc, disp)
}

func (c *Processor) op194E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A6,(%d,A4)\n", pc, disp)
}

func (c *Processor) op194F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A7,(%d,A4)\n", pc, disp)
}

func (c *Processor) op1950() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A0),(%d,A4)\n", pc, disp)
}

func (c *Processor) op1951() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A1),(%d,A4)\n", pc, disp)
}

func (c *Processor) op1952() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A2),(%d,A4)\n", pc, disp)
}

func (c *Processor) op1953() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A3),(%d,A4)\n", pc, disp)
}

func (c *Processor) op1954() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A4),(%d,A4)\n", pc, disp)
}

func (c *Processor) op1955() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A5),(%d,A4)\n", pc, disp)
}

func (c *Processor) op1956() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A6),(%d,A4)\n", pc, disp)
}

func (c *Processor) op1957() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A7),(%d,A4)\n", pc, disp)
}

func (c *Processor) op1978() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A4)\n", pc, v, disp)
}

func (c *Processor) op1979() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A4)\n", pc, v, disp)
}

func (c *Processor) op197C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w #$%X,(%d,A4)\n", pc, v, disp)
}

func (c *Processor) op1A00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w D0,D5\n", pc)
}

func (c *Processor) op1A01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w D1,D5\n", pc)
}

func (c *Processor) op1A02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w D2,D5\n", pc)
}

func (c *Processor) op1A03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w D3,D5\n", pc)
}

func (c *Processor) op1A04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w D4,D5\n", pc)
}

func (c *Processor) op1A05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w D5,D5\n", pc)
}

func (c *Processor) op1A06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w D6,D5\n", pc)
}

func (c *Processor) op1A07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w D7,D5\n", pc)
}

func (c *Processor) op1A08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w A0,D5\n", pc)
}

func (c *Processor) op1A09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w A1,D5\n", pc)
}

func (c *Processor) op1A0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w A2,D5\n", pc)
}

func (c *Processor) op1A0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w A3,D5\n", pc)
}

func (c *Processor) op1A0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w A4,D5\n", pc)
}

func (c *Processor) op1A0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w A5,D5\n", pc)
}

func (c *Processor) op1A0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w A6,D5\n", pc)
}

func (c *Processor) op1A0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[5] = uint32(v)
	c.tracef("%04X move.w A7,D5\n", pc)
}

func (c *Processor) op1A10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[5] = uint32(v)
	c.tracef("%04X move.w (A0),D5\n", pc)
}

func (c *Processor) op1A11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[5] = uint32(v)
	c.tracef("%04X move.w (A1),D5\n", pc)
}

func (c *Processor) op1A12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[5] = uint32(v)
	c.tracef("%04X move.w (A2),D5\n", pc)
}

func (c *Processor) op1A13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[5] = uint32(v)
	c.tracef("%04X move.w (A3),D5\n", pc)
}

func (c *Processor) op1A14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[5] = uint32(v)
	c.tracef("%04X move.w (A4),D5\n", pc)
}

func (c *Processor) op1A15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[5] = uint32(v)
	c.tracef("%04X move.w (A5),D5\n", pc)
}

func (c *Processor) op1A16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[5] = uint32(v)
	c.tracef("%04X move.w (A6),D5\n", pc)
}

func (c *Processor) op1A17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[5] = uint32(v)
	c.tracef("%04X move.w (A7),D5\n", pc)
}

func (c *Processor) op1A38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[5] = uint32(v)
	c.tracef("%04X move.w $%X,D5\n", pc, v)
}

func (c *Processor) op1A39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[5] = uint32(v)
	c.tracef("%04X move.w $%X,D5\n", pc, v)
}

func (c *Processor) op1A3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[5] = uint32(v)
	c.tracef("%04X move.w #$%X,D5\n", pc, v)
}

func (c *Processor) op1A40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w D0,A5\n", pc)
}

func (c *Processor) op1A41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w D1,A5\n", pc)
}

func (c *Processor) op1A42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w D2,A5\n", pc)
}

func (c *Processor) op1A43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w D3,A5\n", pc)
}

func (c *Processor) op1A44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w D4,A5\n", pc)
}

func (c *Processor) op1A45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w D5,A5\n", pc)
}

func (c *Processor) op1A46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w D6,A5\n", pc)
}

func (c *Processor) op1A47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w D7,A5\n", pc)
}

func (c *Processor) op1A48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w A0,A5\n", pc)
}

func (c *Processor) op1A49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w A1,A5\n", pc)
}

func (c *Processor) op1A4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w A2,A5\n", pc)
}

func (c *Processor) op1A4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w A3,A5\n", pc)
}

func (c *Processor) op1A4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w A4,A5\n", pc)
}

func (c *Processor) op1A4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w A5,A5\n", pc)
}

func (c *Processor) op1A4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w A6,A5\n", pc)
}

func (c *Processor) op1A4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w A7,A5\n", pc)
}

func (c *Processor) op1A50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w (A0),A5\n", pc)
}

func (c *Processor) op1A51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w (A1),A5\n", pc)
}

func (c *Processor) op1A52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w (A2),A5\n", pc)
}

func (c *Processor) op1A53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w (A3),A5\n", pc)
}

func (c *Processor) op1A54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w (A4),A5\n", pc)
}

func (c *Processor) op1A55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w (A5),A5\n", pc)
}

func (c *Processor) op1A56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w (A6),A5\n", pc)
}

func (c *Processor) op1A57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w (A7),A5\n", pc)
}

func (c *Processor) op1A78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w $%X,A5\n", pc, v)
}

func (c *Processor) op1A79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w $%X,A5\n", pc, v)
}

func (c *Processor) op1A7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] = uint32(v)
	c.tracef("%04X movea.w #$%X,A5\n", pc, v)
}

func (c *Processor) op1A80() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D0,(A5)\n", pc)
}

func (c *Processor) op1A81() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D1,(A5)\n", pc)
}

func (c *Processor) op1A82() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D2,(A5)\n", pc)
}

func (c *Processor) op1A83() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D3,(A5)\n", pc)
}

func (c *Processor) op1A84() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D4,(A5)\n", pc)
}

func (c *Processor) op1A85() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D5,(A5)\n", pc)
}

func (c *Processor) op1A86() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D6,(A5)\n", pc)
}

func (c *Processor) op1A87() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D7,(A5)\n", pc)
}

func (c *Processor) op1A88() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A0,(A5)\n", pc)
}

func (c *Processor) op1A89() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A1,(A5)\n", pc)
}

func (c *Processor) op1A8A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A2,(A5)\n", pc)
}

func (c *Processor) op1A8B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A3,(A5)\n", pc)
}

func (c *Processor) op1A8C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A4,(A5)\n", pc)
}

func (c *Processor) op1A8D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A5,(A5)\n", pc)
}

func (c *Processor) op1A8E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A6,(A5)\n", pc)
}

func (c *Processor) op1A8F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A7,(A5)\n", pc)
}

func (c *Processor) op1A90() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A0),(A5)\n", pc)
}

func (c *Processor) op1A91() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A1),(A5)\n", pc)
}

func (c *Processor) op1A92() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A2),(A5)\n", pc)
}

func (c *Processor) op1A93() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A3),(A5)\n", pc)
}

func (c *Processor) op1A94() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A4),(A5)\n", pc)
}

func (c *Processor) op1A95() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A5),(A5)\n", pc)
}

func (c *Processor) op1A96() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A6),(A5)\n", pc)
}

func (c *Processor) op1A97() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A7),(A5)\n", pc)
}

func (c *Processor) op1AB8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w $%X,(A5)\n", pc, v)
}

func (c *Processor) op1AB9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w $%X,(A5)\n", pc, v)
}

func (c *Processor) op1ABC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w #$%X,(A5)\n", pc, v)
}

func (c *Processor) op1AC0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w D0,(A5)+\n", pc)
}

func (c *Processor) op1AC1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w D1,(A5)+\n", pc)
}

func (c *Processor) op1AC2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w D2,(A5)+\n", pc)
}

func (c *Processor) op1AC3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w D3,(A5)+\n", pc)
}

func (c *Processor) op1AC4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w D4,(A5)+\n", pc)
}

func (c *Processor) op1AC5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w D5,(A5)+\n", pc)
}

func (c *Processor) op1AC6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w D6,(A5)+\n", pc)
}

func (c *Processor) op1AC7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w D7,(A5)+\n", pc)
}

func (c *Processor) op1AC8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w A0,(A5)+\n", pc)
}

func (c *Processor) op1AC9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w A1,(A5)+\n", pc)
}

func (c *Processor) op1ACA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w A2,(A5)+\n", pc)
}

func (c *Processor) op1ACB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w A3,(A5)+\n", pc)
}

func (c *Processor) op1ACC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w A4,(A5)+\n", pc)
}

func (c *Processor) op1ACD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w A5,(A5)+\n", pc)
}

func (c *Processor) op1ACE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w A6,(A5)+\n", pc)
}

func (c *Processor) op1ACF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w A7,(A5)+\n", pc)
}

func (c *Processor) op1AD0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w (A0),(A5)+\n", pc)
}

func (c *Processor) op1AD1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w (A1),(A5)+\n", pc)
}

func (c *Processor) op1AD2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w (A2),(A5)+\n", pc)
}

func (c *Processor) op1AD3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w (A3),(A5)+\n", pc)
}

func (c *Processor) op1AD4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w (A4),(A5)+\n", pc)
}

func (c *Processor) op1AD5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w (A5),(A5)+\n", pc)
}

func (c *Processor) op1AD6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w (A6),(A5)+\n", pc)
}

func (c *Processor) op1AD7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w (A7),(A5)+\n", pc)
}

func (c *Processor) op1AF8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w $%X,(A5)+\n", pc, v)
}

func (c *Processor) op1AF9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w $%X,(A5)+\n", pc, v)
}

func (c *Processor) op1AFC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.w #$%X,(A5)+\n", pc, v)
}

func (c *Processor) op1B00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D0,-(A5)\n", pc)
}

func (c *Processor) op1B01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D1,-(A5)\n", pc)
}

func (c *Processor) op1B02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D2,-(A5)\n", pc)
}

func (c *Processor) op1B03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D3,-(A5)\n", pc)
}

func (c *Processor) op1B04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D4,-(A5)\n", pc)
}

func (c *Processor) op1B05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D5,-(A5)\n", pc)
}

func (c *Processor) op1B06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D6,-(A5)\n", pc)
}

func (c *Processor) op1B07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w D7,-(A5)\n", pc)
}

func (c *Processor) op1B08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A0,-(A5)\n", pc)
}

func (c *Processor) op1B09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A1,-(A5)\n", pc)
}

func (c *Processor) op1B0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A2,-(A5)\n", pc)
}

func (c *Processor) op1B0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A3,-(A5)\n", pc)
}

func (c *Processor) op1B0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A4,-(A5)\n", pc)
}

func (c *Processor) op1B0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A5,-(A5)\n", pc)
}

func (c *Processor) op1B0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A6,-(A5)\n", pc)
}

func (c *Processor) op1B0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w A7,-(A5)\n", pc)
}

func (c *Processor) op1B10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A0),-(A5)\n", pc)
}

func (c *Processor) op1B11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A1),-(A5)\n", pc)
}

func (c *Processor) op1B12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A2),-(A5)\n", pc)
}

func (c *Processor) op1B13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A3),-(A5)\n", pc)
}

func (c *Processor) op1B14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A4),-(A5)\n", pc)
}

func (c *Processor) op1B15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A5),-(A5)\n", pc)
}

func (c *Processor) op1B16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A6),-(A5)\n", pc)
}

func (c *Processor) op1B17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w (A7),-(A5)\n", pc)
}

func (c *Processor) op1B38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w $%X,-(A5)\n", pc, v)
}

func (c *Processor) op1B39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w $%X,-(A5)\n", pc, v)
}

func (c *Processor) op1B3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.w #$%X,-(A5)\n", pc, v)
}

func (c *Processor) op1B40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D0,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D1,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D2,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D3,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D4,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D5,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D6,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D7,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A0,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A1,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A2,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A3,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A4,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A5,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A6,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A7,(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A0),(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A1),(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A2),(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A3),(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A4),(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A5),(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A6),(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A7),(%d,A5)\n", pc, disp)
}

func (c *Processor) op1B78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A5)\n", pc, v, disp)
}

func (c *Processor) op1B79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A5)\n", pc, v, disp)
}

func (c *Processor) op1B7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w #$%X,(%d,A5)\n", pc, v, disp)
}

func (c *Processor) op1C00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w D0,D6\n", pc)
}

func (c *Processor) op1C01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w D1,D6\n", pc)
}

func (c *Processor) op1C02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w D2,D6\n", pc)
}

func (c *Processor) op1C03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w D3,D6\n", pc)
}

func (c *Processor) op1C04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w D4,D6\n", pc)
}

func (c *Processor) op1C05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w D5,D6\n", pc)
}

func (c *Processor) op1C06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w D6,D6\n", pc)
}

func (c *Processor) op1C07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w D7,D6\n", pc)
}

func (c *Processor) op1C08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w A0,D6\n", pc)
}

func (c *Processor) op1C09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w A1,D6\n", pc)
}

func (c *Processor) op1C0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w A2,D6\n", pc)
}

func (c *Processor) op1C0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w A3,D6\n", pc)
}

func (c *Processor) op1C0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w A4,D6\n", pc)
}

func (c *Processor) op1C0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w A5,D6\n", pc)
}

func (c *Processor) op1C0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w A6,D6\n", pc)
}

func (c *Processor) op1C0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[6] = uint32(v)
	c.tracef("%04X move.w A7,D6\n", pc)
}

func (c *Processor) op1C10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[6] = uint32(v)
	c.tracef("%04X move.w (A0),D6\n", pc)
}

func (c *Processor) op1C11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[6] = uint32(v)
	c.tracef("%04X move.w (A1),D6\n", pc)
}

func (c *Processor) op1C12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[6] = uint32(v)
	c.tracef("%04X move.w (A2),D6\n", pc)
}

func (c *Processor) op1C13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[6] = uint32(v)
	c.tracef("%04X move.w (A3),D6\n", pc)
}

func (c *Processor) op1C14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[6] = uint32(v)
	c.tracef("%04X move.w (A4),D6\n", pc)
}

func (c *Processor) op1C15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[6] = uint32(v)
	c.tracef("%04X move.w (A5),D6\n", pc)
}

func (c *Processor) op1C16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[6] = uint32(v)
	c.tracef("%04X move.w (A6),D6\n", pc)
}

func (c *Processor) op1C17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[6] = uint32(v)
	c.tracef("%04X move.w (A7),D6\n", pc)
}

func (c *Processor) op1C38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[6] = uint32(v)
	c.tracef("%04X move.w $%X,D6\n", pc, v)
}

func (c *Processor) op1C39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[6] = uint32(v)
	c.tracef("%04X move.w $%X,D6\n", pc, v)
}

func (c *Processor) op1C3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[6] = uint32(v)
	c.tracef("%04X move.w #$%X,D6\n", pc, v)
}

func (c *Processor) op1C40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w D0,A6\n", pc)
}

func (c *Processor) op1C41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w D1,A6\n", pc)
}

func (c *Processor) op1C42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w D2,A6\n", pc)
}

func (c *Processor) op1C43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w D3,A6\n", pc)
}

func (c *Processor) op1C44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w D4,A6\n", pc)
}

func (c *Processor) op1C45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w D5,A6\n", pc)
}

func (c *Processor) op1C46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w D6,A6\n", pc)
}

func (c *Processor) op1C47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w D7,A6\n", pc)
}

func (c *Processor) op1C48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w A0,A6\n", pc)
}

func (c *Processor) op1C49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w A1,A6\n", pc)
}

func (c *Processor) op1C4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w A2,A6\n", pc)
}

func (c *Processor) op1C4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w A3,A6\n", pc)
}

func (c *Processor) op1C4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w A4,A6\n", pc)
}

func (c *Processor) op1C4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w A5,A6\n", pc)
}

func (c *Processor) op1C4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w A6,A6\n", pc)
}

func (c *Processor) op1C4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w A7,A6\n", pc)
}

func (c *Processor) op1C50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w (A0),A6\n", pc)
}

func (c *Processor) op1C51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w (A1),A6\n", pc)
}

func (c *Processor) op1C52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w (A2),A6\n", pc)
}

func (c *Processor) op1C53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w (A3),A6\n", pc)
}

func (c *Processor) op1C54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w (A4),A6\n", pc)
}

func (c *Processor) op1C55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w (A5),A6\n", pc)
}

func (c *Processor) op1C56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w (A6),A6\n", pc)
}

func (c *Processor) op1C57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w (A7),A6\n", pc)
}

func (c *Processor) op1C78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w $%X,A6\n", pc, v)
}

func (c *Processor) op1C79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w $%X,A6\n", pc, v)
}

func (c *Processor) op1C7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] = uint32(v)
	c.tracef("%04X movea.w #$%X,A6\n", pc, v)
}

func (c *Processor) op1C80() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D0,(A6)\n", pc)
}

func (c *Processor) op1C81() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D1,(A6)\n", pc)
}

func (c *Processor) op1C82() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D2,(A6)\n", pc)
}

func (c *Processor) op1C83() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D3,(A6)\n", pc)
}

func (c *Processor) op1C84() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D4,(A6)\n", pc)
}

func (c *Processor) op1C85() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D5,(A6)\n", pc)
}

func (c *Processor) op1C86() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D6,(A6)\n", pc)
}

func (c *Processor) op1C87() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D7,(A6)\n", pc)
}

func (c *Processor) op1C88() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A0,(A6)\n", pc)
}

func (c *Processor) op1C89() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A1,(A6)\n", pc)
}

func (c *Processor) op1C8A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A2,(A6)\n", pc)
}

func (c *Processor) op1C8B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A3,(A6)\n", pc)
}

func (c *Processor) op1C8C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A4,(A6)\n", pc)
}

func (c *Processor) op1C8D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A5,(A6)\n", pc)
}

func (c *Processor) op1C8E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A6,(A6)\n", pc)
}

func (c *Processor) op1C8F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A7,(A6)\n", pc)
}

func (c *Processor) op1C90() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A0),(A6)\n", pc)
}

func (c *Processor) op1C91() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A1),(A6)\n", pc)
}

func (c *Processor) op1C92() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A2),(A6)\n", pc)
}

func (c *Processor) op1C93() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A3),(A6)\n", pc)
}

func (c *Processor) op1C94() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A4),(A6)\n", pc)
}

func (c *Processor) op1C95() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A5),(A6)\n", pc)
}

func (c *Processor) op1C96() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A6),(A6)\n", pc)
}

func (c *Processor) op1C97() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A7),(A6)\n", pc)
}

func (c *Processor) op1CB8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w $%X,(A6)\n", pc, v)
}

func (c *Processor) op1CB9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w $%X,(A6)\n", pc, v)
}

func (c *Processor) op1CBC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w #$%X,(A6)\n", pc, v)
}

func (c *Processor) op1CC0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w D0,(A6)+\n", pc)
}

func (c *Processor) op1CC1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w D1,(A6)+\n", pc)
}

func (c *Processor) op1CC2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w D2,(A6)+\n", pc)
}

func (c *Processor) op1CC3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w D3,(A6)+\n", pc)
}

func (c *Processor) op1CC4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w D4,(A6)+\n", pc)
}

func (c *Processor) op1CC5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w D5,(A6)+\n", pc)
}

func (c *Processor) op1CC6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w D6,(A6)+\n", pc)
}

func (c *Processor) op1CC7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w D7,(A6)+\n", pc)
}

func (c *Processor) op1CC8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w A0,(A6)+\n", pc)
}

func (c *Processor) op1CC9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w A1,(A6)+\n", pc)
}

func (c *Processor) op1CCA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w A2,(A6)+\n", pc)
}

func (c *Processor) op1CCB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w A3,(A6)+\n", pc)
}

func (c *Processor) op1CCC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w A4,(A6)+\n", pc)
}

func (c *Processor) op1CCD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w A5,(A6)+\n", pc)
}

func (c *Processor) op1CCE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w A6,(A6)+\n", pc)
}

func (c *Processor) op1CCF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w A7,(A6)+\n", pc)
}

func (c *Processor) op1CD0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w (A0),(A6)+\n", pc)
}

func (c *Processor) op1CD1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w (A1),(A6)+\n", pc)
}

func (c *Processor) op1CD2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w (A2),(A6)+\n", pc)
}

func (c *Processor) op1CD3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w (A3),(A6)+\n", pc)
}

func (c *Processor) op1CD4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w (A4),(A6)+\n", pc)
}

func (c *Processor) op1CD5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w (A5),(A6)+\n", pc)
}

func (c *Processor) op1CD6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w (A6),(A6)+\n", pc)
}

func (c *Processor) op1CD7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w (A7),(A6)+\n", pc)
}

func (c *Processor) op1CF8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w $%X,(A6)+\n", pc, v)
}

func (c *Processor) op1CF9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w $%X,(A6)+\n", pc, v)
}

func (c *Processor) op1CFC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.w #$%X,(A6)+\n", pc, v)
}

func (c *Processor) op1D00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D0,-(A6)\n", pc)
}

func (c *Processor) op1D01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D1,-(A6)\n", pc)
}

func (c *Processor) op1D02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D2,-(A6)\n", pc)
}

func (c *Processor) op1D03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D3,-(A6)\n", pc)
}

func (c *Processor) op1D04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D4,-(A6)\n", pc)
}

func (c *Processor) op1D05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D5,-(A6)\n", pc)
}

func (c *Processor) op1D06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D6,-(A6)\n", pc)
}

func (c *Processor) op1D07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w D7,-(A6)\n", pc)
}

func (c *Processor) op1D08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A0,-(A6)\n", pc)
}

func (c *Processor) op1D09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A1,-(A6)\n", pc)
}

func (c *Processor) op1D0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A2,-(A6)\n", pc)
}

func (c *Processor) op1D0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A3,-(A6)\n", pc)
}

func (c *Processor) op1D0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A4,-(A6)\n", pc)
}

func (c *Processor) op1D0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A5,-(A6)\n", pc)
}

func (c *Processor) op1D0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A6,-(A6)\n", pc)
}

func (c *Processor) op1D0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w A7,-(A6)\n", pc)
}

func (c *Processor) op1D10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A0),-(A6)\n", pc)
}

func (c *Processor) op1D11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A1),-(A6)\n", pc)
}

func (c *Processor) op1D12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A2),-(A6)\n", pc)
}

func (c *Processor) op1D13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A3),-(A6)\n", pc)
}

func (c *Processor) op1D14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A4),-(A6)\n", pc)
}

func (c *Processor) op1D15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A5),-(A6)\n", pc)
}

func (c *Processor) op1D16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A6),-(A6)\n", pc)
}

func (c *Processor) op1D17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w (A7),-(A6)\n", pc)
}

func (c *Processor) op1D38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w $%X,-(A6)\n", pc, v)
}

func (c *Processor) op1D39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w $%X,-(A6)\n", pc, v)
}

func (c *Processor) op1D3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.w #$%X,-(A6)\n", pc, v)
}

func (c *Processor) op1D40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D0,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D1,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D2,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D3,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D4,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D5,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D6,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D7,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A0,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A1,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A2,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A3,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A4,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A5,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A6,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A7,(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A0),(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A1),(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A2),(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A3),(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A4),(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A5),(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A6),(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A7),(%d,A6)\n", pc, disp)
}

func (c *Processor) op1D78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A6)\n", pc, v, disp)
}

func (c *Processor) op1D79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A6)\n", pc, v, disp)
}

func (c *Processor) op1D7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w #$%X,(%d,A6)\n", pc, v, disp)
}

func (c *Processor) op1E00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w D0,D7\n", pc)
}

func (c *Processor) op1E01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w D1,D7\n", pc)
}

func (c *Processor) op1E02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w D2,D7\n", pc)
}

func (c *Processor) op1E03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w D3,D7\n", pc)
}

func (c *Processor) op1E04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w D4,D7\n", pc)
}

func (c *Processor) op1E05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w D5,D7\n", pc)
}

func (c *Processor) op1E06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w D6,D7\n", pc)
}

func (c *Processor) op1E07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w D7,D7\n", pc)
}

func (c *Processor) op1E08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w A0,D7\n", pc)
}

func (c *Processor) op1E09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w A1,D7\n", pc)
}

func (c *Processor) op1E0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w A2,D7\n", pc)
}

func (c *Processor) op1E0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w A3,D7\n", pc)
}

func (c *Processor) op1E0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w A4,D7\n", pc)
}

func (c *Processor) op1E0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w A5,D7\n", pc)
}

func (c *Processor) op1E0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w A6,D7\n", pc)
}

func (c *Processor) op1E0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[7] = uint32(v)
	c.tracef("%04X move.w A7,D7\n", pc)
}

func (c *Processor) op1E10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[7] = uint32(v)
	c.tracef("%04X move.w (A0),D7\n", pc)
}

func (c *Processor) op1E11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[7] = uint32(v)
	c.tracef("%04X move.w (A1),D7\n", pc)
}

func (c *Processor) op1E12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[7] = uint32(v)
	c.tracef("%04X move.w (A2),D7\n", pc)
}

func (c *Processor) op1E13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[7] = uint32(v)
	c.tracef("%04X move.w (A3),D7\n", pc)
}

func (c *Processor) op1E14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[7] = uint32(v)
	c.tracef("%04X move.w (A4),D7\n", pc)
}

func (c *Processor) op1E15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[7] = uint32(v)
	c.tracef("%04X move.w (A5),D7\n", pc)
}

func (c *Processor) op1E16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[7] = uint32(v)
	c.tracef("%04X move.w (A6),D7\n", pc)
}

func (c *Processor) op1E17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.D[7] = uint32(v)
	c.tracef("%04X move.w (A7),D7\n", pc)
}

func (c *Processor) op1E38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[7] = uint32(v)
	c.tracef("%04X move.w $%X,D7\n", pc, v)
}

func (c *Processor) op1E39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[7] = uint32(v)
	c.tracef("%04X move.w $%X,D7\n", pc, v)
}

func (c *Processor) op1E3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[7] = uint32(v)
	c.tracef("%04X move.w #$%X,D7\n", pc, v)
}

func (c *Processor) op1E40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w D0,A7\n", pc)
}

func (c *Processor) op1E41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w D1,A7\n", pc)
}

func (c *Processor) op1E42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w D2,A7\n", pc)
}

func (c *Processor) op1E43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w D3,A7\n", pc)
}

func (c *Processor) op1E44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w D4,A7\n", pc)
}

func (c *Processor) op1E45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w D5,A7\n", pc)
}

func (c *Processor) op1E46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w D6,A7\n", pc)
}

func (c *Processor) op1E47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w D7,A7\n", pc)
}

func (c *Processor) op1E48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w A0,A7\n", pc)
}

func (c *Processor) op1E49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w A1,A7\n", pc)
}

func (c *Processor) op1E4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w A2,A7\n", pc)
}

func (c *Processor) op1E4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w A3,A7\n", pc)
}

func (c *Processor) op1E4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w A4,A7\n", pc)
}

func (c *Processor) op1E4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w A5,A7\n", pc)
}

func (c *Processor) op1E4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w A6,A7\n", pc)
}

func (c *Processor) op1E4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w A7,A7\n", pc)
}

func (c *Processor) op1E50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w (A0),A7\n", pc)
}

func (c *Processor) op1E51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w (A1),A7\n", pc)
}

func (c *Processor) op1E52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w (A2),A7\n", pc)
}

func (c *Processor) op1E53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w (A3),A7\n", pc)
}

func (c *Processor) op1E54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w (A4),A7\n", pc)
}

func (c *Processor) op1E55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w (A5),A7\n", pc)
}

func (c *Processor) op1E56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w (A6),A7\n", pc)
}

func (c *Processor) op1E57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w (A7),A7\n", pc)
}

func (c *Processor) op1E78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w $%X,A7\n", pc, v)
}

func (c *Processor) op1E79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w $%X,A7\n", pc, v)
}

func (c *Processor) op1E7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] = uint32(v)
	c.tracef("%04X movea.w #$%X,A7\n", pc, v)
}

func (c *Processor) op1E80() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D0,(A7)\n", pc)
}

func (c *Processor) op1E81() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D1,(A7)\n", pc)
}

func (c *Processor) op1E82() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D2,(A7)\n", pc)
}

func (c *Processor) op1E83() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D3,(A7)\n", pc)
}

func (c *Processor) op1E84() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D4,(A7)\n", pc)
}

func (c *Processor) op1E85() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D5,(A7)\n", pc)
}

func (c *Processor) op1E86() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D6,(A7)\n", pc)
}

func (c *Processor) op1E87() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D7,(A7)\n", pc)
}

func (c *Processor) op1E88() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A0,(A7)\n", pc)
}

func (c *Processor) op1E89() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A1,(A7)\n", pc)
}

func (c *Processor) op1E8A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A2,(A7)\n", pc)
}

func (c *Processor) op1E8B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A3,(A7)\n", pc)
}

func (c *Processor) op1E8C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A4,(A7)\n", pc)
}

func (c *Processor) op1E8D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A5,(A7)\n", pc)
}

func (c *Processor) op1E8E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A6,(A7)\n", pc)
}

func (c *Processor) op1E8F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A7,(A7)\n", pc)
}

func (c *Processor) op1E90() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A0),(A7)\n", pc)
}

func (c *Processor) op1E91() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A1),(A7)\n", pc)
}

func (c *Processor) op1E92() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A2),(A7)\n", pc)
}

func (c *Processor) op1E93() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A3),(A7)\n", pc)
}

func (c *Processor) op1E94() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A4),(A7)\n", pc)
}

func (c *Processor) op1E95() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A5),(A7)\n", pc)
}

func (c *Processor) op1E96() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A6),(A7)\n", pc)
}

func (c *Processor) op1E97() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A7),(A7)\n", pc)
}

func (c *Processor) op1EB8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w $%X,(A7)\n", pc, v)
}

func (c *Processor) op1EB9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w $%X,(A7)\n", pc, v)
}

func (c *Processor) op1EBC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w #$%X,(A7)\n", pc, v)
}

func (c *Processor) op1EC0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w D0,(A7)+\n", pc)
}

func (c *Processor) op1EC1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w D1,(A7)+\n", pc)
}

func (c *Processor) op1EC2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w D2,(A7)+\n", pc)
}

func (c *Processor) op1EC3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w D3,(A7)+\n", pc)
}

func (c *Processor) op1EC4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w D4,(A7)+\n", pc)
}

func (c *Processor) op1EC5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w D5,(A7)+\n", pc)
}

func (c *Processor) op1EC6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w D6,(A7)+\n", pc)
}

func (c *Processor) op1EC7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w D7,(A7)+\n", pc)
}

func (c *Processor) op1EC8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w A0,(A7)+\n", pc)
}

func (c *Processor) op1EC9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w A1,(A7)+\n", pc)
}

func (c *Processor) op1ECA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w A2,(A7)+\n", pc)
}

func (c *Processor) op1ECB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w A3,(A7)+\n", pc)
}

func (c *Processor) op1ECC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w A4,(A7)+\n", pc)
}

func (c *Processor) op1ECD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w A5,(A7)+\n", pc)
}

func (c *Processor) op1ECE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w A6,(A7)+\n", pc)
}

func (c *Processor) op1ECF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w A7,(A7)+\n", pc)
}

func (c *Processor) op1ED0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w (A0),(A7)+\n", pc)
}

func (c *Processor) op1ED1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w (A1),(A7)+\n", pc)
}

func (c *Processor) op1ED2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w (A2),(A7)+\n", pc)
}

func (c *Processor) op1ED3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w (A3),(A7)+\n", pc)
}

func (c *Processor) op1ED4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w (A4),(A7)+\n", pc)
}

func (c *Processor) op1ED5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w (A5),(A7)+\n", pc)
}

func (c *Processor) op1ED6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w (A6),(A7)+\n", pc)
}

func (c *Processor) op1ED7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w (A7),(A7)+\n", pc)
}

func (c *Processor) op1EF8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w $%X,(A7)+\n", pc, v)
}

func (c *Processor) op1EF9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w $%X,(A7)+\n", pc, v)
}

func (c *Processor) op1EFC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.w #$%X,(A7)+\n", pc, v)
}

func (c *Processor) op1F00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D0,-(A7)\n", pc)
}

func (c *Processor) op1F01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D1,-(A7)\n", pc)
}

func (c *Processor) op1F02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D2,-(A7)\n", pc)
}

func (c *Processor) op1F03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D3,-(A7)\n", pc)
}

func (c *Processor) op1F04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D4,-(A7)\n", pc)
}

func (c *Processor) op1F05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D5,-(A7)\n", pc)
}

func (c *Processor) op1F06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D6,-(A7)\n", pc)
}

func (c *Processor) op1F07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w D7,-(A7)\n", pc)
}

func (c *Processor) op1F08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A0,-(A7)\n", pc)
}

func (c *Processor) op1F09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A1,-(A7)\n", pc)
}

func (c *Processor) op1F0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A2,-(A7)\n", pc)
}

func (c *Processor) op1F0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A3,-(A7)\n", pc)
}

func (c *Processor) op1F0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A4,-(A7)\n", pc)
}

func (c *Processor) op1F0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A5,-(A7)\n", pc)
}

func (c *Processor) op1F0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A6,-(A7)\n", pc)
}

func (c *Processor) op1F0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w A7,-(A7)\n", pc)
}

func (c *Processor) op1F10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A0),-(A7)\n", pc)
}

func (c *Processor) op1F11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A1),-(A7)\n", pc)
}

func (c *Processor) op1F12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A2),-(A7)\n", pc)
}

func (c *Processor) op1F13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A3),-(A7)\n", pc)
}

func (c *Processor) op1F14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A4),-(A7)\n", pc)
}

func (c *Processor) op1F15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A5),-(A7)\n", pc)
}

func (c *Processor) op1F16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A6),-(A7)\n", pc)
}

func (c *Processor) op1F17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w (A7),-(A7)\n", pc)
}

func (c *Processor) op1F38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w $%X,-(A7)\n", pc, v)
}

func (c *Processor) op1F39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w $%X,-(A7)\n", pc, v)
}

func (c *Processor) op1F3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.w #$%X,-(A7)\n", pc, v)
}

func (c *Processor) op1F40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D0,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D1,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D2,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D3,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D4,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D5,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D6,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w D7,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A0,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A1,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A2,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A3,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A4,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A5,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A6,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w A7,(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A0),(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A1),(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A2),(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A3),(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A4),(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A5),(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A6),(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:2])
	if c.err != nil {
		return
	}
	v := uint16(c.buf[0])<<8 | uint16(c.buf[1])
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w (A7),(%d,A7)\n", pc, disp)
}

func (c *Processor) op1F78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A7)\n", pc, v, disp)
}

func (c *Processor) op1F79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w $%X,(%d,A7)\n", pc, v, disp)
}

func (c *Processor) op1F7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.w #$%X,(%d,A7)\n", pc, v, disp)
}

func (c *Processor) op2000() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l D0,D0\n", pc)
}

func (c *Processor) op2001() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l D1,D0\n", pc)
}

func (c *Processor) op2002() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l D2,D0\n", pc)
}

func (c *Processor) op2003() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l D3,D0\n", pc)
}

func (c *Processor) op2004() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l D4,D0\n", pc)
}

func (c *Processor) op2005() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l D5,D0\n", pc)
}

func (c *Processor) op2006() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l D6,D0\n", pc)
}

func (c *Processor) op2007() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l D7,D0\n", pc)
}

func (c *Processor) op2008() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l A0,D0\n", pc)
}

func (c *Processor) op2009() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l A1,D0\n", pc)
}

func (c *Processor) op200A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l A2,D0\n", pc)
}

func (c *Processor) op200B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l A3,D0\n", pc)
}

func (c *Processor) op200C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l A4,D0\n", pc)
}

func (c *Processor) op200D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l A5,D0\n", pc)
}

func (c *Processor) op200E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l A6,D0\n", pc)
}

func (c *Processor) op200F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[0] = uint32(v)
	c.tracef("%04X move.l A7,D0\n", pc)
}

func (c *Processor) op2010() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[0] = uint32(v)
	c.tracef("%04X move.l (A0),D0\n", pc)
}

func (c *Processor) op2011() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[0] = uint32(v)
	c.tracef("%04X move.l (A1),D0\n", pc)
}

func (c *Processor) op2012() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[0] = uint32(v)
	c.tracef("%04X move.l (A2),D0\n", pc)
}

func (c *Processor) op2013() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[0] = uint32(v)
	c.tracef("%04X move.l (A3),D0\n", pc)
}

func (c *Processor) op2014() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[0] = uint32(v)
	c.tracef("%04X move.l (A4),D0\n", pc)
}

func (c *Processor) op2015() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[0] = uint32(v)
	c.tracef("%04X move.l (A5),D0\n", pc)
}

func (c *Processor) op2016() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[0] = uint32(v)
	c.tracef("%04X move.l (A6),D0\n", pc)
}

func (c *Processor) op2017() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[0] = uint32(v)
	c.tracef("%04X move.l (A7),D0\n", pc)
}

func (c *Processor) op2038() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[0] = uint32(v)
	c.tracef("%04X move.l $%X,D0\n", pc, v)
}

func (c *Processor) op2039() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[0] = uint32(v)
	c.tracef("%04X move.l $%X,D0\n", pc, v)
}

func (c *Processor) op203C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[0] = uint32(v)
	c.tracef("%04X move.l #$%X,D0\n", pc, v)
}

func (c *Processor) op2040() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l D0,A0\n", pc)
}

func (c *Processor) op2041() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l D1,A0\n", pc)
}

func (c *Processor) op2042() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l D2,A0\n", pc)
}

func (c *Processor) op2043() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l D3,A0\n", pc)
}

func (c *Processor) op2044() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l D4,A0\n", pc)
}

func (c *Processor) op2045() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l D5,A0\n", pc)
}

func (c *Processor) op2046() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l D6,A0\n", pc)
}

func (c *Processor) op2047() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l D7,A0\n", pc)
}

func (c *Processor) op2048() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l A0,A0\n", pc)
}

func (c *Processor) op2049() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l A1,A0\n", pc)
}

func (c *Processor) op204A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l A2,A0\n", pc)
}

func (c *Processor) op204B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l A3,A0\n", pc)
}

func (c *Processor) op204C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l A4,A0\n", pc)
}

func (c *Processor) op204D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l A5,A0\n", pc)
}

func (c *Processor) op204E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l A6,A0\n", pc)
}

func (c *Processor) op204F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l A7,A0\n", pc)
}

func (c *Processor) op2050() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l (A0),A0\n", pc)
}

func (c *Processor) op2051() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l (A1),A0\n", pc)
}

func (c *Processor) op2052() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l (A2),A0\n", pc)
}

func (c *Processor) op2053() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l (A3),A0\n", pc)
}

func (c *Processor) op2054() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l (A4),A0\n", pc)
}

func (c *Processor) op2055() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l (A5),A0\n", pc)
}

func (c *Processor) op2056() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l (A6),A0\n", pc)
}

func (c *Processor) op2057() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l (A7),A0\n", pc)
}

func (c *Processor) op2078() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l $%X,A0\n", pc, v)
}

func (c *Processor) op2079() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l $%X,A0\n", pc, v)
}

func (c *Processor) op207C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] = uint32(v)
	c.tracef("%04X movea.l #$%X,A0\n", pc, v)
}

func (c *Processor) op2080() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D0,(A0)\n", pc)
}

func (c *Processor) op2081() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D1,(A0)\n", pc)
}

func (c *Processor) op2082() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D2,(A0)\n", pc)
}

func (c *Processor) op2083() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D3,(A0)\n", pc)
}

func (c *Processor) op2084() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D4,(A0)\n", pc)
}

func (c *Processor) op2085() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D5,(A0)\n", pc)
}

func (c *Processor) op2086() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D6,(A0)\n", pc)
}

func (c *Processor) op2087() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D7,(A0)\n", pc)
}

func (c *Processor) op2088() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A0,(A0)\n", pc)
}

func (c *Processor) op2089() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A1,(A0)\n", pc)
}

func (c *Processor) op208A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A2,(A0)\n", pc)
}

func (c *Processor) op208B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A3,(A0)\n", pc)
}

func (c *Processor) op208C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A4,(A0)\n", pc)
}

func (c *Processor) op208D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A5,(A0)\n", pc)
}

func (c *Processor) op208E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A6,(A0)\n", pc)
}

func (c *Processor) op208F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A7,(A0)\n", pc)
}

func (c *Processor) op2090() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A0),(A0)\n", pc)
}

func (c *Processor) op2091() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A1),(A0)\n", pc)
}

func (c *Processor) op2092() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A2),(A0)\n", pc)
}

func (c *Processor) op2093() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A3),(A0)\n", pc)
}

func (c *Processor) op2094() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A4),(A0)\n", pc)
}

func (c *Processor) op2095() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A5),(A0)\n", pc)
}

func (c *Processor) op2096() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A6),(A0)\n", pc)
}

func (c *Processor) op2097() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A7),(A0)\n", pc)
}

func (c *Processor) op20B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l $%X,(A0)\n", pc, v)
}

func (c *Processor) op20B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l $%X,(A0)\n", pc, v)
}

func (c *Processor) op20BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l #$%X,(A0)\n", pc, v)
}

func (c *Processor) op20C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l D0,(A0)+\n", pc)
}

func (c *Processor) op20C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l D1,(A0)+\n", pc)
}

func (c *Processor) op20C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l D2,(A0)+\n", pc)
}

func (c *Processor) op20C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l D3,(A0)+\n", pc)
}

func (c *Processor) op20C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l D4,(A0)+\n", pc)
}

func (c *Processor) op20C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l D5,(A0)+\n", pc)
}

func (c *Processor) op20C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l D6,(A0)+\n", pc)
}

func (c *Processor) op20C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l D7,(A0)+\n", pc)
}

func (c *Processor) op20C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l A0,(A0)+\n", pc)
}

func (c *Processor) op20C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l A1,(A0)+\n", pc)
}

func (c *Processor) op20CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l A2,(A0)+\n", pc)
}

func (c *Processor) op20CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l A3,(A0)+\n", pc)
}

func (c *Processor) op20CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l A4,(A0)+\n", pc)
}

func (c *Processor) op20CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l A5,(A0)+\n", pc)
}

func (c *Processor) op20CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l A6,(A0)+\n", pc)
}

func (c *Processor) op20CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l A7,(A0)+\n", pc)
}

func (c *Processor) op20D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l (A0),(A0)+\n", pc)
}

func (c *Processor) op20D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l (A1),(A0)+\n", pc)
}

func (c *Processor) op20D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l (A2),(A0)+\n", pc)
}

func (c *Processor) op20D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l (A3),(A0)+\n", pc)
}

func (c *Processor) op20D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l (A4),(A0)+\n", pc)
}

func (c *Processor) op20D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l (A5),(A0)+\n", pc)
}

func (c *Processor) op20D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l (A6),(A0)+\n", pc)
}

func (c *Processor) op20D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l (A7),(A0)+\n", pc)
}

func (c *Processor) op20F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l $%X,(A0)+\n", pc, v)
}

func (c *Processor) op20F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l $%X,(A0)+\n", pc, v)
}

func (c *Processor) op20FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[0], uint32(v))
	c.A[0] += 4
	c.tracef("%04X move.l #$%X,(A0)+\n", pc, v)
}

func (c *Processor) op2100() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D0,-(A0)\n", pc)
}

func (c *Processor) op2101() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D1,-(A0)\n", pc)
}

func (c *Processor) op2102() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D2,-(A0)\n", pc)
}

func (c *Processor) op2103() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D3,-(A0)\n", pc)
}

func (c *Processor) op2104() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D4,-(A0)\n", pc)
}

func (c *Processor) op2105() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D5,-(A0)\n", pc)
}

func (c *Processor) op2106() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D6,-(A0)\n", pc)
}

func (c *Processor) op2107() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l D7,-(A0)\n", pc)
}

func (c *Processor) op2108() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A0,-(A0)\n", pc)
}

func (c *Processor) op2109() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A1,-(A0)\n", pc)
}

func (c *Processor) op210A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A2,-(A0)\n", pc)
}

func (c *Processor) op210B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A3,-(A0)\n", pc)
}

func (c *Processor) op210C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A4,-(A0)\n", pc)
}

func (c *Processor) op210D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A5,-(A0)\n", pc)
}

func (c *Processor) op210E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A6,-(A0)\n", pc)
}

func (c *Processor) op210F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l A7,-(A0)\n", pc)
}

func (c *Processor) op2110() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A0),-(A0)\n", pc)
}

func (c *Processor) op2111() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A1),-(A0)\n", pc)
}

func (c *Processor) op2112() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A2),-(A0)\n", pc)
}

func (c *Processor) op2113() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A3),-(A0)\n", pc)
}

func (c *Processor) op2114() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A4),-(A0)\n", pc)
}

func (c *Processor) op2115() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A5),-(A0)\n", pc)
}

func (c *Processor) op2116() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A6),-(A0)\n", pc)
}

func (c *Processor) op2117() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l (A7),-(A0)\n", pc)
}

func (c *Processor) op2138() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l $%X,-(A0)\n", pc, v)
}

func (c *Processor) op2139() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l $%X,-(A0)\n", pc, v)
}

func (c *Processor) op213C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[0] -= 4
	c.writeLong(c.A[0], uint32(v))
	c.tracef("%04X move.l #$%X,-(A0)\n", pc, v)
}

func (c *Processor) op2140() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D0,(%d,A0)\n", pc, disp)
}

func (c *Processor) op2141() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D1,(%d,A0)\n", pc, disp)
}

func (c *Processor) op2142() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D2,(%d,A0)\n", pc, disp)
}

func (c *Processor) op2143() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D3,(%d,A0)\n", pc, disp)
}

func (c *Processor) op2144() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D4,(%d,A0)\n", pc, disp)
}

func (c *Processor) op2145() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D5,(%d,A0)\n", pc, disp)
}

func (c *Processor) op2146() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D6,(%d,A0)\n", pc, disp)
}

func (c *Processor) op2147() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D7,(%d,A0)\n", pc, disp)
}

func (c *Processor) op2148() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A0,(%d,A0)\n", pc, disp)
}

func (c *Processor) op2149() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A1,(%d,A0)\n", pc, disp)
}

func (c *Processor) op214A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A2,(%d,A0)\n", pc, disp)
}

func (c *Processor) op214B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A3,(%d,A0)\n", pc, disp)
}

func (c *Processor) op214C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A4,(%d,A0)\n", pc, disp)
}

func (c *Processor) op214D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A5,(%d,A0)\n", pc, disp)
}

func (c *Processor) op214E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A6,(%d,A0)\n", pc, disp)
}

func (c *Processor) op214F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A7,(%d,A0)\n", pc, disp)
}

func (c *Processor) op2150() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A0),(%d,A0)\n", pc, disp)
}

func (c *Processor) op2151() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A1),(%d,A0)\n", pc, disp)
}

func (c *Processor) op2152() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A2),(%d,A0)\n", pc, disp)
}

func (c *Processor) op2153() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A3),(%d,A0)\n", pc, disp)
}

func (c *Processor) op2154() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A4),(%d,A0)\n", pc, disp)
}

func (c *Processor) op2155() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A5),(%d,A0)\n", pc, disp)
}

func (c *Processor) op2156() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A6),(%d,A0)\n", pc, disp)
}

func (c *Processor) op2157() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A7),(%d,A0)\n", pc, disp)
}

func (c *Processor) op2178() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A0)\n", pc, v, disp)
}

func (c *Processor) op2179() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A0)\n", pc, v, disp)
}

func (c *Processor) op217C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[0]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l #$%X,(%d,A0)\n", pc, v, disp)
}

func (c *Processor) op21C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D0,$%X\n", pc, addr)
}

func (c *Processor) op21C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D1,$%X\n", pc, addr)
}

func (c *Processor) op21C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D2,$%X\n", pc, addr)
}

func (c *Processor) op21C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D3,$%X\n", pc, addr)
}

func (c *Processor) op21C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D4,$%X\n", pc, addr)
}

func (c *Processor) op21C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D5,$%X\n", pc, addr)
}

func (c *Processor) op21C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D6,$%X\n", pc, addr)
}

func (c *Processor) op21C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D7,$%X\n", pc, addr)
}

func (c *Processor) op21C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A0,$%X\n", pc, addr)
}

func (c *Processor) op21C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A1,$%X\n", pc, addr)
}

func (c *Processor) op21CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A2,$%X\n", pc, addr)
}

func (c *Processor) op21CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A3,$%X\n", pc, addr)
}

func (c *Processor) op21CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A4,$%X\n", pc, addr)
}

func (c *Processor) op21CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A5,$%X\n", pc, addr)
}

func (c *Processor) op21CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A6,$%X\n", pc, addr)
}

func (c *Processor) op21CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A7,$%X\n", pc, addr)
}

func (c *Processor) op21D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A0),$%X\n", pc, addr)
}

func (c *Processor) op21D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A1),$%X\n", pc, addr)
}

func (c *Processor) op21D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A2),$%X\n", pc, addr)
}

func (c *Processor) op21D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A3),$%X\n", pc, addr)
}

func (c *Processor) op21D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A4),$%X\n", pc, addr)
}

func (c *Processor) op21D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A5),$%X\n", pc, addr)
}

func (c *Processor) op21D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A6),$%X\n", pc, addr)
}

func (c *Processor) op21D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A7),$%X\n", pc, addr)
}

func (c *Processor) op21F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,$%X\n", pc, v, addr)
}

func (c *Processor) op21F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,$%X\n", pc, v, addr)
}

func (c *Processor) op21FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x01)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l #$%X,$%X\n", pc, v, addr)
}

func (c *Processor) op2200() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l D0,D1\n", pc)
}

func (c *Processor) op2201() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l D1,D1\n", pc)
}

func (c *Processor) op2202() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l D2,D1\n", pc)
}

func (c *Processor) op2203() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l D3,D1\n", pc)
}

func (c *Processor) op2204() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l D4,D1\n", pc)
}

func (c *Processor) op2205() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l D5,D1\n", pc)
}

func (c *Processor) op2206() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l D6,D1\n", pc)
}

func (c *Processor) op2207() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l D7,D1\n", pc)
}

func (c *Processor) op2208() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l A0,D1\n", pc)
}

func (c *Processor) op2209() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l A1,D1\n", pc)
}

func (c *Processor) op220A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l A2,D1\n", pc)
}

func (c *Processor) op220B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l A3,D1\n", pc)
}

func (c *Processor) op220C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l A4,D1\n", pc)
}

func (c *Processor) op220D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l A5,D1\n", pc)
}

func (c *Processor) op220E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l A6,D1\n", pc)
}

func (c *Processor) op220F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[1] = uint32(v)
	c.tracef("%04X move.l A7,D1\n", pc)
}

func (c *Processor) op2210() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[1] = uint32(v)
	c.tracef("%04X move.l (A0),D1\n", pc)
}

func (c *Processor) op2211() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[1] = uint32(v)
	c.tracef("%04X move.l (A1),D1\n", pc)
}

func (c *Processor) op2212() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[1] = uint32(v)
	c.tracef("%04X move.l (A2),D1\n", pc)
}

func (c *Processor) op2213() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[1] = uint32(v)
	c.tracef("%04X move.l (A3),D1\n", pc)
}

func (c *Processor) op2214() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[1] = uint32(v)
	c.tracef("%04X move.l (A4),D1\n", pc)
}

func (c *Processor) op2215() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[1] = uint32(v)
	c.tracef("%04X move.l (A5),D1\n", pc)
}

func (c *Processor) op2216() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[1] = uint32(v)
	c.tracef("%04X move.l (A6),D1\n", pc)
}

func (c *Processor) op2217() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[1] = uint32(v)
	c.tracef("%04X move.l (A7),D1\n", pc)
}

func (c *Processor) op2238() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[1] = uint32(v)
	c.tracef("%04X move.l $%X,D1\n", pc, v)
}

func (c *Processor) op2239() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[1] = uint32(v)
	c.tracef("%04X move.l $%X,D1\n", pc, v)
}

func (c *Processor) op223C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[1] = uint32(v)
	c.tracef("%04X move.l #$%X,D1\n", pc, v)
}

func (c *Processor) op2240() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l D0,A1\n", pc)
}

func (c *Processor) op2241() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l D1,A1\n", pc)
}

func (c *Processor) op2242() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l D2,A1\n", pc)
}

func (c *Processor) op2243() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l D3,A1\n", pc)
}

func (c *Processor) op2244() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l D4,A1\n", pc)
}

func (c *Processor) op2245() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l D5,A1\n", pc)
}

func (c *Processor) op2246() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l D6,A1\n", pc)
}

func (c *Processor) op2247() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l D7,A1\n", pc)
}

func (c *Processor) op2248() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l A0,A1\n", pc)
}

func (c *Processor) op2249() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l A1,A1\n", pc)
}

func (c *Processor) op224A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l A2,A1\n", pc)
}

func (c *Processor) op224B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l A3,A1\n", pc)
}

func (c *Processor) op224C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l A4,A1\n", pc)
}

func (c *Processor) op224D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l A5,A1\n", pc)
}

func (c *Processor) op224E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l A6,A1\n", pc)
}

func (c *Processor) op224F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l A7,A1\n", pc)
}

func (c *Processor) op2250() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l (A0),A1\n", pc)
}

func (c *Processor) op2251() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l (A1),A1\n", pc)
}

func (c *Processor) op2252() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l (A2),A1\n", pc)
}

func (c *Processor) op2253() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l (A3),A1\n", pc)
}

func (c *Processor) op2254() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l (A4),A1\n", pc)
}

func (c *Processor) op2255() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l (A5),A1\n", pc)
}

func (c *Processor) op2256() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l (A6),A1\n", pc)
}

func (c *Processor) op2257() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l (A7),A1\n", pc)
}

func (c *Processor) op2278() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l $%X,A1\n", pc, v)
}

func (c *Processor) op2279() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l $%X,A1\n", pc, v)
}

func (c *Processor) op227C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] = uint32(v)
	c.tracef("%04X movea.l #$%X,A1\n", pc, v)
}

func (c *Processor) op2280() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D0,(A1)\n", pc)
}

func (c *Processor) op2281() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D1,(A1)\n", pc)
}

func (c *Processor) op2282() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D2,(A1)\n", pc)
}

func (c *Processor) op2283() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D3,(A1)\n", pc)
}

func (c *Processor) op2284() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D4,(A1)\n", pc)
}

func (c *Processor) op2285() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D5,(A1)\n", pc)
}

func (c *Processor) op2286() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D6,(A1)\n", pc)
}

func (c *Processor) op2287() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D7,(A1)\n", pc)
}

func (c *Processor) op2288() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A0,(A1)\n", pc)
}

func (c *Processor) op2289() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A1,(A1)\n", pc)
}

func (c *Processor) op228A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A2,(A1)\n", pc)
}

func (c *Processor) op228B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A3,(A1)\n", pc)
}

func (c *Processor) op228C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A4,(A1)\n", pc)
}

func (c *Processor) op228D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A5,(A1)\n", pc)
}

func (c *Processor) op228E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A6,(A1)\n", pc)
}

func (c *Processor) op228F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A7,(A1)\n", pc)
}

func (c *Processor) op2290() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A0),(A1)\n", pc)
}

func (c *Processor) op2291() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A1),(A1)\n", pc)
}

func (c *Processor) op2292() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A2),(A1)\n", pc)
}

func (c *Processor) op2293() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A3),(A1)\n", pc)
}

func (c *Processor) op2294() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A4),(A1)\n", pc)
}

func (c *Processor) op2295() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A5),(A1)\n", pc)
}

func (c *Processor) op2296() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A6),(A1)\n", pc)
}

func (c *Processor) op2297() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A7),(A1)\n", pc)
}

func (c *Processor) op22B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l $%X,(A1)\n", pc, v)
}

func (c *Processor) op22B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l $%X,(A1)\n", pc, v)
}

func (c *Processor) op22BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l #$%X,(A1)\n", pc, v)
}

func (c *Processor) op22C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l D0,(A1)+\n", pc)
}

func (c *Processor) op22C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l D1,(A1)+\n", pc)
}

func (c *Processor) op22C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l D2,(A1)+\n", pc)
}

func (c *Processor) op22C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l D3,(A1)+\n", pc)
}

func (c *Processor) op22C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l D4,(A1)+\n", pc)
}

func (c *Processor) op22C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l D5,(A1)+\n", pc)
}

func (c *Processor) op22C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l D6,(A1)+\n", pc)
}

func (c *Processor) op22C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l D7,(A1)+\n", pc)
}

func (c *Processor) op22C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l A0,(A1)+\n", pc)
}

func (c *Processor) op22C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l A1,(A1)+\n", pc)
}

func (c *Processor) op22CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l A2,(A1)+\n", pc)
}

func (c *Processor) op22CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l A3,(A1)+\n", pc)
}

func (c *Processor) op22CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l A4,(A1)+\n", pc)
}

func (c *Processor) op22CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l A5,(A1)+\n", pc)
}

func (c *Processor) op22CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l A6,(A1)+\n", pc)
}

func (c *Processor) op22CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l A7,(A1)+\n", pc)
}

func (c *Processor) op22D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l (A0),(A1)+\n", pc)
}

func (c *Processor) op22D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l (A1),(A1)+\n", pc)
}

func (c *Processor) op22D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l (A2),(A1)+\n", pc)
}

func (c *Processor) op22D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l (A3),(A1)+\n", pc)
}

func (c *Processor) op22D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l (A4),(A1)+\n", pc)
}

func (c *Processor) op22D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l (A5),(A1)+\n", pc)
}

func (c *Processor) op22D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l (A6),(A1)+\n", pc)
}

func (c *Processor) op22D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l (A7),(A1)+\n", pc)
}

func (c *Processor) op22F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l $%X,(A1)+\n", pc, v)
}

func (c *Processor) op22F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l $%X,(A1)+\n", pc, v)
}

func (c *Processor) op22FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[1], uint32(v))
	c.A[1] += 4
	c.tracef("%04X move.l #$%X,(A1)+\n", pc, v)
}

func (c *Processor) op2300() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D0,-(A1)\n", pc)
}

func (c *Processor) op2301() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D1,-(A1)\n", pc)
}

func (c *Processor) op2302() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D2,-(A1)\n", pc)
}

func (c *Processor) op2303() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D3,-(A1)\n", pc)
}

func (c *Processor) op2304() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D4,-(A1)\n", pc)
}

func (c *Processor) op2305() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D5,-(A1)\n", pc)
}

func (c *Processor) op2306() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D6,-(A1)\n", pc)
}

func (c *Processor) op2307() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l D7,-(A1)\n", pc)
}

func (c *Processor) op2308() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A0,-(A1)\n", pc)
}

func (c *Processor) op2309() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A1,-(A1)\n", pc)
}

func (c *Processor) op230A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A2,-(A1)\n", pc)
}

func (c *Processor) op230B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A3,-(A1)\n", pc)
}

func (c *Processor) op230C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A4,-(A1)\n", pc)
}

func (c *Processor) op230D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A5,-(A1)\n", pc)
}

func (c *Processor) op230E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A6,-(A1)\n", pc)
}

func (c *Processor) op230F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l A7,-(A1)\n", pc)
}

func (c *Processor) op2310() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A0),-(A1)\n", pc)
}

func (c *Processor) op2311() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A1),-(A1)\n", pc)
}

func (c *Processor) op2312() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A2),-(A1)\n", pc)
}

func (c *Processor) op2313() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A3),-(A1)\n", pc)
}

func (c *Processor) op2314() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A4),-(A1)\n", pc)
}

func (c *Processor) op2315() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A5),-(A1)\n", pc)
}

func (c *Processor) op2316() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A6),-(A1)\n", pc)
}

func (c *Processor) op2317() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l (A7),-(A1)\n", pc)
}

func (c *Processor) op2338() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l $%X,-(A1)\n", pc, v)
}

func (c *Processor) op2339() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l $%X,-(A1)\n", pc, v)
}

func (c *Processor) op233C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[1] -= 4
	c.writeLong(c.A[1], uint32(v))
	c.tracef("%04X move.l #$%X,-(A1)\n", pc, v)
}

func (c *Processor) op2340() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D0,(%d,A1)\n", pc, disp)
}

func (c *Processor) op2341() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D1,(%d,A1)\n", pc, disp)
}

func (c *Processor) op2342() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D2,(%d,A1)\n", pc, disp)
}

func (c *Processor) op2343() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D3,(%d,A1)\n", pc, disp)
}

func (c *Processor) op2344() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D4,(%d,A1)\n", pc, disp)
}

func (c *Processor) op2345() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D5,(%d,A1)\n", pc, disp)
}

func (c *Processor) op2346() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D6,(%d,A1)\n", pc, disp)
}

func (c *Processor) op2347() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D7,(%d,A1)\n", pc, disp)
}

func (c *Processor) op2348() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A0,(%d,A1)\n", pc, disp)
}

func (c *Processor) op2349() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A1,(%d,A1)\n", pc, disp)
}

func (c *Processor) op234A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A2,(%d,A1)\n", pc, disp)
}

func (c *Processor) op234B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A3,(%d,A1)\n", pc, disp)
}

func (c *Processor) op234C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A4,(%d,A1)\n", pc, disp)
}

func (c *Processor) op234D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A5,(%d,A1)\n", pc, disp)
}

func (c *Processor) op234E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A6,(%d,A1)\n", pc, disp)
}

func (c *Processor) op234F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A7,(%d,A1)\n", pc, disp)
}

func (c *Processor) op2350() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A0),(%d,A1)\n", pc, disp)
}

func (c *Processor) op2351() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A1),(%d,A1)\n", pc, disp)
}

func (c *Processor) op2352() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A2),(%d,A1)\n", pc, disp)
}

func (c *Processor) op2353() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A3),(%d,A1)\n", pc, disp)
}

func (c *Processor) op2354() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A4),(%d,A1)\n", pc, disp)
}

func (c *Processor) op2355() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A5),(%d,A1)\n", pc, disp)
}

func (c *Processor) op2356() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A6),(%d,A1)\n", pc, disp)
}

func (c *Processor) op2357() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A7),(%d,A1)\n", pc, disp)
}

func (c *Processor) op2378() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A1)\n", pc, v, disp)
}

func (c *Processor) op2379() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A1)\n", pc, v, disp)
}

func (c *Processor) op237C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[1]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l #$%X,(%d,A1)\n", pc, v, disp)
}

func (c *Processor) op23C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D0,$%X\n", pc, addr)
}

func (c *Processor) op23C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D1,$%X\n", pc, addr)
}

func (c *Processor) op23C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D2,$%X\n", pc, addr)
}

func (c *Processor) op23C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D3,$%X\n", pc, addr)
}

func (c *Processor) op23C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D4,$%X\n", pc, addr)
}

func (c *Processor) op23C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D5,$%X\n", pc, addr)
}

func (c *Processor) op23C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D6,$%X\n", pc, addr)
}

func (c *Processor) op23C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D7,$%X\n", pc, addr)
}

func (c *Processor) op23C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A0,$%X\n", pc, addr)
}

func (c *Processor) op23C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A1,$%X\n", pc, addr)
}

func (c *Processor) op23CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A2,$%X\n", pc, addr)
}

func (c *Processor) op23CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A3,$%X\n", pc, addr)
}

func (c *Processor) op23CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A4,$%X\n", pc, addr)
}

func (c *Processor) op23CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A5,$%X\n", pc, addr)
}

func (c *Processor) op23CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A6,$%X\n", pc, addr)
}

func (c *Processor) op23CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A7,$%X\n", pc, addr)
}

func (c *Processor) op23D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A0),$%X\n", pc, addr)
}

func (c *Processor) op23D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A1),$%X\n", pc, addr)
}

func (c *Processor) op23D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A2),$%X\n", pc, addr)
}

func (c *Processor) op23D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A3),$%X\n", pc, addr)
}

func (c *Processor) op23D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A4),$%X\n", pc, addr)
}

func (c *Processor) op23D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A5),$%X\n", pc, addr)
}

func (c *Processor) op23D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A6),$%X\n", pc, addr)
}

func (c *Processor) op23D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A7),$%X\n", pc, addr)
}

func (c *Processor) op23F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,$%X\n", pc, v, addr)
}

func (c *Processor) op23F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,$%X\n", pc, v, addr)
}

func (c *Processor) op23FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	addr := c.readImm(0x02)
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l #$%X,$%X\n", pc, v, addr)
}

func (c *Processor) op2400() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l D0,D2\n", pc)
}

func (c *Processor) op2401() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l D1,D2\n", pc)
}

func (c *Processor) op2402() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l D2,D2\n", pc)
}

func (c *Processor) op2403() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l D3,D2\n", pc)
}

func (c *Processor) op2404() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l D4,D2\n", pc)
}

func (c *Processor) op2405() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l D5,D2\n", pc)
}

func (c *Processor) op2406() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l D6,D2\n", pc)
}

func (c *Processor) op2407() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l D7,D2\n", pc)
}

func (c *Processor) op2408() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l A0,D2\n", pc)
}

func (c *Processor) op2409() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l A1,D2\n", pc)
}

func (c *Processor) op240A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l A2,D2\n", pc)
}

func (c *Processor) op240B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l A3,D2\n", pc)
}

func (c *Processor) op240C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l A4,D2\n", pc)
}

func (c *Processor) op240D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l A5,D2\n", pc)
}

func (c *Processor) op240E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l A6,D2\n", pc)
}

func (c *Processor) op240F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[2] = uint32(v)
	c.tracef("%04X move.l A7,D2\n", pc)
}

func (c *Processor) op2410() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[2] = uint32(v)
	c.tracef("%04X move.l (A0),D2\n", pc)
}

func (c *Processor) op2411() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[2] = uint32(v)
	c.tracef("%04X move.l (A1),D2\n", pc)
}

func (c *Processor) op2412() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[2] = uint32(v)
	c.tracef("%04X move.l (A2),D2\n", pc)
}

func (c *Processor) op2413() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[2] = uint32(v)
	c.tracef("%04X move.l (A3),D2\n", pc)
}

func (c *Processor) op2414() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[2] = uint32(v)
	c.tracef("%04X move.l (A4),D2\n", pc)
}

func (c *Processor) op2415() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[2] = uint32(v)
	c.tracef("%04X move.l (A5),D2\n", pc)
}

func (c *Processor) op2416() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[2] = uint32(v)
	c.tracef("%04X move.l (A6),D2\n", pc)
}

func (c *Processor) op2417() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[2] = uint32(v)
	c.tracef("%04X move.l (A7),D2\n", pc)
}

func (c *Processor) op2438() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[2] = uint32(v)
	c.tracef("%04X move.l $%X,D2\n", pc, v)
}

func (c *Processor) op2439() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[2] = uint32(v)
	c.tracef("%04X move.l $%X,D2\n", pc, v)
}

func (c *Processor) op243C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[2] = uint32(v)
	c.tracef("%04X move.l #$%X,D2\n", pc, v)
}

func (c *Processor) op2440() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l D0,A2\n", pc)
}

func (c *Processor) op2441() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l D1,A2\n", pc)
}

func (c *Processor) op2442() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l D2,A2\n", pc)
}

func (c *Processor) op2443() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l D3,A2\n", pc)
}

func (c *Processor) op2444() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l D4,A2\n", pc)
}

func (c *Processor) op2445() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l D5,A2\n", pc)
}

func (c *Processor) op2446() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l D6,A2\n", pc)
}

func (c *Processor) op2447() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l D7,A2\n", pc)
}

func (c *Processor) op2448() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l A0,A2\n", pc)
}

func (c *Processor) op2449() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l A1,A2\n", pc)
}

func (c *Processor) op244A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l A2,A2\n", pc)
}

func (c *Processor) op244B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l A3,A2\n", pc)
}

func (c *Processor) op244C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l A4,A2\n", pc)
}

func (c *Processor) op244D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l A5,A2\n", pc)
}

func (c *Processor) op244E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l A6,A2\n", pc)
}

func (c *Processor) op244F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l A7,A2\n", pc)
}

func (c *Processor) op2450() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l (A0),A2\n", pc)
}

func (c *Processor) op2451() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l (A1),A2\n", pc)
}

func (c *Processor) op2452() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l (A2),A2\n", pc)
}

func (c *Processor) op2453() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l (A3),A2\n", pc)
}

func (c *Processor) op2454() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l (A4),A2\n", pc)
}

func (c *Processor) op2455() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l (A5),A2\n", pc)
}

func (c *Processor) op2456() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l (A6),A2\n", pc)
}

func (c *Processor) op2457() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l (A7),A2\n", pc)
}

func (c *Processor) op2478() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l $%X,A2\n", pc, v)
}

func (c *Processor) op2479() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l $%X,A2\n", pc, v)
}

func (c *Processor) op247C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] = uint32(v)
	c.tracef("%04X movea.l #$%X,A2\n", pc, v)
}

func (c *Processor) op2480() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D0,(A2)\n", pc)
}

func (c *Processor) op2481() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D1,(A2)\n", pc)
}

func (c *Processor) op2482() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D2,(A2)\n", pc)
}

func (c *Processor) op2483() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D3,(A2)\n", pc)
}

func (c *Processor) op2484() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D4,(A2)\n", pc)
}

func (c *Processor) op2485() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D5,(A2)\n", pc)
}

func (c *Processor) op2486() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D6,(A2)\n", pc)
}

func (c *Processor) op2487() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D7,(A2)\n", pc)
}

func (c *Processor) op2488() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A0,(A2)\n", pc)
}

func (c *Processor) op2489() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A1,(A2)\n", pc)
}

func (c *Processor) op248A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A2,(A2)\n", pc)
}

func (c *Processor) op248B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A3,(A2)\n", pc)
}

func (c *Processor) op248C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A4,(A2)\n", pc)
}

func (c *Processor) op248D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A5,(A2)\n", pc)
}

func (c *Processor) op248E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A6,(A2)\n", pc)
}

func (c *Processor) op248F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A7,(A2)\n", pc)
}

func (c *Processor) op2490() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A0),(A2)\n", pc)
}

func (c *Processor) op2491() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A1),(A2)\n", pc)
}

func (c *Processor) op2492() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A2),(A2)\n", pc)
}

func (c *Processor) op2493() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A3),(A2)\n", pc)
}

func (c *Processor) op2494() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A4),(A2)\n", pc)
}

func (c *Processor) op2495() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A5),(A2)\n", pc)
}

func (c *Processor) op2496() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A6),(A2)\n", pc)
}

func (c *Processor) op2497() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A7),(A2)\n", pc)
}

func (c *Processor) op24B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l $%X,(A2)\n", pc, v)
}

func (c *Processor) op24B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l $%X,(A2)\n", pc, v)
}

func (c *Processor) op24BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l #$%X,(A2)\n", pc, v)
}

func (c *Processor) op24C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l D0,(A2)+\n", pc)
}

func (c *Processor) op24C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l D1,(A2)+\n", pc)
}

func (c *Processor) op24C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l D2,(A2)+\n", pc)
}

func (c *Processor) op24C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l D3,(A2)+\n", pc)
}

func (c *Processor) op24C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l D4,(A2)+\n", pc)
}

func (c *Processor) op24C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l D5,(A2)+\n", pc)
}

func (c *Processor) op24C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l D6,(A2)+\n", pc)
}

func (c *Processor) op24C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l D7,(A2)+\n", pc)
}

func (c *Processor) op24C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l A0,(A2)+\n", pc)
}

func (c *Processor) op24C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l A1,(A2)+\n", pc)
}

func (c *Processor) op24CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l A2,(A2)+\n", pc)
}

func (c *Processor) op24CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l A3,(A2)+\n", pc)
}

func (c *Processor) op24CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l A4,(A2)+\n", pc)
}

func (c *Processor) op24CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l A5,(A2)+\n", pc)
}

func (c *Processor) op24CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l A6,(A2)+\n", pc)
}

func (c *Processor) op24CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l A7,(A2)+\n", pc)
}

func (c *Processor) op24D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l (A0),(A2)+\n", pc)
}

func (c *Processor) op24D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l (A1),(A2)+\n", pc)
}

func (c *Processor) op24D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l (A2),(A2)+\n", pc)
}

func (c *Processor) op24D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l (A3),(A2)+\n", pc)
}

func (c *Processor) op24D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l (A4),(A2)+\n", pc)
}

func (c *Processor) op24D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l (A5),(A2)+\n", pc)
}

func (c *Processor) op24D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l (A6),(A2)+\n", pc)
}

func (c *Processor) op24D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l (A7),(A2)+\n", pc)
}

func (c *Processor) op24F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l $%X,(A2)+\n", pc, v)
}

func (c *Processor) op24F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l $%X,(A2)+\n", pc, v)
}

func (c *Processor) op24FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[2], uint32(v))
	c.A[2] += 4
	c.tracef("%04X move.l #$%X,(A2)+\n", pc, v)
}

func (c *Processor) op2500() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D0,-(A2)\n", pc)
}

func (c *Processor) op2501() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D1,-(A2)\n", pc)
}

func (c *Processor) op2502() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D2,-(A2)\n", pc)
}

func (c *Processor) op2503() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D3,-(A2)\n", pc)
}

func (c *Processor) op2504() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D4,-(A2)\n", pc)
}

func (c *Processor) op2505() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D5,-(A2)\n", pc)
}

func (c *Processor) op2506() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D6,-(A2)\n", pc)
}

func (c *Processor) op2507() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l D7,-(A2)\n", pc)
}

func (c *Processor) op2508() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A0,-(A2)\n", pc)
}

func (c *Processor) op2509() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A1,-(A2)\n", pc)
}

func (c *Processor) op250A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A2,-(A2)\n", pc)
}

func (c *Processor) op250B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A3,-(A2)\n", pc)
}

func (c *Processor) op250C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A4,-(A2)\n", pc)
}

func (c *Processor) op250D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A5,-(A2)\n", pc)
}

func (c *Processor) op250E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A6,-(A2)\n", pc)
}

func (c *Processor) op250F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l A7,-(A2)\n", pc)
}

func (c *Processor) op2510() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A0),-(A2)\n", pc)
}

func (c *Processor) op2511() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A1),-(A2)\n", pc)
}

func (c *Processor) op2512() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A2),-(A2)\n", pc)
}

func (c *Processor) op2513() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A3),-(A2)\n", pc)
}

func (c *Processor) op2514() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A4),-(A2)\n", pc)
}

func (c *Processor) op2515() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A5),-(A2)\n", pc)
}

func (c *Processor) op2516() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A6),-(A2)\n", pc)
}

func (c *Processor) op2517() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l (A7),-(A2)\n", pc)
}

func (c *Processor) op2538() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l $%X,-(A2)\n", pc, v)
}

func (c *Processor) op2539() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l $%X,-(A2)\n", pc, v)
}

func (c *Processor) op253C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[2] -= 4
	c.writeLong(c.A[2], uint32(v))
	c.tracef("%04X move.l #$%X,-(A2)\n", pc, v)
}

func (c *Processor) op2540() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D0,(%d,A2)\n", pc, disp)
}

func (c *Processor) op2541() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D1,(%d,A2)\n", pc, disp)
}

func (c *Processor) op2542() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D2,(%d,A2)\n", pc, disp)
}

func (c *Processor) op2543() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D3,(%d,A2)\n", pc, disp)
}

func (c *Processor) op2544() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D4,(%d,A2)\n", pc, disp)
}

func (c *Processor) op2545() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D5,(%d,A2)\n", pc, disp)
}

func (c *Processor) op2546() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D6,(%d,A2)\n", pc, disp)
}

func (c *Processor) op2547() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D7,(%d,A2)\n", pc, disp)
}

func (c *Processor) op2548() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A0,(%d,A2)\n", pc, disp)
}

func (c *Processor) op2549() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A1,(%d,A2)\n", pc, disp)
}

func (c *Processor) op254A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A2,(%d,A2)\n", pc, disp)
}

func (c *Processor) op254B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A3,(%d,A2)\n", pc, disp)
}

func (c *Processor) op254C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A4,(%d,A2)\n", pc, disp)
}

func (c *Processor) op254D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A5,(%d,A2)\n", pc, disp)
}

func (c *Processor) op254E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A6,(%d,A2)\n", pc, disp)
}

func (c *Processor) op254F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A7,(%d,A2)\n", pc, disp)
}

func (c *Processor) op2550() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A0),(%d,A2)\n", pc, disp)
}

func (c *Processor) op2551() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A1),(%d,A2)\n", pc, disp)
}

func (c *Processor) op2552() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A2),(%d,A2)\n", pc, disp)
}

func (c *Processor) op2553() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A3),(%d,A2)\n", pc, disp)
}

func (c *Processor) op2554() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A4),(%d,A2)\n", pc, disp)
}

func (c *Processor) op2555() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A5),(%d,A2)\n", pc, disp)
}

func (c *Processor) op2556() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A6),(%d,A2)\n", pc, disp)
}

func (c *Processor) op2557() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A7),(%d,A2)\n", pc, disp)
}

func (c *Processor) op2578() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A2)\n", pc, v, disp)
}

func (c *Processor) op2579() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A2)\n", pc, v, disp)
}

func (c *Processor) op257C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[2]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l #$%X,(%d,A2)\n", pc, v, disp)
}

func (c *Processor) op2600() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l D0,D3\n", pc)
}

func (c *Processor) op2601() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l D1,D3\n", pc)
}

func (c *Processor) op2602() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l D2,D3\n", pc)
}

func (c *Processor) op2603() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l D3,D3\n", pc)
}

func (c *Processor) op2604() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l D4,D3\n", pc)
}

func (c *Processor) op2605() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l D5,D3\n", pc)
}

func (c *Processor) op2606() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l D6,D3\n", pc)
}

func (c *Processor) op2607() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l D7,D3\n", pc)
}

func (c *Processor) op2608() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l A0,D3\n", pc)
}

func (c *Processor) op2609() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l A1,D3\n", pc)
}

func (c *Processor) op260A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l A2,D3\n", pc)
}

func (c *Processor) op260B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l A3,D3\n", pc)
}

func (c *Processor) op260C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l A4,D3\n", pc)
}

func (c *Processor) op260D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l A5,D3\n", pc)
}

func (c *Processor) op260E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l A6,D3\n", pc)
}

func (c *Processor) op260F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[3] = uint32(v)
	c.tracef("%04X move.l A7,D3\n", pc)
}

func (c *Processor) op2610() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[3] = uint32(v)
	c.tracef("%04X move.l (A0),D3\n", pc)
}

func (c *Processor) op2611() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[3] = uint32(v)
	c.tracef("%04X move.l (A1),D3\n", pc)
}

func (c *Processor) op2612() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[3] = uint32(v)
	c.tracef("%04X move.l (A2),D3\n", pc)
}

func (c *Processor) op2613() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[3] = uint32(v)
	c.tracef("%04X move.l (A3),D3\n", pc)
}

func (c *Processor) op2614() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[3] = uint32(v)
	c.tracef("%04X move.l (A4),D3\n", pc)
}

func (c *Processor) op2615() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[3] = uint32(v)
	c.tracef("%04X move.l (A5),D3\n", pc)
}

func (c *Processor) op2616() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[3] = uint32(v)
	c.tracef("%04X move.l (A6),D3\n", pc)
}

func (c *Processor) op2617() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[3] = uint32(v)
	c.tracef("%04X move.l (A7),D3\n", pc)
}

func (c *Processor) op2638() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[3] = uint32(v)
	c.tracef("%04X move.l $%X,D3\n", pc, v)
}

func (c *Processor) op2639() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[3] = uint32(v)
	c.tracef("%04X move.l $%X,D3\n", pc, v)
}

func (c *Processor) op263C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[3] = uint32(v)
	c.tracef("%04X move.l #$%X,D3\n", pc, v)
}

func (c *Processor) op2640() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l D0,A3\n", pc)
}

func (c *Processor) op2641() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l D1,A3\n", pc)
}

func (c *Processor) op2642() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l D2,A3\n", pc)
}

func (c *Processor) op2643() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l D3,A3\n", pc)
}

func (c *Processor) op2644() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l D4,A3\n", pc)
}

func (c *Processor) op2645() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l D5,A3\n", pc)
}

func (c *Processor) op2646() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l D6,A3\n", pc)
}

func (c *Processor) op2647() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l D7,A3\n", pc)
}

func (c *Processor) op2648() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l A0,A3\n", pc)
}

func (c *Processor) op2649() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l A1,A3\n", pc)
}

func (c *Processor) op264A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l A2,A3\n", pc)
}

func (c *Processor) op264B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l A3,A3\n", pc)
}

func (c *Processor) op264C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l A4,A3\n", pc)
}

func (c *Processor) op264D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l A5,A3\n", pc)
}

func (c *Processor) op264E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l A6,A3\n", pc)
}

func (c *Processor) op264F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l A7,A3\n", pc)
}

func (c *Processor) op2650() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l (A0),A3\n", pc)
}

func (c *Processor) op2651() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l (A1),A3\n", pc)
}

func (c *Processor) op2652() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l (A2),A3\n", pc)
}

func (c *Processor) op2653() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l (A3),A3\n", pc)
}

func (c *Processor) op2654() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l (A4),A3\n", pc)
}

func (c *Processor) op2655() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l (A5),A3\n", pc)
}

func (c *Processor) op2656() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l (A6),A3\n", pc)
}

func (c *Processor) op2657() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l (A7),A3\n", pc)
}

func (c *Processor) op2678() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l $%X,A3\n", pc, v)
}

func (c *Processor) op2679() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l $%X,A3\n", pc, v)
}

func (c *Processor) op267C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] = uint32(v)
	c.tracef("%04X movea.l #$%X,A3\n", pc, v)
}

func (c *Processor) op2680() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D0,(A3)\n", pc)
}

func (c *Processor) op2681() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D1,(A3)\n", pc)
}

func (c *Processor) op2682() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D2,(A3)\n", pc)
}

func (c *Processor) op2683() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D3,(A3)\n", pc)
}

func (c *Processor) op2684() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D4,(A3)\n", pc)
}

func (c *Processor) op2685() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D5,(A3)\n", pc)
}

func (c *Processor) op2686() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D6,(A3)\n", pc)
}

func (c *Processor) op2687() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D7,(A3)\n", pc)
}

func (c *Processor) op2688() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A0,(A3)\n", pc)
}

func (c *Processor) op2689() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A1,(A3)\n", pc)
}

func (c *Processor) op268A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A2,(A3)\n", pc)
}

func (c *Processor) op268B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A3,(A3)\n", pc)
}

func (c *Processor) op268C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A4,(A3)\n", pc)
}

func (c *Processor) op268D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A5,(A3)\n", pc)
}

func (c *Processor) op268E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A6,(A3)\n", pc)
}

func (c *Processor) op268F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A7,(A3)\n", pc)
}

func (c *Processor) op2690() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A0),(A3)\n", pc)
}

func (c *Processor) op2691() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A1),(A3)\n", pc)
}

func (c *Processor) op2692() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A2),(A3)\n", pc)
}

func (c *Processor) op2693() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A3),(A3)\n", pc)
}

func (c *Processor) op2694() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A4),(A3)\n", pc)
}

func (c *Processor) op2695() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A5),(A3)\n", pc)
}

func (c *Processor) op2696() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A6),(A3)\n", pc)
}

func (c *Processor) op2697() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A7),(A3)\n", pc)
}

func (c *Processor) op26B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l $%X,(A3)\n", pc, v)
}

func (c *Processor) op26B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l $%X,(A3)\n", pc, v)
}

func (c *Processor) op26BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l #$%X,(A3)\n", pc, v)
}

func (c *Processor) op26C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l D0,(A3)+\n", pc)
}

func (c *Processor) op26C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l D1,(A3)+\n", pc)
}

func (c *Processor) op26C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l D2,(A3)+\n", pc)
}

func (c *Processor) op26C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l D3,(A3)+\n", pc)
}

func (c *Processor) op26C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l D4,(A3)+\n", pc)
}

func (c *Processor) op26C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l D5,(A3)+\n", pc)
}

func (c *Processor) op26C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l D6,(A3)+\n", pc)
}

func (c *Processor) op26C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l D7,(A3)+\n", pc)
}

func (c *Processor) op26C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l A0,(A3)+\n", pc)
}

func (c *Processor) op26C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l A1,(A3)+\n", pc)
}

func (c *Processor) op26CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l A2,(A3)+\n", pc)
}

func (c *Processor) op26CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l A3,(A3)+\n", pc)
}

func (c *Processor) op26CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l A4,(A3)+\n", pc)
}

func (c *Processor) op26CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l A5,(A3)+\n", pc)
}

func (c *Processor) op26CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l A6,(A3)+\n", pc)
}

func (c *Processor) op26CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l A7,(A3)+\n", pc)
}

func (c *Processor) op26D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l (A0),(A3)+\n", pc)
}

func (c *Processor) op26D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l (A1),(A3)+\n", pc)
}

func (c *Processor) op26D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l (A2),(A3)+\n", pc)
}

func (c *Processor) op26D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l (A3),(A3)+\n", pc)
}

func (c *Processor) op26D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l (A4),(A3)+\n", pc)
}

func (c *Processor) op26D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l (A5),(A3)+\n", pc)
}

func (c *Processor) op26D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l (A6),(A3)+\n", pc)
}

func (c *Processor) op26D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l (A7),(A3)+\n", pc)
}

func (c *Processor) op26F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l $%X,(A3)+\n", pc, v)
}

func (c *Processor) op26F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l $%X,(A3)+\n", pc, v)
}

func (c *Processor) op26FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[3], uint32(v))
	c.A[3] += 4
	c.tracef("%04X move.l #$%X,(A3)+\n", pc, v)
}

func (c *Processor) op2700() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D0,-(A3)\n", pc)
}

func (c *Processor) op2701() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D1,-(A3)\n", pc)
}

func (c *Processor) op2702() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D2,-(A3)\n", pc)
}

func (c *Processor) op2703() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D3,-(A3)\n", pc)
}

func (c *Processor) op2704() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D4,-(A3)\n", pc)
}

func (c *Processor) op2705() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D5,-(A3)\n", pc)
}

func (c *Processor) op2706() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D6,-(A3)\n", pc)
}

func (c *Processor) op2707() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l D7,-(A3)\n", pc)
}

func (c *Processor) op2708() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A0,-(A3)\n", pc)
}

func (c *Processor) op2709() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A1,-(A3)\n", pc)
}

func (c *Processor) op270A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A2,-(A3)\n", pc)
}

func (c *Processor) op270B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A3,-(A3)\n", pc)
}

func (c *Processor) op270C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A4,-(A3)\n", pc)
}

func (c *Processor) op270D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A5,-(A3)\n", pc)
}

func (c *Processor) op270E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A6,-(A3)\n", pc)
}

func (c *Processor) op270F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l A7,-(A3)\n", pc)
}

func (c *Processor) op2710() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A0),-(A3)\n", pc)
}

func (c *Processor) op2711() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A1),-(A3)\n", pc)
}

func (c *Processor) op2712() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A2),-(A3)\n", pc)
}

func (c *Processor) op2713() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A3),-(A3)\n", pc)
}

func (c *Processor) op2714() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A4),-(A3)\n", pc)
}

func (c *Processor) op2715() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A5),-(A3)\n", pc)
}

func (c *Processor) op2716() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A6),-(A3)\n", pc)
}

func (c *Processor) op2717() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l (A7),-(A3)\n", pc)
}

func (c *Processor) op2738() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l $%X,-(A3)\n", pc, v)
}

func (c *Processor) op2739() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l $%X,-(A3)\n", pc, v)
}

func (c *Processor) op273C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[3] -= 4
	c.writeLong(c.A[3], uint32(v))
	c.tracef("%04X move.l #$%X,-(A3)\n", pc, v)
}

func (c *Processor) op2740() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D0,(%d,A3)\n", pc, disp)
}

func (c *Processor) op2741() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D1,(%d,A3)\n", pc, disp)
}

func (c *Processor) op2742() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D2,(%d,A3)\n", pc, disp)
}

func (c *Processor) op2743() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D3,(%d,A3)\n", pc, disp)
}

func (c *Processor) op2744() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D4,(%d,A3)\n", pc, disp)
}

func (c *Processor) op2745() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D5,(%d,A3)\n", pc, disp)
}

func (c *Processor) op2746() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D6,(%d,A3)\n", pc, disp)
}

func (c *Processor) op2747() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D7,(%d,A3)\n", pc, disp)
}

func (c *Processor) op2748() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A0,(%d,A3)\n", pc, disp)
}

func (c *Processor) op2749() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A1,(%d,A3)\n", pc, disp)
}

func (c *Processor) op274A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A2,(%d,A3)\n", pc, disp)
}

func (c *Processor) op274B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A3,(%d,A3)\n", pc, disp)
}

func (c *Processor) op274C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A4,(%d,A3)\n", pc, disp)
}

func (c *Processor) op274D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A5,(%d,A3)\n", pc, disp)
}

func (c *Processor) op274E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A6,(%d,A3)\n", pc, disp)
}

func (c *Processor) op274F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A7,(%d,A3)\n", pc, disp)
}

func (c *Processor) op2750() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A0),(%d,A3)\n", pc, disp)
}

func (c *Processor) op2751() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A1),(%d,A3)\n", pc, disp)
}

func (c *Processor) op2752() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A2),(%d,A3)\n", pc, disp)
}

func (c *Processor) op2753() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A3),(%d,A3)\n", pc, disp)
}

func (c *Processor) op2754() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A4),(%d,A3)\n", pc, disp)
}

func (c *Processor) op2755() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A5),(%d,A3)\n", pc, disp)
}

func (c *Processor) op2756() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A6),(%d,A3)\n", pc, disp)
}

func (c *Processor) op2757() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A7),(%d,A3)\n", pc, disp)
}

func (c *Processor) op2778() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A3)\n", pc, v, disp)
}

func (c *Processor) op2779() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A3)\n", pc, v, disp)
}

func (c *Processor) op277C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[3]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l #$%X,(%d,A3)\n", pc, v, disp)
}

func (c *Processor) op2800() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l D0,D4\n", pc)
}

func (c *Processor) op2801() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l D1,D4\n", pc)
}

func (c *Processor) op2802() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l D2,D4\n", pc)
}

func (c *Processor) op2803() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l D3,D4\n", pc)
}

func (c *Processor) op2804() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l D4,D4\n", pc)
}

func (c *Processor) op2805() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l D5,D4\n", pc)
}

func (c *Processor) op2806() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l D6,D4\n", pc)
}

func (c *Processor) op2807() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l D7,D4\n", pc)
}

func (c *Processor) op2808() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l A0,D4\n", pc)
}

func (c *Processor) op2809() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l A1,D4\n", pc)
}

func (c *Processor) op280A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l A2,D4\n", pc)
}

func (c *Processor) op280B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l A3,D4\n", pc)
}

func (c *Processor) op280C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l A4,D4\n", pc)
}

func (c *Processor) op280D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l A5,D4\n", pc)
}

func (c *Processor) op280E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l A6,D4\n", pc)
}

func (c *Processor) op280F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[4] = uint32(v)
	c.tracef("%04X move.l A7,D4\n", pc)
}

func (c *Processor) op2810() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[4] = uint32(v)
	c.tracef("%04X move.l (A0),D4\n", pc)
}

func (c *Processor) op2811() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[4] = uint32(v)
	c.tracef("%04X move.l (A1),D4\n", pc)
}

func (c *Processor) op2812() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[4] = uint32(v)
	c.tracef("%04X move.l (A2),D4\n", pc)
}

func (c *Processor) op2813() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[4] = uint32(v)
	c.tracef("%04X move.l (A3),D4\n", pc)
}

func (c *Processor) op2814() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[4] = uint32(v)
	c.tracef("%04X move.l (A4),D4\n", pc)
}

func (c *Processor) op2815() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[4] = uint32(v)
	c.tracef("%04X move.l (A5),D4\n", pc)
}

func (c *Processor) op2816() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[4] = uint32(v)
	c.tracef("%04X move.l (A6),D4\n", pc)
}

func (c *Processor) op2817() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[4] = uint32(v)
	c.tracef("%04X move.l (A7),D4\n", pc)
}

func (c *Processor) op2838() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[4] = uint32(v)
	c.tracef("%04X move.l $%X,D4\n", pc, v)
}

func (c *Processor) op2839() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[4] = uint32(v)
	c.tracef("%04X move.l $%X,D4\n", pc, v)
}

func (c *Processor) op283C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[4] = uint32(v)
	c.tracef("%04X move.l #$%X,D4\n", pc, v)
}

func (c *Processor) op2840() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l D0,A4\n", pc)
}

func (c *Processor) op2841() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l D1,A4\n", pc)
}

func (c *Processor) op2842() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l D2,A4\n", pc)
}

func (c *Processor) op2843() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l D3,A4\n", pc)
}

func (c *Processor) op2844() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l D4,A4\n", pc)
}

func (c *Processor) op2845() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l D5,A4\n", pc)
}

func (c *Processor) op2846() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l D6,A4\n", pc)
}

func (c *Processor) op2847() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l D7,A4\n", pc)
}

func (c *Processor) op2848() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l A0,A4\n", pc)
}

func (c *Processor) op2849() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l A1,A4\n", pc)
}

func (c *Processor) op284A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l A2,A4\n", pc)
}

func (c *Processor) op284B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l A3,A4\n", pc)
}

func (c *Processor) op284C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l A4,A4\n", pc)
}

func (c *Processor) op284D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l A5,A4\n", pc)
}

func (c *Processor) op284E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l A6,A4\n", pc)
}

func (c *Processor) op284F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l A7,A4\n", pc)
}

func (c *Processor) op2850() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l (A0),A4\n", pc)
}

func (c *Processor) op2851() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l (A1),A4\n", pc)
}

func (c *Processor) op2852() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l (A2),A4\n", pc)
}

func (c *Processor) op2853() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l (A3),A4\n", pc)
}

func (c *Processor) op2854() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l (A4),A4\n", pc)
}

func (c *Processor) op2855() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l (A5),A4\n", pc)
}

func (c *Processor) op2856() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l (A6),A4\n", pc)
}

func (c *Processor) op2857() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l (A7),A4\n", pc)
}

func (c *Processor) op2878() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l $%X,A4\n", pc, v)
}

func (c *Processor) op2879() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l $%X,A4\n", pc, v)
}

func (c *Processor) op287C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] = uint32(v)
	c.tracef("%04X movea.l #$%X,A4\n", pc, v)
}

func (c *Processor) op2880() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D0,(A4)\n", pc)
}

func (c *Processor) op2881() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D1,(A4)\n", pc)
}

func (c *Processor) op2882() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D2,(A4)\n", pc)
}

func (c *Processor) op2883() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D3,(A4)\n", pc)
}

func (c *Processor) op2884() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D4,(A4)\n", pc)
}

func (c *Processor) op2885() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D5,(A4)\n", pc)
}

func (c *Processor) op2886() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D6,(A4)\n", pc)
}

func (c *Processor) op2887() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D7,(A4)\n", pc)
}

func (c *Processor) op2888() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A0,(A4)\n", pc)
}

func (c *Processor) op2889() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A1,(A4)\n", pc)
}

func (c *Processor) op288A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A2,(A4)\n", pc)
}

func (c *Processor) op288B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A3,(A4)\n", pc)
}

func (c *Processor) op288C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A4,(A4)\n", pc)
}

func (c *Processor) op288D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A5,(A4)\n", pc)
}

func (c *Processor) op288E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A6,(A4)\n", pc)
}

func (c *Processor) op288F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A7,(A4)\n", pc)
}

func (c *Processor) op2890() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A0),(A4)\n", pc)
}

func (c *Processor) op2891() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A1),(A4)\n", pc)
}

func (c *Processor) op2892() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A2),(A4)\n", pc)
}

func (c *Processor) op2893() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A3),(A4)\n", pc)
}

func (c *Processor) op2894() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A4),(A4)\n", pc)
}

func (c *Processor) op2895() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A5),(A4)\n", pc)
}

func (c *Processor) op2896() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A6),(A4)\n", pc)
}

func (c *Processor) op2897() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A7),(A4)\n", pc)
}

func (c *Processor) op28B8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l $%X,(A4)\n", pc, v)
}

func (c *Processor) op28B9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l $%X,(A4)\n", pc, v)
}

func (c *Processor) op28BC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l #$%X,(A4)\n", pc, v)
}

func (c *Processor) op28C0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l D0,(A4)+\n", pc)
}

func (c *Processor) op28C1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l D1,(A4)+\n", pc)
}

func (c *Processor) op28C2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l D2,(A4)+\n", pc)
}

func (c *Processor) op28C3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l D3,(A4)+\n", pc)
}

func (c *Processor) op28C4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l D4,(A4)+\n", pc)
}

func (c *Processor) op28C5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l D5,(A4)+\n", pc)
}

func (c *Processor) op28C6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l D6,(A4)+\n", pc)
}

func (c *Processor) op28C7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l D7,(A4)+\n", pc)
}

func (c *Processor) op28C8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l A0,(A4)+\n", pc)
}

func (c *Processor) op28C9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l A1,(A4)+\n", pc)
}

func (c *Processor) op28CA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l A2,(A4)+\n", pc)
}

func (c *Processor) op28CB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l A3,(A4)+\n", pc)
}

func (c *Processor) op28CC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l A4,(A4)+\n", pc)
}

func (c *Processor) op28CD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l A5,(A4)+\n", pc)
}

func (c *Processor) op28CE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l A6,(A4)+\n", pc)
}

func (c *Processor) op28CF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l A7,(A4)+\n", pc)
}

func (c *Processor) op28D0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l (A0),(A4)+\n", pc)
}

func (c *Processor) op28D1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l (A1),(A4)+\n", pc)
}

func (c *Processor) op28D2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l (A2),(A4)+\n", pc)
}

func (c *Processor) op28D3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l (A3),(A4)+\n", pc)
}

func (c *Processor) op28D4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l (A4),(A4)+\n", pc)
}

func (c *Processor) op28D5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l (A5),(A4)+\n", pc)
}

func (c *Processor) op28D6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l (A6),(A4)+\n", pc)
}

func (c *Processor) op28D7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l (A7),(A4)+\n", pc)
}

func (c *Processor) op28F8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l $%X,(A4)+\n", pc, v)
}

func (c *Processor) op28F9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l $%X,(A4)+\n", pc, v)
}

func (c *Processor) op28FC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[4], uint32(v))
	c.A[4] += 4
	c.tracef("%04X move.l #$%X,(A4)+\n", pc, v)
}

func (c *Processor) op2900() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D0,-(A4)\n", pc)
}

func (c *Processor) op2901() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D1,-(A4)\n", pc)
}

func (c *Processor) op2902() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D2,-(A4)\n", pc)
}

func (c *Processor) op2903() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D3,-(A4)\n", pc)
}

func (c *Processor) op2904() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D4,-(A4)\n", pc)
}

func (c *Processor) op2905() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D5,-(A4)\n", pc)
}

func (c *Processor) op2906() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D6,-(A4)\n", pc)
}

func (c *Processor) op2907() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l D7,-(A4)\n", pc)
}

func (c *Processor) op2908() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A0,-(A4)\n", pc)
}

func (c *Processor) op2909() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A1,-(A4)\n", pc)
}

func (c *Processor) op290A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A2,-(A4)\n", pc)
}

func (c *Processor) op290B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A3,-(A4)\n", pc)
}

func (c *Processor) op290C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A4,-(A4)\n", pc)
}

func (c *Processor) op290D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A5,-(A4)\n", pc)
}

func (c *Processor) op290E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A6,-(A4)\n", pc)
}

func (c *Processor) op290F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l A7,-(A4)\n", pc)
}

func (c *Processor) op2910() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A0),-(A4)\n", pc)
}

func (c *Processor) op2911() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A1),-(A4)\n", pc)
}

func (c *Processor) op2912() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A2),-(A4)\n", pc)
}

func (c *Processor) op2913() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A3),-(A4)\n", pc)
}

func (c *Processor) op2914() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A4),-(A4)\n", pc)
}

func (c *Processor) op2915() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A5),-(A4)\n", pc)
}

func (c *Processor) op2916() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A6),-(A4)\n", pc)
}

func (c *Processor) op2917() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l (A7),-(A4)\n", pc)
}

func (c *Processor) op2938() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l $%X,-(A4)\n", pc, v)
}

func (c *Processor) op2939() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l $%X,-(A4)\n", pc, v)
}

func (c *Processor) op293C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[4] -= 4
	c.writeLong(c.A[4], uint32(v))
	c.tracef("%04X move.l #$%X,-(A4)\n", pc, v)
}

func (c *Processor) op2940() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D0,(%d,A4)\n", pc, disp)
}

func (c *Processor) op2941() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D1,(%d,A4)\n", pc, disp)
}

func (c *Processor) op2942() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D2,(%d,A4)\n", pc, disp)
}

func (c *Processor) op2943() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D3,(%d,A4)\n", pc, disp)
}

func (c *Processor) op2944() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D4,(%d,A4)\n", pc, disp)
}

func (c *Processor) op2945() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D5,(%d,A4)\n", pc, disp)
}

func (c *Processor) op2946() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D6,(%d,A4)\n", pc, disp)
}

func (c *Processor) op2947() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D7,(%d,A4)\n", pc, disp)
}

func (c *Processor) op2948() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A0,(%d,A4)\n", pc, disp)
}

func (c *Processor) op2949() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A1,(%d,A4)\n", pc, disp)
}

func (c *Processor) op294A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A2,(%d,A4)\n", pc, disp)
}

func (c *Processor) op294B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A3,(%d,A4)\n", pc, disp)
}

func (c *Processor) op294C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A4,(%d,A4)\n", pc, disp)
}

func (c *Processor) op294D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A5,(%d,A4)\n", pc, disp)
}

func (c *Processor) op294E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A6,(%d,A4)\n", pc, disp)
}

func (c *Processor) op294F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A7,(%d,A4)\n", pc, disp)
}

func (c *Processor) op2950() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A0),(%d,A4)\n", pc, disp)
}

func (c *Processor) op2951() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A1),(%d,A4)\n", pc, disp)
}

func (c *Processor) op2952() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A2),(%d,A4)\n", pc, disp)
}

func (c *Processor) op2953() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A3),(%d,A4)\n", pc, disp)
}

func (c *Processor) op2954() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A4),(%d,A4)\n", pc, disp)
}

func (c *Processor) op2955() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A5),(%d,A4)\n", pc, disp)
}

func (c *Processor) op2956() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A6),(%d,A4)\n", pc, disp)
}

func (c *Processor) op2957() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A7),(%d,A4)\n", pc, disp)
}

func (c *Processor) op2978() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A4)\n", pc, v, disp)
}

func (c *Processor) op2979() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A4)\n", pc, v, disp)
}

func (c *Processor) op297C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[4]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l #$%X,(%d,A4)\n", pc, v, disp)
}

func (c *Processor) op2A00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l D0,D5\n", pc)
}

func (c *Processor) op2A01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l D1,D5\n", pc)
}

func (c *Processor) op2A02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l D2,D5\n", pc)
}

func (c *Processor) op2A03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l D3,D5\n", pc)
}

func (c *Processor) op2A04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l D4,D5\n", pc)
}

func (c *Processor) op2A05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l D5,D5\n", pc)
}

func (c *Processor) op2A06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l D6,D5\n", pc)
}

func (c *Processor) op2A07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l D7,D5\n", pc)
}

func (c *Processor) op2A08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l A0,D5\n", pc)
}

func (c *Processor) op2A09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l A1,D5\n", pc)
}

func (c *Processor) op2A0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l A2,D5\n", pc)
}

func (c *Processor) op2A0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l A3,D5\n", pc)
}

func (c *Processor) op2A0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l A4,D5\n", pc)
}

func (c *Processor) op2A0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l A5,D5\n", pc)
}

func (c *Processor) op2A0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l A6,D5\n", pc)
}

func (c *Processor) op2A0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[5] = uint32(v)
	c.tracef("%04X move.l A7,D5\n", pc)
}

func (c *Processor) op2A10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[5] = uint32(v)
	c.tracef("%04X move.l (A0),D5\n", pc)
}

func (c *Processor) op2A11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[5] = uint32(v)
	c.tracef("%04X move.l (A1),D5\n", pc)
}

func (c *Processor) op2A12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[5] = uint32(v)
	c.tracef("%04X move.l (A2),D5\n", pc)
}

func (c *Processor) op2A13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[5] = uint32(v)
	c.tracef("%04X move.l (A3),D5\n", pc)
}

func (c *Processor) op2A14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[5] = uint32(v)
	c.tracef("%04X move.l (A4),D5\n", pc)
}

func (c *Processor) op2A15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[5] = uint32(v)
	c.tracef("%04X move.l (A5),D5\n", pc)
}

func (c *Processor) op2A16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[5] = uint32(v)
	c.tracef("%04X move.l (A6),D5\n", pc)
}

func (c *Processor) op2A17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[5] = uint32(v)
	c.tracef("%04X move.l (A7),D5\n", pc)
}

func (c *Processor) op2A38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[5] = uint32(v)
	c.tracef("%04X move.l $%X,D5\n", pc, v)
}

func (c *Processor) op2A39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[5] = uint32(v)
	c.tracef("%04X move.l $%X,D5\n", pc, v)
}

func (c *Processor) op2A3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[5] = uint32(v)
	c.tracef("%04X move.l #$%X,D5\n", pc, v)
}

func (c *Processor) op2A40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l D0,A5\n", pc)
}

func (c *Processor) op2A41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l D1,A5\n", pc)
}

func (c *Processor) op2A42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l D2,A5\n", pc)
}

func (c *Processor) op2A43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l D3,A5\n", pc)
}

func (c *Processor) op2A44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l D4,A5\n", pc)
}

func (c *Processor) op2A45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l D5,A5\n", pc)
}

func (c *Processor) op2A46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l D6,A5\n", pc)
}

func (c *Processor) op2A47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l D7,A5\n", pc)
}

func (c *Processor) op2A48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l A0,A5\n", pc)
}

func (c *Processor) op2A49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l A1,A5\n", pc)
}

func (c *Processor) op2A4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l A2,A5\n", pc)
}

func (c *Processor) op2A4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l A3,A5\n", pc)
}

func (c *Processor) op2A4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l A4,A5\n", pc)
}

func (c *Processor) op2A4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l A5,A5\n", pc)
}

func (c *Processor) op2A4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l A6,A5\n", pc)
}

func (c *Processor) op2A4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l A7,A5\n", pc)
}

func (c *Processor) op2A50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l (A0),A5\n", pc)
}

func (c *Processor) op2A51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l (A1),A5\n", pc)
}

func (c *Processor) op2A52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l (A2),A5\n", pc)
}

func (c *Processor) op2A53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l (A3),A5\n", pc)
}

func (c *Processor) op2A54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l (A4),A5\n", pc)
}

func (c *Processor) op2A55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l (A5),A5\n", pc)
}

func (c *Processor) op2A56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l (A6),A5\n", pc)
}

func (c *Processor) op2A57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l (A7),A5\n", pc)
}

func (c *Processor) op2A78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l $%X,A5\n", pc, v)
}

func (c *Processor) op2A79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l $%X,A5\n", pc, v)
}

func (c *Processor) op2A7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] = uint32(v)
	c.tracef("%04X movea.l #$%X,A5\n", pc, v)
}

func (c *Processor) op2A80() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D0,(A5)\n", pc)
}

func (c *Processor) op2A81() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D1,(A5)\n", pc)
}

func (c *Processor) op2A82() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D2,(A5)\n", pc)
}

func (c *Processor) op2A83() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D3,(A5)\n", pc)
}

func (c *Processor) op2A84() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D4,(A5)\n", pc)
}

func (c *Processor) op2A85() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D5,(A5)\n", pc)
}

func (c *Processor) op2A86() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D6,(A5)\n", pc)
}

func (c *Processor) op2A87() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D7,(A5)\n", pc)
}

func (c *Processor) op2A88() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A0,(A5)\n", pc)
}

func (c *Processor) op2A89() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A1,(A5)\n", pc)
}

func (c *Processor) op2A8A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A2,(A5)\n", pc)
}

func (c *Processor) op2A8B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A3,(A5)\n", pc)
}

func (c *Processor) op2A8C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A4,(A5)\n", pc)
}

func (c *Processor) op2A8D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A5,(A5)\n", pc)
}

func (c *Processor) op2A8E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A6,(A5)\n", pc)
}

func (c *Processor) op2A8F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A7,(A5)\n", pc)
}

func (c *Processor) op2A90() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A0),(A5)\n", pc)
}

func (c *Processor) op2A91() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A1),(A5)\n", pc)
}

func (c *Processor) op2A92() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A2),(A5)\n", pc)
}

func (c *Processor) op2A93() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A3),(A5)\n", pc)
}

func (c *Processor) op2A94() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A4),(A5)\n", pc)
}

func (c *Processor) op2A95() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A5),(A5)\n", pc)
}

func (c *Processor) op2A96() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A6),(A5)\n", pc)
}

func (c *Processor) op2A97() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A7),(A5)\n", pc)
}

func (c *Processor) op2AB8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l $%X,(A5)\n", pc, v)
}

func (c *Processor) op2AB9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l $%X,(A5)\n", pc, v)
}

func (c *Processor) op2ABC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l #$%X,(A5)\n", pc, v)
}

func (c *Processor) op2AC0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l D0,(A5)+\n", pc)
}

func (c *Processor) op2AC1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l D1,(A5)+\n", pc)
}

func (c *Processor) op2AC2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l D2,(A5)+\n", pc)
}

func (c *Processor) op2AC3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l D3,(A5)+\n", pc)
}

func (c *Processor) op2AC4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l D4,(A5)+\n", pc)
}

func (c *Processor) op2AC5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l D5,(A5)+\n", pc)
}

func (c *Processor) op2AC6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l D6,(A5)+\n", pc)
}

func (c *Processor) op2AC7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l D7,(A5)+\n", pc)
}

func (c *Processor) op2AC8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l A0,(A5)+\n", pc)
}

func (c *Processor) op2AC9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l A1,(A5)+\n", pc)
}

func (c *Processor) op2ACA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l A2,(A5)+\n", pc)
}

func (c *Processor) op2ACB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l A3,(A5)+\n", pc)
}

func (c *Processor) op2ACC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l A4,(A5)+\n", pc)
}

func (c *Processor) op2ACD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l A5,(A5)+\n", pc)
}

func (c *Processor) op2ACE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l A6,(A5)+\n", pc)
}

func (c *Processor) op2ACF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l A7,(A5)+\n", pc)
}

func (c *Processor) op2AD0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l (A0),(A5)+\n", pc)
}

func (c *Processor) op2AD1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l (A1),(A5)+\n", pc)
}

func (c *Processor) op2AD2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l (A2),(A5)+\n", pc)
}

func (c *Processor) op2AD3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l (A3),(A5)+\n", pc)
}

func (c *Processor) op2AD4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l (A4),(A5)+\n", pc)
}

func (c *Processor) op2AD5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l (A5),(A5)+\n", pc)
}

func (c *Processor) op2AD6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l (A6),(A5)+\n", pc)
}

func (c *Processor) op2AD7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l (A7),(A5)+\n", pc)
}

func (c *Processor) op2AF8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l $%X,(A5)+\n", pc, v)
}

func (c *Processor) op2AF9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l $%X,(A5)+\n", pc, v)
}

func (c *Processor) op2AFC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[5], uint32(v))
	c.A[5] += 4
	c.tracef("%04X move.l #$%X,(A5)+\n", pc, v)
}

func (c *Processor) op2B00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D0,-(A5)\n", pc)
}

func (c *Processor) op2B01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D1,-(A5)\n", pc)
}

func (c *Processor) op2B02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D2,-(A5)\n", pc)
}

func (c *Processor) op2B03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D3,-(A5)\n", pc)
}

func (c *Processor) op2B04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D4,-(A5)\n", pc)
}

func (c *Processor) op2B05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D5,-(A5)\n", pc)
}

func (c *Processor) op2B06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D6,-(A5)\n", pc)
}

func (c *Processor) op2B07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l D7,-(A5)\n", pc)
}

func (c *Processor) op2B08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A0,-(A5)\n", pc)
}

func (c *Processor) op2B09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A1,-(A5)\n", pc)
}

func (c *Processor) op2B0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A2,-(A5)\n", pc)
}

func (c *Processor) op2B0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A3,-(A5)\n", pc)
}

func (c *Processor) op2B0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A4,-(A5)\n", pc)
}

func (c *Processor) op2B0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A5,-(A5)\n", pc)
}

func (c *Processor) op2B0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A6,-(A5)\n", pc)
}

func (c *Processor) op2B0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l A7,-(A5)\n", pc)
}

func (c *Processor) op2B10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A0),-(A5)\n", pc)
}

func (c *Processor) op2B11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A1),-(A5)\n", pc)
}

func (c *Processor) op2B12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A2),-(A5)\n", pc)
}

func (c *Processor) op2B13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A3),-(A5)\n", pc)
}

func (c *Processor) op2B14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A4),-(A5)\n", pc)
}

func (c *Processor) op2B15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A5),-(A5)\n", pc)
}

func (c *Processor) op2B16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A6),-(A5)\n", pc)
}

func (c *Processor) op2B17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l (A7),-(A5)\n", pc)
}

func (c *Processor) op2B38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l $%X,-(A5)\n", pc, v)
}

func (c *Processor) op2B39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l $%X,-(A5)\n", pc, v)
}

func (c *Processor) op2B3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[5] -= 4
	c.writeLong(c.A[5], uint32(v))
	c.tracef("%04X move.l #$%X,-(A5)\n", pc, v)
}

func (c *Processor) op2B40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D0,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D1,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D2,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D3,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D4,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D5,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D6,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D7,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A0,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A1,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A2,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A3,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A4,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A5,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A6,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A7,(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A0),(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A1),(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A2),(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A3),(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A4),(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A5),(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A6),(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A7),(%d,A5)\n", pc, disp)
}

func (c *Processor) op2B78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A5)\n", pc, v, disp)
}

func (c *Processor) op2B79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A5)\n", pc, v, disp)
}

func (c *Processor) op2B7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[5]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l #$%X,(%d,A5)\n", pc, v, disp)
}

func (c *Processor) op2C00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l D0,D6\n", pc)
}

func (c *Processor) op2C01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l D1,D6\n", pc)
}

func (c *Processor) op2C02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l D2,D6\n", pc)
}

func (c *Processor) op2C03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l D3,D6\n", pc)
}

func (c *Processor) op2C04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l D4,D6\n", pc)
}

func (c *Processor) op2C05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l D5,D6\n", pc)
}

func (c *Processor) op2C06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l D6,D6\n", pc)
}

func (c *Processor) op2C07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l D7,D6\n", pc)
}

func (c *Processor) op2C08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l A0,D6\n", pc)
}

func (c *Processor) op2C09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l A1,D6\n", pc)
}

func (c *Processor) op2C0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l A2,D6\n", pc)
}

func (c *Processor) op2C0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l A3,D6\n", pc)
}

func (c *Processor) op2C0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l A4,D6\n", pc)
}

func (c *Processor) op2C0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l A5,D6\n", pc)
}

func (c *Processor) op2C0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l A6,D6\n", pc)
}

func (c *Processor) op2C0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[6] = uint32(v)
	c.tracef("%04X move.l A7,D6\n", pc)
}

func (c *Processor) op2C10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[6] = uint32(v)
	c.tracef("%04X move.l (A0),D6\n", pc)
}

func (c *Processor) op2C11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[6] = uint32(v)
	c.tracef("%04X move.l (A1),D6\n", pc)
}

func (c *Processor) op2C12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[6] = uint32(v)
	c.tracef("%04X move.l (A2),D6\n", pc)
}

func (c *Processor) op2C13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[6] = uint32(v)
	c.tracef("%04X move.l (A3),D6\n", pc)
}

func (c *Processor) op2C14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[6] = uint32(v)
	c.tracef("%04X move.l (A4),D6\n", pc)
}

func (c *Processor) op2C15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[6] = uint32(v)
	c.tracef("%04X move.l (A5),D6\n", pc)
}

func (c *Processor) op2C16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[6] = uint32(v)
	c.tracef("%04X move.l (A6),D6\n", pc)
}

func (c *Processor) op2C17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[6] = uint32(v)
	c.tracef("%04X move.l (A7),D6\n", pc)
}

func (c *Processor) op2C38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[6] = uint32(v)
	c.tracef("%04X move.l $%X,D6\n", pc, v)
}

func (c *Processor) op2C39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[6] = uint32(v)
	c.tracef("%04X move.l $%X,D6\n", pc, v)
}

func (c *Processor) op2C3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[6] = uint32(v)
	c.tracef("%04X move.l #$%X,D6\n", pc, v)
}

func (c *Processor) op2C40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l D0,A6\n", pc)
}

func (c *Processor) op2C41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l D1,A6\n", pc)
}

func (c *Processor) op2C42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l D2,A6\n", pc)
}

func (c *Processor) op2C43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l D3,A6\n", pc)
}

func (c *Processor) op2C44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l D4,A6\n", pc)
}

func (c *Processor) op2C45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l D5,A6\n", pc)
}

func (c *Processor) op2C46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l D6,A6\n", pc)
}

func (c *Processor) op2C47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l D7,A6\n", pc)
}

func (c *Processor) op2C48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l A0,A6\n", pc)
}

func (c *Processor) op2C49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l A1,A6\n", pc)
}

func (c *Processor) op2C4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l A2,A6\n", pc)
}

func (c *Processor) op2C4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l A3,A6\n", pc)
}

func (c *Processor) op2C4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l A4,A6\n", pc)
}

func (c *Processor) op2C4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l A5,A6\n", pc)
}

func (c *Processor) op2C4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l A6,A6\n", pc)
}

func (c *Processor) op2C4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l A7,A6\n", pc)
}

func (c *Processor) op2C50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l (A0),A6\n", pc)
}

func (c *Processor) op2C51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l (A1),A6\n", pc)
}

func (c *Processor) op2C52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l (A2),A6\n", pc)
}

func (c *Processor) op2C53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l (A3),A6\n", pc)
}

func (c *Processor) op2C54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l (A4),A6\n", pc)
}

func (c *Processor) op2C55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l (A5),A6\n", pc)
}

func (c *Processor) op2C56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l (A6),A6\n", pc)
}

func (c *Processor) op2C57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l (A7),A6\n", pc)
}

func (c *Processor) op2C78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l $%X,A6\n", pc, v)
}

func (c *Processor) op2C79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l $%X,A6\n", pc, v)
}

func (c *Processor) op2C7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] = uint32(v)
	c.tracef("%04X movea.l #$%X,A6\n", pc, v)
}

func (c *Processor) op2C80() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D0,(A6)\n", pc)
}

func (c *Processor) op2C81() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D1,(A6)\n", pc)
}

func (c *Processor) op2C82() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D2,(A6)\n", pc)
}

func (c *Processor) op2C83() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D3,(A6)\n", pc)
}

func (c *Processor) op2C84() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D4,(A6)\n", pc)
}

func (c *Processor) op2C85() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D5,(A6)\n", pc)
}

func (c *Processor) op2C86() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D6,(A6)\n", pc)
}

func (c *Processor) op2C87() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D7,(A6)\n", pc)
}

func (c *Processor) op2C88() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A0,(A6)\n", pc)
}

func (c *Processor) op2C89() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A1,(A6)\n", pc)
}

func (c *Processor) op2C8A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A2,(A6)\n", pc)
}

func (c *Processor) op2C8B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A3,(A6)\n", pc)
}

func (c *Processor) op2C8C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A4,(A6)\n", pc)
}

func (c *Processor) op2C8D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A5,(A6)\n", pc)
}

func (c *Processor) op2C8E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A6,(A6)\n", pc)
}

func (c *Processor) op2C8F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A7,(A6)\n", pc)
}

func (c *Processor) op2C90() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A0),(A6)\n", pc)
}

func (c *Processor) op2C91() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A1),(A6)\n", pc)
}

func (c *Processor) op2C92() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A2),(A6)\n", pc)
}

func (c *Processor) op2C93() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A3),(A6)\n", pc)
}

func (c *Processor) op2C94() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A4),(A6)\n", pc)
}

func (c *Processor) op2C95() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A5),(A6)\n", pc)
}

func (c *Processor) op2C96() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A6),(A6)\n", pc)
}

func (c *Processor) op2C97() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A7),(A6)\n", pc)
}

func (c *Processor) op2CB8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l $%X,(A6)\n", pc, v)
}

func (c *Processor) op2CB9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l $%X,(A6)\n", pc, v)
}

func (c *Processor) op2CBC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l #$%X,(A6)\n", pc, v)
}

func (c *Processor) op2CC0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l D0,(A6)+\n", pc)
}

func (c *Processor) op2CC1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l D1,(A6)+\n", pc)
}

func (c *Processor) op2CC2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l D2,(A6)+\n", pc)
}

func (c *Processor) op2CC3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l D3,(A6)+\n", pc)
}

func (c *Processor) op2CC4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l D4,(A6)+\n", pc)
}

func (c *Processor) op2CC5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l D5,(A6)+\n", pc)
}

func (c *Processor) op2CC6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l D6,(A6)+\n", pc)
}

func (c *Processor) op2CC7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l D7,(A6)+\n", pc)
}

func (c *Processor) op2CC8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l A0,(A6)+\n", pc)
}

func (c *Processor) op2CC9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l A1,(A6)+\n", pc)
}

func (c *Processor) op2CCA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l A2,(A6)+\n", pc)
}

func (c *Processor) op2CCB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l A3,(A6)+\n", pc)
}

func (c *Processor) op2CCC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l A4,(A6)+\n", pc)
}

func (c *Processor) op2CCD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l A5,(A6)+\n", pc)
}

func (c *Processor) op2CCE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l A6,(A6)+\n", pc)
}

func (c *Processor) op2CCF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l A7,(A6)+\n", pc)
}

func (c *Processor) op2CD0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l (A0),(A6)+\n", pc)
}

func (c *Processor) op2CD1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l (A1),(A6)+\n", pc)
}

func (c *Processor) op2CD2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l (A2),(A6)+\n", pc)
}

func (c *Processor) op2CD3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l (A3),(A6)+\n", pc)
}

func (c *Processor) op2CD4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l (A4),(A6)+\n", pc)
}

func (c *Processor) op2CD5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l (A5),(A6)+\n", pc)
}

func (c *Processor) op2CD6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l (A6),(A6)+\n", pc)
}

func (c *Processor) op2CD7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l (A7),(A6)+\n", pc)
}

func (c *Processor) op2CF8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l $%X,(A6)+\n", pc, v)
}

func (c *Processor) op2CF9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l $%X,(A6)+\n", pc, v)
}

func (c *Processor) op2CFC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[6], uint32(v))
	c.A[6] += 4
	c.tracef("%04X move.l #$%X,(A6)+\n", pc, v)
}

func (c *Processor) op2D00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D0,-(A6)\n", pc)
}

func (c *Processor) op2D01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D1,-(A6)\n", pc)
}

func (c *Processor) op2D02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D2,-(A6)\n", pc)
}

func (c *Processor) op2D03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D3,-(A6)\n", pc)
}

func (c *Processor) op2D04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D4,-(A6)\n", pc)
}

func (c *Processor) op2D05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D5,-(A6)\n", pc)
}

func (c *Processor) op2D06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D6,-(A6)\n", pc)
}

func (c *Processor) op2D07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l D7,-(A6)\n", pc)
}

func (c *Processor) op2D08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A0,-(A6)\n", pc)
}

func (c *Processor) op2D09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A1,-(A6)\n", pc)
}

func (c *Processor) op2D0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A2,-(A6)\n", pc)
}

func (c *Processor) op2D0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A3,-(A6)\n", pc)
}

func (c *Processor) op2D0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A4,-(A6)\n", pc)
}

func (c *Processor) op2D0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A5,-(A6)\n", pc)
}

func (c *Processor) op2D0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A6,-(A6)\n", pc)
}

func (c *Processor) op2D0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l A7,-(A6)\n", pc)
}

func (c *Processor) op2D10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A0),-(A6)\n", pc)
}

func (c *Processor) op2D11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A1),-(A6)\n", pc)
}

func (c *Processor) op2D12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A2),-(A6)\n", pc)
}

func (c *Processor) op2D13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A3),-(A6)\n", pc)
}

func (c *Processor) op2D14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A4),-(A6)\n", pc)
}

func (c *Processor) op2D15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A5),-(A6)\n", pc)
}

func (c *Processor) op2D16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A6),-(A6)\n", pc)
}

func (c *Processor) op2D17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l (A7),-(A6)\n", pc)
}

func (c *Processor) op2D38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l $%X,-(A6)\n", pc, v)
}

func (c *Processor) op2D39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l $%X,-(A6)\n", pc, v)
}

func (c *Processor) op2D3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[6] -= 4
	c.writeLong(c.A[6], uint32(v))
	c.tracef("%04X move.l #$%X,-(A6)\n", pc, v)
}

func (c *Processor) op2D40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D0,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D1,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D2,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D3,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D4,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D5,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D6,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D7,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A0,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A1,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A2,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A3,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A4,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A5,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A6,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A7,(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A0),(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A1),(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A2),(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A3),(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A4),(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A5),(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A6),(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A7),(%d,A6)\n", pc, disp)
}

func (c *Processor) op2D78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A6)\n", pc, v, disp)
}

func (c *Processor) op2D79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A6)\n", pc, v, disp)
}

func (c *Processor) op2D7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[6]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l #$%X,(%d,A6)\n", pc, v, disp)
}

func (c *Processor) op2E00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l D0,D7\n", pc)
}

func (c *Processor) op2E01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l D1,D7\n", pc)
}

func (c *Processor) op2E02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l D2,D7\n", pc)
}

func (c *Processor) op2E03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l D3,D7\n", pc)
}

func (c *Processor) op2E04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l D4,D7\n", pc)
}

func (c *Processor) op2E05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l D5,D7\n", pc)
}

func (c *Processor) op2E06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l D6,D7\n", pc)
}

func (c *Processor) op2E07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l D7,D7\n", pc)
}

func (c *Processor) op2E08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l A0,D7\n", pc)
}

func (c *Processor) op2E09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l A1,D7\n", pc)
}

func (c *Processor) op2E0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l A2,D7\n", pc)
}

func (c *Processor) op2E0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l A3,D7\n", pc)
}

func (c *Processor) op2E0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l A4,D7\n", pc)
}

func (c *Processor) op2E0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l A5,D7\n", pc)
}

func (c *Processor) op2E0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l A6,D7\n", pc)
}

func (c *Processor) op2E0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.D[7] = uint32(v)
	c.tracef("%04X move.l A7,D7\n", pc)
}

func (c *Processor) op2E10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[7] = uint32(v)
	c.tracef("%04X move.l (A0),D7\n", pc)
}

func (c *Processor) op2E11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[7] = uint32(v)
	c.tracef("%04X move.l (A1),D7\n", pc)
}

func (c *Processor) op2E12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[7] = uint32(v)
	c.tracef("%04X move.l (A2),D7\n", pc)
}

func (c *Processor) op2E13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[7] = uint32(v)
	c.tracef("%04X move.l (A3),D7\n", pc)
}

func (c *Processor) op2E14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[7] = uint32(v)
	c.tracef("%04X move.l (A4),D7\n", pc)
}

func (c *Processor) op2E15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[7] = uint32(v)
	c.tracef("%04X move.l (A5),D7\n", pc)
}

func (c *Processor) op2E16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[7] = uint32(v)
	c.tracef("%04X move.l (A6),D7\n", pc)
}

func (c *Processor) op2E17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.D[7] = uint32(v)
	c.tracef("%04X move.l (A7),D7\n", pc)
}

func (c *Processor) op2E38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[7] = uint32(v)
	c.tracef("%04X move.l $%X,D7\n", pc, v)
}

func (c *Processor) op2E39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[7] = uint32(v)
	c.tracef("%04X move.l $%X,D7\n", pc, v)
}

func (c *Processor) op2E3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.D[7] = uint32(v)
	c.tracef("%04X move.l #$%X,D7\n", pc, v)
}

func (c *Processor) op2E40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l D0,A7\n", pc)
}

func (c *Processor) op2E41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l D1,A7\n", pc)
}

func (c *Processor) op2E42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l D2,A7\n", pc)
}

func (c *Processor) op2E43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l D3,A7\n", pc)
}

func (c *Processor) op2E44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l D4,A7\n", pc)
}

func (c *Processor) op2E45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l D5,A7\n", pc)
}

func (c *Processor) op2E46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l D6,A7\n", pc)
}

func (c *Processor) op2E47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l D7,A7\n", pc)
}

func (c *Processor) op2E48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l A0,A7\n", pc)
}

func (c *Processor) op2E49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l A1,A7\n", pc)
}

func (c *Processor) op2E4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l A2,A7\n", pc)
}

func (c *Processor) op2E4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l A3,A7\n", pc)
}

func (c *Processor) op2E4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l A4,A7\n", pc)
}

func (c *Processor) op2E4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l A5,A7\n", pc)
}

func (c *Processor) op2E4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l A6,A7\n", pc)
}

func (c *Processor) op2E4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l A7,A7\n", pc)
}

func (c *Processor) op2E50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l (A0),A7\n", pc)
}

func (c *Processor) op2E51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l (A1),A7\n", pc)
}

func (c *Processor) op2E52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l (A2),A7\n", pc)
}

func (c *Processor) op2E53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l (A3),A7\n", pc)
}

func (c *Processor) op2E54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l (A4),A7\n", pc)
}

func (c *Processor) op2E55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l (A5),A7\n", pc)
}

func (c *Processor) op2E56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l (A6),A7\n", pc)
}

func (c *Processor) op2E57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l (A7),A7\n", pc)
}

func (c *Processor) op2E78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l $%X,A7\n", pc, v)
}

func (c *Processor) op2E79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l $%X,A7\n", pc, v)
}

func (c *Processor) op2E7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] = uint32(v)
	c.tracef("%04X movea.l #$%X,A7\n", pc, v)
}

func (c *Processor) op2E80() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D0,(A7)\n", pc)
}

func (c *Processor) op2E81() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D1,(A7)\n", pc)
}

func (c *Processor) op2E82() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D2,(A7)\n", pc)
}

func (c *Processor) op2E83() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D3,(A7)\n", pc)
}

func (c *Processor) op2E84() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D4,(A7)\n", pc)
}

func (c *Processor) op2E85() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D5,(A7)\n", pc)
}

func (c *Processor) op2E86() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D6,(A7)\n", pc)
}

func (c *Processor) op2E87() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D7,(A7)\n", pc)
}

func (c *Processor) op2E88() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A0,(A7)\n", pc)
}

func (c *Processor) op2E89() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A1,(A7)\n", pc)
}

func (c *Processor) op2E8A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A2,(A7)\n", pc)
}

func (c *Processor) op2E8B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A3,(A7)\n", pc)
}

func (c *Processor) op2E8C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A4,(A7)\n", pc)
}

func (c *Processor) op2E8D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A5,(A7)\n", pc)
}

func (c *Processor) op2E8E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A6,(A7)\n", pc)
}

func (c *Processor) op2E8F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A7,(A7)\n", pc)
}

func (c *Processor) op2E90() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A0),(A7)\n", pc)
}

func (c *Processor) op2E91() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A1),(A7)\n", pc)
}

func (c *Processor) op2E92() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A2),(A7)\n", pc)
}

func (c *Processor) op2E93() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A3),(A7)\n", pc)
}

func (c *Processor) op2E94() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A4),(A7)\n", pc)
}

func (c *Processor) op2E95() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A5),(A7)\n", pc)
}

func (c *Processor) op2E96() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A6),(A7)\n", pc)
}

func (c *Processor) op2E97() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A7),(A7)\n", pc)
}

func (c *Processor) op2EB8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l $%X,(A7)\n", pc, v)
}

func (c *Processor) op2EB9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l $%X,(A7)\n", pc, v)
}

func (c *Processor) op2EBC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l #$%X,(A7)\n", pc, v)
}

func (c *Processor) op2EC0() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l D0,(A7)+\n", pc)
}

func (c *Processor) op2EC1() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l D1,(A7)+\n", pc)
}

func (c *Processor) op2EC2() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l D2,(A7)+\n", pc)
}

func (c *Processor) op2EC3() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l D3,(A7)+\n", pc)
}

func (c *Processor) op2EC4() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l D4,(A7)+\n", pc)
}

func (c *Processor) op2EC5() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l D5,(A7)+\n", pc)
}

func (c *Processor) op2EC6() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l D6,(A7)+\n", pc)
}

func (c *Processor) op2EC7() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l D7,(A7)+\n", pc)
}

func (c *Processor) op2EC8() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l A0,(A7)+\n", pc)
}

func (c *Processor) op2EC9() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l A1,(A7)+\n", pc)
}

func (c *Processor) op2ECA() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l A2,(A7)+\n", pc)
}

func (c *Processor) op2ECB() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l A3,(A7)+\n", pc)
}

func (c *Processor) op2ECC() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l A4,(A7)+\n", pc)
}

func (c *Processor) op2ECD() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l A5,(A7)+\n", pc)
}

func (c *Processor) op2ECE() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l A6,(A7)+\n", pc)
}

func (c *Processor) op2ECF() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l A7,(A7)+\n", pc)
}

func (c *Processor) op2ED0() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l (A0),(A7)+\n", pc)
}

func (c *Processor) op2ED1() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l (A1),(A7)+\n", pc)
}

func (c *Processor) op2ED2() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l (A2),(A7)+\n", pc)
}

func (c *Processor) op2ED3() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l (A3),(A7)+\n", pc)
}

func (c *Processor) op2ED4() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l (A4),(A7)+\n", pc)
}

func (c *Processor) op2ED5() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l (A5),(A7)+\n", pc)
}

func (c *Processor) op2ED6() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l (A6),(A7)+\n", pc)
}

func (c *Processor) op2ED7() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l (A7),(A7)+\n", pc)
}

func (c *Processor) op2EF8() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l $%X,(A7)+\n", pc, v)
}

func (c *Processor) op2EF9() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l $%X,(A7)+\n", pc, v)
}

func (c *Processor) op2EFC() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.writeLong(c.A[7], uint32(v))
	c.A[7] += 4
	c.tracef("%04X move.l #$%X,(A7)+\n", pc, v)
}

func (c *Processor) op2F00() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D0,-(A7)\n", pc)
}

func (c *Processor) op2F01() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D1,-(A7)\n", pc)
}

func (c *Processor) op2F02() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D2,-(A7)\n", pc)
}

func (c *Processor) op2F03() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D3,-(A7)\n", pc)
}

func (c *Processor) op2F04() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D4,-(A7)\n", pc)
}

func (c *Processor) op2F05() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D5,-(A7)\n", pc)
}

func (c *Processor) op2F06() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D6,-(A7)\n", pc)
}

func (c *Processor) op2F07() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l D7,-(A7)\n", pc)
}

func (c *Processor) op2F08() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A0,-(A7)\n", pc)
}

func (c *Processor) op2F09() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A1,-(A7)\n", pc)
}

func (c *Processor) op2F0A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A2,-(A7)\n", pc)
}

func (c *Processor) op2F0B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A3,-(A7)\n", pc)
}

func (c *Processor) op2F0C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A4,-(A7)\n", pc)
}

func (c *Processor) op2F0D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A5,-(A7)\n", pc)
}

func (c *Processor) op2F0E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A6,-(A7)\n", pc)
}

func (c *Processor) op2F0F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l A7,-(A7)\n", pc)
}

func (c *Processor) op2F10() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A0),-(A7)\n", pc)
}

func (c *Processor) op2F11() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A1),-(A7)\n", pc)
}

func (c *Processor) op2F12() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A2),-(A7)\n", pc)
}

func (c *Processor) op2F13() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A3),-(A7)\n", pc)
}

func (c *Processor) op2F14() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A4),-(A7)\n", pc)
}

func (c *Processor) op2F15() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A5),-(A7)\n", pc)
}

func (c *Processor) op2F16() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A6),-(A7)\n", pc)
}

func (c *Processor) op2F17() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l (A7),-(A7)\n", pc)
}

func (c *Processor) op2F38() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l $%X,-(A7)\n", pc, v)
}

func (c *Processor) op2F39() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l $%X,-(A7)\n", pc, v)
}

func (c *Processor) op2F3C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	c.A[7] -= 4
	c.writeLong(c.A[7], uint32(v))
	c.tracef("%04X move.l #$%X,-(A7)\n", pc, v)
}

func (c *Processor) op2F40() {
	pc := c.PC
	c.PC += 2
	v := c.D[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D0,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F41() {
	pc := c.PC
	c.PC += 2
	v := c.D[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D1,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F42() {
	pc := c.PC
	c.PC += 2
	v := c.D[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D2,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F43() {
	pc := c.PC
	c.PC += 2
	v := c.D[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D3,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F44() {
	pc := c.PC
	c.PC += 2
	v := c.D[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D4,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F45() {
	pc := c.PC
	c.PC += 2
	v := c.D[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D5,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F46() {
	pc := c.PC
	c.PC += 2
	v := c.D[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D6,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F47() {
	pc := c.PC
	c.PC += 2
	v := c.D[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l D7,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F48() {
	pc := c.PC
	c.PC += 2
	v := c.A[0]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A0,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F49() {
	pc := c.PC
	c.PC += 2
	v := c.A[1]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A1,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F4A() {
	pc := c.PC
	c.PC += 2
	v := c.A[2]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A2,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F4B() {
	pc := c.PC
	c.PC += 2
	v := c.A[3]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A3,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F4C() {
	pc := c.PC
	c.PC += 2
	v := c.A[4]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A4,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F4D() {
	pc := c.PC
	c.PC += 2
	v := c.A[5]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A5,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F4E() {
	pc := c.PC
	c.PC += 2
	v := c.A[6]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A6,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F4F() {
	pc := c.PC
	c.PC += 2
	v := c.A[7]
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l A7,(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F50() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[0]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A0),(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F51() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[1]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A1),(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F52() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[2]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A2),(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F53() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[3]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A3),(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F54() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[4]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A4),(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F55() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[5]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A5),(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F56() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[6]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A6),(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F57() {
	pc := c.PC
	c.PC += 2
	_, c.err = c.M.Read(int(c.A[7]), c.buf[:4])
	if c.err != nil {
		return
	}
	v := uint32(c.buf[3]) | uint32(c.buf[2])<<8 | uint32(c.buf[1])<<16 | uint32(c.buf[0])<<24
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l (A7),(%d,A7)\n", pc, disp)
}

func (c *Processor) op2F78() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A7)\n", pc, v, disp)
}

func (c *Processor) op2F79() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l $%X,(%d,A7)\n", pc, v, disp)
}

func (c *Processor) op2F7C() {
	pc := c.PC
	c.PC += 2
	v := c.readImm(0x02)
	disp := c.readImm(0x01)
	addr := disp + c.A[7]
	c.writeLong(addr, uint32(v))
	c.tracef("%04X move.l #$%X,(%d,A7)\n", pc, v, disp)
}

func (c *Processor) mapFn(op uint16) func() {
	table := map[uint16]func(){
		0x0000: c.op0000,
		0x0001: c.op0001,
		0x0002: c.op0002,
		0x0003: c.op0003,
		0x0004: c.op0004,
		0x0005: c.op0005,
		0x0006: c.op0006,
		0x0007: c.op0007,
		0x0008: c.op0008,
		0x0009: c.op0009,
		0x000A: c.op000A,
		0x000B: c.op000B,
		0x000C: c.op000C,
		0x000D: c.op000D,
		0x000E: c.op000E,
		0x000F: c.op000F,
		0x0010: c.op0010,
		0x0011: c.op0011,
		0x0012: c.op0012,
		0x0013: c.op0013,
		0x0014: c.op0014,
		0x0015: c.op0015,
		0x0016: c.op0016,
		0x0017: c.op0017,
		0x0038: c.op0038,
		0x0039: c.op0039,
		0x003C: c.op003C,
		0x0040: c.op0040,
		0x0041: c.op0041,
		0x0042: c.op0042,
		0x0043: c.op0043,
		0x0044: c.op0044,
		0x0045: c.op0045,
		0x0046: c.op0046,
		0x0047: c.op0047,
		0x0048: c.op0048,
		0x0049: c.op0049,
		0x004A: c.op004A,
		0x004B: c.op004B,
		0x004C: c.op004C,
		0x004D: c.op004D,
		0x004E: c.op004E,
		0x004F: c.op004F,
		0x0050: c.op0050,
		0x0051: c.op0051,
		0x0052: c.op0052,
		0x0053: c.op0053,
		0x0054: c.op0054,
		0x0055: c.op0055,
		0x0056: c.op0056,
		0x0057: c.op0057,
		0x0078: c.op0078,
		0x0079: c.op0079,
		0x007C: c.op007C,
		0x0080: c.op0080,
		0x0081: c.op0081,
		0x0082: c.op0082,
		0x0083: c.op0083,
		0x0084: c.op0084,
		0x0085: c.op0085,
		0x0086: c.op0086,
		0x0087: c.op0087,
		0x0088: c.op0088,
		0x0089: c.op0089,
		0x008A: c.op008A,
		0x008B: c.op008B,
		0x008C: c.op008C,
		0x008D: c.op008D,
		0x008E: c.op008E,
		0x008F: c.op008F,
		0x0090: c.op0090,
		0x0091: c.op0091,
		0x0092: c.op0092,
		0x0093: c.op0093,
		0x0094: c.op0094,
		0x0095: c.op0095,
		0x0096: c.op0096,
		0x0097: c.op0097,
		0x00B8: c.op00B8,
		0x00B9: c.op00B9,
		0x00BC: c.op00BC,
		0x00C0: c.op00C0,
		0x00C1: c.op00C1,
		0x00C2: c.op00C2,
		0x00C3: c.op00C3,
		0x00C4: c.op00C4,
		0x00C5: c.op00C5,
		0x00C6: c.op00C6,
		0x00C7: c.op00C7,
		0x00C8: c.op00C8,
		0x00C9: c.op00C9,
		0x00CA: c.op00CA,
		0x00CB: c.op00CB,
		0x00CC: c.op00CC,
		0x00CD: c.op00CD,
		0x00CE: c.op00CE,
		0x00CF: c.op00CF,
		0x00D0: c.op00D0,
		0x00D1: c.op00D1,
		0x00D2: c.op00D2,
		0x00D3: c.op00D3,
		0x00D4: c.op00D4,
		0x00D5: c.op00D5,
		0x00D6: c.op00D6,
		0x00D7: c.op00D7,
		0x00F8: c.op00F8,
		0x00F9: c.op00F9,
		0x00FC: c.op00FC,
		0x0100: c.op0100,
		0x0101: c.op0101,
		0x0102: c.op0102,
		0x0103: c.op0103,
		0x0104: c.op0104,
		0x0105: c.op0105,
		0x0106: c.op0106,
		0x0107: c.op0107,
		0x0108: c.op0108,
		0x0109: c.op0109,
		0x010A: c.op010A,
		0x010B: c.op010B,
		0x010C: c.op010C,
		0x010D: c.op010D,
		0x010E: c.op010E,
		0x010F: c.op010F,
		0x0110: c.op0110,
		0x0111: c.op0111,
		0x0112: c.op0112,
		0x0113: c.op0113,
		0x0114: c.op0114,
		0x0115: c.op0115,
		0x0116: c.op0116,
		0x0117: c.op0117,
		0x0138: c.op0138,
		0x0139: c.op0139,
		0x013C: c.op013C,
		0x0140: c.op0140,
		0x0141: c.op0141,
		0x0142: c.op0142,
		0x0143: c.op0143,
		0x0144: c.op0144,
		0x0145: c.op0145,
		0x0146: c.op0146,
		0x0147: c.op0147,
		0x0148: c.op0148,
		0x0149: c.op0149,
		0x014A: c.op014A,
		0x014B: c.op014B,
		0x014C: c.op014C,
		0x014D: c.op014D,
		0x014E: c.op014E,
		0x014F: c.op014F,
		0x0150: c.op0150,
		0x0151: c.op0151,
		0x0152: c.op0152,
		0x0153: c.op0153,
		0x0154: c.op0154,
		0x0155: c.op0155,
		0x0156: c.op0156,
		0x0157: c.op0157,
		0x0178: c.op0178,
		0x0179: c.op0179,
		0x017C: c.op017C,
		0x01C0: c.op01C0,
		0x01C1: c.op01C1,
		0x01C2: c.op01C2,
		0x01C3: c.op01C3,
		0x01C4: c.op01C4,
		0x01C5: c.op01C5,
		0x01C6: c.op01C6,
		0x01C7: c.op01C7,
		0x01C8: c.op01C8,
		0x01C9: c.op01C9,
		0x01CA: c.op01CA,
		0x01CB: c.op01CB,
		0x01CC: c.op01CC,
		0x01CD: c.op01CD,
		0x01CE: c.op01CE,
		0x01CF: c.op01CF,
		0x01D0: c.op01D0,
		0x01D1: c.op01D1,
		0x01D2: c.op01D2,
		0x01D3: c.op01D3,
		0x01D4: c.op01D4,
		0x01D5: c.op01D5,
		0x01D6: c.op01D6,
		0x01D7: c.op01D7,
		0x01F8: c.op01F8,
		0x01F9: c.op01F9,
		0x01FC: c.op01FC,
		0x0200: c.op0200,
		0x0201: c.op0201,
		0x0202: c.op0202,
		0x0203: c.op0203,
		0x0204: c.op0204,
		0x0205: c.op0205,
		0x0206: c.op0206,
		0x0207: c.op0207,
		0x0208: c.op0208,
		0x0209: c.op0209,
		0x020A: c.op020A,
		0x020B: c.op020B,
		0x020C: c.op020C,
		0x020D: c.op020D,
		0x020E: c.op020E,
		0x020F: c.op020F,
		0x0210: c.op0210,
		0x0211: c.op0211,
		0x0212: c.op0212,
		0x0213: c.op0213,
		0x0214: c.op0214,
		0x0215: c.op0215,
		0x0216: c.op0216,
		0x0217: c.op0217,
		0x0238: c.op0238,
		0x0239: c.op0239,
		0x023C: c.op023C,
		0x0240: c.op0240,
		0x0241: c.op0241,
		0x0242: c.op0242,
		0x0243: c.op0243,
		0x0244: c.op0244,
		0x0245: c.op0245,
		0x0246: c.op0246,
		0x0247: c.op0247,
		0x0248: c.op0248,
		0x0249: c.op0249,
		0x024A: c.op024A,
		0x024B: c.op024B,
		0x024C: c.op024C,
		0x024D: c.op024D,
		0x024E: c.op024E,
		0x024F: c.op024F,
		0x0250: c.op0250,
		0x0251: c.op0251,
		0x0252: c.op0252,
		0x0253: c.op0253,
		0x0254: c.op0254,
		0x0255: c.op0255,
		0x0256: c.op0256,
		0x0257: c.op0257,
		0x0278: c.op0278,
		0x0279: c.op0279,
		0x027C: c.op027C,
		0x0280: c.op0280,
		0x0281: c.op0281,
		0x0282: c.op0282,
		0x0283: c.op0283,
		0x0284: c.op0284,
		0x0285: c.op0285,
		0x0286: c.op0286,
		0x0287: c.op0287,
		0x0288: c.op0288,
		0x0289: c.op0289,
		0x028A: c.op028A,
		0x028B: c.op028B,
		0x028C: c.op028C,
		0x028D: c.op028D,
		0x028E: c.op028E,
		0x028F: c.op028F,
		0x0290: c.op0290,
		0x0291: c.op0291,
		0x0292: c.op0292,
		0x0293: c.op0293,
		0x0294: c.op0294,
		0x0295: c.op0295,
		0x0296: c.op0296,
		0x0297: c.op0297,
		0x02B8: c.op02B8,
		0x02B9: c.op02B9,
		0x02BC: c.op02BC,
		0x02C0: c.op02C0,
		0x02C1: c.op02C1,
		0x02C2: c.op02C2,
		0x02C3: c.op02C3,
		0x02C4: c.op02C4,
		0x02C5: c.op02C5,
		0x02C6: c.op02C6,
		0x02C7: c.op02C7,
		0x02C8: c.op02C8,
		0x02C9: c.op02C9,
		0x02CA: c.op02CA,
		0x02CB: c.op02CB,
		0x02CC: c.op02CC,
		0x02CD: c.op02CD,
		0x02CE: c.op02CE,
		0x02CF: c.op02CF,
		0x02D0: c.op02D0,
		0x02D1: c.op02D1,
		0x02D2: c.op02D2,
		0x02D3: c.op02D3,
		0x02D4: c.op02D4,
		0x02D5: c.op02D5,
		0x02D6: c.op02D6,
		0x02D7: c.op02D7,
		0x02F8: c.op02F8,
		0x02F9: c.op02F9,
		0x02FC: c.op02FC,
		0x0300: c.op0300,
		0x0301: c.op0301,
		0x0302: c.op0302,
		0x0303: c.op0303,
		0x0304: c.op0304,
		0x0305: c.op0305,
		0x0306: c.op0306,
		0x0307: c.op0307,
		0x0308: c.op0308,
		0x0309: c.op0309,
		0x030A: c.op030A,
		0x030B: c.op030B,
		0x030C: c.op030C,
		0x030D: c.op030D,
		0x030E: c.op030E,
		0x030F: c.op030F,
		0x0310: c.op0310,
		0x0311: c.op0311,
		0x0312: c.op0312,
		0x0313: c.op0313,
		0x0314: c.op0314,
		0x0315: c.op0315,
		0x0316: c.op0316,
		0x0317: c.op0317,
		0x0338: c.op0338,
		0x0339: c.op0339,
		0x033C: c.op033C,
		0x0340: c.op0340,
		0x0341: c.op0341,
		0x0342: c.op0342,
		0x0343: c.op0343,
		0x0344: c.op0344,
		0x0345: c.op0345,
		0x0346: c.op0346,
		0x0347: c.op0347,
		0x0348: c.op0348,
		0x0349: c.op0349,
		0x034A: c.op034A,
		0x034B: c.op034B,
		0x034C: c.op034C,
		0x034D: c.op034D,
		0x034E: c.op034E,
		0x034F: c.op034F,
		0x0350: c.op0350,
		0x0351: c.op0351,
		0x0352: c.op0352,
		0x0353: c.op0353,
		0x0354: c.op0354,
		0x0355: c.op0355,
		0x0356: c.op0356,
		0x0357: c.op0357,
		0x0378: c.op0378,
		0x0379: c.op0379,
		0x037C: c.op037C,
		0x03C0: c.op03C0,
		0x03C1: c.op03C1,
		0x03C2: c.op03C2,
		0x03C3: c.op03C3,
		0x03C4: c.op03C4,
		0x03C5: c.op03C5,
		0x03C6: c.op03C6,
		0x03C7: c.op03C7,
		0x03C8: c.op03C8,
		0x03C9: c.op03C9,
		0x03CA: c.op03CA,
		0x03CB: c.op03CB,
		0x03CC: c.op03CC,
		0x03CD: c.op03CD,
		0x03CE: c.op03CE,
		0x03CF: c.op03CF,
		0x03D0: c.op03D0,
		0x03D1: c.op03D1,
		0x03D2: c.op03D2,
		0x03D3: c.op03D3,
		0x03D4: c.op03D4,
		0x03D5: c.op03D5,
		0x03D6: c.op03D6,
		0x03D7: c.op03D7,
		0x03F8: c.op03F8,
		0x03F9: c.op03F9,
		0x03FC: c.op03FC,
		0x0400: c.op0400,
		0x0401: c.op0401,
		0x0402: c.op0402,
		0x0403: c.op0403,
		0x0404: c.op0404,
		0x0405: c.op0405,
		0x0406: c.op0406,
		0x0407: c.op0407,
		0x0408: c.op0408,
		0x0409: c.op0409,
		0x040A: c.op040A,
		0x040B: c.op040B,
		0x040C: c.op040C,
		0x040D: c.op040D,
		0x040E: c.op040E,
		0x040F: c.op040F,
		0x0410: c.op0410,
		0x0411: c.op0411,
		0x0412: c.op0412,
		0x0413: c.op0413,
		0x0414: c.op0414,
		0x0415: c.op0415,
		0x0416: c.op0416,
		0x0417: c.op0417,
		0x0438: c.op0438,
		0x0439: c.op0439,
		0x043C: c.op043C,
		0x0440: c.op0440,
		0x0441: c.op0441,
		0x0442: c.op0442,
		0x0443: c.op0443,
		0x0444: c.op0444,
		0x0445: c.op0445,
		0x0446: c.op0446,
		0x0447: c.op0447,
		0x0448: c.op0448,
		0x0449: c.op0449,
		0x044A: c.op044A,
		0x044B: c.op044B,
		0x044C: c.op044C,
		0x044D: c.op044D,
		0x044E: c.op044E,
		0x044F: c.op044F,
		0x0450: c.op0450,
		0x0451: c.op0451,
		0x0452: c.op0452,
		0x0453: c.op0453,
		0x0454: c.op0454,
		0x0455: c.op0455,
		0x0456: c.op0456,
		0x0457: c.op0457,
		0x0478: c.op0478,
		0x0479: c.op0479,
		0x047C: c.op047C,
		0x0480: c.op0480,
		0x0481: c.op0481,
		0x0482: c.op0482,
		0x0483: c.op0483,
		0x0484: c.op0484,
		0x0485: c.op0485,
		0x0486: c.op0486,
		0x0487: c.op0487,
		0x0488: c.op0488,
		0x0489: c.op0489,
		0x048A: c.op048A,
		0x048B: c.op048B,
		0x048C: c.op048C,
		0x048D: c.op048D,
		0x048E: c.op048E,
		0x048F: c.op048F,
		0x0490: c.op0490,
		0x0491: c.op0491,
		0x0492: c.op0492,
		0x0493: c.op0493,
		0x0494: c.op0494,
		0x0495: c.op0495,
		0x0496: c.op0496,
		0x0497: c.op0497,
		0x04B8: c.op04B8,
		0x04B9: c.op04B9,
		0x04BC: c.op04BC,
		0x04C0: c.op04C0,
		0x04C1: c.op04C1,
		0x04C2: c.op04C2,
		0x04C3: c.op04C3,
		0x04C4: c.op04C4,
		0x04C5: c.op04C5,
		0x04C6: c.op04C6,
		0x04C7: c.op04C7,
		0x04C8: c.op04C8,
		0x04C9: c.op04C9,
		0x04CA: c.op04CA,
		0x04CB: c.op04CB,
		0x04CC: c.op04CC,
		0x04CD: c.op04CD,
		0x04CE: c.op04CE,
		0x04CF: c.op04CF,
		0x04D0: c.op04D0,
		0x04D1: c.op04D1,
		0x04D2: c.op04D2,
		0x04D3: c.op04D3,
		0x04D4: c.op04D4,
		0x04D5: c.op04D5,
		0x04D6: c.op04D6,
		0x04D7: c.op04D7,
		0x04F8: c.op04F8,
		0x04F9: c.op04F9,
		0x04FC: c.op04FC,
		0x0500: c.op0500,
		0x0501: c.op0501,
		0x0502: c.op0502,
		0x0503: c.op0503,
		0x0504: c.op0504,
		0x0505: c.op0505,
		0x0506: c.op0506,
		0x0507: c.op0507,
		0x0508: c.op0508,
		0x0509: c.op0509,
		0x050A: c.op050A,
		0x050B: c.op050B,
		0x050C: c.op050C,
		0x050D: c.op050D,
		0x050E: c.op050E,
		0x050F: c.op050F,
		0x0510: c.op0510,
		0x0511: c.op0511,
		0x0512: c.op0512,
		0x0513: c.op0513,
		0x0514: c.op0514,
		0x0515: c.op0515,
		0x0516: c.op0516,
		0x0517: c.op0517,
		0x0538: c.op0538,
		0x0539: c.op0539,
		0x053C: c.op053C,
		0x0540: c.op0540,
		0x0541: c.op0541,
		0x0542: c.op0542,
		0x0543: c.op0543,
		0x0544: c.op0544,
		0x0545: c.op0545,
		0x0546: c.op0546,
		0x0547: c.op0547,
		0x0548: c.op0548,
		0x0549: c.op0549,
		0x054A: c.op054A,
		0x054B: c.op054B,
		0x054C: c.op054C,
		0x054D: c.op054D,
		0x054E: c.op054E,
		0x054F: c.op054F,
		0x0550: c.op0550,
		0x0551: c.op0551,
		0x0552: c.op0552,
		0x0553: c.op0553,
		0x0554: c.op0554,
		0x0555: c.op0555,
		0x0556: c.op0556,
		0x0557: c.op0557,
		0x0578: c.op0578,
		0x0579: c.op0579,
		0x057C: c.op057C,
		0x0600: c.op0600,
		0x0601: c.op0601,
		0x0602: c.op0602,
		0x0603: c.op0603,
		0x0604: c.op0604,
		0x0605: c.op0605,
		0x0606: c.op0606,
		0x0607: c.op0607,
		0x0608: c.op0608,
		0x0609: c.op0609,
		0x060A: c.op060A,
		0x060B: c.op060B,
		0x060C: c.op060C,
		0x060D: c.op060D,
		0x060E: c.op060E,
		0x060F: c.op060F,
		0x0610: c.op0610,
		0x0611: c.op0611,
		0x0612: c.op0612,
		0x0613: c.op0613,
		0x0614: c.op0614,
		0x0615: c.op0615,
		0x0616: c.op0616,
		0x0617: c.op0617,
		0x0638: c.op0638,
		0x0639: c.op0639,
		0x063C: c.op063C,
		0x0640: c.op0640,
		0x0641: c.op0641,
		0x0642: c.op0642,
		0x0643: c.op0643,
		0x0644: c.op0644,
		0x0645: c.op0645,
		0x0646: c.op0646,
		0x0647: c.op0647,
		0x0648: c.op0648,
		0x0649: c.op0649,
		0x064A: c.op064A,
		0x064B: c.op064B,
		0x064C: c.op064C,
		0x064D: c.op064D,
		0x064E: c.op064E,
		0x064F: c.op064F,
		0x0650: c.op0650,
		0x0651: c.op0651,
		0x0652: c.op0652,
		0x0653: c.op0653,
		0x0654: c.op0654,
		0x0655: c.op0655,
		0x0656: c.op0656,
		0x0657: c.op0657,
		0x0678: c.op0678,
		0x0679: c.op0679,
		0x067C: c.op067C,
		0x0680: c.op0680,
		0x0681: c.op0681,
		0x0682: c.op0682,
		0x0683: c.op0683,
		0x0684: c.op0684,
		0x0685: c.op0685,
		0x0686: c.op0686,
		0x0687: c.op0687,
		0x0688: c.op0688,
		0x0689: c.op0689,
		0x068A: c.op068A,
		0x068B: c.op068B,
		0x068C: c.op068C,
		0x068D: c.op068D,
		0x068E: c.op068E,
		0x068F: c.op068F,
		0x0690: c.op0690,
		0x0691: c.op0691,
		0x0692: c.op0692,
		0x0693: c.op0693,
		0x0694: c.op0694,
		0x0695: c.op0695,
		0x0696: c.op0696,
		0x0697: c.op0697,
		0x06B8: c.op06B8,
		0x06B9: c.op06B9,
		0x06BC: c.op06BC,
		0x06C0: c.op06C0,
		0x06C1: c.op06C1,
		0x06C2: c.op06C2,
		0x06C3: c.op06C3,
		0x06C4: c.op06C4,
		0x06C5: c.op06C5,
		0x06C6: c.op06C6,
		0x06C7: c.op06C7,
		0x06C8: c.op06C8,
		0x06C9: c.op06C9,
		0x06CA: c.op06CA,
		0x06CB: c.op06CB,
		0x06CC: c.op06CC,
		0x06CD: c.op06CD,
		0x06CE: c.op06CE,
		0x06CF: c.op06CF,
		0x06D0: c.op06D0,
		0x06D1: c.op06D1,
		0x06D2: c.op06D2,
		0x06D3: c.op06D3,
		0x06D4: c.op06D4,
		0x06D5: c.op06D5,
		0x06D6: c.op06D6,
		0x06D7: c.op06D7,
		0x06F8: c.op06F8,
		0x06F9: c.op06F9,
		0x06FC: c.op06FC,
		0x0700: c.op0700,
		0x0701: c.op0701,
		0x0702: c.op0702,
		0x0703: c.op0703,
		0x0704: c.op0704,
		0x0705: c.op0705,
		0x0706: c.op0706,
		0x0707: c.op0707,
		0x0708: c.op0708,
		0x0709: c.op0709,
		0x070A: c.op070A,
		0x070B: c.op070B,
		0x070C: c.op070C,
		0x070D: c.op070D,
		0x070E: c.op070E,
		0x070F: c.op070F,
		0x0710: c.op0710,
		0x0711: c.op0711,
		0x0712: c.op0712,
		0x0713: c.op0713,
		0x0714: c.op0714,
		0x0715: c.op0715,
		0x0716: c.op0716,
		0x0717: c.op0717,
		0x0738: c.op0738,
		0x0739: c.op0739,
		0x073C: c.op073C,
		0x0740: c.op0740,
		0x0741: c.op0741,
		0x0742: c.op0742,
		0x0743: c.op0743,
		0x0744: c.op0744,
		0x0745: c.op0745,
		0x0746: c.op0746,
		0x0747: c.op0747,
		0x0748: c.op0748,
		0x0749: c.op0749,
		0x074A: c.op074A,
		0x074B: c.op074B,
		0x074C: c.op074C,
		0x074D: c.op074D,
		0x074E: c.op074E,
		0x074F: c.op074F,
		0x0750: c.op0750,
		0x0751: c.op0751,
		0x0752: c.op0752,
		0x0753: c.op0753,
		0x0754: c.op0754,
		0x0755: c.op0755,
		0x0756: c.op0756,
		0x0757: c.op0757,
		0x0778: c.op0778,
		0x0779: c.op0779,
		0x077C: c.op077C,
		0x0800: c.op0800,
		0x0801: c.op0801,
		0x0802: c.op0802,
		0x0803: c.op0803,
		0x0804: c.op0804,
		0x0805: c.op0805,
		0x0806: c.op0806,
		0x0807: c.op0807,
		0x0808: c.op0808,
		0x0809: c.op0809,
		0x080A: c.op080A,
		0x080B: c.op080B,
		0x080C: c.op080C,
		0x080D: c.op080D,
		0x080E: c.op080E,
		0x080F: c.op080F,
		0x0810: c.op0810,
		0x0811: c.op0811,
		0x0812: c.op0812,
		0x0813: c.op0813,
		0x0814: c.op0814,
		0x0815: c.op0815,
		0x0816: c.op0816,
		0x0817: c.op0817,
		0x0838: c.op0838,
		0x0839: c.op0839,
		0x083C: c.op083C,
		0x0840: c.op0840,
		0x0841: c.op0841,
		0x0842: c.op0842,
		0x0843: c.op0843,
		0x0844: c.op0844,
		0x0845: c.op0845,
		0x0846: c.op0846,
		0x0847: c.op0847,
		0x0848: c.op0848,
		0x0849: c.op0849,
		0x084A: c.op084A,
		0x084B: c.op084B,
		0x084C: c.op084C,
		0x084D: c.op084D,
		0x084E: c.op084E,
		0x084F: c.op084F,
		0x0850: c.op0850,
		0x0851: c.op0851,
		0x0852: c.op0852,
		0x0853: c.op0853,
		0x0854: c.op0854,
		0x0855: c.op0855,
		0x0856: c.op0856,
		0x0857: c.op0857,
		0x0878: c.op0878,
		0x0879: c.op0879,
		0x087C: c.op087C,
		0x0880: c.op0880,
		0x0881: c.op0881,
		0x0882: c.op0882,
		0x0883: c.op0883,
		0x0884: c.op0884,
		0x0885: c.op0885,
		0x0886: c.op0886,
		0x0887: c.op0887,
		0x0888: c.op0888,
		0x0889: c.op0889,
		0x088A: c.op088A,
		0x088B: c.op088B,
		0x088C: c.op088C,
		0x088D: c.op088D,
		0x088E: c.op088E,
		0x088F: c.op088F,
		0x0890: c.op0890,
		0x0891: c.op0891,
		0x0892: c.op0892,
		0x0893: c.op0893,
		0x0894: c.op0894,
		0x0895: c.op0895,
		0x0896: c.op0896,
		0x0897: c.op0897,
		0x08B8: c.op08B8,
		0x08B9: c.op08B9,
		0x08BC: c.op08BC,
		0x08C0: c.op08C0,
		0x08C1: c.op08C1,
		0x08C2: c.op08C2,
		0x08C3: c.op08C3,
		0x08C4: c.op08C4,
		0x08C5: c.op08C5,
		0x08C6: c.op08C6,
		0x08C7: c.op08C7,
		0x08C8: c.op08C8,
		0x08C9: c.op08C9,
		0x08CA: c.op08CA,
		0x08CB: c.op08CB,
		0x08CC: c.op08CC,
		0x08CD: c.op08CD,
		0x08CE: c.op08CE,
		0x08CF: c.op08CF,
		0x08D0: c.op08D0,
		0x08D1: c.op08D1,
		0x08D2: c.op08D2,
		0x08D3: c.op08D3,
		0x08D4: c.op08D4,
		0x08D5: c.op08D5,
		0x08D6: c.op08D6,
		0x08D7: c.op08D7,
		0x08F8: c.op08F8,
		0x08F9: c.op08F9,
		0x08FC: c.op08FC,
		0x0900: c.op0900,
		0x0901: c.op0901,
		0x0902: c.op0902,
		0x0903: c.op0903,
		0x0904: c.op0904,
		0x0905: c.op0905,
		0x0906: c.op0906,
		0x0907: c.op0907,
		0x0908: c.op0908,
		0x0909: c.op0909,
		0x090A: c.op090A,
		0x090B: c.op090B,
		0x090C: c.op090C,
		0x090D: c.op090D,
		0x090E: c.op090E,
		0x090F: c.op090F,
		0x0910: c.op0910,
		0x0911: c.op0911,
		0x0912: c.op0912,
		0x0913: c.op0913,
		0x0914: c.op0914,
		0x0915: c.op0915,
		0x0916: c.op0916,
		0x0917: c.op0917,
		0x0938: c.op0938,
		0x0939: c.op0939,
		0x093C: c.op093C,
		0x0940: c.op0940,
		0x0941: c.op0941,
		0x0942: c.op0942,
		0x0943: c.op0943,
		0x0944: c.op0944,
		0x0945: c.op0945,
		0x0946: c.op0946,
		0x0947: c.op0947,
		0x0948: c.op0948,
		0x0949: c.op0949,
		0x094A: c.op094A,
		0x094B: c.op094B,
		0x094C: c.op094C,
		0x094D: c.op094D,
		0x094E: c.op094E,
		0x094F: c.op094F,
		0x0950: c.op0950,
		0x0951: c.op0951,
		0x0952: c.op0952,
		0x0953: c.op0953,
		0x0954: c.op0954,
		0x0955: c.op0955,
		0x0956: c.op0956,
		0x0957: c.op0957,
		0x0978: c.op0978,
		0x0979: c.op0979,
		0x097C: c.op097C,
		0x0A00: c.op0A00,
		0x0A01: c.op0A01,
		0x0A02: c.op0A02,
		0x0A03: c.op0A03,
		0x0A04: c.op0A04,
		0x0A05: c.op0A05,
		0x0A06: c.op0A06,
		0x0A07: c.op0A07,
		0x0A08: c.op0A08,
		0x0A09: c.op0A09,
		0x0A0A: c.op0A0A,
		0x0A0B: c.op0A0B,
		0x0A0C: c.op0A0C,
		0x0A0D: c.op0A0D,
		0x0A0E: c.op0A0E,
		0x0A0F: c.op0A0F,
		0x0A10: c.op0A10,
		0x0A11: c.op0A11,
		0x0A12: c.op0A12,
		0x0A13: c.op0A13,
		0x0A14: c.op0A14,
		0x0A15: c.op0A15,
		0x0A16: c.op0A16,
		0x0A17: c.op0A17,
		0x0A38: c.op0A38,
		0x0A39: c.op0A39,
		0x0A3C: c.op0A3C,
		0x0A40: c.op0A40,
		0x0A41: c.op0A41,
		0x0A42: c.op0A42,
		0x0A43: c.op0A43,
		0x0A44: c.op0A44,
		0x0A45: c.op0A45,
		0x0A46: c.op0A46,
		0x0A47: c.op0A47,
		0x0A48: c.op0A48,
		0x0A49: c.op0A49,
		0x0A4A: c.op0A4A,
		0x0A4B: c.op0A4B,
		0x0A4C: c.op0A4C,
		0x0A4D: c.op0A4D,
		0x0A4E: c.op0A4E,
		0x0A4F: c.op0A4F,
		0x0A50: c.op0A50,
		0x0A51: c.op0A51,
		0x0A52: c.op0A52,
		0x0A53: c.op0A53,
		0x0A54: c.op0A54,
		0x0A55: c.op0A55,
		0x0A56: c.op0A56,
		0x0A57: c.op0A57,
		0x0A78: c.op0A78,
		0x0A79: c.op0A79,
		0x0A7C: c.op0A7C,
		0x0A80: c.op0A80,
		0x0A81: c.op0A81,
		0x0A82: c.op0A82,
		0x0A83: c.op0A83,
		0x0A84: c.op0A84,
		0x0A85: c.op0A85,
		0x0A86: c.op0A86,
		0x0A87: c.op0A87,
		0x0A88: c.op0A88,
		0x0A89: c.op0A89,
		0x0A8A: c.op0A8A,
		0x0A8B: c.op0A8B,
		0x0A8C: c.op0A8C,
		0x0A8D: c.op0A8D,
		0x0A8E: c.op0A8E,
		0x0A8F: c.op0A8F,
		0x0A90: c.op0A90,
		0x0A91: c.op0A91,
		0x0A92: c.op0A92,
		0x0A93: c.op0A93,
		0x0A94: c.op0A94,
		0x0A95: c.op0A95,
		0x0A96: c.op0A96,
		0x0A97: c.op0A97,
		0x0AB8: c.op0AB8,
		0x0AB9: c.op0AB9,
		0x0ABC: c.op0ABC,
		0x0AC0: c.op0AC0,
		0x0AC1: c.op0AC1,
		0x0AC2: c.op0AC2,
		0x0AC3: c.op0AC3,
		0x0AC4: c.op0AC4,
		0x0AC5: c.op0AC5,
		0x0AC6: c.op0AC6,
		0x0AC7: c.op0AC7,
		0x0AC8: c.op0AC8,
		0x0AC9: c.op0AC9,
		0x0ACA: c.op0ACA,
		0x0ACB: c.op0ACB,
		0x0ACC: c.op0ACC,
		0x0ACD: c.op0ACD,
		0x0ACE: c.op0ACE,
		0x0ACF: c.op0ACF,
		0x0AD0: c.op0AD0,
		0x0AD1: c.op0AD1,
		0x0AD2: c.op0AD2,
		0x0AD3: c.op0AD3,
		0x0AD4: c.op0AD4,
		0x0AD5: c.op0AD5,
		0x0AD6: c.op0AD6,
		0x0AD7: c.op0AD7,
		0x0AF8: c.op0AF8,
		0x0AF9: c.op0AF9,
		0x0AFC: c.op0AFC,
		0x0B00: c.op0B00,
		0x0B01: c.op0B01,
		0x0B02: c.op0B02,
		0x0B03: c.op0B03,
		0x0B04: c.op0B04,
		0x0B05: c.op0B05,
		0x0B06: c.op0B06,
		0x0B07: c.op0B07,
		0x0B08: c.op0B08,
		0x0B09: c.op0B09,
		0x0B0A: c.op0B0A,
		0x0B0B: c.op0B0B,
		0x0B0C: c.op0B0C,
		0x0B0D: c.op0B0D,
		0x0B0E: c.op0B0E,
		0x0B0F: c.op0B0F,
		0x0B10: c.op0B10,
		0x0B11: c.op0B11,
		0x0B12: c.op0B12,
		0x0B13: c.op0B13,
		0x0B14: c.op0B14,
		0x0B15: c.op0B15,
		0x0B16: c.op0B16,
		0x0B17: c.op0B17,
		0x0B38: c.op0B38,
		0x0B39: c.op0B39,
		0x0B3C: c.op0B3C,
		0x0B40: c.op0B40,
		0x0B41: c.op0B41,
		0x0B42: c.op0B42,
		0x0B43: c.op0B43,
		0x0B44: c.op0B44,
		0x0B45: c.op0B45,
		0x0B46: c.op0B46,
		0x0B47: c.op0B47,
		0x0B48: c.op0B48,
		0x0B49: c.op0B49,
		0x0B4A: c.op0B4A,
		0x0B4B: c.op0B4B,
		0x0B4C: c.op0B4C,
		0x0B4D: c.op0B4D,
		0x0B4E: c.op0B4E,
		0x0B4F: c.op0B4F,
		0x0B50: c.op0B50,
		0x0B51: c.op0B51,
		0x0B52: c.op0B52,
		0x0B53: c.op0B53,
		0x0B54: c.op0B54,
		0x0B55: c.op0B55,
		0x0B56: c.op0B56,
		0x0B57: c.op0B57,
		0x0B78: c.op0B78,
		0x0B79: c.op0B79,
		0x0B7C: c.op0B7C,
		0x0C00: c.op0C00,
		0x0C01: c.op0C01,
		0x0C02: c.op0C02,
		0x0C03: c.op0C03,
		0x0C04: c.op0C04,
		0x0C05: c.op0C05,
		0x0C06: c.op0C06,
		0x0C07: c.op0C07,
		0x0C08: c.op0C08,
		0x0C09: c.op0C09,
		0x0C0A: c.op0C0A,
		0x0C0B: c.op0C0B,
		0x0C0C: c.op0C0C,
		0x0C0D: c.op0C0D,
		0x0C0E: c.op0C0E,
		0x0C0F: c.op0C0F,
		0x0C10: c.op0C10,
		0x0C11: c.op0C11,
		0x0C12: c.op0C12,
		0x0C13: c.op0C13,
		0x0C14: c.op0C14,
		0x0C15: c.op0C15,
		0x0C16: c.op0C16,
		0x0C17: c.op0C17,
		0x0C38: c.op0C38,
		0x0C39: c.op0C39,
		0x0C3C: c.op0C3C,
		0x0C40: c.op0C40,
		0x0C41: c.op0C41,
		0x0C42: c.op0C42,
		0x0C43: c.op0C43,
		0x0C44: c.op0C44,
		0x0C45: c.op0C45,
		0x0C46: c.op0C46,
		0x0C47: c.op0C47,
		0x0C48: c.op0C48,
		0x0C49: c.op0C49,
		0x0C4A: c.op0C4A,
		0x0C4B: c.op0C4B,
		0x0C4C: c.op0C4C,
		0x0C4D: c.op0C4D,
		0x0C4E: c.op0C4E,
		0x0C4F: c.op0C4F,
		0x0C50: c.op0C50,
		0x0C51: c.op0C51,
		0x0C52: c.op0C52,
		0x0C53: c.op0C53,
		0x0C54: c.op0C54,
		0x0C55: c.op0C55,
		0x0C56: c.op0C56,
		0x0C57: c.op0C57,
		0x0C78: c.op0C78,
		0x0C79: c.op0C79,
		0x0C7C: c.op0C7C,
		0x0C80: c.op0C80,
		0x0C81: c.op0C81,
		0x0C82: c.op0C82,
		0x0C83: c.op0C83,
		0x0C84: c.op0C84,
		0x0C85: c.op0C85,
		0x0C86: c.op0C86,
		0x0C87: c.op0C87,
		0x0C88: c.op0C88,
		0x0C89: c.op0C89,
		0x0C8A: c.op0C8A,
		0x0C8B: c.op0C8B,
		0x0C8C: c.op0C8C,
		0x0C8D: c.op0C8D,
		0x0C8E: c.op0C8E,
		0x0C8F: c.op0C8F,
		0x0C90: c.op0C90,
		0x0C91: c.op0C91,
		0x0C92: c.op0C92,
		0x0C93: c.op0C93,
		0x0C94: c.op0C94,
		0x0C95: c.op0C95,
		0x0C96: c.op0C96,
		0x0C97: c.op0C97,
		0x0CB8: c.op0CB8,
		0x0CB9: c.op0CB9,
		0x0CBC: c.op0CBC,
		0x0CC0: c.op0CC0,
		0x0CC1: c.op0CC1,
		0x0CC2: c.op0CC2,
		0x0CC3: c.op0CC3,
		0x0CC4: c.op0CC4,
		0x0CC5: c.op0CC5,
		0x0CC6: c.op0CC6,
		0x0CC7: c.op0CC7,
		0x0CC8: c.op0CC8,
		0x0CC9: c.op0CC9,
		0x0CCA: c.op0CCA,
		0x0CCB: c.op0CCB,
		0x0CCC: c.op0CCC,
		0x0CCD: c.op0CCD,
		0x0CCE: c.op0CCE,
		0x0CCF: c.op0CCF,
		0x0CD0: c.op0CD0,
		0x0CD1: c.op0CD1,
		0x0CD2: c.op0CD2,
		0x0CD3: c.op0CD3,
		0x0CD4: c.op0CD4,
		0x0CD5: c.op0CD5,
		0x0CD6: c.op0CD6,
		0x0CD7: c.op0CD7,
		0x0CF8: c.op0CF8,
		0x0CF9: c.op0CF9,
		0x0CFC: c.op0CFC,
		0x0D00: c.op0D00,
		0x0D01: c.op0D01,
		0x0D02: c.op0D02,
		0x0D03: c.op0D03,
		0x0D04: c.op0D04,
		0x0D05: c.op0D05,
		0x0D06: c.op0D06,
		0x0D07: c.op0D07,
		0x0D08: c.op0D08,
		0x0D09: c.op0D09,
		0x0D0A: c.op0D0A,
		0x0D0B: c.op0D0B,
		0x0D0C: c.op0D0C,
		0x0D0D: c.op0D0D,
		0x0D0E: c.op0D0E,
		0x0D0F: c.op0D0F,
		0x0D10: c.op0D10,
		0x0D11: c.op0D11,
		0x0D12: c.op0D12,
		0x0D13: c.op0D13,
		0x0D14: c.op0D14,
		0x0D15: c.op0D15,
		0x0D16: c.op0D16,
		0x0D17: c.op0D17,
		0x0D38: c.op0D38,
		0x0D39: c.op0D39,
		0x0D3C: c.op0D3C,
		0x0D40: c.op0D40,
		0x0D41: c.op0D41,
		0x0D42: c.op0D42,
		0x0D43: c.op0D43,
		0x0D44: c.op0D44,
		0x0D45: c.op0D45,
		0x0D46: c.op0D46,
		0x0D47: c.op0D47,
		0x0D48: c.op0D48,
		0x0D49: c.op0D49,
		0x0D4A: c.op0D4A,
		0x0D4B: c.op0D4B,
		0x0D4C: c.op0D4C,
		0x0D4D: c.op0D4D,
		0x0D4E: c.op0D4E,
		0x0D4F: c.op0D4F,
		0x0D50: c.op0D50,
		0x0D51: c.op0D51,
		0x0D52: c.op0D52,
		0x0D53: c.op0D53,
		0x0D54: c.op0D54,
		0x0D55: c.op0D55,
		0x0D56: c.op0D56,
		0x0D57: c.op0D57,
		0x0D78: c.op0D78,
		0x0D79: c.op0D79,
		0x0D7C: c.op0D7C,
		0x0E00: c.op0E00,
		0x0E01: c.op0E01,
		0x0E02: c.op0E02,
		0x0E03: c.op0E03,
		0x0E04: c.op0E04,
		0x0E05: c.op0E05,
		0x0E06: c.op0E06,
		0x0E07: c.op0E07,
		0x0E08: c.op0E08,
		0x0E09: c.op0E09,
		0x0E0A: c.op0E0A,
		0x0E0B: c.op0E0B,
		0x0E0C: c.op0E0C,
		0x0E0D: c.op0E0D,
		0x0E0E: c.op0E0E,
		0x0E0F: c.op0E0F,
		0x0E10: c.op0E10,
		0x0E11: c.op0E11,
		0x0E12: c.op0E12,
		0x0E13: c.op0E13,
		0x0E14: c.op0E14,
		0x0E15: c.op0E15,
		0x0E16: c.op0E16,
		0x0E17: c.op0E17,
		0x0E38: c.op0E38,
		0x0E39: c.op0E39,
		0x0E3C: c.op0E3C,
		0x0E40: c.op0E40,
		0x0E41: c.op0E41,
		0x0E42: c.op0E42,
		0x0E43: c.op0E43,
		0x0E44: c.op0E44,
		0x0E45: c.op0E45,
		0x0E46: c.op0E46,
		0x0E47: c.op0E47,
		0x0E48: c.op0E48,
		0x0E49: c.op0E49,
		0x0E4A: c.op0E4A,
		0x0E4B: c.op0E4B,
		0x0E4C: c.op0E4C,
		0x0E4D: c.op0E4D,
		0x0E4E: c.op0E4E,
		0x0E4F: c.op0E4F,
		0x0E50: c.op0E50,
		0x0E51: c.op0E51,
		0x0E52: c.op0E52,
		0x0E53: c.op0E53,
		0x0E54: c.op0E54,
		0x0E55: c.op0E55,
		0x0E56: c.op0E56,
		0x0E57: c.op0E57,
		0x0E78: c.op0E78,
		0x0E79: c.op0E79,
		0x0E7C: c.op0E7C,
		0x0E80: c.op0E80,
		0x0E81: c.op0E81,
		0x0E82: c.op0E82,
		0x0E83: c.op0E83,
		0x0E84: c.op0E84,
		0x0E85: c.op0E85,
		0x0E86: c.op0E86,
		0x0E87: c.op0E87,
		0x0E88: c.op0E88,
		0x0E89: c.op0E89,
		0x0E8A: c.op0E8A,
		0x0E8B: c.op0E8B,
		0x0E8C: c.op0E8C,
		0x0E8D: c.op0E8D,
		0x0E8E: c.op0E8E,
		0x0E8F: c.op0E8F,
		0x0E90: c.op0E90,
		0x0E91: c.op0E91,
		0x0E92: c.op0E92,
		0x0E93: c.op0E93,
		0x0E94: c.op0E94,
		0x0E95: c.op0E95,
		0x0E96: c.op0E96,
		0x0E97: c.op0E97,
		0x0EB8: c.op0EB8,
		0x0EB9: c.op0EB9,
		0x0EBC: c.op0EBC,
		0x0EC0: c.op0EC0,
		0x0EC1: c.op0EC1,
		0x0EC2: c.op0EC2,
		0x0EC3: c.op0EC3,
		0x0EC4: c.op0EC4,
		0x0EC5: c.op0EC5,
		0x0EC6: c.op0EC6,
		0x0EC7: c.op0EC7,
		0x0EC8: c.op0EC8,
		0x0EC9: c.op0EC9,
		0x0ECA: c.op0ECA,
		0x0ECB: c.op0ECB,
		0x0ECC: c.op0ECC,
		0x0ECD: c.op0ECD,
		0x0ECE: c.op0ECE,
		0x0ECF: c.op0ECF,
		0x0ED0: c.op0ED0,
		0x0ED1: c.op0ED1,
		0x0ED2: c.op0ED2,
		0x0ED3: c.op0ED3,
		0x0ED4: c.op0ED4,
		0x0ED5: c.op0ED5,
		0x0ED6: c.op0ED6,
		0x0ED7: c.op0ED7,
		0x0EF8: c.op0EF8,
		0x0EF9: c.op0EF9,
		0x0EFC: c.op0EFC,
		0x0F00: c.op0F00,
		0x0F01: c.op0F01,
		0x0F02: c.op0F02,
		0x0F03: c.op0F03,
		0x0F04: c.op0F04,
		0x0F05: c.op0F05,
		0x0F06: c.op0F06,
		0x0F07: c.op0F07,
		0x0F08: c.op0F08,
		0x0F09: c.op0F09,
		0x0F0A: c.op0F0A,
		0x0F0B: c.op0F0B,
		0x0F0C: c.op0F0C,
		0x0F0D: c.op0F0D,
		0x0F0E: c.op0F0E,
		0x0F0F: c.op0F0F,
		0x0F10: c.op0F10,
		0x0F11: c.op0F11,
		0x0F12: c.op0F12,
		0x0F13: c.op0F13,
		0x0F14: c.op0F14,
		0x0F15: c.op0F15,
		0x0F16: c.op0F16,
		0x0F17: c.op0F17,
		0x0F38: c.op0F38,
		0x0F39: c.op0F39,
		0x0F3C: c.op0F3C,
		0x0F40: c.op0F40,
		0x0F41: c.op0F41,
		0x0F42: c.op0F42,
		0x0F43: c.op0F43,
		0x0F44: c.op0F44,
		0x0F45: c.op0F45,
		0x0F46: c.op0F46,
		0x0F47: c.op0F47,
		0x0F48: c.op0F48,
		0x0F49: c.op0F49,
		0x0F4A: c.op0F4A,
		0x0F4B: c.op0F4B,
		0x0F4C: c.op0F4C,
		0x0F4D: c.op0F4D,
		0x0F4E: c.op0F4E,
		0x0F4F: c.op0F4F,
		0x0F50: c.op0F50,
		0x0F51: c.op0F51,
		0x0F52: c.op0F52,
		0x0F53: c.op0F53,
		0x0F54: c.op0F54,
		0x0F55: c.op0F55,
		0x0F56: c.op0F56,
		0x0F57: c.op0F57,
		0x0F78: c.op0F78,
		0x0F79: c.op0F79,
		0x0F7C: c.op0F7C,
		0x1000: c.op1000,
		0x1001: c.op1001,
		0x1002: c.op1002,
		0x1003: c.op1003,
		0x1004: c.op1004,
		0x1005: c.op1005,
		0x1006: c.op1006,
		0x1007: c.op1007,
		0x1008: c.op1008,
		0x1009: c.op1009,
		0x100A: c.op100A,
		0x100B: c.op100B,
		0x100C: c.op100C,
		0x100D: c.op100D,
		0x100E: c.op100E,
		0x100F: c.op100F,
		0x1010: c.op1010,
		0x1011: c.op1011,
		0x1012: c.op1012,
		0x1013: c.op1013,
		0x1014: c.op1014,
		0x1015: c.op1015,
		0x1016: c.op1016,
		0x1017: c.op1017,
		0x1038: c.op1038,
		0x1039: c.op1039,
		0x103C: c.op103C,
		0x1040: c.op1040,
		0x1041: c.op1041,
		0x1042: c.op1042,
		0x1043: c.op1043,
		0x1044: c.op1044,
		0x1045: c.op1045,
		0x1046: c.op1046,
		0x1047: c.op1047,
		0x1048: c.op1048,
		0x1049: c.op1049,
		0x104A: c.op104A,
		0x104B: c.op104B,
		0x104C: c.op104C,
		0x104D: c.op104D,
		0x104E: c.op104E,
		0x104F: c.op104F,
		0x1050: c.op1050,
		0x1051: c.op1051,
		0x1052: c.op1052,
		0x1053: c.op1053,
		0x1054: c.op1054,
		0x1055: c.op1055,
		0x1056: c.op1056,
		0x1057: c.op1057,
		0x1078: c.op1078,
		0x1079: c.op1079,
		0x107C: c.op107C,
		0x1080: c.op1080,
		0x1081: c.op1081,
		0x1082: c.op1082,
		0x1083: c.op1083,
		0x1084: c.op1084,
		0x1085: c.op1085,
		0x1086: c.op1086,
		0x1087: c.op1087,
		0x1088: c.op1088,
		0x1089: c.op1089,
		0x108A: c.op108A,
		0x108B: c.op108B,
		0x108C: c.op108C,
		0x108D: c.op108D,
		0x108E: c.op108E,
		0x108F: c.op108F,
		0x1090: c.op1090,
		0x1091: c.op1091,
		0x1092: c.op1092,
		0x1093: c.op1093,
		0x1094: c.op1094,
		0x1095: c.op1095,
		0x1096: c.op1096,
		0x1097: c.op1097,
		0x10B8: c.op10B8,
		0x10B9: c.op10B9,
		0x10BC: c.op10BC,
		0x10C0: c.op10C0,
		0x10C1: c.op10C1,
		0x10C2: c.op10C2,
		0x10C3: c.op10C3,
		0x10C4: c.op10C4,
		0x10C5: c.op10C5,
		0x10C6: c.op10C6,
		0x10C7: c.op10C7,
		0x10C8: c.op10C8,
		0x10C9: c.op10C9,
		0x10CA: c.op10CA,
		0x10CB: c.op10CB,
		0x10CC: c.op10CC,
		0x10CD: c.op10CD,
		0x10CE: c.op10CE,
		0x10CF: c.op10CF,
		0x10D0: c.op10D0,
		0x10D1: c.op10D1,
		0x10D2: c.op10D2,
		0x10D3: c.op10D3,
		0x10D4: c.op10D4,
		0x10D5: c.op10D5,
		0x10D6: c.op10D6,
		0x10D7: c.op10D7,
		0x10F8: c.op10F8,
		0x10F9: c.op10F9,
		0x10FC: c.op10FC,
		0x1100: c.op1100,
		0x1101: c.op1101,
		0x1102: c.op1102,
		0x1103: c.op1103,
		0x1104: c.op1104,
		0x1105: c.op1105,
		0x1106: c.op1106,
		0x1107: c.op1107,
		0x1108: c.op1108,
		0x1109: c.op1109,
		0x110A: c.op110A,
		0x110B: c.op110B,
		0x110C: c.op110C,
		0x110D: c.op110D,
		0x110E: c.op110E,
		0x110F: c.op110F,
		0x1110: c.op1110,
		0x1111: c.op1111,
		0x1112: c.op1112,
		0x1113: c.op1113,
		0x1114: c.op1114,
		0x1115: c.op1115,
		0x1116: c.op1116,
		0x1117: c.op1117,
		0x1138: c.op1138,
		0x1139: c.op1139,
		0x113C: c.op113C,
		0x1140: c.op1140,
		0x1141: c.op1141,
		0x1142: c.op1142,
		0x1143: c.op1143,
		0x1144: c.op1144,
		0x1145: c.op1145,
		0x1146: c.op1146,
		0x1147: c.op1147,
		0x1148: c.op1148,
		0x1149: c.op1149,
		0x114A: c.op114A,
		0x114B: c.op114B,
		0x114C: c.op114C,
		0x114D: c.op114D,
		0x114E: c.op114E,
		0x114F: c.op114F,
		0x1150: c.op1150,
		0x1151: c.op1151,
		0x1152: c.op1152,
		0x1153: c.op1153,
		0x1154: c.op1154,
		0x1155: c.op1155,
		0x1156: c.op1156,
		0x1157: c.op1157,
		0x1178: c.op1178,
		0x1179: c.op1179,
		0x117C: c.op117C,
		0x11C0: c.op11C0,
		0x11C1: c.op11C1,
		0x11C2: c.op11C2,
		0x11C3: c.op11C3,
		0x11C4: c.op11C4,
		0x11C5: c.op11C5,
		0x11C6: c.op11C6,
		0x11C7: c.op11C7,
		0x11C8: c.op11C8,
		0x11C9: c.op11C9,
		0x11CA: c.op11CA,
		0x11CB: c.op11CB,
		0x11CC: c.op11CC,
		0x11CD: c.op11CD,
		0x11CE: c.op11CE,
		0x11CF: c.op11CF,
		0x11D0: c.op11D0,
		0x11D1: c.op11D1,
		0x11D2: c.op11D2,
		0x11D3: c.op11D3,
		0x11D4: c.op11D4,
		0x11D5: c.op11D5,
		0x11D6: c.op11D6,
		0x11D7: c.op11D7,
		0x11F8: c.op11F8,
		0x11F9: c.op11F9,
		0x11FC: c.op11FC,
		0x1200: c.op1200,
		0x1201: c.op1201,
		0x1202: c.op1202,
		0x1203: c.op1203,
		0x1204: c.op1204,
		0x1205: c.op1205,
		0x1206: c.op1206,
		0x1207: c.op1207,
		0x1208: c.op1208,
		0x1209: c.op1209,
		0x120A: c.op120A,
		0x120B: c.op120B,
		0x120C: c.op120C,
		0x120D: c.op120D,
		0x120E: c.op120E,
		0x120F: c.op120F,
		0x1210: c.op1210,
		0x1211: c.op1211,
		0x1212: c.op1212,
		0x1213: c.op1213,
		0x1214: c.op1214,
		0x1215: c.op1215,
		0x1216: c.op1216,
		0x1217: c.op1217,
		0x1238: c.op1238,
		0x1239: c.op1239,
		0x123C: c.op123C,
		0x1240: c.op1240,
		0x1241: c.op1241,
		0x1242: c.op1242,
		0x1243: c.op1243,
		0x1244: c.op1244,
		0x1245: c.op1245,
		0x1246: c.op1246,
		0x1247: c.op1247,
		0x1248: c.op1248,
		0x1249: c.op1249,
		0x124A: c.op124A,
		0x124B: c.op124B,
		0x124C: c.op124C,
		0x124D: c.op124D,
		0x124E: c.op124E,
		0x124F: c.op124F,
		0x1250: c.op1250,
		0x1251: c.op1251,
		0x1252: c.op1252,
		0x1253: c.op1253,
		0x1254: c.op1254,
		0x1255: c.op1255,
		0x1256: c.op1256,
		0x1257: c.op1257,
		0x1278: c.op1278,
		0x1279: c.op1279,
		0x127C: c.op127C,
		0x1280: c.op1280,
		0x1281: c.op1281,
		0x1282: c.op1282,
		0x1283: c.op1283,
		0x1284: c.op1284,
		0x1285: c.op1285,
		0x1286: c.op1286,
		0x1287: c.op1287,
		0x1288: c.op1288,
		0x1289: c.op1289,
		0x128A: c.op128A,
		0x128B: c.op128B,
		0x128C: c.op128C,
		0x128D: c.op128D,
		0x128E: c.op128E,
		0x128F: c.op128F,
		0x1290: c.op1290,
		0x1291: c.op1291,
		0x1292: c.op1292,
		0x1293: c.op1293,
		0x1294: c.op1294,
		0x1295: c.op1295,
		0x1296: c.op1296,
		0x1297: c.op1297,
		0x12B8: c.op12B8,
		0x12B9: c.op12B9,
		0x12BC: c.op12BC,
		0x12C0: c.op12C0,
		0x12C1: c.op12C1,
		0x12C2: c.op12C2,
		0x12C3: c.op12C3,
		0x12C4: c.op12C4,
		0x12C5: c.op12C5,
		0x12C6: c.op12C6,
		0x12C7: c.op12C7,
		0x12C8: c.op12C8,
		0x12C9: c.op12C9,
		0x12CA: c.op12CA,
		0x12CB: c.op12CB,
		0x12CC: c.op12CC,
		0x12CD: c.op12CD,
		0x12CE: c.op12CE,
		0x12CF: c.op12CF,
		0x12D0: c.op12D0,
		0x12D1: c.op12D1,
		0x12D2: c.op12D2,
		0x12D3: c.op12D3,
		0x12D4: c.op12D4,
		0x12D5: c.op12D5,
		0x12D6: c.op12D6,
		0x12D7: c.op12D7,
		0x12F8: c.op12F8,
		0x12F9: c.op12F9,
		0x12FC: c.op12FC,
		0x1300: c.op1300,
		0x1301: c.op1301,
		0x1302: c.op1302,
		0x1303: c.op1303,
		0x1304: c.op1304,
		0x1305: c.op1305,
		0x1306: c.op1306,
		0x1307: c.op1307,
		0x1308: c.op1308,
		0x1309: c.op1309,
		0x130A: c.op130A,
		0x130B: c.op130B,
		0x130C: c.op130C,
		0x130D: c.op130D,
		0x130E: c.op130E,
		0x130F: c.op130F,
		0x1310: c.op1310,
		0x1311: c.op1311,
		0x1312: c.op1312,
		0x1313: c.op1313,
		0x1314: c.op1314,
		0x1315: c.op1315,
		0x1316: c.op1316,
		0x1317: c.op1317,
		0x1338: c.op1338,
		0x1339: c.op1339,
		0x133C: c.op133C,
		0x1340: c.op1340,
		0x1341: c.op1341,
		0x1342: c.op1342,
		0x1343: c.op1343,
		0x1344: c.op1344,
		0x1345: c.op1345,
		0x1346: c.op1346,
		0x1347: c.op1347,
		0x1348: c.op1348,
		0x1349: c.op1349,
		0x134A: c.op134A,
		0x134B: c.op134B,
		0x134C: c.op134C,
		0x134D: c.op134D,
		0x134E: c.op134E,
		0x134F: c.op134F,
		0x1350: c.op1350,
		0x1351: c.op1351,
		0x1352: c.op1352,
		0x1353: c.op1353,
		0x1354: c.op1354,
		0x1355: c.op1355,
		0x1356: c.op1356,
		0x1357: c.op1357,
		0x1378: c.op1378,
		0x1379: c.op1379,
		0x137C: c.op137C,
		0x13C0: c.op13C0,
		0x13C1: c.op13C1,
		0x13C2: c.op13C2,
		0x13C3: c.op13C3,
		0x13C4: c.op13C4,
		0x13C5: c.op13C5,
		0x13C6: c.op13C6,
		0x13C7: c.op13C7,
		0x13C8: c.op13C8,
		0x13C9: c.op13C9,
		0x13CA: c.op13CA,
		0x13CB: c.op13CB,
		0x13CC: c.op13CC,
		0x13CD: c.op13CD,
		0x13CE: c.op13CE,
		0x13CF: c.op13CF,
		0x13D0: c.op13D0,
		0x13D1: c.op13D1,
		0x13D2: c.op13D2,
		0x13D3: c.op13D3,
		0x13D4: c.op13D4,
		0x13D5: c.op13D5,
		0x13D6: c.op13D6,
		0x13D7: c.op13D7,
		0x13F8: c.op13F8,
		0x13F9: c.op13F9,
		0x13FC: c.op13FC,
		0x1400: c.op1400,
		0x1401: c.op1401,
		0x1402: c.op1402,
		0x1403: c.op1403,
		0x1404: c.op1404,
		0x1405: c.op1405,
		0x1406: c.op1406,
		0x1407: c.op1407,
		0x1408: c.op1408,
		0x1409: c.op1409,
		0x140A: c.op140A,
		0x140B: c.op140B,
		0x140C: c.op140C,
		0x140D: c.op140D,
		0x140E: c.op140E,
		0x140F: c.op140F,
		0x1410: c.op1410,
		0x1411: c.op1411,
		0x1412: c.op1412,
		0x1413: c.op1413,
		0x1414: c.op1414,
		0x1415: c.op1415,
		0x1416: c.op1416,
		0x1417: c.op1417,
		0x1438: c.op1438,
		0x1439: c.op1439,
		0x143C: c.op143C,
		0x1440: c.op1440,
		0x1441: c.op1441,
		0x1442: c.op1442,
		0x1443: c.op1443,
		0x1444: c.op1444,
		0x1445: c.op1445,
		0x1446: c.op1446,
		0x1447: c.op1447,
		0x1448: c.op1448,
		0x1449: c.op1449,
		0x144A: c.op144A,
		0x144B: c.op144B,
		0x144C: c.op144C,
		0x144D: c.op144D,
		0x144E: c.op144E,
		0x144F: c.op144F,
		0x1450: c.op1450,
		0x1451: c.op1451,
		0x1452: c.op1452,
		0x1453: c.op1453,
		0x1454: c.op1454,
		0x1455: c.op1455,
		0x1456: c.op1456,
		0x1457: c.op1457,
		0x1478: c.op1478,
		0x1479: c.op1479,
		0x147C: c.op147C,
		0x1480: c.op1480,
		0x1481: c.op1481,
		0x1482: c.op1482,
		0x1483: c.op1483,
		0x1484: c.op1484,
		0x1485: c.op1485,
		0x1486: c.op1486,
		0x1487: c.op1487,
		0x1488: c.op1488,
		0x1489: c.op1489,
		0x148A: c.op148A,
		0x148B: c.op148B,
		0x148C: c.op148C,
		0x148D: c.op148D,
		0x148E: c.op148E,
		0x148F: c.op148F,
		0x1490: c.op1490,
		0x1491: c.op1491,
		0x1492: c.op1492,
		0x1493: c.op1493,
		0x1494: c.op1494,
		0x1495: c.op1495,
		0x1496: c.op1496,
		0x1497: c.op1497,
		0x14B8: c.op14B8,
		0x14B9: c.op14B9,
		0x14BC: c.op14BC,
		0x14C0: c.op14C0,
		0x14C1: c.op14C1,
		0x14C2: c.op14C2,
		0x14C3: c.op14C3,
		0x14C4: c.op14C4,
		0x14C5: c.op14C5,
		0x14C6: c.op14C6,
		0x14C7: c.op14C7,
		0x14C8: c.op14C8,
		0x14C9: c.op14C9,
		0x14CA: c.op14CA,
		0x14CB: c.op14CB,
		0x14CC: c.op14CC,
		0x14CD: c.op14CD,
		0x14CE: c.op14CE,
		0x14CF: c.op14CF,
		0x14D0: c.op14D0,
		0x14D1: c.op14D1,
		0x14D2: c.op14D2,
		0x14D3: c.op14D3,
		0x14D4: c.op14D4,
		0x14D5: c.op14D5,
		0x14D6: c.op14D6,
		0x14D7: c.op14D7,
		0x14F8: c.op14F8,
		0x14F9: c.op14F9,
		0x14FC: c.op14FC,
		0x1500: c.op1500,
		0x1501: c.op1501,
		0x1502: c.op1502,
		0x1503: c.op1503,
		0x1504: c.op1504,
		0x1505: c.op1505,
		0x1506: c.op1506,
		0x1507: c.op1507,
		0x1508: c.op1508,
		0x1509: c.op1509,
		0x150A: c.op150A,
		0x150B: c.op150B,
		0x150C: c.op150C,
		0x150D: c.op150D,
		0x150E: c.op150E,
		0x150F: c.op150F,
		0x1510: c.op1510,
		0x1511: c.op1511,
		0x1512: c.op1512,
		0x1513: c.op1513,
		0x1514: c.op1514,
		0x1515: c.op1515,
		0x1516: c.op1516,
		0x1517: c.op1517,
		0x1538: c.op1538,
		0x1539: c.op1539,
		0x153C: c.op153C,
		0x1540: c.op1540,
		0x1541: c.op1541,
		0x1542: c.op1542,
		0x1543: c.op1543,
		0x1544: c.op1544,
		0x1545: c.op1545,
		0x1546: c.op1546,
		0x1547: c.op1547,
		0x1548: c.op1548,
		0x1549: c.op1549,
		0x154A: c.op154A,
		0x154B: c.op154B,
		0x154C: c.op154C,
		0x154D: c.op154D,
		0x154E: c.op154E,
		0x154F: c.op154F,
		0x1550: c.op1550,
		0x1551: c.op1551,
		0x1552: c.op1552,
		0x1553: c.op1553,
		0x1554: c.op1554,
		0x1555: c.op1555,
		0x1556: c.op1556,
		0x1557: c.op1557,
		0x1578: c.op1578,
		0x1579: c.op1579,
		0x157C: c.op157C,
		0x1600: c.op1600,
		0x1601: c.op1601,
		0x1602: c.op1602,
		0x1603: c.op1603,
		0x1604: c.op1604,
		0x1605: c.op1605,
		0x1606: c.op1606,
		0x1607: c.op1607,
		0x1608: c.op1608,
		0x1609: c.op1609,
		0x160A: c.op160A,
		0x160B: c.op160B,
		0x160C: c.op160C,
		0x160D: c.op160D,
		0x160E: c.op160E,
		0x160F: c.op160F,
		0x1610: c.op1610,
		0x1611: c.op1611,
		0x1612: c.op1612,
		0x1613: c.op1613,
		0x1614: c.op1614,
		0x1615: c.op1615,
		0x1616: c.op1616,
		0x1617: c.op1617,
		0x1638: c.op1638,
		0x1639: c.op1639,
		0x163C: c.op163C,
		0x1640: c.op1640,
		0x1641: c.op1641,
		0x1642: c.op1642,
		0x1643: c.op1643,
		0x1644: c.op1644,
		0x1645: c.op1645,
		0x1646: c.op1646,
		0x1647: c.op1647,
		0x1648: c.op1648,
		0x1649: c.op1649,
		0x164A: c.op164A,
		0x164B: c.op164B,
		0x164C: c.op164C,
		0x164D: c.op164D,
		0x164E: c.op164E,
		0x164F: c.op164F,
		0x1650: c.op1650,
		0x1651: c.op1651,
		0x1652: c.op1652,
		0x1653: c.op1653,
		0x1654: c.op1654,
		0x1655: c.op1655,
		0x1656: c.op1656,
		0x1657: c.op1657,
		0x1678: c.op1678,
		0x1679: c.op1679,
		0x167C: c.op167C,
		0x1680: c.op1680,
		0x1681: c.op1681,
		0x1682: c.op1682,
		0x1683: c.op1683,
		0x1684: c.op1684,
		0x1685: c.op1685,
		0x1686: c.op1686,
		0x1687: c.op1687,
		0x1688: c.op1688,
		0x1689: c.op1689,
		0x168A: c.op168A,
		0x168B: c.op168B,
		0x168C: c.op168C,
		0x168D: c.op168D,
		0x168E: c.op168E,
		0x168F: c.op168F,
		0x1690: c.op1690,
		0x1691: c.op1691,
		0x1692: c.op1692,
		0x1693: c.op1693,
		0x1694: c.op1694,
		0x1695: c.op1695,
		0x1696: c.op1696,
		0x1697: c.op1697,
		0x16B8: c.op16B8,
		0x16B9: c.op16B9,
		0x16BC: c.op16BC,
		0x16C0: c.op16C0,
		0x16C1: c.op16C1,
		0x16C2: c.op16C2,
		0x16C3: c.op16C3,
		0x16C4: c.op16C4,
		0x16C5: c.op16C5,
		0x16C6: c.op16C6,
		0x16C7: c.op16C7,
		0x16C8: c.op16C8,
		0x16C9: c.op16C9,
		0x16CA: c.op16CA,
		0x16CB: c.op16CB,
		0x16CC: c.op16CC,
		0x16CD: c.op16CD,
		0x16CE: c.op16CE,
		0x16CF: c.op16CF,
		0x16D0: c.op16D0,
		0x16D1: c.op16D1,
		0x16D2: c.op16D2,
		0x16D3: c.op16D3,
		0x16D4: c.op16D4,
		0x16D5: c.op16D5,
		0x16D6: c.op16D6,
		0x16D7: c.op16D7,
		0x16F8: c.op16F8,
		0x16F9: c.op16F9,
		0x16FC: c.op16FC,
		0x1700: c.op1700,
		0x1701: c.op1701,
		0x1702: c.op1702,
		0x1703: c.op1703,
		0x1704: c.op1704,
		0x1705: c.op1705,
		0x1706: c.op1706,
		0x1707: c.op1707,
		0x1708: c.op1708,
		0x1709: c.op1709,
		0x170A: c.op170A,
		0x170B: c.op170B,
		0x170C: c.op170C,
		0x170D: c.op170D,
		0x170E: c.op170E,
		0x170F: c.op170F,
		0x1710: c.op1710,
		0x1711: c.op1711,
		0x1712: c.op1712,
		0x1713: c.op1713,
		0x1714: c.op1714,
		0x1715: c.op1715,
		0x1716: c.op1716,
		0x1717: c.op1717,
		0x1738: c.op1738,
		0x1739: c.op1739,
		0x173C: c.op173C,
		0x1740: c.op1740,
		0x1741: c.op1741,
		0x1742: c.op1742,
		0x1743: c.op1743,
		0x1744: c.op1744,
		0x1745: c.op1745,
		0x1746: c.op1746,
		0x1747: c.op1747,
		0x1748: c.op1748,
		0x1749: c.op1749,
		0x174A: c.op174A,
		0x174B: c.op174B,
		0x174C: c.op174C,
		0x174D: c.op174D,
		0x174E: c.op174E,
		0x174F: c.op174F,
		0x1750: c.op1750,
		0x1751: c.op1751,
		0x1752: c.op1752,
		0x1753: c.op1753,
		0x1754: c.op1754,
		0x1755: c.op1755,
		0x1756: c.op1756,
		0x1757: c.op1757,
		0x1778: c.op1778,
		0x1779: c.op1779,
		0x177C: c.op177C,
		0x1800: c.op1800,
		0x1801: c.op1801,
		0x1802: c.op1802,
		0x1803: c.op1803,
		0x1804: c.op1804,
		0x1805: c.op1805,
		0x1806: c.op1806,
		0x1807: c.op1807,
		0x1808: c.op1808,
		0x1809: c.op1809,
		0x180A: c.op180A,
		0x180B: c.op180B,
		0x180C: c.op180C,
		0x180D: c.op180D,
		0x180E: c.op180E,
		0x180F: c.op180F,
		0x1810: c.op1810,
		0x1811: c.op1811,
		0x1812: c.op1812,
		0x1813: c.op1813,
		0x1814: c.op1814,
		0x1815: c.op1815,
		0x1816: c.op1816,
		0x1817: c.op1817,
		0x1838: c.op1838,
		0x1839: c.op1839,
		0x183C: c.op183C,
		0x1840: c.op1840,
		0x1841: c.op1841,
		0x1842: c.op1842,
		0x1843: c.op1843,
		0x1844: c.op1844,
		0x1845: c.op1845,
		0x1846: c.op1846,
		0x1847: c.op1847,
		0x1848: c.op1848,
		0x1849: c.op1849,
		0x184A: c.op184A,
		0x184B: c.op184B,
		0x184C: c.op184C,
		0x184D: c.op184D,
		0x184E: c.op184E,
		0x184F: c.op184F,
		0x1850: c.op1850,
		0x1851: c.op1851,
		0x1852: c.op1852,
		0x1853: c.op1853,
		0x1854: c.op1854,
		0x1855: c.op1855,
		0x1856: c.op1856,
		0x1857: c.op1857,
		0x1878: c.op1878,
		0x1879: c.op1879,
		0x187C: c.op187C,
		0x1880: c.op1880,
		0x1881: c.op1881,
		0x1882: c.op1882,
		0x1883: c.op1883,
		0x1884: c.op1884,
		0x1885: c.op1885,
		0x1886: c.op1886,
		0x1887: c.op1887,
		0x1888: c.op1888,
		0x1889: c.op1889,
		0x188A: c.op188A,
		0x188B: c.op188B,
		0x188C: c.op188C,
		0x188D: c.op188D,
		0x188E: c.op188E,
		0x188F: c.op188F,
		0x1890: c.op1890,
		0x1891: c.op1891,
		0x1892: c.op1892,
		0x1893: c.op1893,
		0x1894: c.op1894,
		0x1895: c.op1895,
		0x1896: c.op1896,
		0x1897: c.op1897,
		0x18B8: c.op18B8,
		0x18B9: c.op18B9,
		0x18BC: c.op18BC,
		0x18C0: c.op18C0,
		0x18C1: c.op18C1,
		0x18C2: c.op18C2,
		0x18C3: c.op18C3,
		0x18C4: c.op18C4,
		0x18C5: c.op18C5,
		0x18C6: c.op18C6,
		0x18C7: c.op18C7,
		0x18C8: c.op18C8,
		0x18C9: c.op18C9,
		0x18CA: c.op18CA,
		0x18CB: c.op18CB,
		0x18CC: c.op18CC,
		0x18CD: c.op18CD,
		0x18CE: c.op18CE,
		0x18CF: c.op18CF,
		0x18D0: c.op18D0,
		0x18D1: c.op18D1,
		0x18D2: c.op18D2,
		0x18D3: c.op18D3,
		0x18D4: c.op18D4,
		0x18D5: c.op18D5,
		0x18D6: c.op18D6,
		0x18D7: c.op18D7,
		0x18F8: c.op18F8,
		0x18F9: c.op18F9,
		0x18FC: c.op18FC,
		0x1900: c.op1900,
		0x1901: c.op1901,
		0x1902: c.op1902,
		0x1903: c.op1903,
		0x1904: c.op1904,
		0x1905: c.op1905,
		0x1906: c.op1906,
		0x1907: c.op1907,
		0x1908: c.op1908,
		0x1909: c.op1909,
		0x190A: c.op190A,
		0x190B: c.op190B,
		0x190C: c.op190C,
		0x190D: c.op190D,
		0x190E: c.op190E,
		0x190F: c.op190F,
		0x1910: c.op1910,
		0x1911: c.op1911,
		0x1912: c.op1912,
		0x1913: c.op1913,
		0x1914: c.op1914,
		0x1915: c.op1915,
		0x1916: c.op1916,
		0x1917: c.op1917,
		0x1938: c.op1938,
		0x1939: c.op1939,
		0x193C: c.op193C,
		0x1940: c.op1940,
		0x1941: c.op1941,
		0x1942: c.op1942,
		0x1943: c.op1943,
		0x1944: c.op1944,
		0x1945: c.op1945,
		0x1946: c.op1946,
		0x1947: c.op1947,
		0x1948: c.op1948,
		0x1949: c.op1949,
		0x194A: c.op194A,
		0x194B: c.op194B,
		0x194C: c.op194C,
		0x194D: c.op194D,
		0x194E: c.op194E,
		0x194F: c.op194F,
		0x1950: c.op1950,
		0x1951: c.op1951,
		0x1952: c.op1952,
		0x1953: c.op1953,
		0x1954: c.op1954,
		0x1955: c.op1955,
		0x1956: c.op1956,
		0x1957: c.op1957,
		0x1978: c.op1978,
		0x1979: c.op1979,
		0x197C: c.op197C,
		0x1A00: c.op1A00,
		0x1A01: c.op1A01,
		0x1A02: c.op1A02,
		0x1A03: c.op1A03,
		0x1A04: c.op1A04,
		0x1A05: c.op1A05,
		0x1A06: c.op1A06,
		0x1A07: c.op1A07,
		0x1A08: c.op1A08,
		0x1A09: c.op1A09,
		0x1A0A: c.op1A0A,
		0x1A0B: c.op1A0B,
		0x1A0C: c.op1A0C,
		0x1A0D: c.op1A0D,
		0x1A0E: c.op1A0E,
		0x1A0F: c.op1A0F,
		0x1A10: c.op1A10,
		0x1A11: c.op1A11,
		0x1A12: c.op1A12,
		0x1A13: c.op1A13,
		0x1A14: c.op1A14,
		0x1A15: c.op1A15,
		0x1A16: c.op1A16,
		0x1A17: c.op1A17,
		0x1A38: c.op1A38,
		0x1A39: c.op1A39,
		0x1A3C: c.op1A3C,
		0x1A40: c.op1A40,
		0x1A41: c.op1A41,
		0x1A42: c.op1A42,
		0x1A43: c.op1A43,
		0x1A44: c.op1A44,
		0x1A45: c.op1A45,
		0x1A46: c.op1A46,
		0x1A47: c.op1A47,
		0x1A48: c.op1A48,
		0x1A49: c.op1A49,
		0x1A4A: c.op1A4A,
		0x1A4B: c.op1A4B,
		0x1A4C: c.op1A4C,
		0x1A4D: c.op1A4D,
		0x1A4E: c.op1A4E,
		0x1A4F: c.op1A4F,
		0x1A50: c.op1A50,
		0x1A51: c.op1A51,
		0x1A52: c.op1A52,
		0x1A53: c.op1A53,
		0x1A54: c.op1A54,
		0x1A55: c.op1A55,
		0x1A56: c.op1A56,
		0x1A57: c.op1A57,
		0x1A78: c.op1A78,
		0x1A79: c.op1A79,
		0x1A7C: c.op1A7C,
		0x1A80: c.op1A80,
		0x1A81: c.op1A81,
		0x1A82: c.op1A82,
		0x1A83: c.op1A83,
		0x1A84: c.op1A84,
		0x1A85: c.op1A85,
		0x1A86: c.op1A86,
		0x1A87: c.op1A87,
		0x1A88: c.op1A88,
		0x1A89: c.op1A89,
		0x1A8A: c.op1A8A,
		0x1A8B: c.op1A8B,
		0x1A8C: c.op1A8C,
		0x1A8D: c.op1A8D,
		0x1A8E: c.op1A8E,
		0x1A8F: c.op1A8F,
		0x1A90: c.op1A90,
		0x1A91: c.op1A91,
		0x1A92: c.op1A92,
		0x1A93: c.op1A93,
		0x1A94: c.op1A94,
		0x1A95: c.op1A95,
		0x1A96: c.op1A96,
		0x1A97: c.op1A97,
		0x1AB8: c.op1AB8,
		0x1AB9: c.op1AB9,
		0x1ABC: c.op1ABC,
		0x1AC0: c.op1AC0,
		0x1AC1: c.op1AC1,
		0x1AC2: c.op1AC2,
		0x1AC3: c.op1AC3,
		0x1AC4: c.op1AC4,
		0x1AC5: c.op1AC5,
		0x1AC6: c.op1AC6,
		0x1AC7: c.op1AC7,
		0x1AC8: c.op1AC8,
		0x1AC9: c.op1AC9,
		0x1ACA: c.op1ACA,
		0x1ACB: c.op1ACB,
		0x1ACC: c.op1ACC,
		0x1ACD: c.op1ACD,
		0x1ACE: c.op1ACE,
		0x1ACF: c.op1ACF,
		0x1AD0: c.op1AD0,
		0x1AD1: c.op1AD1,
		0x1AD2: c.op1AD2,
		0x1AD3: c.op1AD3,
		0x1AD4: c.op1AD4,
		0x1AD5: c.op1AD5,
		0x1AD6: c.op1AD6,
		0x1AD7: c.op1AD7,
		0x1AF8: c.op1AF8,
		0x1AF9: c.op1AF9,
		0x1AFC: c.op1AFC,
		0x1B00: c.op1B00,
		0x1B01: c.op1B01,
		0x1B02: c.op1B02,
		0x1B03: c.op1B03,
		0x1B04: c.op1B04,
		0x1B05: c.op1B05,
		0x1B06: c.op1B06,
		0x1B07: c.op1B07,
		0x1B08: c.op1B08,
		0x1B09: c.op1B09,
		0x1B0A: c.op1B0A,
		0x1B0B: c.op1B0B,
		0x1B0C: c.op1B0C,
		0x1B0D: c.op1B0D,
		0x1B0E: c.op1B0E,
		0x1B0F: c.op1B0F,
		0x1B10: c.op1B10,
		0x1B11: c.op1B11,
		0x1B12: c.op1B12,
		0x1B13: c.op1B13,
		0x1B14: c.op1B14,
		0x1B15: c.op1B15,
		0x1B16: c.op1B16,
		0x1B17: c.op1B17,
		0x1B38: c.op1B38,
		0x1B39: c.op1B39,
		0x1B3C: c.op1B3C,
		0x1B40: c.op1B40,
		0x1B41: c.op1B41,
		0x1B42: c.op1B42,
		0x1B43: c.op1B43,
		0x1B44: c.op1B44,
		0x1B45: c.op1B45,
		0x1B46: c.op1B46,
		0x1B47: c.op1B47,
		0x1B48: c.op1B48,
		0x1B49: c.op1B49,
		0x1B4A: c.op1B4A,
		0x1B4B: c.op1B4B,
		0x1B4C: c.op1B4C,
		0x1B4D: c.op1B4D,
		0x1B4E: c.op1B4E,
		0x1B4F: c.op1B4F,
		0x1B50: c.op1B50,
		0x1B51: c.op1B51,
		0x1B52: c.op1B52,
		0x1B53: c.op1B53,
		0x1B54: c.op1B54,
		0x1B55: c.op1B55,
		0x1B56: c.op1B56,
		0x1B57: c.op1B57,
		0x1B78: c.op1B78,
		0x1B79: c.op1B79,
		0x1B7C: c.op1B7C,
		0x1C00: c.op1C00,
		0x1C01: c.op1C01,
		0x1C02: c.op1C02,
		0x1C03: c.op1C03,
		0x1C04: c.op1C04,
		0x1C05: c.op1C05,
		0x1C06: c.op1C06,
		0x1C07: c.op1C07,
		0x1C08: c.op1C08,
		0x1C09: c.op1C09,
		0x1C0A: c.op1C0A,
		0x1C0B: c.op1C0B,
		0x1C0C: c.op1C0C,
		0x1C0D: c.op1C0D,
		0x1C0E: c.op1C0E,
		0x1C0F: c.op1C0F,
		0x1C10: c.op1C10,
		0x1C11: c.op1C11,
		0x1C12: c.op1C12,
		0x1C13: c.op1C13,
		0x1C14: c.op1C14,
		0x1C15: c.op1C15,
		0x1C16: c.op1C16,
		0x1C17: c.op1C17,
		0x1C38: c.op1C38,
		0x1C39: c.op1C39,
		0x1C3C: c.op1C3C,
		0x1C40: c.op1C40,
		0x1C41: c.op1C41,
		0x1C42: c.op1C42,
		0x1C43: c.op1C43,
		0x1C44: c.op1C44,
		0x1C45: c.op1C45,
		0x1C46: c.op1C46,
		0x1C47: c.op1C47,
		0x1C48: c.op1C48,
		0x1C49: c.op1C49,
		0x1C4A: c.op1C4A,
		0x1C4B: c.op1C4B,
		0x1C4C: c.op1C4C,
		0x1C4D: c.op1C4D,
		0x1C4E: c.op1C4E,
		0x1C4F: c.op1C4F,
		0x1C50: c.op1C50,
		0x1C51: c.op1C51,
		0x1C52: c.op1C52,
		0x1C53: c.op1C53,
		0x1C54: c.op1C54,
		0x1C55: c.op1C55,
		0x1C56: c.op1C56,
		0x1C57: c.op1C57,
		0x1C78: c.op1C78,
		0x1C79: c.op1C79,
		0x1C7C: c.op1C7C,
		0x1C80: c.op1C80,
		0x1C81: c.op1C81,
		0x1C82: c.op1C82,
		0x1C83: c.op1C83,
		0x1C84: c.op1C84,
		0x1C85: c.op1C85,
		0x1C86: c.op1C86,
		0x1C87: c.op1C87,
		0x1C88: c.op1C88,
		0x1C89: c.op1C89,
		0x1C8A: c.op1C8A,
		0x1C8B: c.op1C8B,
		0x1C8C: c.op1C8C,
		0x1C8D: c.op1C8D,
		0x1C8E: c.op1C8E,
		0x1C8F: c.op1C8F,
		0x1C90: c.op1C90,
		0x1C91: c.op1C91,
		0x1C92: c.op1C92,
		0x1C93: c.op1C93,
		0x1C94: c.op1C94,
		0x1C95: c.op1C95,
		0x1C96: c.op1C96,
		0x1C97: c.op1C97,
		0x1CB8: c.op1CB8,
		0x1CB9: c.op1CB9,
		0x1CBC: c.op1CBC,
		0x1CC0: c.op1CC0,
		0x1CC1: c.op1CC1,
		0x1CC2: c.op1CC2,
		0x1CC3: c.op1CC3,
		0x1CC4: c.op1CC4,
		0x1CC5: c.op1CC5,
		0x1CC6: c.op1CC6,
		0x1CC7: c.op1CC7,
		0x1CC8: c.op1CC8,
		0x1CC9: c.op1CC9,
		0x1CCA: c.op1CCA,
		0x1CCB: c.op1CCB,
		0x1CCC: c.op1CCC,
		0x1CCD: c.op1CCD,
		0x1CCE: c.op1CCE,
		0x1CCF: c.op1CCF,
		0x1CD0: c.op1CD0,
		0x1CD1: c.op1CD1,
		0x1CD2: c.op1CD2,
		0x1CD3: c.op1CD3,
		0x1CD4: c.op1CD4,
		0x1CD5: c.op1CD5,
		0x1CD6: c.op1CD6,
		0x1CD7: c.op1CD7,
		0x1CF8: c.op1CF8,
		0x1CF9: c.op1CF9,
		0x1CFC: c.op1CFC,
		0x1D00: c.op1D00,
		0x1D01: c.op1D01,
		0x1D02: c.op1D02,
		0x1D03: c.op1D03,
		0x1D04: c.op1D04,
		0x1D05: c.op1D05,
		0x1D06: c.op1D06,
		0x1D07: c.op1D07,
		0x1D08: c.op1D08,
		0x1D09: c.op1D09,
		0x1D0A: c.op1D0A,
		0x1D0B: c.op1D0B,
		0x1D0C: c.op1D0C,
		0x1D0D: c.op1D0D,
		0x1D0E: c.op1D0E,
		0x1D0F: c.op1D0F,
		0x1D10: c.op1D10,
		0x1D11: c.op1D11,
		0x1D12: c.op1D12,
		0x1D13: c.op1D13,
		0x1D14: c.op1D14,
		0x1D15: c.op1D15,
		0x1D16: c.op1D16,
		0x1D17: c.op1D17,
		0x1D38: c.op1D38,
		0x1D39: c.op1D39,
		0x1D3C: c.op1D3C,
		0x1D40: c.op1D40,
		0x1D41: c.op1D41,
		0x1D42: c.op1D42,
		0x1D43: c.op1D43,
		0x1D44: c.op1D44,
		0x1D45: c.op1D45,
		0x1D46: c.op1D46,
		0x1D47: c.op1D47,
		0x1D48: c.op1D48,
		0x1D49: c.op1D49,
		0x1D4A: c.op1D4A,
		0x1D4B: c.op1D4B,
		0x1D4C: c.op1D4C,
		0x1D4D: c.op1D4D,
		0x1D4E: c.op1D4E,
		0x1D4F: c.op1D4F,
		0x1D50: c.op1D50,
		0x1D51: c.op1D51,
		0x1D52: c.op1D52,
		0x1D53: c.op1D53,
		0x1D54: c.op1D54,
		0x1D55: c.op1D55,
		0x1D56: c.op1D56,
		0x1D57: c.op1D57,
		0x1D78: c.op1D78,
		0x1D79: c.op1D79,
		0x1D7C: c.op1D7C,
		0x1E00: c.op1E00,
		0x1E01: c.op1E01,
		0x1E02: c.op1E02,
		0x1E03: c.op1E03,
		0x1E04: c.op1E04,
		0x1E05: c.op1E05,
		0x1E06: c.op1E06,
		0x1E07: c.op1E07,
		0x1E08: c.op1E08,
		0x1E09: c.op1E09,
		0x1E0A: c.op1E0A,
		0x1E0B: c.op1E0B,
		0x1E0C: c.op1E0C,
		0x1E0D: c.op1E0D,
		0x1E0E: c.op1E0E,
		0x1E0F: c.op1E0F,
		0x1E10: c.op1E10,
		0x1E11: c.op1E11,
		0x1E12: c.op1E12,
		0x1E13: c.op1E13,
		0x1E14: c.op1E14,
		0x1E15: c.op1E15,
		0x1E16: c.op1E16,
		0x1E17: c.op1E17,
		0x1E38: c.op1E38,
		0x1E39: c.op1E39,
		0x1E3C: c.op1E3C,
		0x1E40: c.op1E40,
		0x1E41: c.op1E41,
		0x1E42: c.op1E42,
		0x1E43: c.op1E43,
		0x1E44: c.op1E44,
		0x1E45: c.op1E45,
		0x1E46: c.op1E46,
		0x1E47: c.op1E47,
		0x1E48: c.op1E48,
		0x1E49: c.op1E49,
		0x1E4A: c.op1E4A,
		0x1E4B: c.op1E4B,
		0x1E4C: c.op1E4C,
		0x1E4D: c.op1E4D,
		0x1E4E: c.op1E4E,
		0x1E4F: c.op1E4F,
		0x1E50: c.op1E50,
		0x1E51: c.op1E51,
		0x1E52: c.op1E52,
		0x1E53: c.op1E53,
		0x1E54: c.op1E54,
		0x1E55: c.op1E55,
		0x1E56: c.op1E56,
		0x1E57: c.op1E57,
		0x1E78: c.op1E78,
		0x1E79: c.op1E79,
		0x1E7C: c.op1E7C,
		0x1E80: c.op1E80,
		0x1E81: c.op1E81,
		0x1E82: c.op1E82,
		0x1E83: c.op1E83,
		0x1E84: c.op1E84,
		0x1E85: c.op1E85,
		0x1E86: c.op1E86,
		0x1E87: c.op1E87,
		0x1E88: c.op1E88,
		0x1E89: c.op1E89,
		0x1E8A: c.op1E8A,
		0x1E8B: c.op1E8B,
		0x1E8C: c.op1E8C,
		0x1E8D: c.op1E8D,
		0x1E8E: c.op1E8E,
		0x1E8F: c.op1E8F,
		0x1E90: c.op1E90,
		0x1E91: c.op1E91,
		0x1E92: c.op1E92,
		0x1E93: c.op1E93,
		0x1E94: c.op1E94,
		0x1E95: c.op1E95,
		0x1E96: c.op1E96,
		0x1E97: c.op1E97,
		0x1EB8: c.op1EB8,
		0x1EB9: c.op1EB9,
		0x1EBC: c.op1EBC,
		0x1EC0: c.op1EC0,
		0x1EC1: c.op1EC1,
		0x1EC2: c.op1EC2,
		0x1EC3: c.op1EC3,
		0x1EC4: c.op1EC4,
		0x1EC5: c.op1EC5,
		0x1EC6: c.op1EC6,
		0x1EC7: c.op1EC7,
		0x1EC8: c.op1EC8,
		0x1EC9: c.op1EC9,
		0x1ECA: c.op1ECA,
		0x1ECB: c.op1ECB,
		0x1ECC: c.op1ECC,
		0x1ECD: c.op1ECD,
		0x1ECE: c.op1ECE,
		0x1ECF: c.op1ECF,
		0x1ED0: c.op1ED0,
		0x1ED1: c.op1ED1,
		0x1ED2: c.op1ED2,
		0x1ED3: c.op1ED3,
		0x1ED4: c.op1ED4,
		0x1ED5: c.op1ED5,
		0x1ED6: c.op1ED6,
		0x1ED7: c.op1ED7,
		0x1EF8: c.op1EF8,
		0x1EF9: c.op1EF9,
		0x1EFC: c.op1EFC,
		0x1F00: c.op1F00,
		0x1F01: c.op1F01,
		0x1F02: c.op1F02,
		0x1F03: c.op1F03,
		0x1F04: c.op1F04,
		0x1F05: c.op1F05,
		0x1F06: c.op1F06,
		0x1F07: c.op1F07,
		0x1F08: c.op1F08,
		0x1F09: c.op1F09,
		0x1F0A: c.op1F0A,
		0x1F0B: c.op1F0B,
		0x1F0C: c.op1F0C,
		0x1F0D: c.op1F0D,
		0x1F0E: c.op1F0E,
		0x1F0F: c.op1F0F,
		0x1F10: c.op1F10,
		0x1F11: c.op1F11,
		0x1F12: c.op1F12,
		0x1F13: c.op1F13,
		0x1F14: c.op1F14,
		0x1F15: c.op1F15,
		0x1F16: c.op1F16,
		0x1F17: c.op1F17,
		0x1F38: c.op1F38,
		0x1F39: c.op1F39,
		0x1F3C: c.op1F3C,
		0x1F40: c.op1F40,
		0x1F41: c.op1F41,
		0x1F42: c.op1F42,
		0x1F43: c.op1F43,
		0x1F44: c.op1F44,
		0x1F45: c.op1F45,
		0x1F46: c.op1F46,
		0x1F47: c.op1F47,
		0x1F48: c.op1F48,
		0x1F49: c.op1F49,
		0x1F4A: c.op1F4A,
		0x1F4B: c.op1F4B,
		0x1F4C: c.op1F4C,
		0x1F4D: c.op1F4D,
		0x1F4E: c.op1F4E,
		0x1F4F: c.op1F4F,
		0x1F50: c.op1F50,
		0x1F51: c.op1F51,
		0x1F52: c.op1F52,
		0x1F53: c.op1F53,
		0x1F54: c.op1F54,
		0x1F55: c.op1F55,
		0x1F56: c.op1F56,
		0x1F57: c.op1F57,
		0x1F78: c.op1F78,
		0x1F79: c.op1F79,
		0x1F7C: c.op1F7C,
		0x2000: c.op2000,
		0x2001: c.op2001,
		0x2002: c.op2002,
		0x2003: c.op2003,
		0x2004: c.op2004,
		0x2005: c.op2005,
		0x2006: c.op2006,
		0x2007: c.op2007,
		0x2008: c.op2008,
		0x2009: c.op2009,
		0x200A: c.op200A,
		0x200B: c.op200B,
		0x200C: c.op200C,
		0x200D: c.op200D,
		0x200E: c.op200E,
		0x200F: c.op200F,
		0x2010: c.op2010,
		0x2011: c.op2011,
		0x2012: c.op2012,
		0x2013: c.op2013,
		0x2014: c.op2014,
		0x2015: c.op2015,
		0x2016: c.op2016,
		0x2017: c.op2017,
		0x2038: c.op2038,
		0x2039: c.op2039,
		0x203C: c.op203C,
		0x2040: c.op2040,
		0x2041: c.op2041,
		0x2042: c.op2042,
		0x2043: c.op2043,
		0x2044: c.op2044,
		0x2045: c.op2045,
		0x2046: c.op2046,
		0x2047: c.op2047,
		0x2048: c.op2048,
		0x2049: c.op2049,
		0x204A: c.op204A,
		0x204B: c.op204B,
		0x204C: c.op204C,
		0x204D: c.op204D,
		0x204E: c.op204E,
		0x204F: c.op204F,
		0x2050: c.op2050,
		0x2051: c.op2051,
		0x2052: c.op2052,
		0x2053: c.op2053,
		0x2054: c.op2054,
		0x2055: c.op2055,
		0x2056: c.op2056,
		0x2057: c.op2057,
		0x2078: c.op2078,
		0x2079: c.op2079,
		0x207C: c.op207C,
		0x2080: c.op2080,
		0x2081: c.op2081,
		0x2082: c.op2082,
		0x2083: c.op2083,
		0x2084: c.op2084,
		0x2085: c.op2085,
		0x2086: c.op2086,
		0x2087: c.op2087,
		0x2088: c.op2088,
		0x2089: c.op2089,
		0x208A: c.op208A,
		0x208B: c.op208B,
		0x208C: c.op208C,
		0x208D: c.op208D,
		0x208E: c.op208E,
		0x208F: c.op208F,
		0x2090: c.op2090,
		0x2091: c.op2091,
		0x2092: c.op2092,
		0x2093: c.op2093,
		0x2094: c.op2094,
		0x2095: c.op2095,
		0x2096: c.op2096,
		0x2097: c.op2097,
		0x20B8: c.op20B8,
		0x20B9: c.op20B9,
		0x20BC: c.op20BC,
		0x20C0: c.op20C0,
		0x20C1: c.op20C1,
		0x20C2: c.op20C2,
		0x20C3: c.op20C3,
		0x20C4: c.op20C4,
		0x20C5: c.op20C5,
		0x20C6: c.op20C6,
		0x20C7: c.op20C7,
		0x20C8: c.op20C8,
		0x20C9: c.op20C9,
		0x20CA: c.op20CA,
		0x20CB: c.op20CB,
		0x20CC: c.op20CC,
		0x20CD: c.op20CD,
		0x20CE: c.op20CE,
		0x20CF: c.op20CF,
		0x20D0: c.op20D0,
		0x20D1: c.op20D1,
		0x20D2: c.op20D2,
		0x20D3: c.op20D3,
		0x20D4: c.op20D4,
		0x20D5: c.op20D5,
		0x20D6: c.op20D6,
		0x20D7: c.op20D7,
		0x20F8: c.op20F8,
		0x20F9: c.op20F9,
		0x20FC: c.op20FC,
		0x2100: c.op2100,
		0x2101: c.op2101,
		0x2102: c.op2102,
		0x2103: c.op2103,
		0x2104: c.op2104,
		0x2105: c.op2105,
		0x2106: c.op2106,
		0x2107: c.op2107,
		0x2108: c.op2108,
		0x2109: c.op2109,
		0x210A: c.op210A,
		0x210B: c.op210B,
		0x210C: c.op210C,
		0x210D: c.op210D,
		0x210E: c.op210E,
		0x210F: c.op210F,
		0x2110: c.op2110,
		0x2111: c.op2111,
		0x2112: c.op2112,
		0x2113: c.op2113,
		0x2114: c.op2114,
		0x2115: c.op2115,
		0x2116: c.op2116,
		0x2117: c.op2117,
		0x2138: c.op2138,
		0x2139: c.op2139,
		0x213C: c.op213C,
		0x2140: c.op2140,
		0x2141: c.op2141,
		0x2142: c.op2142,
		0x2143: c.op2143,
		0x2144: c.op2144,
		0x2145: c.op2145,
		0x2146: c.op2146,
		0x2147: c.op2147,
		0x2148: c.op2148,
		0x2149: c.op2149,
		0x214A: c.op214A,
		0x214B: c.op214B,
		0x214C: c.op214C,
		0x214D: c.op214D,
		0x214E: c.op214E,
		0x214F: c.op214F,
		0x2150: c.op2150,
		0x2151: c.op2151,
		0x2152: c.op2152,
		0x2153: c.op2153,
		0x2154: c.op2154,
		0x2155: c.op2155,
		0x2156: c.op2156,
		0x2157: c.op2157,
		0x2178: c.op2178,
		0x2179: c.op2179,
		0x217C: c.op217C,
		0x21C0: c.op21C0,
		0x21C1: c.op21C1,
		0x21C2: c.op21C2,
		0x21C3: c.op21C3,
		0x21C4: c.op21C4,
		0x21C5: c.op21C5,
		0x21C6: c.op21C6,
		0x21C7: c.op21C7,
		0x21C8: c.op21C8,
		0x21C9: c.op21C9,
		0x21CA: c.op21CA,
		0x21CB: c.op21CB,
		0x21CC: c.op21CC,
		0x21CD: c.op21CD,
		0x21CE: c.op21CE,
		0x21CF: c.op21CF,
		0x21D0: c.op21D0,
		0x21D1: c.op21D1,
		0x21D2: c.op21D2,
		0x21D3: c.op21D3,
		0x21D4: c.op21D4,
		0x21D5: c.op21D5,
		0x21D6: c.op21D6,
		0x21D7: c.op21D7,
		0x21F8: c.op21F8,
		0x21F9: c.op21F9,
		0x21FC: c.op21FC,
		0x2200: c.op2200,
		0x2201: c.op2201,
		0x2202: c.op2202,
		0x2203: c.op2203,
		0x2204: c.op2204,
		0x2205: c.op2205,
		0x2206: c.op2206,
		0x2207: c.op2207,
		0x2208: c.op2208,
		0x2209: c.op2209,
		0x220A: c.op220A,
		0x220B: c.op220B,
		0x220C: c.op220C,
		0x220D: c.op220D,
		0x220E: c.op220E,
		0x220F: c.op220F,
		0x2210: c.op2210,
		0x2211: c.op2211,
		0x2212: c.op2212,
		0x2213: c.op2213,
		0x2214: c.op2214,
		0x2215: c.op2215,
		0x2216: c.op2216,
		0x2217: c.op2217,
		0x2238: c.op2238,
		0x2239: c.op2239,
		0x223C: c.op223C,
		0x2240: c.op2240,
		0x2241: c.op2241,
		0x2242: c.op2242,
		0x2243: c.op2243,
		0x2244: c.op2244,
		0x2245: c.op2245,
		0x2246: c.op2246,
		0x2247: c.op2247,
		0x2248: c.op2248,
		0x2249: c.op2249,
		0x224A: c.op224A,
		0x224B: c.op224B,
		0x224C: c.op224C,
		0x224D: c.op224D,
		0x224E: c.op224E,
		0x224F: c.op224F,
		0x2250: c.op2250,
		0x2251: c.op2251,
		0x2252: c.op2252,
		0x2253: c.op2253,
		0x2254: c.op2254,
		0x2255: c.op2255,
		0x2256: c.op2256,
		0x2257: c.op2257,
		0x2278: c.op2278,
		0x2279: c.op2279,
		0x227C: c.op227C,
		0x2280: c.op2280,
		0x2281: c.op2281,
		0x2282: c.op2282,
		0x2283: c.op2283,
		0x2284: c.op2284,
		0x2285: c.op2285,
		0x2286: c.op2286,
		0x2287: c.op2287,
		0x2288: c.op2288,
		0x2289: c.op2289,
		0x228A: c.op228A,
		0x228B: c.op228B,
		0x228C: c.op228C,
		0x228D: c.op228D,
		0x228E: c.op228E,
		0x228F: c.op228F,
		0x2290: c.op2290,
		0x2291: c.op2291,
		0x2292: c.op2292,
		0x2293: c.op2293,
		0x2294: c.op2294,
		0x2295: c.op2295,
		0x2296: c.op2296,
		0x2297: c.op2297,
		0x22B8: c.op22B8,
		0x22B9: c.op22B9,
		0x22BC: c.op22BC,
		0x22C0: c.op22C0,
		0x22C1: c.op22C1,
		0x22C2: c.op22C2,
		0x22C3: c.op22C3,
		0x22C4: c.op22C4,
		0x22C5: c.op22C5,
		0x22C6: c.op22C6,
		0x22C7: c.op22C7,
		0x22C8: c.op22C8,
		0x22C9: c.op22C9,
		0x22CA: c.op22CA,
		0x22CB: c.op22CB,
		0x22CC: c.op22CC,
		0x22CD: c.op22CD,
		0x22CE: c.op22CE,
		0x22CF: c.op22CF,
		0x22D0: c.op22D0,
		0x22D1: c.op22D1,
		0x22D2: c.op22D2,
		0x22D3: c.op22D3,
		0x22D4: c.op22D4,
		0x22D5: c.op22D5,
		0x22D6: c.op22D6,
		0x22D7: c.op22D7,
		0x22F8: c.op22F8,
		0x22F9: c.op22F9,
		0x22FC: c.op22FC,
		0x2300: c.op2300,
		0x2301: c.op2301,
		0x2302: c.op2302,
		0x2303: c.op2303,
		0x2304: c.op2304,
		0x2305: c.op2305,
		0x2306: c.op2306,
		0x2307: c.op2307,
		0x2308: c.op2308,
		0x2309: c.op2309,
		0x230A: c.op230A,
		0x230B: c.op230B,
		0x230C: c.op230C,
		0x230D: c.op230D,
		0x230E: c.op230E,
		0x230F: c.op230F,
		0x2310: c.op2310,
		0x2311: c.op2311,
		0x2312: c.op2312,
		0x2313: c.op2313,
		0x2314: c.op2314,
		0x2315: c.op2315,
		0x2316: c.op2316,
		0x2317: c.op2317,
		0x2338: c.op2338,
		0x2339: c.op2339,
		0x233C: c.op233C,
		0x2340: c.op2340,
		0x2341: c.op2341,
		0x2342: c.op2342,
		0x2343: c.op2343,
		0x2344: c.op2344,
		0x2345: c.op2345,
		0x2346: c.op2346,
		0x2347: c.op2347,
		0x2348: c.op2348,
		0x2349: c.op2349,
		0x234A: c.op234A,
		0x234B: c.op234B,
		0x234C: c.op234C,
		0x234D: c.op234D,
		0x234E: c.op234E,
		0x234F: c.op234F,
		0x2350: c.op2350,
		0x2351: c.op2351,
		0x2352: c.op2352,
		0x2353: c.op2353,
		0x2354: c.op2354,
		0x2355: c.op2355,
		0x2356: c.op2356,
		0x2357: c.op2357,
		0x2378: c.op2378,
		0x2379: c.op2379,
		0x237C: c.op237C,
		0x23C0: c.op23C0,
		0x23C1: c.op23C1,
		0x23C2: c.op23C2,
		0x23C3: c.op23C3,
		0x23C4: c.op23C4,
		0x23C5: c.op23C5,
		0x23C6: c.op23C6,
		0x23C7: c.op23C7,
		0x23C8: c.op23C8,
		0x23C9: c.op23C9,
		0x23CA: c.op23CA,
		0x23CB: c.op23CB,
		0x23CC: c.op23CC,
		0x23CD: c.op23CD,
		0x23CE: c.op23CE,
		0x23CF: c.op23CF,
		0x23D0: c.op23D0,
		0x23D1: c.op23D1,
		0x23D2: c.op23D2,
		0x23D3: c.op23D3,
		0x23D4: c.op23D4,
		0x23D5: c.op23D5,
		0x23D6: c.op23D6,
		0x23D7: c.op23D7,
		0x23F8: c.op23F8,
		0x23F9: c.op23F9,
		0x23FC: c.op23FC,
		0x2400: c.op2400,
		0x2401: c.op2401,
		0x2402: c.op2402,
		0x2403: c.op2403,
		0x2404: c.op2404,
		0x2405: c.op2405,
		0x2406: c.op2406,
		0x2407: c.op2407,
		0x2408: c.op2408,
		0x2409: c.op2409,
		0x240A: c.op240A,
		0x240B: c.op240B,
		0x240C: c.op240C,
		0x240D: c.op240D,
		0x240E: c.op240E,
		0x240F: c.op240F,
		0x2410: c.op2410,
		0x2411: c.op2411,
		0x2412: c.op2412,
		0x2413: c.op2413,
		0x2414: c.op2414,
		0x2415: c.op2415,
		0x2416: c.op2416,
		0x2417: c.op2417,
		0x2438: c.op2438,
		0x2439: c.op2439,
		0x243C: c.op243C,
		0x2440: c.op2440,
		0x2441: c.op2441,
		0x2442: c.op2442,
		0x2443: c.op2443,
		0x2444: c.op2444,
		0x2445: c.op2445,
		0x2446: c.op2446,
		0x2447: c.op2447,
		0x2448: c.op2448,
		0x2449: c.op2449,
		0x244A: c.op244A,
		0x244B: c.op244B,
		0x244C: c.op244C,
		0x244D: c.op244D,
		0x244E: c.op244E,
		0x244F: c.op244F,
		0x2450: c.op2450,
		0x2451: c.op2451,
		0x2452: c.op2452,
		0x2453: c.op2453,
		0x2454: c.op2454,
		0x2455: c.op2455,
		0x2456: c.op2456,
		0x2457: c.op2457,
		0x2478: c.op2478,
		0x2479: c.op2479,
		0x247C: c.op247C,
		0x2480: c.op2480,
		0x2481: c.op2481,
		0x2482: c.op2482,
		0x2483: c.op2483,
		0x2484: c.op2484,
		0x2485: c.op2485,
		0x2486: c.op2486,
		0x2487: c.op2487,
		0x2488: c.op2488,
		0x2489: c.op2489,
		0x248A: c.op248A,
		0x248B: c.op248B,
		0x248C: c.op248C,
		0x248D: c.op248D,
		0x248E: c.op248E,
		0x248F: c.op248F,
		0x2490: c.op2490,
		0x2491: c.op2491,
		0x2492: c.op2492,
		0x2493: c.op2493,
		0x2494: c.op2494,
		0x2495: c.op2495,
		0x2496: c.op2496,
		0x2497: c.op2497,
		0x24B8: c.op24B8,
		0x24B9: c.op24B9,
		0x24BC: c.op24BC,
		0x24C0: c.op24C0,
		0x24C1: c.op24C1,
		0x24C2: c.op24C2,
		0x24C3: c.op24C3,
		0x24C4: c.op24C4,
		0x24C5: c.op24C5,
		0x24C6: c.op24C6,
		0x24C7: c.op24C7,
		0x24C8: c.op24C8,
		0x24C9: c.op24C9,
		0x24CA: c.op24CA,
		0x24CB: c.op24CB,
		0x24CC: c.op24CC,
		0x24CD: c.op24CD,
		0x24CE: c.op24CE,
		0x24CF: c.op24CF,
		0x24D0: c.op24D0,
		0x24D1: c.op24D1,
		0x24D2: c.op24D2,
		0x24D3: c.op24D3,
		0x24D4: c.op24D4,
		0x24D5: c.op24D5,
		0x24D6: c.op24D6,
		0x24D7: c.op24D7,
		0x24F8: c.op24F8,
		0x24F9: c.op24F9,
		0x24FC: c.op24FC,
		0x2500: c.op2500,
		0x2501: c.op2501,
		0x2502: c.op2502,
		0x2503: c.op2503,
		0x2504: c.op2504,
		0x2505: c.op2505,
		0x2506: c.op2506,
		0x2507: c.op2507,
		0x2508: c.op2508,
		0x2509: c.op2509,
		0x250A: c.op250A,
		0x250B: c.op250B,
		0x250C: c.op250C,
		0x250D: c.op250D,
		0x250E: c.op250E,
		0x250F: c.op250F,
		0x2510: c.op2510,
		0x2511: c.op2511,
		0x2512: c.op2512,
		0x2513: c.op2513,
		0x2514: c.op2514,
		0x2515: c.op2515,
		0x2516: c.op2516,
		0x2517: c.op2517,
		0x2538: c.op2538,
		0x2539: c.op2539,
		0x253C: c.op253C,
		0x2540: c.op2540,
		0x2541: c.op2541,
		0x2542: c.op2542,
		0x2543: c.op2543,
		0x2544: c.op2544,
		0x2545: c.op2545,
		0x2546: c.op2546,
		0x2547: c.op2547,
		0x2548: c.op2548,
		0x2549: c.op2549,
		0x254A: c.op254A,
		0x254B: c.op254B,
		0x254C: c.op254C,
		0x254D: c.op254D,
		0x254E: c.op254E,
		0x254F: c.op254F,
		0x2550: c.op2550,
		0x2551: c.op2551,
		0x2552: c.op2552,
		0x2553: c.op2553,
		0x2554: c.op2554,
		0x2555: c.op2555,
		0x2556: c.op2556,
		0x2557: c.op2557,
		0x2578: c.op2578,
		0x2579: c.op2579,
		0x257C: c.op257C,
		0x2600: c.op2600,
		0x2601: c.op2601,
		0x2602: c.op2602,
		0x2603: c.op2603,
		0x2604: c.op2604,
		0x2605: c.op2605,
		0x2606: c.op2606,
		0x2607: c.op2607,
		0x2608: c.op2608,
		0x2609: c.op2609,
		0x260A: c.op260A,
		0x260B: c.op260B,
		0x260C: c.op260C,
		0x260D: c.op260D,
		0x260E: c.op260E,
		0x260F: c.op260F,
		0x2610: c.op2610,
		0x2611: c.op2611,
		0x2612: c.op2612,
		0x2613: c.op2613,
		0x2614: c.op2614,
		0x2615: c.op2615,
		0x2616: c.op2616,
		0x2617: c.op2617,
		0x2638: c.op2638,
		0x2639: c.op2639,
		0x263C: c.op263C,
		0x2640: c.op2640,
		0x2641: c.op2641,
		0x2642: c.op2642,
		0x2643: c.op2643,
		0x2644: c.op2644,
		0x2645: c.op2645,
		0x2646: c.op2646,
		0x2647: c.op2647,
		0x2648: c.op2648,
		0x2649: c.op2649,
		0x264A: c.op264A,
		0x264B: c.op264B,
		0x264C: c.op264C,
		0x264D: c.op264D,
		0x264E: c.op264E,
		0x264F: c.op264F,
		0x2650: c.op2650,
		0x2651: c.op2651,
		0x2652: c.op2652,
		0x2653: c.op2653,
		0x2654: c.op2654,
		0x2655: c.op2655,
		0x2656: c.op2656,
		0x2657: c.op2657,
		0x2678: c.op2678,
		0x2679: c.op2679,
		0x267C: c.op267C,
		0x2680: c.op2680,
		0x2681: c.op2681,
		0x2682: c.op2682,
		0x2683: c.op2683,
		0x2684: c.op2684,
		0x2685: c.op2685,
		0x2686: c.op2686,
		0x2687: c.op2687,
		0x2688: c.op2688,
		0x2689: c.op2689,
		0x268A: c.op268A,
		0x268B: c.op268B,
		0x268C: c.op268C,
		0x268D: c.op268D,
		0x268E: c.op268E,
		0x268F: c.op268F,
		0x2690: c.op2690,
		0x2691: c.op2691,
		0x2692: c.op2692,
		0x2693: c.op2693,
		0x2694: c.op2694,
		0x2695: c.op2695,
		0x2696: c.op2696,
		0x2697: c.op2697,
		0x26B8: c.op26B8,
		0x26B9: c.op26B9,
		0x26BC: c.op26BC,
		0x26C0: c.op26C0,
		0x26C1: c.op26C1,
		0x26C2: c.op26C2,
		0x26C3: c.op26C3,
		0x26C4: c.op26C4,
		0x26C5: c.op26C5,
		0x26C6: c.op26C6,
		0x26C7: c.op26C7,
		0x26C8: c.op26C8,
		0x26C9: c.op26C9,
		0x26CA: c.op26CA,
		0x26CB: c.op26CB,
		0x26CC: c.op26CC,
		0x26CD: c.op26CD,
		0x26CE: c.op26CE,
		0x26CF: c.op26CF,
		0x26D0: c.op26D0,
		0x26D1: c.op26D1,
		0x26D2: c.op26D2,
		0x26D3: c.op26D3,
		0x26D4: c.op26D4,
		0x26D5: c.op26D5,
		0x26D6: c.op26D6,
		0x26D7: c.op26D7,
		0x26F8: c.op26F8,
		0x26F9: c.op26F9,
		0x26FC: c.op26FC,
		0x2700: c.op2700,
		0x2701: c.op2701,
		0x2702: c.op2702,
		0x2703: c.op2703,
		0x2704: c.op2704,
		0x2705: c.op2705,
		0x2706: c.op2706,
		0x2707: c.op2707,
		0x2708: c.op2708,
		0x2709: c.op2709,
		0x270A: c.op270A,
		0x270B: c.op270B,
		0x270C: c.op270C,
		0x270D: c.op270D,
		0x270E: c.op270E,
		0x270F: c.op270F,
		0x2710: c.op2710,
		0x2711: c.op2711,
		0x2712: c.op2712,
		0x2713: c.op2713,
		0x2714: c.op2714,
		0x2715: c.op2715,
		0x2716: c.op2716,
		0x2717: c.op2717,
		0x2738: c.op2738,
		0x2739: c.op2739,
		0x273C: c.op273C,
		0x2740: c.op2740,
		0x2741: c.op2741,
		0x2742: c.op2742,
		0x2743: c.op2743,
		0x2744: c.op2744,
		0x2745: c.op2745,
		0x2746: c.op2746,
		0x2747: c.op2747,
		0x2748: c.op2748,
		0x2749: c.op2749,
		0x274A: c.op274A,
		0x274B: c.op274B,
		0x274C: c.op274C,
		0x274D: c.op274D,
		0x274E: c.op274E,
		0x274F: c.op274F,
		0x2750: c.op2750,
		0x2751: c.op2751,
		0x2752: c.op2752,
		0x2753: c.op2753,
		0x2754: c.op2754,
		0x2755: c.op2755,
		0x2756: c.op2756,
		0x2757: c.op2757,
		0x2778: c.op2778,
		0x2779: c.op2779,
		0x277C: c.op277C,
		0x2800: c.op2800,
		0x2801: c.op2801,
		0x2802: c.op2802,
		0x2803: c.op2803,
		0x2804: c.op2804,
		0x2805: c.op2805,
		0x2806: c.op2806,
		0x2807: c.op2807,
		0x2808: c.op2808,
		0x2809: c.op2809,
		0x280A: c.op280A,
		0x280B: c.op280B,
		0x280C: c.op280C,
		0x280D: c.op280D,
		0x280E: c.op280E,
		0x280F: c.op280F,
		0x2810: c.op2810,
		0x2811: c.op2811,
		0x2812: c.op2812,
		0x2813: c.op2813,
		0x2814: c.op2814,
		0x2815: c.op2815,
		0x2816: c.op2816,
		0x2817: c.op2817,
		0x2838: c.op2838,
		0x2839: c.op2839,
		0x283C: c.op283C,
		0x2840: c.op2840,
		0x2841: c.op2841,
		0x2842: c.op2842,
		0x2843: c.op2843,
		0x2844: c.op2844,
		0x2845: c.op2845,
		0x2846: c.op2846,
		0x2847: c.op2847,
		0x2848: c.op2848,
		0x2849: c.op2849,
		0x284A: c.op284A,
		0x284B: c.op284B,
		0x284C: c.op284C,
		0x284D: c.op284D,
		0x284E: c.op284E,
		0x284F: c.op284F,
		0x2850: c.op2850,
		0x2851: c.op2851,
		0x2852: c.op2852,
		0x2853: c.op2853,
		0x2854: c.op2854,
		0x2855: c.op2855,
		0x2856: c.op2856,
		0x2857: c.op2857,
		0x2878: c.op2878,
		0x2879: c.op2879,
		0x287C: c.op287C,
		0x2880: c.op2880,
		0x2881: c.op2881,
		0x2882: c.op2882,
		0x2883: c.op2883,
		0x2884: c.op2884,
		0x2885: c.op2885,
		0x2886: c.op2886,
		0x2887: c.op2887,
		0x2888: c.op2888,
		0x2889: c.op2889,
		0x288A: c.op288A,
		0x288B: c.op288B,
		0x288C: c.op288C,
		0x288D: c.op288D,
		0x288E: c.op288E,
		0x288F: c.op288F,
		0x2890: c.op2890,
		0x2891: c.op2891,
		0x2892: c.op2892,
		0x2893: c.op2893,
		0x2894: c.op2894,
		0x2895: c.op2895,
		0x2896: c.op2896,
		0x2897: c.op2897,
		0x28B8: c.op28B8,
		0x28B9: c.op28B9,
		0x28BC: c.op28BC,
		0x28C0: c.op28C0,
		0x28C1: c.op28C1,
		0x28C2: c.op28C2,
		0x28C3: c.op28C3,
		0x28C4: c.op28C4,
		0x28C5: c.op28C5,
		0x28C6: c.op28C6,
		0x28C7: c.op28C7,
		0x28C8: c.op28C8,
		0x28C9: c.op28C9,
		0x28CA: c.op28CA,
		0x28CB: c.op28CB,
		0x28CC: c.op28CC,
		0x28CD: c.op28CD,
		0x28CE: c.op28CE,
		0x28CF: c.op28CF,
		0x28D0: c.op28D0,
		0x28D1: c.op28D1,
		0x28D2: c.op28D2,
		0x28D3: c.op28D3,
		0x28D4: c.op28D4,
		0x28D5: c.op28D5,
		0x28D6: c.op28D6,
		0x28D7: c.op28D7,
		0x28F8: c.op28F8,
		0x28F9: c.op28F9,
		0x28FC: c.op28FC,
		0x2900: c.op2900,
		0x2901: c.op2901,
		0x2902: c.op2902,
		0x2903: c.op2903,
		0x2904: c.op2904,
		0x2905: c.op2905,
		0x2906: c.op2906,
		0x2907: c.op2907,
		0x2908: c.op2908,
		0x2909: c.op2909,
		0x290A: c.op290A,
		0x290B: c.op290B,
		0x290C: c.op290C,
		0x290D: c.op290D,
		0x290E: c.op290E,
		0x290F: c.op290F,
		0x2910: c.op2910,
		0x2911: c.op2911,
		0x2912: c.op2912,
		0x2913: c.op2913,
		0x2914: c.op2914,
		0x2915: c.op2915,
		0x2916: c.op2916,
		0x2917: c.op2917,
		0x2938: c.op2938,
		0x2939: c.op2939,
		0x293C: c.op293C,
		0x2940: c.op2940,
		0x2941: c.op2941,
		0x2942: c.op2942,
		0x2943: c.op2943,
		0x2944: c.op2944,
		0x2945: c.op2945,
		0x2946: c.op2946,
		0x2947: c.op2947,
		0x2948: c.op2948,
		0x2949: c.op2949,
		0x294A: c.op294A,
		0x294B: c.op294B,
		0x294C: c.op294C,
		0x294D: c.op294D,
		0x294E: c.op294E,
		0x294F: c.op294F,
		0x2950: c.op2950,
		0x2951: c.op2951,
		0x2952: c.op2952,
		0x2953: c.op2953,
		0x2954: c.op2954,
		0x2955: c.op2955,
		0x2956: c.op2956,
		0x2957: c.op2957,
		0x2978: c.op2978,
		0x2979: c.op2979,
		0x297C: c.op297C,
		0x2A00: c.op2A00,
		0x2A01: c.op2A01,
		0x2A02: c.op2A02,
		0x2A03: c.op2A03,
		0x2A04: c.op2A04,
		0x2A05: c.op2A05,
		0x2A06: c.op2A06,
		0x2A07: c.op2A07,
		0x2A08: c.op2A08,
		0x2A09: c.op2A09,
		0x2A0A: c.op2A0A,
		0x2A0B: c.op2A0B,
		0x2A0C: c.op2A0C,
		0x2A0D: c.op2A0D,
		0x2A0E: c.op2A0E,
		0x2A0F: c.op2A0F,
		0x2A10: c.op2A10,
		0x2A11: c.op2A11,
		0x2A12: c.op2A12,
		0x2A13: c.op2A13,
		0x2A14: c.op2A14,
		0x2A15: c.op2A15,
		0x2A16: c.op2A16,
		0x2A17: c.op2A17,
		0x2A38: c.op2A38,
		0x2A39: c.op2A39,
		0x2A3C: c.op2A3C,
		0x2A40: c.op2A40,
		0x2A41: c.op2A41,
		0x2A42: c.op2A42,
		0x2A43: c.op2A43,
		0x2A44: c.op2A44,
		0x2A45: c.op2A45,
		0x2A46: c.op2A46,
		0x2A47: c.op2A47,
		0x2A48: c.op2A48,
		0x2A49: c.op2A49,
		0x2A4A: c.op2A4A,
		0x2A4B: c.op2A4B,
		0x2A4C: c.op2A4C,
		0x2A4D: c.op2A4D,
		0x2A4E: c.op2A4E,
		0x2A4F: c.op2A4F,
		0x2A50: c.op2A50,
		0x2A51: c.op2A51,
		0x2A52: c.op2A52,
		0x2A53: c.op2A53,
		0x2A54: c.op2A54,
		0x2A55: c.op2A55,
		0x2A56: c.op2A56,
		0x2A57: c.op2A57,
		0x2A78: c.op2A78,
		0x2A79: c.op2A79,
		0x2A7C: c.op2A7C,
		0x2A80: c.op2A80,
		0x2A81: c.op2A81,
		0x2A82: c.op2A82,
		0x2A83: c.op2A83,
		0x2A84: c.op2A84,
		0x2A85: c.op2A85,
		0x2A86: c.op2A86,
		0x2A87: c.op2A87,
		0x2A88: c.op2A88,
		0x2A89: c.op2A89,
		0x2A8A: c.op2A8A,
		0x2A8B: c.op2A8B,
		0x2A8C: c.op2A8C,
		0x2A8D: c.op2A8D,
		0x2A8E: c.op2A8E,
		0x2A8F: c.op2A8F,
		0x2A90: c.op2A90,
		0x2A91: c.op2A91,
		0x2A92: c.op2A92,
		0x2A93: c.op2A93,
		0x2A94: c.op2A94,
		0x2A95: c.op2A95,
		0x2A96: c.op2A96,
		0x2A97: c.op2A97,
		0x2AB8: c.op2AB8,
		0x2AB9: c.op2AB9,
		0x2ABC: c.op2ABC,
		0x2AC0: c.op2AC0,
		0x2AC1: c.op2AC1,
		0x2AC2: c.op2AC2,
		0x2AC3: c.op2AC3,
		0x2AC4: c.op2AC4,
		0x2AC5: c.op2AC5,
		0x2AC6: c.op2AC6,
		0x2AC7: c.op2AC7,
		0x2AC8: c.op2AC8,
		0x2AC9: c.op2AC9,
		0x2ACA: c.op2ACA,
		0x2ACB: c.op2ACB,
		0x2ACC: c.op2ACC,
		0x2ACD: c.op2ACD,
		0x2ACE: c.op2ACE,
		0x2ACF: c.op2ACF,
		0x2AD0: c.op2AD0,
		0x2AD1: c.op2AD1,
		0x2AD2: c.op2AD2,
		0x2AD3: c.op2AD3,
		0x2AD4: c.op2AD4,
		0x2AD5: c.op2AD5,
		0x2AD6: c.op2AD6,
		0x2AD7: c.op2AD7,
		0x2AF8: c.op2AF8,
		0x2AF9: c.op2AF9,
		0x2AFC: c.op2AFC,
		0x2B00: c.op2B00,
		0x2B01: c.op2B01,
		0x2B02: c.op2B02,
		0x2B03: c.op2B03,
		0x2B04: c.op2B04,
		0x2B05: c.op2B05,
		0x2B06: c.op2B06,
		0x2B07: c.op2B07,
		0x2B08: c.op2B08,
		0x2B09: c.op2B09,
		0x2B0A: c.op2B0A,
		0x2B0B: c.op2B0B,
		0x2B0C: c.op2B0C,
		0x2B0D: c.op2B0D,
		0x2B0E: c.op2B0E,
		0x2B0F: c.op2B0F,
		0x2B10: c.op2B10,
		0x2B11: c.op2B11,
		0x2B12: c.op2B12,
		0x2B13: c.op2B13,
		0x2B14: c.op2B14,
		0x2B15: c.op2B15,
		0x2B16: c.op2B16,
		0x2B17: c.op2B17,
		0x2B38: c.op2B38,
		0x2B39: c.op2B39,
		0x2B3C: c.op2B3C,
		0x2B40: c.op2B40,
		0x2B41: c.op2B41,
		0x2B42: c.op2B42,
		0x2B43: c.op2B43,
		0x2B44: c.op2B44,
		0x2B45: c.op2B45,
		0x2B46: c.op2B46,
		0x2B47: c.op2B47,
		0x2B48: c.op2B48,
		0x2B49: c.op2B49,
		0x2B4A: c.op2B4A,
		0x2B4B: c.op2B4B,
		0x2B4C: c.op2B4C,
		0x2B4D: c.op2B4D,
		0x2B4E: c.op2B4E,
		0x2B4F: c.op2B4F,
		0x2B50: c.op2B50,
		0x2B51: c.op2B51,
		0x2B52: c.op2B52,
		0x2B53: c.op2B53,
		0x2B54: c.op2B54,
		0x2B55: c.op2B55,
		0x2B56: c.op2B56,
		0x2B57: c.op2B57,
		0x2B78: c.op2B78,
		0x2B79: c.op2B79,
		0x2B7C: c.op2B7C,
		0x2C00: c.op2C00,
		0x2C01: c.op2C01,
		0x2C02: c.op2C02,
		0x2C03: c.op2C03,
		0x2C04: c.op2C04,
		0x2C05: c.op2C05,
		0x2C06: c.op2C06,
		0x2C07: c.op2C07,
		0x2C08: c.op2C08,
		0x2C09: c.op2C09,
		0x2C0A: c.op2C0A,
		0x2C0B: c.op2C0B,
		0x2C0C: c.op2C0C,
		0x2C0D: c.op2C0D,
		0x2C0E: c.op2C0E,
		0x2C0F: c.op2C0F,
		0x2C10: c.op2C10,
		0x2C11: c.op2C11,
		0x2C12: c.op2C12,
		0x2C13: c.op2C13,
		0x2C14: c.op2C14,
		0x2C15: c.op2C15,
		0x2C16: c.op2C16,
		0x2C17: c.op2C17,
		0x2C38: c.op2C38,
		0x2C39: c.op2C39,
		0x2C3C: c.op2C3C,
		0x2C40: c.op2C40,
		0x2C41: c.op2C41,
		0x2C42: c.op2C42,
		0x2C43: c.op2C43,
		0x2C44: c.op2C44,
		0x2C45: c.op2C45,
		0x2C46: c.op2C46,
		0x2C47: c.op2C47,
		0x2C48: c.op2C48,
		0x2C49: c.op2C49,
		0x2C4A: c.op2C4A,
		0x2C4B: c.op2C4B,
		0x2C4C: c.op2C4C,
		0x2C4D: c.op2C4D,
		0x2C4E: c.op2C4E,
		0x2C4F: c.op2C4F,
		0x2C50: c.op2C50,
		0x2C51: c.op2C51,
		0x2C52: c.op2C52,
		0x2C53: c.op2C53,
		0x2C54: c.op2C54,
		0x2C55: c.op2C55,
		0x2C56: c.op2C56,
		0x2C57: c.op2C57,
		0x2C78: c.op2C78,
		0x2C79: c.op2C79,
		0x2C7C: c.op2C7C,
		0x2C80: c.op2C80,
		0x2C81: c.op2C81,
		0x2C82: c.op2C82,
		0x2C83: c.op2C83,
		0x2C84: c.op2C84,
		0x2C85: c.op2C85,
		0x2C86: c.op2C86,
		0x2C87: c.op2C87,
		0x2C88: c.op2C88,
		0x2C89: c.op2C89,
		0x2C8A: c.op2C8A,
		0x2C8B: c.op2C8B,
		0x2C8C: c.op2C8C,
		0x2C8D: c.op2C8D,
		0x2C8E: c.op2C8E,
		0x2C8F: c.op2C8F,
		0x2C90: c.op2C90,
		0x2C91: c.op2C91,
		0x2C92: c.op2C92,
		0x2C93: c.op2C93,
		0x2C94: c.op2C94,
		0x2C95: c.op2C95,
		0x2C96: c.op2C96,
		0x2C97: c.op2C97,
		0x2CB8: c.op2CB8,
		0x2CB9: c.op2CB9,
		0x2CBC: c.op2CBC,
		0x2CC0: c.op2CC0,
		0x2CC1: c.op2CC1,
		0x2CC2: c.op2CC2,
		0x2CC3: c.op2CC3,
		0x2CC4: c.op2CC4,
		0x2CC5: c.op2CC5,
		0x2CC6: c.op2CC6,
		0x2CC7: c.op2CC7,
		0x2CC8: c.op2CC8,
		0x2CC9: c.op2CC9,
		0x2CCA: c.op2CCA,
		0x2CCB: c.op2CCB,
		0x2CCC: c.op2CCC,
		0x2CCD: c.op2CCD,
		0x2CCE: c.op2CCE,
		0x2CCF: c.op2CCF,
		0x2CD0: c.op2CD0,
		0x2CD1: c.op2CD1,
		0x2CD2: c.op2CD2,
		0x2CD3: c.op2CD3,
		0x2CD4: c.op2CD4,
		0x2CD5: c.op2CD5,
		0x2CD6: c.op2CD6,
		0x2CD7: c.op2CD7,
		0x2CF8: c.op2CF8,
		0x2CF9: c.op2CF9,
		0x2CFC: c.op2CFC,
		0x2D00: c.op2D00,
		0x2D01: c.op2D01,
		0x2D02: c.op2D02,
		0x2D03: c.op2D03,
		0x2D04: c.op2D04,
		0x2D05: c.op2D05,
		0x2D06: c.op2D06,
		0x2D07: c.op2D07,
		0x2D08: c.op2D08,
		0x2D09: c.op2D09,
		0x2D0A: c.op2D0A,
		0x2D0B: c.op2D0B,
		0x2D0C: c.op2D0C,
		0x2D0D: c.op2D0D,
		0x2D0E: c.op2D0E,
		0x2D0F: c.op2D0F,
		0x2D10: c.op2D10,
		0x2D11: c.op2D11,
		0x2D12: c.op2D12,
		0x2D13: c.op2D13,
		0x2D14: c.op2D14,
		0x2D15: c.op2D15,
		0x2D16: c.op2D16,
		0x2D17: c.op2D17,
		0x2D38: c.op2D38,
		0x2D39: c.op2D39,
		0x2D3C: c.op2D3C,
		0x2D40: c.op2D40,
		0x2D41: c.op2D41,
		0x2D42: c.op2D42,
		0x2D43: c.op2D43,
		0x2D44: c.op2D44,
		0x2D45: c.op2D45,
		0x2D46: c.op2D46,
		0x2D47: c.op2D47,
		0x2D48: c.op2D48,
		0x2D49: c.op2D49,
		0x2D4A: c.op2D4A,
		0x2D4B: c.op2D4B,
		0x2D4C: c.op2D4C,
		0x2D4D: c.op2D4D,
		0x2D4E: c.op2D4E,
		0x2D4F: c.op2D4F,
		0x2D50: c.op2D50,
		0x2D51: c.op2D51,
		0x2D52: c.op2D52,
		0x2D53: c.op2D53,
		0x2D54: c.op2D54,
		0x2D55: c.op2D55,
		0x2D56: c.op2D56,
		0x2D57: c.op2D57,
		0x2D78: c.op2D78,
		0x2D79: c.op2D79,
		0x2D7C: c.op2D7C,
		0x2E00: c.op2E00,
		0x2E01: c.op2E01,
		0x2E02: c.op2E02,
		0x2E03: c.op2E03,
		0x2E04: c.op2E04,
		0x2E05: c.op2E05,
		0x2E06: c.op2E06,
		0x2E07: c.op2E07,
		0x2E08: c.op2E08,
		0x2E09: c.op2E09,
		0x2E0A: c.op2E0A,
		0x2E0B: c.op2E0B,
		0x2E0C: c.op2E0C,
		0x2E0D: c.op2E0D,
		0x2E0E: c.op2E0E,
		0x2E0F: c.op2E0F,
		0x2E10: c.op2E10,
		0x2E11: c.op2E11,
		0x2E12: c.op2E12,
		0x2E13: c.op2E13,
		0x2E14: c.op2E14,
		0x2E15: c.op2E15,
		0x2E16: c.op2E16,
		0x2E17: c.op2E17,
		0x2E38: c.op2E38,
		0x2E39: c.op2E39,
		0x2E3C: c.op2E3C,
		0x2E40: c.op2E40,
		0x2E41: c.op2E41,
		0x2E42: c.op2E42,
		0x2E43: c.op2E43,
		0x2E44: c.op2E44,
		0x2E45: c.op2E45,
		0x2E46: c.op2E46,
		0x2E47: c.op2E47,
		0x2E48: c.op2E48,
		0x2E49: c.op2E49,
		0x2E4A: c.op2E4A,
		0x2E4B: c.op2E4B,
		0x2E4C: c.op2E4C,
		0x2E4D: c.op2E4D,
		0x2E4E: c.op2E4E,
		0x2E4F: c.op2E4F,
		0x2E50: c.op2E50,
		0x2E51: c.op2E51,
		0x2E52: c.op2E52,
		0x2E53: c.op2E53,
		0x2E54: c.op2E54,
		0x2E55: c.op2E55,
		0x2E56: c.op2E56,
		0x2E57: c.op2E57,
		0x2E78: c.op2E78,
		0x2E79: c.op2E79,
		0x2E7C: c.op2E7C,
		0x2E80: c.op2E80,
		0x2E81: c.op2E81,
		0x2E82: c.op2E82,
		0x2E83: c.op2E83,
		0x2E84: c.op2E84,
		0x2E85: c.op2E85,
		0x2E86: c.op2E86,
		0x2E87: c.op2E87,
		0x2E88: c.op2E88,
		0x2E89: c.op2E89,
		0x2E8A: c.op2E8A,
		0x2E8B: c.op2E8B,
		0x2E8C: c.op2E8C,
		0x2E8D: c.op2E8D,
		0x2E8E: c.op2E8E,
		0x2E8F: c.op2E8F,
		0x2E90: c.op2E90,
		0x2E91: c.op2E91,
		0x2E92: c.op2E92,
		0x2E93: c.op2E93,
		0x2E94: c.op2E94,
		0x2E95: c.op2E95,
		0x2E96: c.op2E96,
		0x2E97: c.op2E97,
		0x2EB8: c.op2EB8,
		0x2EB9: c.op2EB9,
		0x2EBC: c.op2EBC,
		0x2EC0: c.op2EC0,
		0x2EC1: c.op2EC1,
		0x2EC2: c.op2EC2,
		0x2EC3: c.op2EC3,
		0x2EC4: c.op2EC4,
		0x2EC5: c.op2EC5,
		0x2EC6: c.op2EC6,
		0x2EC7: c.op2EC7,
		0x2EC8: c.op2EC8,
		0x2EC9: c.op2EC9,
		0x2ECA: c.op2ECA,
		0x2ECB: c.op2ECB,
		0x2ECC: c.op2ECC,
		0x2ECD: c.op2ECD,
		0x2ECE: c.op2ECE,
		0x2ECF: c.op2ECF,
		0x2ED0: c.op2ED0,
		0x2ED1: c.op2ED1,
		0x2ED2: c.op2ED2,
		0x2ED3: c.op2ED3,
		0x2ED4: c.op2ED4,
		0x2ED5: c.op2ED5,
		0x2ED6: c.op2ED6,
		0x2ED7: c.op2ED7,
		0x2EF8: c.op2EF8,
		0x2EF9: c.op2EF9,
		0x2EFC: c.op2EFC,
		0x2F00: c.op2F00,
		0x2F01: c.op2F01,
		0x2F02: c.op2F02,
		0x2F03: c.op2F03,
		0x2F04: c.op2F04,
		0x2F05: c.op2F05,
		0x2F06: c.op2F06,
		0x2F07: c.op2F07,
		0x2F08: c.op2F08,
		0x2F09: c.op2F09,
		0x2F0A: c.op2F0A,
		0x2F0B: c.op2F0B,
		0x2F0C: c.op2F0C,
		0x2F0D: c.op2F0D,
		0x2F0E: c.op2F0E,
		0x2F0F: c.op2F0F,
		0x2F10: c.op2F10,
		0x2F11: c.op2F11,
		0x2F12: c.op2F12,
		0x2F13: c.op2F13,
		0x2F14: c.op2F14,
		0x2F15: c.op2F15,
		0x2F16: c.op2F16,
		0x2F17: c.op2F17,
		0x2F38: c.op2F38,
		0x2F39: c.op2F39,
		0x2F3C: c.op2F3C,
		0x2F40: c.op2F40,
		0x2F41: c.op2F41,
		0x2F42: c.op2F42,
		0x2F43: c.op2F43,
		0x2F44: c.op2F44,
		0x2F45: c.op2F45,
		0x2F46: c.op2F46,
		0x2F47: c.op2F47,
		0x2F48: c.op2F48,
		0x2F49: c.op2F49,
		0x2F4A: c.op2F4A,
		0x2F4B: c.op2F4B,
		0x2F4C: c.op2F4C,
		0x2F4D: c.op2F4D,
		0x2F4E: c.op2F4E,
		0x2F4F: c.op2F4F,
		0x2F50: c.op2F50,
		0x2F51: c.op2F51,
		0x2F52: c.op2F52,
		0x2F53: c.op2F53,
		0x2F54: c.op2F54,
		0x2F55: c.op2F55,
		0x2F56: c.op2F56,
		0x2F57: c.op2F57,
		0x2F78: c.op2F78,
		0x2F79: c.op2F79,
		0x2F7C: c.op2F7C,
	}
	return table[op]
}
