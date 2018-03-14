package m68k

// A stepFunc is a function that executes a single processor instruction and
// returns trace information. Any errors should be returned by Processor.err.
type stepFunc func(*Processor) *stepTrace

// An opcode maps an opcode to a stepFunc. Opcodes are matched by applying each
// mask and comparing the result to the base.
type opcode struct {
	base, mask uint16
	f          stepFunc
}

var opcodes = [][]opcode{
	[]opcode{ /* 0x0... */
		{0x0000, 0xFF00, opOri},
		{0x0100, 0xF1C0, opBtst},
		{0x0200, 0xFF00, opAndi},
		{0x0400, 0xFF00, opSubi},
		{0x0600, 0xFF00, opAddi},
		{0x0A00, 0xFF00, opEori},
	},
	[]opcode{ /* 0x1... */
		{0x1000, 0xF000, opMove},
	},
	[]opcode{ /* 0x2... */
		{0x2000, 0xF000, opMove},
	},
	[]opcode{ /* 0x3... */

		{0x3000, 0xF000, opMove},
	},
	[]opcode{ /* 0x4... */
		{0x41C0, 0xF1C0, opLea},
		{0x4880, 0xFB80, opMovem},
		{0x4A00, 0xFF00, opTst},
		{0x4E40, 0xFFF0, opTrap},
		{0x4E60, 0xFFF0, opMoveUsp},
		{0x4E72, 0xFFFF, opStop},
		{0x4E75, 0xFFFF, opRts},
		{0x4E80, 0xFFC0, opJsr},
	},
	[]opcode{ /* 0x5... */
		{0x50C8, 0xF0F8, opDBcc},
	},
	[]opcode{ /* 0x6... */
		{0x6000, 0xF000, opBcc},
	},
	[]opcode{ /* 0x7... */
		{0x7000, 0xF100, opMoveq},
	},
	[]opcode{ /* 0x8... */ },
	[]opcode{ /* 0x9... */ },
	[]opcode{ /* 0xA... */ },
	[]opcode{ /* 0xB... */ },
	[]opcode{ /* 0xC... */
		{0xC100, 0xF1F0, opAbcd},
	},
	[]opcode{ /* 0xD... */
		{0xD000, 0xF000, opAdd},
	},
	[]opcode{ /* 0xE... */ },
	[]opcode{ /* 0xF... */ },
}

func dispatch(op uint16) stepFunc {
	o := op >> 12
	// TODO: switch to faster binary search
	for i := 0; i < len(opcodes[o]); i++ {
		if op&opcodes[o][i].mask == opcodes[o][i].base {
			return opcodes[o][i].f
		}
	}
	return nil
}
