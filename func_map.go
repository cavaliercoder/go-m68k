package m68k

type funcMap [][]stepFunc

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

var defaultFuncMap = funcMap{
	[]stepFunc{ // 0x0n
		opOri, nil, opAndi, nil, nil, nil, nil, nil,
		nil, nil, nil, nil, nil, nil, nil, nil,
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
}
