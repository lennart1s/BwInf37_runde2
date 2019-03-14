package lib

type Graph struct {
	Nodes []*Node
}

type Node struct {
	X float64
	Y float64

	ID string

	Edges []*Edge

	ShortestPath   float64
	ShortestParent *Node
}

type Edge struct {
	*Node

	ID string

	Weight float64
}
