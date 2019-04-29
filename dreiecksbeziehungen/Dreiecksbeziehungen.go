package main

import (
	"BwInf37_runde2/dreiecksbeziehungen/files"
	"BwInf37_runde2/dreiecksbeziehungen/lib"
	"BwInf37_runde2/dreiecksbeziehungen/rendering"
	"BwInf37_runde2/dreiecksbeziehungen/solution"
	"fmt"
	"math"
	"os"
)

func main() {
	args := os.Args[1:]
	triangles := files.Load(args[0])

	solution.Solve(triangles)

	rendering.RenderTriangles(triangles, "./triOutput.png")

	fmt.Println("-------------")
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
