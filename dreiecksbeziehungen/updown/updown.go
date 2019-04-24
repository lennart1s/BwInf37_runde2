package updown

import (
	"BwInf37_runde2/dreiecksbeziehungen/lib"
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
			t.Rotate(lib.Corner((angleIndex)%3), -a)

			groundTriangle(t)
		} else {

		}
	}

}

func horizontalAppendToTriangles(t *lib.Triangle, others []*lib.Triangle) {

}

func collidesWithOthers(t *lib.Triangle, others []*lib.Triangle) {

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
