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

func NewMatrix(size int) *Matrix {
	data := make([][]float64, size)
	for i := 0; i < size; i++ {
		data[i] = make([]float64, size)
	}
	return &Matrix{W: size, H: size, Data: data}
}
