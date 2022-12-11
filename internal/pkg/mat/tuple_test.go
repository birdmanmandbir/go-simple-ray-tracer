package mat

import (
	"github.com/stretchr/testify/assert"
	"math"
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
	t3 := Add(*t1, *t2)
	assert.Equal(t, 1.0, t3.Data[0])
	assert.Equal(t, 1.0, t3.Data[1])
	assert.Equal(t, 6.0, t3.Data[2])
	assert.Equal(t, 1.0, t3.Data[3])
}

func TestSub(t *testing.T) {
	type args struct {
		t1 Tuple4
		t2 Tuple4
	}
	tests := []struct {
		name string
		args args
		want *Tuple4
	}{
		{name: "P-P", args: args{t1: *NewPoint(3, 2, 1), t2: *NewPoint(5, 6, 7)}, want: NewVector(-2, -4, -6)},
		{name: "P-V", args: args{t1: *NewPoint(3, 2, 1), t2: *NewVector(5, 6, 7)}, want: NewPoint(-2, -4, -6)},
		{name: "V-V", args: args{t1: *NewVector(3, 2, 1), t2: *NewVector(5, 6, 7)}, want: NewVector(-2, -4, -6)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Sub(tt.args.t1, tt.args.t2), "Sub(%v, %v)", tt.args.t1, tt.args.t2)
		})
	}
}

func TestNegate(t *testing.T) {
	expect := NewTuple4(-1, 2, -3, 4)
	input := Negate(*NewTuple4(1, -2, 3, -4))
	assert.Equal(t, expect, input)
}

func TestMultiplyByScalar(t *testing.T) {
	type args struct {
		t      Tuple4
		scalar float64
	}
	tests := []struct {
		name string
		args args
		want *Tuple4
	}{
		{name: "scalar", args: args{t: *NewTuple4(1, -2, 3, -4), scalar: 3.5}, want: NewTuple4(3.5, -7, 10.5, -14)},
		{name: "fraction", args: args{t: *NewTuple4(1, -2, 3, -4), scalar: 0.5}, want: NewTuple4(0.5, -1, 1.5, -2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, MultiplyByScalar(tt.args.t, tt.args.scalar), "MultiplyByScalar(%v, %v)", tt.args.t, tt.args.scalar)
		})
	}
}

func TestDivideByScalar(t *testing.T) {
	type args struct {
		t      Tuple4
		scalar float64
	}
	tests := []struct {
		name string
		args args
		want *Tuple4
	}{
		{name: "divide", args: args{t: *NewTuple4(1, -2, 3, -4), scalar: 2}, want: NewTuple4(0.5, -1, 1.5, -2)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, DivideByScalar(tt.args.t, tt.args.scalar), "DivideByScalar(%v, %v)", tt.args.t, tt.args.scalar)
		})
	}
}

func TestMagnitude(t *testing.T) {
	type args struct {
		t Tuple4
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{args: args{t: *NewVector(1, 0, 0)}, want: 1},
		{args: args{t: *NewVector(0, 1, 0)}, want: 1},
		{args: args{t: *NewVector(0, 0, 1)}, want: 1},
		{args: args{t: *NewVector(1, 2, 3)}, want: math.Sqrt(14)},
		{args: args{t: *NewVector(-1, -2, -3)}, want: math.Sqrt(14)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Magnitude(tt.args.t), "Magnitude(%v)", tt.args.t)
		})
	}
}

func TestNormalize(t *testing.T) {
	type args struct {
		t Tuple4
	}
	tests := []struct {
		name string
		args args
		want *Tuple4
	}{
		{args: args{t: *NewVector(4, 0, 0)}, want: NewVector(1, 0, 0)},
		{args: args{t: *NewVector(1, 2, 3)}, want: NewVector(0.26726, 0.53452, 0.80178)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Truef(t, TupleEquals(*tt.want, *Normalize(tt.args.t)), "Normalize(%v)", tt.args.t)
		})
	}
}

func TestNormalizeMagnitudeIsOne(t *testing.T) {
	input := *NewVector(1, 2, 3)
	assert.Equal(t, 1.0, Magnitude(*Normalize(input)))
}

func TestDot(t *testing.T) {
	type args struct {
		t1 Tuple4
		t2 Tuple4
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{args: args{t1: *NewVector(1, 2, 3), t2: *NewVector(2, 3, 4)}, want: 20.0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Dot(tt.args.t1, tt.args.t2), "Dot(%v, %v)", tt.args.t1, tt.args.t2)
		})
	}
}

func TestCross(t *testing.T) {
	type args struct {
		t1 Tuple4
		t2 Tuple4
	}
	tests := []struct {
		name string
		args args
		want *Tuple4
	}{
		{args: args{t1: *NewVector(1, 2, 3), t2: *NewVector(2, 3, 4)}, want: NewVector(-1, 2, -1)},
		{args: args{t1: *NewVector(2, 3, 4), t2: *NewVector(1, 2, 3)}, want: NewVector(1, -2, 1)},
		{args: args{t1: *NewVector(1, 0, 0), t2: *NewVector(0, 1, 0)}, want: NewVector(0, 0, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, Cross(tt.args.t1, tt.args.t2), "Cross(%v, %v)", tt.args.t1, tt.args.t2)
		})
	}
}
