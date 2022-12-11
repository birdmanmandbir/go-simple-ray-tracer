package mat

import "math"

type Tuple4 struct {
	Data []float64
}

func (t Tuple4) IsPoint() bool {
	return t.Get(3) == 1.0
}

func (t Tuple4) IsVector() bool {
	return t.Get(3) == 0.0
}

func (t Tuple4) Get(row int) float64 {
	return t.Data[row]
}

func (t Tuple4) GetX() float64 {
	return t.Data[0]
}

func (t Tuple4) GetY() float64 {
	return t.Data[1]
}

func (t Tuple4) GetZ() float64 {
	return t.Data[2]
}

func (t Tuple4) GetW() float64 {
	return t.Data[3]
}

func NewPoint(x, y, z float64) *Tuple4 {
	return NewTuple4(x, y, z, 1.0)
}

func NewVector(x, y, z float64) *Tuple4 {
	return NewTuple4(x, y, z, 0.0)
}

func NewTuple4(x, y, z, w float64) *Tuple4 {
	return &Tuple4{Data: []float64{x, y, z, w}}
}

func NewEmptyTuple4() *Tuple4 {
	return &Tuple4{Data: make([]float64, 4)}
}

func Add(t1, t2 Tuple4) *Tuple4 {
	t3 := NewEmptyTuple4()
	for i := 0; i < 4; i++ {
		t3.Data[i] = t1.Data[i] + t2.Data[i]
	}
	return t3
}

func Sub(t1, t2 Tuple4) *Tuple4 {
	t3 := NewEmptyTuple4()
	for i := 0; i < 4; i++ {
		t3.Data[i] = t1.Data[i] - t2.Data[i]
	}
	return t3
}

func Negate(t Tuple4) *Tuple4 {
	return Sub(*NewVector(0, 0, 0), t)
}

func MultiplyByScalar(t Tuple4, scalar float64) *Tuple4 {
	t1 := NewEmptyTuple4()
	for i := 0; i < 4; i++ {
		t1.Data[i] = t.Data[i] * scalar
	}
	return t1
}

func DivideByScalar(t Tuple4, scalar float64) *Tuple4 {
	return MultiplyByScalar(t, 1/scalar)
}

func Magnitude(t Tuple4) float64 {
	powSum := math.Pow(t.Data[0], 2) + math.Pow(t.Data[1], 2) + math.Pow(t.Data[2], 2) +
		math.Pow(t.Data[3], 2)
	return math.Sqrt(powSum)
}

func Normalize(t Tuple4) *Tuple4 {
	t1 := NewEmptyTuple4()
	magnitude := Magnitude(t)
	for i := 0; i < 4; i++ {
		t1.Data[i] = t.Data[i] / magnitude
	}
	return t1
}

func TupleEquals(t1, t2 Tuple4) bool {
	return AlmostEqualFloat64(t1.Data[0], t2.Data[0]) &&
		AlmostEqualFloat64(t1.Data[1], t2.Data[1]) &&
		AlmostEqualFloat64(t1.Data[2], t2.Data[2]) &&
		AlmostEqualFloat64(t1.Data[3], t2.Data[3])
}

func Dot(t1, t2 Tuple4) float64 {
	return t1.Get(0)*t2.Get(0) +
		t1.Get(1)*t2.Get(1) +
		t1.Get(2)*t2.Get(2) +
		t1.Get(3)*t2.Get(3)
}

func Cross(t1, t2 Tuple4) *Tuple4 {
	t3 := NewEmptyTuple4()
	t3.Data[0] = t1.GetY()*t2.GetZ() - t1.GetZ()*t2.GetY()
	t3.Data[1] = t1.GetZ()*t2.GetX() - t1.GetX()*t2.GetZ()
	t3.Data[2] = t1.GetX()*t2.GetY() - t1.GetY()*t2.GetX()
	return t3
}
