package lib

type Triangle struct {
	A *Vector
	B *Vector
	C *Vector

	Alpha float64
	Beta  float64
	Gamma float64

	Info map[string]string
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
		c.Add(&backTransl)
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
	for i := 0; i < len(corners); i++ {
		a += corners[i].X*corners[(i+1)%3].Y - corners[(i+1)%3].X*corners[i].Y
	}

	return a < 0
}

func (t *Triangle) ArrangeAntiClockwise() {
	if t.IsClockwise() {
		*t.A, *t.C = *t.C, *t.A
		t.Alpha, t.Gamma = t.Gamma, t.Alpha
	}
}

func (t *Triangle) Move(dx float64, dy float64) {
	dv := Vector{dx, dy}
	for _, c := range t.Corners() {
		c.Add(&dv)
	}
}

func (t *Triangle) Collides(o *Triangle) bool {
	lines := triangleToLines(t)
	oLines := triangleToLines(o)

	for _, l := range lines {
		for _, ol := range oLines {
			if LineSegementIntersection(l, ol) {
				return true
			}
		}
	}

	return false
}

func triangleToLines(t *Triangle) []*Line {
	var lines []*Line
	cs := t.Corners()
	for ci, c := range cs {
		lines = append(lines, &Line{A: *c, B: *cs[(ci+1)%3]})
	}

	return lines
}
