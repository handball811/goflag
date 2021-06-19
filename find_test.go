package goflag_test

import (
	"testing"

	. "github.com/handball811/goflag"
	"github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
	// when
	ch := int(Find(0xfffffffffffffffe))

	//then
	assert.Equal(t, 0, ch)
}

func TestFind2(t *testing.T) {
	// when
	ch := int(Find(0xfffffffffffff0ff))

	//then
	assert.Equal(t, 8, ch)
}

func TestFind3(t *testing.T) {
	// when
	ch := int(Find(0xffffff0fff0fffff))

	//then
	assert.Equal(t, 20, ch)
}

func TestFind4(t *testing.T) {
	// when
	ch := int(Find(0xfffffff0ffffffff))

	//then
	assert.Equal(t, 32, ch)
}
