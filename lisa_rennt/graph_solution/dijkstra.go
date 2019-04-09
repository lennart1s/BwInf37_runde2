package graph_solution

import (
	"BwInf37_runde2/lisa_rennt/lib"
	"math"
)

func Dijkstra(g *lib.Graph) *lib.Node {
	var finish *lib.Node

	minWeight := math.MaxFloat64
	for _, n := range g.Nodes {
		for _, e := range n.Edges {
			if e.Weight < minWeight {
				minWeight = e.Weight
			}
		}
	}
	for _, n := range g.Nodes {
		if n.Info["Type"] != "start" {
			n.ShortestPath = math.MaxFloat64
		}
		/*if minWeight < 0 {
			for _, e := range n.Edges {
				e.Weight -= minWeight
			}
		}*/
	}

	for q := g.Nodes; len(q) > 0; {
		var min *lib.Node
		var minI int
		for i, n := range q {
			if min == nil || n.ShortestPath < min.ShortestPath {
				min = n
				minI = i
			}
		}
		if min.Info["Type"] == "finish" {
			finish = min
			break
		}
		q = append(q[:minI], q[minI+1:]...)
		for _, e := range min.Edges {
			if min.ShortestPath+e.Weight < e.ShortestPath {
				e.ShortestPath = min.ShortestPath + e.Weight
				e.ShortestParent = min
			}
		}
	}

	/*for _, n := range g.Nodes {
		for _, e := range n.Edges {
			e.Weight += minWeight
		}
	}*/
	return finish

	/*	q := g.Nodes
		for len(q) > 0 {
			var min *lib.Node
			minI := 0
			for i, n := range q {
				if min == nil || n.ShortestPath < min.ShortestPath {
					min = n
					minI = i
				}
			}
			if min.Info["Type"] == "finish" {
				return min
			}
			q = append(q[:minI], q[minI+1:]...)

			for _, e := range min.Edges {
				for _, o := range q {
					if e.Node == o {
						if min.ShortestPath+e.Weight < e.ShortestPath {
							e.ShortestPath = min.ShortestPath + e.Weight
							e.ShortestParent = min
						}
						break
					}
				}

			}
		}

		var min *lib.Node
		for _, n := range g.Nodes {
			if n.Info["Type"] == "finish" && (min == nil || n.ShortestPath < min.ShortestPath) {
				min = n
			}
		}
		return min*/

	/*var toGo []*lib.Node
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
				e.Node.ShortestParent = n
				e.Node.ShortestPath = n.ShortestPath + e.Weight
				toGo = append(toGo, e.Node)
			}
		}

		if len(toGo) > 1 {
			toGo = append(toGo[:nIndex], toGo[nIndex+1:]...)
		} else {
			println("no path found")
			return nil
		}
	}*/
}
