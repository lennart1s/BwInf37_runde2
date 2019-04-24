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

func (t *Triangle) Rotate(rc Corner, angle float64) {
	corners := t.Corners()

	translator := NegateVector(corners[rc])
	backTransl := *corners[rc]

	for _, c := range corners {
		c.Add(translator)
		c.Rotate(angle)
		//fmt.Println(c.X)
		c.Add(&backTransl)
		//fmt.Println(backTransl.X, "+")
		//fmt.Println("->", c.X)
	}
}

func (t *Triangle) Corners() []*Vector {
	return []*Vector{t.A, t.B, t.C}
}

func (t *Triangle) Angles() []float64 {
	return []float64{t.Alpha, t.Beta, t.Gamma}
}

func (t *Triangle) IsClockwise() bool {
	corners := t.Corners()
	a := 0.0
	for i := 0; i < len(corners)-1; i++ {
		a += corners[i].X*corners[i+1].Y - corners[i+1].X*corners[i].Y
	}
	a *= 0.5

	return a < 0
}

func (t *Triangle) ArrangeAntiClockwise() {
	if t.IsClockwise() {
		t.A, t.C = t.C, t.A
	}
}

func (t *Triangle) Move(dx float64, dy float64) {
	dv := Vector{dx, dy}
	for _, c := range t.Corners() {
		c.Add(&dv)
	}
}
