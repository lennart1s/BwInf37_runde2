package lib

type Graph struct {
	Nodes []*Node
}

type Node struct {
	*Vertex

	Info map[string]string

	Edges []*Edge

	ShortestPath   float64
	ShortestParent *Node
}

type Edge struct {
	*Node

	Info map[string]string

	Weight float64
}
