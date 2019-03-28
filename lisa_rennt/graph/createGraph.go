package graph

import (
	"BwInf37_runde2/lisa_rennt/lib"
	"math"
	"strconv"
)

const (
	lisaVel_KMH = 15.0
	busVel_KMH  = 30.0
)

func Create(home lib.Vertex, polygons []*lib.Polygon) lib.Graph {
	for _, p := range polygons {
		p.ArrangeClockwise()
	}

	g := lib.Graph{}
	g.Nodes = append(g.Nodes, &lib.Node{Vertex: &lib.Vertex{X: home.X, Y: home.Y}})
	g.Nodes = append(g.Nodes, verticesToNodes(polygons)...)

	var borders []*lib.Line
	borders = append(borders, getAllLines(polygons)...)

	idealAngle := math.Atan(lisaVel_KMH / math.Sqrt(math.Pow(busVel_KMH, 2)-math.Pow(lisaVel_KMH, 2)))

	for _, n := range g.Nodes {
		for _, o := range g.Nodes {
			if n == o {
				continue
			}
			if !belongToSamePolygon(n, o) {
				if canReach(n, o, borders) {
					n.Edges = append(n.Edges, &lib.Edge{Node: o, Weight: distance(n, o)})
				}
			} else {
				pi, _ := strconv.Atoi(n.Info["p"])
				if canReach(n, o, borders) && canReachP(n, o, polygons[pi]) {
					n.Edges = append(n.Edges, &lib.Edge{Node: o, Weight: distance(n, o)})
				}
			}
		}
		// Check if can reach optimal busY
		idealY := n.Y + math.Tan(idealAngle)*n.X
		borderNode := lib.Node{Vertex: &lib.Vertex{X: 0, Y: idealY}}
		if canReach(n, &borderNode, borders) {
			n.Edges = append(n.Edges, &lib.Edge{Node: &borderNode, Weight: 0}) // TODO: add real weight
		}
	}

	return g
}

func canReach(n *lib.Node, o *lib.Node, borders []*lib.Line) bool {
	ray := lib.Line{A: lib.Vertex{X: n.X, Y: n.Y}, B: lib.Vertex{X: o.X, Y: o.Y}}

	for _, b := range borders {
		if isNodeOnLine(n, b) || isNodeOnLine(o, b) {
			continue
		}

		if lib.LineSegementIntersection(&ray, b) {
			return false
		}
	}

	return true
}

func canReachP(n *lib.Node, o *lib.Node, p *lib.Polygon) bool {
	canReach := false

	var prevV lib.Vertex
	f := func(v *lib.Vertex) bool {
		transl := lib.Vertex{X: v.X - n.X, Y: v.Y - n.Y}
		dp := dotProd(prevV, transl)
		prevV = transl
		if dp < 0 {
			return true
		} else if v.X == o.X && v.Y == o.Y {
			canReach = true
			return true
		}
		return false
	}
	f2 := func(v *lib.Vertex) bool {
		transl := lib.Vertex{X: v.X - n.X, Y: v.Y - n.Y}
		dp := dotProd(prevV, transl)
		prevV = transl
		if dp > 0 {
			return true
		} else if v.X == o.X && v.Y == o.Y {
			canReach = true
			return true
		}
		return false
	}
	p.ForAllVertices(lib.Vertex{X: n.X, Y: n.Y}, false, f)
	if canReach {
		return true
	}
	prevV = lib.Vertex{X: 0, Y: 0}
	p.ForAllVertices(lib.Vertex{X: n.X, Y: n.Y}, true, f2)
	return canReach
}

func distance(n *lib.Node, o *lib.Node) float64 {
	return math.Sqrt(math.Pow(n.X-o.X, 2) + math.Pow(n.Y-o.Y, 2))
}

func belongToSamePolygon(n *lib.Node, o *lib.Node) bool {
	return n.Info["p"] == o.Info["p"]
}

func isNodeOnLine(n *lib.Node, l *lib.Line) bool {
	return n.Info["p"] == l.Info["p"] && (n.Info["v"] == l.Info["v1"] || n.Info["v"] == l.Info["v2"])
}

func areNeightbors(n *lib.Node, o *lib.Node, vertices []lib.Vertex) bool {
	for vi, v := range vertices {
		if (v.X == n.X && v.Y == n.Y) || (v.X == o.X && v.Y == o.Y) {
			j := vi + 1
			if j >= len(vertices) {
				j = 0
			}
			return (vertices[j].X == n.X && vertices[j].Y == n.Y) || (vertices[j].X == o.X && vertices[j].Y == o.Y)
		}
	}

	return false
}

func getAllLines(polygons []*lib.Polygon) []*lib.Line {
	var lines []*lib.Line

	for ip, p := range polygons {
		for i := 0; i < len(p.Vertices); i++ {
			j := i + 1
			if j >= len(p.Vertices) {
				j = 0
			}
			lines = append(lines, &lib.Line{A: p.Vertices[i], B: p.Vertices[j], Info: map[string]string{"p": strconv.Itoa(ip),
				"v1": strconv.Itoa(i), "v2": strconv.Itoa(j)}})
		}
	}

	return lines
}

func verticesToNodes(polygons []*lib.Polygon) []*lib.Node {
	var nodes []*lib.Node
	for ip, p := range polygons {
		for iv, _ := range p.Vertices {
			nodes = append(nodes, &lib.Node{Vertex: &polygons[ip].Vertices[iv], Info: map[string]string{"p": strconv.Itoa(ip), "v": strconv.Itoa(iv)}})
		}
	}

	return nodes
}

func rotateVector(v lib.Vertex, a float64) lib.Vertex {
	x, y := math.Cos(a)*v.X-math.Sin(a)*v.Y, math.Sin(a)*v.X+math.Cos(a)*v.Y
	if math.Abs(x) < 0.0001 {
		x = 0
	}
	if math.Abs(y) < 0.0001 {
		y = 0
	}

	return lib.Vertex{X: x, Y: y}
}

func dotProd(u lib.Vertex, v lib.Vertex) float64 {
	return u.X*v.Y - v.X*u.Y
}

func sign(x float64) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}
