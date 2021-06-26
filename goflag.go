package goflag

type Flag struct {
	flags []uint64
	depth uint
	size  int
}

func NewFlag() *Flag {
	flag, _ := NewFlagWithSize(SIZE_262144)
	return flag
}

func NewFlagWithSize(
	flagSize Size,
) (*Flag, error) {
	fs := uint(flagSize)
	if fs >= uint(SIZE_BORDER) {
		return nil, ErrFlagSize
	}
	// ready for data
	flag := &Flag{
		depth: fs,
		flags: make([]uint64, (1<<(fs*6+6)-1)/63),
		size:  1 << (fs*6 + 6),
	}

	return flag, nil
}

func NewFlags(
	size uint,
) (*Flag, error) {
	var i uint
	var sz uint
	for i = 0; i < SIZE_BORDER; i++ {
		sz = 1 << (i*6 + 6)
		if size <= sz {
			break
		}
	}
	if i >= SIZE_BORDER {
		return nil, ErrFlagSize
	}
	flag, err := NewFlagWithSize(Size(i))
	if err != nil {
		return nil, err
	}
	for k := size; k < sz; k++ {
		flag.Up(int(k))
	}
	flag.size = int(size)
	return flag, nil
}

func (f *Flag) Check(i int) (bool, error) {
	if i < 0 || i >= int(f.size) {
		return false, ErrOutOfBound
	}
	return f.check(i, f.depth, 0), nil
}

// iは該当の高さにおける位置
func (f *Flag) check(i int, d uint, index int) bool {
	dm := depthMetrix[d]
	bit := i >> dm
	if (f.flags[index]>>bit)&1 == 1 {
		return true
	}
	// これ位以上の深さに行かないのでフラグはたっていないことを意味する
	if d == 0 {
		return false
	}
	return f.check(i-(bit<<dm), d-1, (index<<6)+bit+1)
}

func (f *Flag) Up(i int) error {
	if i < 0 || i >= int(f.size) {
		return ErrOutOfBound
	}
	f.up(i, f.depth, 0)
	return nil
}

func (f *Flag) up(i int, d uint, index int) {
	dm := depthMetrix[d]
	bit := i >> dm
	if d == 0 {
		f.flags[index] |= 1 << bit
		f.update(index)
	} else {
		f.up(i-(bit<<dm), d-1, (index<<6)+bit+1)
	}
}

func (f *Flag) update(index int) {
	for index > 0 {
		on := f.flags[index] == FLAG_FULL
		index--
		var flag uint64 = 1 << (index & FLAG_63)
		index >>= 6
		if on == (f.flags[index]&flag > 0) {
			break
		}
		if on {
			f.flags[index] |= flag
		} else {
			f.flags[index] &= ^flag
		}
	}
}

func (f *Flag) Down(i int) error {
	if i < 0 || i >= int(f.size) {
		return ErrOutOfBound
	}
	var index int = (i >> 6) + ((1<<(6*int(f.depth)))-1)/63
	f.flags[index] &= ^(1 << (i & 63))
	f.update(index)
	return nil
}

func (f *Flag) FindAndUp() (int, error) {
	// 空であることを前提とする
	var i int = 0
	var low = -1
	var index int = 0
	for d := int(f.depth); d >= 0; d-- {
		index <<= 6
		index += low
		index++
		t := f.flags[index]
		if t == FLAG_FULL {
			return -1, ErrNoSpace
		}
		low = int(Find(t))
		i <<= 6
		i += low
	}
	f.flags[index] |= 1 << low
	f.update(index)
	return i, nil
}
