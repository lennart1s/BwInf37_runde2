package rendering

import (
	"BwInf37_runde2/lisa_rennt/lib"

	"github.com/llgcode/draw2d/draw2dimg"
)

func RenderGraph(gc *draw2dimg.GraphicContext, g *lib.Graph) {
	gc.SetFillColor(GRAPH)
	gc.SetStrokeColor(GRAPH)
	gc.SetLineWidth(GRAPH_EDGE_WIDTH)

	for _, n := range g.Nodes {
		renderCircle(gc, lib.Vertex{X: n.X, Y: n.Y}, GRAPH_NODE_RADIUS)
		for _, e := range n.Edges {
			gc.MoveTo(n.X, n.Y)
			gc.LineTo(e.X, e.Y)
			gc.FillStroke()
		}
	}

}
