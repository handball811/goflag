package goflag_test

import (
	"testing"

	. "github.com/handball811/goflag"
	"github.com/stretchr/testify/assert"
)

func TestGoFlagCheck(t *testing.T) {
	// setup
	var ok bool
	var err error
	target, _ := NewFlagWithSize(SIZE_4096)

	// when
	ok, err = target.Check(0)

	// then
	assert.Equal(t, false, ok)
	assert.Nil(t, err)
}

func TestGoFlagCheckAndUp(t *testing.T) {
	// setup
	var ok bool
	var err error
	target, _ := NewFlagWithSize(SIZE_4096)

	// when
	target.Up(24)
	ok, err = target.Check(24)

	// then
	assert.Equal(t, true, ok)
	assert.Nil(t, err)
}

func TestGoFlagCheckAndUpAndDown(t *testing.T) {
	// setup
	var ok bool
	var err error
	target, _ := NewFlagWithSize(SIZE_4096)

	// when
	target.Up(24)
	target.Down(24)
	ok, err = target.Check(0)

	// then
	assert.Equal(t, false, ok)
	assert.Nil(t, err)
}

func TestGoFlagCheckError(t *testing.T) {
	// setup
	var ok bool
	var err error
	target, _ := NewFlagWithSize(SIZE_4096)

	// when
	ok, err = target.Check(-1)

	// then
	assert.Equal(t, false, ok)
	assert.Equal(t, ErrOutOfBound, err)
}

func TestGoFlagCheckError2(t *testing.T) {
	// setup
	var ok bool
	var err error
	target, _ := NewFlagWithSize(SIZE_4096)

	// when
	ok, err = target.Check(4096)

	// then
	assert.Equal(t, false, ok)
	assert.Equal(t, ErrOutOfBound, err)
}

func TestGoFlagUpError(t *testing.T) {
	// setup
	var err error
	target, _ := NewFlagWithSize(SIZE_4096)

	// when
	err = target.Up(4096)

	// then
	assert.Equal(t, ErrOutOfBound, err)
}

func TestGoFlagUpError2(t *testing.T) {
	// setup
	var err error
	target, _ := NewFlagWithSize(SIZE_4096)

	// when
	err = target.Up(24)

	// then
	assert.Nil(t, err)
}

func TestGoFlagUpError3(t *testing.T) {
	// setup
	var err error
	target, _ := NewFlagWithSize(SIZE_4096)

	// when
	err = target.Up(-1)

	// then
	assert.Equal(t, ErrOutOfBound, err)
}

func TestGoFlagFindAndUp(t *testing.T) {
	// setup
	var index int
	var ok bool
	var err error
	target, _ := NewFlagWithSize(SIZE_4096)

	// when
	index, err = target.FindAndUp()
	ok, _ = target.Check(index)

	// then
	assert.Equal(t, true, index >= 0)
	assert.Nil(t, err)
	assert.Equal(t, true, ok)
}

func TestGoFlagFindAndUp2(t *testing.T) {
	// setup
	var index int
	var err error
	target, _ := NewFlagWithSize(SIZE_4096)

	// when
	for i := 0; i < 4096; i++ {
		target.FindAndUp()
	}
	index, err = target.FindAndUp()

	// then
	assert.Equal(t, -1, index)
	assert.Equal(t, ErrNoSpace, err)
}

func TestGoFlagFindAndUp3(t *testing.T) {
	// setup
	var index int = 1024
	var err error
	target, _ := NewFlagWithSize(SIZE_4096)

	// when
	for i := 0; i < 4096; i++ {
		target.FindAndUp()
	}
	target.Down(index)
	index, err = target.FindAndUp()

	// then
	assert.Equal(t, 1024, index)
	assert.Nil(t, err)
}

func TestFlagsFindAndUp(t *testing.T) {
	// setup
	var index int = 1024
	var err error
	target, _ := NewFlags(100)

	// when
	for i := 0; i < 99; i++ {
		target.FindAndUp()
	}
	index, err = target.FindAndUp()

	//then
	assert.Equal(t, 99, index)
	assert.Nil(t, err)
}

func TestFlagsFindAndUp2(t *testing.T) {
	// setup
	var index int = 1024
	var err error
	target, _ := NewFlags(100)

	// when
	for i := 0; i < 101; i++ {
		target.FindAndUp()
	}
	index, err = target.FindAndUp()

	//then
	assert.Equal(t, -1, index)
	assert.Equal(t, ErrNoSpace, err)
}

//---------------------------------------------- Benchmark -----------------------------------

func BenchmarkGoFlagFindAndUp4096_0(b *testing.B) {
	target, _ := NewFlagWithSize(SIZE_4096)

	//fill
	for i := 0; i < 0; i++ {
		target.FindAndUp()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index, err := target.FindAndUp()
		if err == nil {
			target.Down(index)
		}
	}
}

func BenchmarkSliceFindAndUp4096_0(b *testing.B) {
	target := make([]bool, 4096)

	//fill
	for i := 0; i < 0; i++ {
		target[i] = true
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var j int = 0
		for ; j < 4096; j++ {
			if !target[j] {
				target[j] = true
				break
			}
		}
		if j < 4096 {
			target[j] = false
		}
	}
}

func BenchmarkGoFlagFindAndUp4096_100(b *testing.B) {
	target, _ := NewFlagWithSize(SIZE_4096)

	//fill
	for i := 0; i < 100; i++ {
		target.FindAndUp()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index, err := target.FindAndUp()
		if err == nil {
			target.Down(index)
		}
	}
}

func BenchmarkSliceFindAndUp4096_100(b *testing.B) {
	target := make([]bool, 4096)

	//fill
	for i := 0; i < 100; i++ {
		target[i] = true
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var j int = 0
		for ; j < 4096; j++ {
			if !target[j] {
				target[j] = true
				break
			}
		}
		if j < 4096 {
			target[j] = false
		}
	}
}

func BenchmarkGoFlagFindAndUp4096_1000(b *testing.B) {
	target, _ := NewFlagWithSize(SIZE_4096)

	//fill
	for i := 0; i < 1000; i++ {
		target.FindAndUp()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index, err := target.FindAndUp()
		if err == nil {
			target.Down(index)
		}
	}
}

func BenchmarkSliceFindAndUp4096_1000(b *testing.B) {
	target := make([]bool, 4096)

	//fill
	for i := 0; i < 1000; i++ {
		target[i] = true
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var j int = 0
		for ; j < 4096; j++ {
			if !target[j] {
				target[j] = true
				break
			}
		}
		if j < 4096 {
			target[j] = false
		}
	}
}

func BenchmarkGoFlagFindAndUp4096_4096(b *testing.B) {
	target, _ := NewFlagWithSize(SIZE_4096)

	//fill
	for i := 0; i < 4096; i++ {
		target.FindAndUp()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index, err := target.FindAndUp()
		if err == nil {
			target.Down(index)
		}
	}
}

func BenchmarkSliceFindAndUp4096_4096(b *testing.B) {
	target := make([]bool, 4096)

	//fill
	for i := 0; i < 4096; i++ {
		target[i] = true
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		var j int = 0
		for ; j < 4096; j++ {
			if !target[j] {
				target[j] = true
				break
			}
		}
		if j < 4096 {
			target[j] = false
		}
	}
}

func BenchmarkGoFlagFindAndUp262144(b *testing.B) {
	target, _ := NewFlagWithSize(SIZE_262144)

	//fill
	for i := 0; i < 131072; i++ {
		target.FindAndUp()
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index, err := target.FindAndUp()
		if err == nil {
			target.Down(index)
		}
	}
}

func BenchmarkGoFlagFindAndUp64(b *testing.B) {
	target, _ := NewFlagWithSize(SIZE_64)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		index, err := target.FindAndUp()
		if err == nil {
			target.Down(index)
		}
	}
}
