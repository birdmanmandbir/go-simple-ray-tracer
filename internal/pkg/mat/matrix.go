package mat

import "math"

type Matrix struct {
	W    int
	H    int
	Data [][]float64
}

func (m *Matrix) Get(x int, y int) float64 {
	if !m.checkVisit(x, y) {
		return math.MaxFloat64
	}
	return m.Data[x][y]
}

func (m *Matrix) Set(x int, y int, val float64) {
	if !m.checkVisit(x, y) {
		return
	}
	m.Data[x][y] = val
}

func (m *Matrix) checkVisit(x int, y int) bool {
	if x < 0 || y < 0 || x >= m.W || y >= m.H {
		println("Visit Matrix: Out of bound")
		return false
	}
	return true
}

func (m *Matrix) Equals(o *Matrix) bool {
	if m.W != o.W || m.H != o.H {
		return false
	}
	for i := 0; i < m.W; i++ {
		for j := 0; j < m.H; j++ {
			if !AlmostEqualFloat64(m.Get(i, j), o.Get(i, j)) {
				return false
			}
		}
	}
	return true
}

func NewMatrixByRaw(raw [][]float64) *Matrix {
	if len(raw) <= 0 {
		println("NewMatrix wrong raw")
		return nil
	}
	return &Matrix{Data: raw, W: len(raw), H: len(raw[0])}
}

func NewMatrix(size int) *Matrix {
	data := make([][]float64, size)
	for i := 0; i < size; i++ {
		data[i] = make([]float64, size)
	}
	return &Matrix{W: size, H: size, Data: data}
}

func GetRowTupleOfMatrix4x4(m Matrix, x int) *Tuple4 {
	return NewTuple4ByRaw(m.Data[x])
}

// GetColTupleOfMatrix4x4
// TODO Better Tuple support n-dim
func GetColTupleOfMatrix4x4(m Matrix, y int) *Tuple4 {
	return NewTuple4(m.Get(0, y), m.Get(1, y), m.Get(2, y), m.Get(3, y))
}

func MultiplyMatrix4x4(m1 Matrix, m2 Matrix) *Matrix {
	m3 := NewMatrix(4)
	for i := 0; i < 4; i++ {
		row := GetRowTupleOfMatrix4x4(m1, i)
		for j := 0; j < 4; j++ {
			col := GetColTupleOfMatrix4x4(m2, j)
			m3.Set(i, j, Dot(*row, *col))
		}
	}
	return m3
}

func MultiplyMatrixByTuple(m Matrix, t Tuple4) *Tuple4 {
	t1 := NewEmptyTuple4()
	for i := 0; i < 4; i++ {
		row := GetRowTupleOfMatrix4x4(m, i)
		t1.Set(i, Dot(*row, t))
	}
	return t1
}

func NewIdentityMatrix(size int) *Matrix {
	m := NewMatrix(size)
	for i := 0; i < size; i++ {
		m.Set(i, i, 1)
	}
	return m
}
