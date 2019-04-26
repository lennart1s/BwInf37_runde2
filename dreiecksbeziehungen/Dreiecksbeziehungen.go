package main

import (
	"BwInf37_runde2/dreiecksbeziehungen/files"
	"BwInf37_runde2/dreiecksbeziehungen/lib"
	"BwInf37_runde2/dreiecksbeziehungen/rendering"
	"BwInf37_runde2/dreiecksbeziehungen/updown"
	"fmt"
	"math"
)

func main() {
	triangles := files.Load("./dreiecksbeziehungen/examples/dreiecke2.txt")

	updown.UpDown(triangles)

	/*for _, t := range triangles {
		found := false
		for _, c := range t.Corners() {
			if c.Y == 0 {
				found = true
			}
		}
		fmt.Println(found)
	}*/

	rendering.RenderTriangles(triangles, "./triOutput.png")

	fmt.Println(getDistance(triangles))
	fmt.Println(eachGrounded(triangles))
}

func getDistance(triangles []*lib.Triangle) float64 {
	leftTRightC := 0.0
	for _, c := range triangles[0].Corners() {
		if c.Y < 0.0001 && c.X > leftTRightC {
			leftTRightC = c.X
		}
	}
	rightTLeftC := math.MaxFloat64
	for _, c := range triangles[len(triangles)-1].Corners() {
		if c.Y < 0.0001 && c.X < rightTLeftC {
			rightTLeftC = c.X
		}
	}

	return rightTLeftC - leftTRightC
}

func eachGrounded(triangles []*lib.Triangle) bool {
triSearch:
	for _, t := range triangles {
		for _, c := range t.Corners() {
			if c.Y < 0.0001 {
				continue triSearch
			}
		}
		return false
	}
	return true
}
