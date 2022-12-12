package mat

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddColor(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	c3 := NewColor(1.6, 0.7, 1.0)
	assert.Equal(t, c3, Add(*c1, *c2))
}

func TestSubColor(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)
	c3 := NewColor(0.2, 0.5, 0.5)
	assert.True(t, c3.Equals(*Sub(*c1, *c2)))
}

func TestColorMulByScalar(t *testing.T) {
	c := NewColor(0.2, 0.3, 0.4)
	c1 := NewColor(0.4, 0.6, 0.8)
	assert.Equal(t, c1, MultiplyByScalar(*c, 2))
}

func TestMultiplyColors(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)
	c3 := NewColor(0.9, 0.2, 0.04)
	assert.True(t, c3.Equals(*Hadamard(*c1, *c2)))
}
