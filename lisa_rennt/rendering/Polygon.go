package rendering

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

type Vertex struct {
	X float64
	Y float64
}
