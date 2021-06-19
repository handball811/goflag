package goflag

type Size uint

const (
	SIZE_64     Size = 0
	SIZE_4096        = 1
	SIZE_262144      = 2
	SIZE_BORDER      = 3

	FLAG_FULL = 0xffffffffffffffff
	FLAG_63   = 0x3f
)

var (
	depthMetrix = [...]int{
		0,
		6,
		12,
	}

	flagMetrix = [...]uint64{
		0x00000000ffffffff,
		0x000000000000ffff,
		0x00000000000000ff,
		0x000000000000000f,
		0x0000000000000003,
		0x0000000000000001,
	}
)
