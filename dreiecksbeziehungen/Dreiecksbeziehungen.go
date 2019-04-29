package main

import (
	"BwInf37_runde2/dreiecksbeziehungen/files"
	"BwInf37_runde2/dreiecksbeziehungen/lib"
	"BwInf37_runde2/dreiecksbeziehungen/rendering"
	"BwInf37_runde2/dreiecksbeziehungen/solution"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]
	triangles := files.Load(args[0])

	solution.Solve(triangles)

	fmt.Printf("Datei:\t\t%v\n", args[0])
	fmt.Printf("Gesamtabstand:\t%vm\n", int(getDistance(triangles)+0.5))
	description := ""
	for _, t := range triangles {
		description += t.Info["ID"]
		for ci, c := range t.Corners() {
			description += " x" + strconv.Itoa(ci+1) + ":" + strconv.Itoa(int(c.X+0.5))
			description += " y" + strconv.Itoa(ci+1) + ":" + strconv.Itoa(int(c.Y+0.5))
		}
		description += "\n"
	}
	fmt.Printf("x- und y-Koordinaten:\n%v", description)

	rendering.RenderTriangles(triangles, "./dreiecksbeziehungen_solution.png")
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
