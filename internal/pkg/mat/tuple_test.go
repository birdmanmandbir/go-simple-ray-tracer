package mat

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPoint(t *testing.T) {
	v := NewPoint(4.3, -4.2, 3.1)
	assert.True(t, v.IsPoint())
}

func TestIsVector(t *testing.T) {
	v := NewVector(4.3, -4.2, 3.1)
	assert.True(t, v.IsVector())
}

func TestAdd(t *testing.T) {
	t1 := NewPoint(3, -2, 5)
	t2 := NewVector(-2, 3, 1)
	t3 := Add(t1, t2)
	assert.Equal(t, 1.0, t3.Data[0])
	assert.Equal(t, 1.0, t3.Data[1])
	assert.Equal(t, 6.0, t3.Data[2])
	assert.Equal(t, 1.0, t3.Data[3])
}

func TestSub(t *testing.T) {
	type args struct {
		t1 *Tuple4
		t2 *Tuple4
	}
	tests := []struct {
		name string
		args args
		want *Tuple4
	}{
		{name: "P-P", args: args{t1: NewPoint(3, 2, 1), t2: NewPoint(5, 6, 7)}, want: NewVector(-2, -4, -6)},
		{name: "P-V", args: args{t1: NewPoint(3, 2, 1), t2: NewVector(5, 6, 7)}, want: NewPoint(-2, -4, -6)},
		{name: "V-V", args: args{t1: NewVector(3, 2, 1), t2: NewVector(5, 6, 7)}, want: NewVector(-2, -4, -6)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Sub(tt.args.t1, tt.args.t2), "Sub(%v, %v)", tt.args.t1, tt.args.t2)
		})
	}
}

func TestNegate(t *testing.T) {
	expect := NewTuple4([]float64{-1, 2, -3, 4})
	input := Negate(NewTuple4([]float64{1, -2, 3, -4}))
	assert.Equal(t, expect, input)
}
