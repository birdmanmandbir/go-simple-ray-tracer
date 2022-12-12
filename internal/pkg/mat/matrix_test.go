package mat

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewMatrix2x2(t *testing.T) {
	m := NewMatrix(2)
	m.Data = [][]float64{
		{-3, 5},
		{1, -2},
	}
	assert.Equal(t, -3.0, m.Get(0, 0))
	assert.Equal(t, 5.0, m.Get(0, 1))
	assert.Equal(t, 1.0, m.Get(1, 0))
	assert.Equal(t, -2.0, m.Get(1, 1))
}

func TestNewMatrix3x3(t *testing.T) {
	m := NewMatrix(3)
	m.Data = [][]float64{
		{-3, 5, 0},
		{1, -2, -7},
		{0, 1, 1},
	}
	assert.Equal(t, -3.0, m.Get(0, 0))
	assert.Equal(t, -2.0, m.Get(1, 1))
	assert.Equal(t, 1.0, m.Get(2, 2))
}

func TestNewMatrix4x4(t *testing.T) {
	m := NewMatrix(4)
	m.Data = [][]float64{
		{1, 2, 3, 4},
		{5.5, 6.5, 7.5, 8.5},
		{9, 10, 11, 12},
		{13.5, 14.5, 15.5, 16.5},
	}
	assert.Equal(t, 1.0, m.Get(0, 0))
	assert.Equal(t, 4.0, m.Get(0, 3))
	assert.Equal(t, 5.5, m.Get(1, 0))
	assert.Equal(t, 7.5, m.Get(1, 2))
	assert.Equal(t, 11.0, m.Get(2, 2))
	assert.Equal(t, 13.5, m.Get(3, 0))
	assert.Equal(t, 15.5, m.Get(3, 2))
}

func TestMatrix_Equals(t *testing.T) {
	type fields struct {
		W    int
		H    int
		Data [][]float64
	}
	type args struct {
		o *Matrix
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{name: "eq", fields: fields{W: 4, H: 4, Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 8, 7, 6},
			{5, 4, 3, 2},
		}}, args: args{o: &Matrix{W: 4, H: 4, Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 8, 7, 6},
			{5, 4, 3, 2},
		}}}, want: true},
		{name: "neq", fields: fields{W: 4, H: 4, Data: [][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 8, 7, 6},
			{5, 4, 3, 2},
		}}, args: args{o: &Matrix{W: 4, H: 4, Data: [][]float64{
			{2, 3, 4, 5},
			{6, 7, 8, 9},
			{8, 7, 6, 5},
			{4, 3, 2, 1},
		}}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Matrix{
				W:    tt.fields.W,
				H:    tt.fields.H,
				Data: tt.fields.Data,
			}
			assert.Equalf(t, tt.want, m.Equals(tt.args.o), "Equals(%v)", tt.args.o)
		})
	}
}