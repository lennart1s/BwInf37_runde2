package lib

type Graph struct {
	Nodes []*Node
}

type Node struct {
	X float64
	Y float64

	Id string

	Edges []*Edge

	ShortestPath   float64
	ShortestParent *Node
}

type Edge struct {
	*Node
	Weight float64
}
