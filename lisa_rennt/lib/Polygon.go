package lib

import (
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
