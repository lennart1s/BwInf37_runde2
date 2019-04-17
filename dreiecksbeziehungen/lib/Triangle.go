package lib

type Triangle struct {
	A *Vector
	B *Vector
	C *Vector

	Alpha float64
	Beta  float64
	Gamma float64
}

type Corner int

const (
	A Corner = 0
	B Corner = 1
	C Corner = 2
)

func (t *Triangle) Rotate(c Corner, angle float64) {
	corners := []*Vector{t.A, t.B, t.C}

	translator := NegateVector(corners[c])
	backTransl := corners[c]

	for _, c := range corners {
		c.Add(translator)
		c.Rotate(angle)
		c.Add(backTransl)
	}
}

func (t *Triangle) Corners() []*Vector {
	return []*Vector{t.A, t.B, t.C}
}

func (t *Triangle) Angles() []float64 {
	return []float64{t.Alpha, t.Beta, t.Gamma}
}
