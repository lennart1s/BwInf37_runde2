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

func RenderPath(gc *draw2dimg.GraphicContext, g *lib.Graph, n *lib.Node) {
	gc.SetFillColor(GRAPH_PATH)
	gc.SetStrokeColor(GRAPH_PATH)
	gc.SetLineWidth(GRAPH_PATH_WIDTH)

	gc.MoveTo(n.X, n.Y)
	drawToParent(gc, n)
	gc.Stroke()
}

func drawToParent(gc *draw2dimg.GraphicContext, n *lib.Node) {
	if n.ShortestParent != nil {
		gc.LineTo(n.ShortestParent.X, n.ShortestParent.Y)
		drawToParent(gc, n.ShortestParent)
	}
}
