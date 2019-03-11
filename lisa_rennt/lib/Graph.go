package lib

type Graph struct {
	Nodes []Node
}

type Node struct {
	Vertex
	Id string

	Edges []Edge

	ShortestPath   float64
	ShortestParent *Node
}

type Edge struct {
	*Node
	Weight float64
}
