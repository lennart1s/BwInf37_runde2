package solution

import (
	"BwInf37_runde2/dreiecksbeziehungen/lib"
	"math"
)

func Solve(triangles []*lib.Triangle) {

	angleSum := 0.0
	for _, t := range triangles {
		angleSum += t.Angles()[getSmallestAngleIndex(t)]
	}

	for ti := 0; ti < len(triangles); ti++ {
		t := triangles[ti]

		prevAngle := 0.0
		if ti == 0 {
			prevAngle = math.Pi
		} else {
			a := getFurthestRightGroundedCorner(triangles[ti-1])
			prevAngle = getSiteGroundAngle(triangles[ti-1].Corners()[a], triangles[ti-1].Corners()[(a+1)%3])
		}

		if prevAngle >= math.Pi/2 || angleSum <= prevAngle { // anlehnen
			ai := getSmallestAngleIndex(t)
			angle := getSiteGroundAngle(t.Corners()[ai], t.Corners()[(ai+2)%3])
			t.Rotate(lib.Corner(ai), prevAngle-angle)
		} else {
			// search
			oIndex := -1
			difference := 0.0
			for oi := ti; oi < len(triangles); oi++ {
				o := triangles[oi]
				oai := getSmallestAngleIndex(o)
				oAngle := o.Angles()[(oai+1)%3]
				diff := math.Abs(prevAngle - oAngle)
				if diff < difference || oIndex == -1 {
					oIndex = oi
					difference = diff
				}
			}
			// switch
			triangles[ti], triangles[oIndex] = triangles[oIndex], triangles[ti]
			t = triangles[ti]
			//rotate
			ai := getSmallestAngleIndex(t)
			if t.Angles()[(ai+1)%3] >= math.Pi/2 { // unten ranziehen
				angle := getSiteGroundAngle(t.Corners()[(ai+1)%3], t.Corners()[ai])
				t.Rotate(lib.Corner((ai+1)%3), prevAngle-angle)
			} else {
				angle := getSiteGroundAngle(t.Corners()[(ai+1)%3], t.Corners()[(ai+2)%3])
				t.Rotate(lib.Corner((ai+1)%3), -angle)
			}
		}

		angleSum -= t.Angles()[getSmallestAngleIndex(t)]
		groundTriangle(t)
		horizontalAppendToTriangles(t, triangles[:ti])

	}

}

func getSiteGroundAngle(a *lib.Vector, b *lib.Vector) float64 {
	translB := lib.AddVector(b, lib.NegateVector(a))

	return translB.GetAngleWithAxis()
}

func getFurthestRightGroundedCorner(t *lib.Triangle) int {
	index := -1
	xVal := 0.0
	for ci, c := range t.Corners() {
		if c.Y < 0.0001 && (index == -1 || c.X > xVal) {
			index = ci
			xVal = c.X
		}
	}

	return index
}

func getSmallestAngleIndex(t *lib.Triangle) int {
	angle := 0.0
	index := -1
	for ai, a := range t.Angles() {
		if a < angle || index == -1 {
			index = ai
			angle = a
		}
	}

	return index
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
