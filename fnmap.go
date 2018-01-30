package m68k

var fnmap = [][]stepFunc{
	[]stepFunc{ // 0x00
		nil, nil, nil, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
	},
	[]stepFunc{ // 0x01
		opMove, opMove, opMove, opMove, opMove, opMove, opMove, opMove,
		opMove, opMove, opMove, opMove, opMove, opMove, opMove, opMove,
	},
	[]stepFunc{ // 0x02
		opMove, opMove, opMove, opMove, opMove, opMove, opMove, opMove,
		opMove, opMove, opMove, opMove, opMove, opMove, opMove, opMove,
	},
	[]stepFunc{ // 0x03
		opMove, opMove, opMove, opMove, opMove, opMove, opMove, opMove,
		opMove, opMove, opMove, opMove, opMove, opMove, opMove, opMove,
	},
}
