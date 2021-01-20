package java

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var hashCode = &HashCode{}

func TestString(t *testing.T) {
	hash, _ := hashCode.Eval("bf498604-9a88-419b-bf1b-7a4c1d3fa5ff")

	assert.Equal(t, 1670200295, hash)
}

func TestInt(t *testing.T) {
	hash, _ := hashCode.Eval(9999)

	assert.Equal(t, 1754688, hash)
}

func TestObject(t *testing.T) {
	hash, _ := hashCode.Eval(struct{ Hello string }{"hello"})

	assert.Equal(t, 567835174, hash)
}
