package m68k

// A stepFunc is a function that executes a single processor instruction and
// returns trace information. Any errors should be returned by Processor.err.
type stepFunc func(*Processor) *stepTrace

// A funcMap maps stepFuncs to operation codes using the first two nibbles of
// the opcode.
type funcMap [][]stepFunc

// Resolve returns the stepFunc associated with the given opcode or nil if none
// exists.
func (f funcMap) Resolve(n uint16) stepFunc {
	i := int(n >> 12)
	if i < len(f) && f[i] != nil {
		j := int(n >> 8 & 0x000F)
		if j < len(f[i]) && f[i][j] != nil {
			return f[i][j]
		}
	}
	return nil
}

// defaultFuncMap maps opcodes to all implemented stepFuncs.
var defaultFuncMap = funcMap{
	[]stepFunc{ // 0x0n
		opOri, nil, opAndi, nil, opSubi, nil, opAddi, nil,
		nil, nil, opEori, nil, nil, nil, nil, nil,
	},
	[]stepFunc{ // 0x1n
		opMove, opMove, opMove, opMove, opMove, opMove, opMove, opMove,
		opMove, opMove, opMove, opMove, opMove, opMove, opMove, opMove,
	},
	[]stepFunc{ // 0x2n
		opMove, opMove, opMove, opMove, opMove, opMove, opMove, opMove,
		opMove, opMove, opMove, opMove, opMove, opMove, opMove, opMove,
	},
	[]stepFunc{ // 0x3n
		opMove, opMove, opMove, opMove, opMove, opMove, opMove, opMove,
		opMove, opMove, opMove, opMove, opMove, opMove, opMove, opMove,
	},
	[]stepFunc{ // 0x4n
		nil, op41, nil, op41, nil, op41, nil, op41,
		opMovem, op41, op4A, op41, opMovem, op41, op4E, op41,
	},
	[]stepFunc{ // 0x5n
		nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
	},
	[]stepFunc{ // 0x6n
		opBcc, opBcc, opBcc, opBcc, opBcc, opBcc, opBcc, opBcc,
		opBcc, opBcc, opBcc, opBcc, opBcc, opBcc, opBcc, opBcc,
	},
	[]stepFunc{ // 0x7n
		opMoveq, opMoveq, opMoveq, opMoveq, opMoveq, opMoveq, opMoveq, opMoveq,
		opMoveq, opMoveq, opMoveq, opMoveq, opMoveq, opMoveq, opMoveq, opMoveq,
	},
	[]stepFunc{ // 0x8n
		nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
	},
	[]stepFunc{ // 0x9n
		nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
	},
	[]stepFunc{ // 0xAn
		nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
	},
	[]stepFunc{ // 0xBn
		nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
	},
	[]stepFunc{ // 0xCn
		opC0, opC0, opC0, opC0, opC0, opC0, opC0, opC0,
		opC0, opC0, opC0, opC0, opC0, opC0, opC0, opC0,
	},
	[]stepFunc{ // 0xDn
		opAdd, opAdd, opAdd, opAdd, opAdd, opAdd, opAdd, opAdd,
		opAdd, opAdd, opAdd, opAdd, opAdd, opAdd, opAdd, opAdd,
	},
	[]stepFunc{ // 0xEn
		nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
	},
	[]stepFunc{ // 0xFn
		nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
	},
}
