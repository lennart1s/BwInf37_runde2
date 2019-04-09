package graph_solution

import (
	"BwInf37_runde2/lisa_rennt/lib"
	"math"
)

func BellmanFord(g *lib.Graph) *lib.Node {
	var finish *lib.Node

	for _, n := range g.Nodes {
		if n.Info["Type"] != "start" {
			n.ShortestPath = math.MaxFloat64
		}
	}

	for i := 0; i < len(g.Nodes)-1; i++ {
		for _, n := range g.Nodes {
			for _, e := range n.Edges {
				if n.ShortestPath+e.Weight < e.ShortestPath {
					e.ShortestPath = n.ShortestPath + e.Weight
					e.ShortestParent = n
				}
			}
		}
	}

	for _, n := range g.Nodes {
		if n.Info["Type"] == "finish" {
			if finish == nil || n.ShortestPath < finish.ShortestPath {
				finish = n
			}
		}
	}

	return finish
}
