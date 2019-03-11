package graph

import "BwInf37_runde2/lisa_rennt/lib"

func Create(home lib.Vertex, polygons []lib.Polygon) lib.Graph {
	g := lib.Graph{}
	g.Nodes = append(g.Nodes, &lib.Node{X: home.X, Y: home.Y})
	g.Nodes = append(g.Nodes, verticesToNodes(polygons)...)

	var lines []*lib.Line
	lines = append(lines, getAllLines(polygons)...)

	for _, n := range g.Nodes {
		for _, o := range g.Nodes {
			if n != o && canReach(n, o, lines) {
				n.Edges = append(n.Edges, &lib.Edge{Node: o, Weight: distance(n, o)})
			}
		}
	}

	return g
}

func canReach(n *lib.Node, o *lib.Node, lines []*lib.Line) bool {
	ray := lib.Line{A: lib.Vertex{X: n.X, Y: n.Y}, B: lib.Vertex{X: o.X, Y: o.Y}}

	for _, l := range lines {
		if lib.LineSegementIntersection(&ray, l) {
			return false
		}
	}

	return true
}

func distance(n *lib.Node, o *lib.Node) float64 {
	return 0
}

func getAllLines(polygons []lib.Polygon) []*lib.Line {
	var lines []*lib.Line

	for _, p := range polygons {
		for i := 0; i < len(p.Vertices); i++ {
			j := i
			if j >= len(p.Vertices) {
				j = 0
			}
			lines = append(lines, &lib.Line{A: p.Vertices[i], B: p.Vertices[j]})
		}
	}

	return lines
}

func verticesToNodes(polygons []lib.Polygon) []*lib.Node {
	var nodes []*lib.Node
	for _, p := range polygons {
		for _, v := range p.Vertices {
			nodes = append(nodes, &lib.Node{X: v.X, Y: v.Y})
		}
	}

	return nodes
}
