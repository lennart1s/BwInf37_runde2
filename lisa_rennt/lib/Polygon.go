package lib

import (
	"fmt"
	"strconv"
	"strings"
)

type Polygon struct {
	Center   Vertex
	Vertices []Vertex
}

func PolygonFromLine(line string) *Polygon {
	p := Polygon{}

	parts := strings.Split(line, " ")[1:]
	if len(parts)%2 != 0 {
		println("Invalid polygon-data: " + line)
		return nil
	}
	for i := 0; i < len(parts)-1; i += 2 {
		x, err := strconv.ParseFloat(parts[i], 64)
		y, err := strconv.ParseFloat(parts[i+1], 64)
		if err != nil {
			println("Error while parsing polygon vertex: " + parts[i] + " " + parts[i+1])
			return nil
		}
		p.Vertices = append(p.Vertices, Vertex{x, y})
	}

	return &p
}

func (p *Polygon) IsClockwise() bool {
	a := 0.0
	for i := 0; i < len(p.Vertices)-1; i++ {
		a += p.Vertices[i].X*p.Vertices[i+1].Y - p.Vertices[i+1].X*p.Vertices[i].Y
	}
	a *= 0.5

	return a < 0
}

func (p *Polygon) ArrangeClockwise() {
	if !p.IsClockwise() {
		verts := len(p.Vertices)
		for i := 0; i < verts/2; i++ {
			p.Vertices[i], p.Vertices[verts-i-1] = p.Vertices[verts-i-1], p.Vertices[i]
		}
	}
}

func (p *Polygon) ForAllVertices(startV Vertex, itBw bool, itF iterationFunc) {
	i := 0
	for {
		if p.Vertices[i].X == startV.X && p.Vertices[i].Y == startV.Y {
			break
		}
		i++
		if i == len(p.Vertices) {
			println("Start-vertex not found.")
			fmt.Println(startV)
			fmt.Println(p.Vertices)
			return
		}
	}
	step := 1
	if itBw {
		step = -1
	}
	first := false
	for !first || !(p.Vertices[i].X == startV.X && p.Vertices[i].Y == startV.Y) {
		first = true
		if itF(&p.Vertices[i]) {
			break
		}

		i += step
		if i >= len(p.Vertices) {
			i = 0
		} else if i < 0 {
			i = len(p.Vertices) - 1
		}
	}

}

type iterationFunc = func(v *Vertex) bool
