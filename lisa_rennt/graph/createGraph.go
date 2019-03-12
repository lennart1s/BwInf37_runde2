package graph

import (
	"BwInf37_runde2/lisa_rennt/lib"
	"math"
	"strconv"
	"strings"
)

func Create(home lib.Vertex, polygons []*lib.Polygon) lib.Graph {
	g := lib.Graph{}
	g.Nodes = append(g.Nodes, &lib.Node{X: home.X, Y: home.Y})
	g.Nodes = append(g.Nodes, verticesToNodes(polygons)...)

	var borders []*lib.Line
	borders = append(borders, getAllLines(polygons)...)

	for _, n := range g.Nodes {
		for _, o := range g.Nodes {
			if !belongToSamePolygon(n, o) {
				if canReach(n, o, borders) {
					n.Edges = append(n.Edges, &lib.Edge{Node: o, Weight: distance(n, o)})
				}
			}
		}
	}

	return g
}

func canReach(n *lib.Node, o *lib.Node, borders []*lib.Line) bool {
	ray := lib.Line{A: lib.Vertex{X: n.X, Y: n.Y}, B: lib.Vertex{X: o.X, Y: o.Y}}

	for _, b := range borders {
		if segmentsHaveCommandPoint(&ray, b) {
			continue
		}

		if lib.LineSegementIntersection(&ray, b) {
			return false
		}
	}

	return true
}

func distance(n *lib.Node, o *lib.Node) float64 {
	return math.Sqrt(math.Pow(n.X-o.X, 2) + math.Pow(n.Y-o.Y, 2))
}

func belongToSamePolygon(n *lib.Node, o *lib.Node) bool {
	return strings.Split(n.Id, ",")[0] == strings.Split(o.Id, ",")[0]
}

func segmentsHaveCommandPoint(l1 *lib.Line, l2 *lib.Line) bool {
	return l1.A == l2.A || l1.A == l2.B ||
		l1.B == l2.A || l1.B == l2.B
}

func getAllLines(polygons []*lib.Polygon) []*lib.Line {
	var lines []*lib.Line

	for _, p := range polygons {
		for i := 0; i < len(p.Vertices); i++ {
			j := i + 1
			if j >= len(p.Vertices) {
				j = 0
			}
			lines = append(lines, &lib.Line{A: p.Vertices[i], B: p.Vertices[j]})
		}
	}

	return lines
}

func addAllVertices(nodes *map[lib.Vertex]*lib.Node, polygons []*lib.Polygon) {
	for _, p := range polygons {
		for _, v := range p.Vertices {
			(*nodes)[v] = &lib.Node{X: v.X, Y: v.Y}
		}
	}
}

func verticesToNodes(polygons []*lib.Polygon) []*lib.Node {
	var nodes []*lib.Node
	for ip, p := range polygons {
		for iv, v := range p.Vertices {
			id := strconv.Itoa(ip) + "," + strconv.Itoa(iv)
			nodes = append(nodes, &lib.Node{X: v.X, Y: v.Y, Id: id})
		}
	}

	return nodes
}
