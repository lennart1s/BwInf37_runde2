package updown

import (
	"BwInf37_runde2/dreiecksbeziehungen/lib"
	"math"
)

func UpDown(triangles []*lib.Triangle) {
	for ti, t := range triangles {
		angle := -1.0
		angleIndex := -1
		for ai, a := range t.Angles() {
			if a < angle || angleIndex == -1 {
				angle = a
				angleIndex = ai
			}
		}

		if ti%2 == 0 {
			a := lib.AddVector(t.Corners()[(angleIndex+2)%3], lib.NegateVector(t.Corners()[(angleIndex+1)%3])).GetAngle(&lib.Vector{1, 0})
			t.Rotate(lib.Corner(angleIndex%3), -a)
		} else {
			a := lib.AddVector(t.Corners()[(angleIndex+2)%3], lib.NegateVector(t.Corners()[(angleIndex)%3])).GetAngle(&lib.Vector{1, 0})
			t.Rotate(lib.Corner(angleIndex%3), (math.Pi/2+angle/2)-a)
		}
		groundTriangle(t)
		horizontalAppendToTriangles(t, triangles[:ti])
	}

}

func horizontalAppendToTriangles(t *lib.Triangle, others []*lib.Triangle) {
	mostRight := 0.0
	mostLeft := math.MaxFloat64
	for _, o := range others {
		for _, c := range o.Corners() {
			if c.X > mostRight {
				mostRight = c.X
			}
		}
	}
	for _, c := range t.Corners() {
		if c.X < mostLeft {
			mostLeft = c.X
		}
	}

	t.Move(mostRight-mostLeft, 0)

	if len(others) == 0 {
		return
	}

	for dx := -20.0; dx < -0.1; {
		t.Move(dx, 0)
		if collidesWithOthers(t, others) {
			t.Move(-dx, 0)
			dx *= 0.5
		}
	}
}

func collidesWithOthers(t *lib.Triangle, others []*lib.Triangle) bool {
	for _, o := range others {
		if t.Collides(o) {
			return true
		}
	}

	return false
}

func groundTriangle(t *lib.Triangle) {
	h := 0.0
	corner := -1
	for ci, c := range t.Corners() {
		if c.Y < h || corner == -1 {
			h = c.Y
			corner = ci
		}
	}

	t.Move(0, -h)
}
