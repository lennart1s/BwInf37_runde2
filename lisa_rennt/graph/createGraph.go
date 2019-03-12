package graph

import (
	"BwInf37_runde2/lisa_rennt/lib"
	"math"
)

func Create(home lib.Vertex, polygons []*lib.Polygon) lib.Graph {
	g := lib.Graph{}
	//g.Nodes = append(g.Nodes, &lib.Node{X: home.X, Y: home.Y})
	//g.Nodes = append(g.Nodes, verticesToNodes(polygons)...)

	//var nodes map[*lib.Vertex]*lib.Node
	nodes := make(map[lib.Vertex]*lib.Node)
	nodes[lib.Vertex{}] = &lib.Node{X: home.X, Y: home.Y}
	addAllVertices(&nodes, polygons)

	var lines []*lib.Line
	lines = append(lines, getAllLines(polygons)...)

	

	for _, p := range polygons {
		for _, v := range p.Vertices {
			for _, op := range polygons {
				if op != p {
					for _, ov := range op.Vertices {
						if canReach(&v, &ov, lines) {
							nodes[v].Edges = append(nodes[v].Edges, &lib.Edge{Node: nodes[ov], Weight: distance(&v, &ov)})
						}
					}
				}
			}
		}
	}

	for _, n := range nodes {
		g.Nodes = append(g.Nodes, n)
	}

	/*for _, n := range g.Nodes {
		for _, o := range g.Nodes {
			if n != o && canReach(n, o, lines) {
				n.Edges = append(n.Edges, &lib.Edge{Node: o, Weight: distance(n, o)})
			}
		}
	}*/

	/*for _, p := range polygons {
		for _, v := range p.Vertices {
			n := lib.Node{X: v.X, Y: v.Y}
			for _, op := range polygons {
				if op != p {
					for _, ov := range op.Vertices {
						if canReach(v, ov, lines) {
							n.Edges = append(n.Edges, lib.Edge{})
						}
					}
				}
			}
			g.Nodes = append(g.Nodes, &n)
		}
	}*/

	return g
}

func canReach(n *lib.Vertex, o *lib.Vertex, lines []*lib.Line) bool {
	ray := lib.Line{A: *n, B: *o}

	for _, l := range lines {
		if l.A.X == ray.A.X && l.A.Y == ray.A.Y {
			continue
		}
		if l.B.X == ray.A.X && l.B.Y == ray.A.Y {
			continue
		}
		if l.A.X == ray.B.X && l.A.Y == ray.B.Y {
			continue
		}
		if l.B.X == ray.B.X && l.B.Y == ray.B.Y {
			continue
		}

		if lib.LineSegementIntersection(&ray, l) {
			return false
		}
	}

	return true
}

func distance(n *lib.Vertex, o *lib.Vertex) float64 {
	return math.Sqrt(math.Pow(n.X-o.X, 2) + math.Pow(n.Y-o.Y, 2))
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

/*func verticesToNodes(polygons []*lib.Polygon) []*lib.Node {
	var nodes []*lib.Node
	for _, p := range polygons {
		for _, v := range p.Vertices {
			nodes = append(nodes, &lib.Node{X: v.X, Y: v.Y})
		}
	}

	return nodes
}*/
