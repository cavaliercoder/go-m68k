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

var opcodes = []opcode{
	{0x0000, 0xFF00, opOri},
	{0x0200, 0xFF00, opAndi},
	{0x0400, 0xFF00, opSubi},
	{0x0600, 0xFF00, opAddi},
	{0x0A00, 0xFF00, opEori},
	{0x1000, 0xF000, opMove},
	{0x2000, 0xF000, opMove},
	{0x3000, 0xF000, opMove},
	{0x41C0, 0xF1C0, opLea},
	{0x4880, 0xFB80, opMovem},
	{0x4A00, 0xFF00, opTst},
	{0x4E40, 0xFFF0, opTrap},
	{0x4E60, 0xFFF0, opMoveUsp},
	{0x4E72, 0xFFFF, opStop},
	{0x4E80, 0xFFC0, opJsr},
	{0x50C8, 0xF0F8, opDBcc},
	{0x6000, 0xF000, opBcc},
	{0x7000, 0xF100, opMoveq},
	{0xC100, 0xF1F0, opAbcd},
	{0xD000, 0xF000, opAdd},
}

// opcodeIx records the index in opcodes of all opcodes starting with each
// possible nibble.
var opcodeIx = []int{
	0,  // 0x0...
	5,  // 0x1...
	6,  // 0x2...
	7,  // 0x3...
	8,  // 0x4...
	15, // 0x5...
	16, // 0x6...
	17, // 0x7...
	17, // 0x8...
	17, // 0x9...
	17, // 0xA...
	17, // 0xB...
	17, // 0xC...
	18, // 0xD...
	18, // 0xE...
	18, // 0xF...
}

func dispatch(op uint16) stepFunc {
	for i := opcodeIx[op>>12]; i < len(opcodes); i++ {
		if op&opcodes[i].mask == opcodes[i].base {
			return opcodes[i].f
		}
	}
	return nil
}
