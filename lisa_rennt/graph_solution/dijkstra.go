package graph_solution

import (
	"BwInf37_runde2/lisa_rennt/lib"
)

func Dijkstra(g *lib.Graph) *lib.Node {
	/*var minWeight float64 = math.MaxFloat64
	for _, n := range g.Nodes {
		for _, e := range n.Edges {
			if e.Weight < minWeight {
				minWeight = e.Weight
			}
		}
	}
	for _, n := range g.Nodes {
		for _, e := range n.Edges {
			e.Weight -= minWeight
		}
	}*/

	var start *lib.Node
	for _, n := range g.Nodes {
		if n.Info["Type"] == "start" {
			start = n
		} else {
			n.ShortestPath = -1
		}
	}

	var toGo []*lib.Node
	toGo = append(toGo, start)

	for {
		var n *lib.Node
		nIndex := -1
		for i, searched := range toGo {
			if n == nil || searched.ShortestPath < n.ShortestPath {
				n = searched
				nIndex = i
			}
		}
		if n.Info["Type"] == "finish" {
			return n
		}

		for _, e := range n.Edges {
			if e.Node.ShortestPath == -1 || n.ShortestPath+e.Weight < e.Node.ShortestPath {
				e.ShortestParent = n
				e.ShortestPath = n.ShortestPath + e.Weight
				toGo = append(toGo, e.Node)
			}
		}

		if len(toGo) > 1 {
			toGo = append(toGo[:nIndex], toGo[nIndex+1:]...)
		} else {
			println("no path found")
			return nil
		}
	}
}
