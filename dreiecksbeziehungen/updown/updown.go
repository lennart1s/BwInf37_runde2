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

		if ti % 2 == 0 {
			
		} else {

		}
	}

}
