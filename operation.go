package goflag

type FlagOp interface {
	// Check if index flag is on
	Check(i int) (bool, error)
	// Force up the index flag
	Up(i int) error
	// Force down the index flag
	Down(i int) error
	// Find down flag and then get index and up the flag
	FindAndUp() (int, error)
}
