package mat

type Tuple4 struct {
	Data []float64
}

func NewPoint(x, y, z float64) *Tuple4 {
	return NewTuple4(x, y, z, 1.0)
}

func NewVector(x, y, z float64) *Tuple4 {
	return NewTuple4(x, y, z, 0.0)
}

func NewTuple4(x, y, z, flag float64) *Tuple4 {
	return &Tuple4{Data: []float64{x, y, z, flag}}
}

func NewEmptyTuple4() *Tuple4 {
	return &Tuple4{Data: make([]float64, 4)}
}

func (t Tuple4) IsPoint() bool {
	return t.Data[3] == 1.0
}

func (t Tuple4) IsVector() bool {
	return t.Data[3] == 0.0
}

func Add(t1 *Tuple4, t2 *Tuple4) *Tuple4 {
	t3 := NewEmptyTuple4()
	for i := 0; i < 4; i++ {
		t3.Data[i] = t1.Data[i] + t2.Data[i]
	}
	return t3
}

func Sub(t1 *Tuple4, t2 *Tuple4) *Tuple4 {
	t3 := NewEmptyTuple4()
	for i := 0; i < 4; i++ {
		t3.Data[i] = t1.Data[i] - t2.Data[i]
	}
	return t3
}

func Negate(t *Tuple4) *Tuple4 {
	return Sub(NewVector(0, 0, 0), t)
}

func MultiplyByScalar(t *Tuple4, scalar float64) *Tuple4 {
	t1 := NewEmptyTuple4()
	for i := 0; i < 4; i++ {
		t1.Data[i] = t.Data[i] * scalar
	}
	return t1
}

func DivideByScalar(t *Tuple4, scalar float64) *Tuple4 {
	return MultiplyByScalar(t, 1/scalar)
}
