package rendering

import (
	"image/color"
	"math"

	"github.com/llgcode/draw2d/draw2dimg"
)

func RenderObstacles(gc *draw2dimg.GraphicContext, polygons ...*Polygon) {
	gc.SetFillColor(OBSTACLE_FILL)
	gc.SetStrokeColor(OBSTACLE_OUTLINE)
	gc.SetLineWidth(OBSTACLE_STROKE)

	for _, p := range polygons {
		gc.MoveTo(p.Center.X+p.Vertices[0].X, p.Center.Y+p.Vertices[0].Y)
		for _, v := range p.Vertices[1:] {
			gc.LineTo(p.Center.X+v.X, p.Center.Y+v.Y)
		}
		gc.Close()
		gc.FillStroke()
	}
}

func RenderBusStop(gc *draw2dimg.GraphicContext) {
	gc.SetFillColor(BUSSTOP_FILL)
	gc.SetStrokeColor(BUSSTOP_OUTLINE)
	gc.SetLineWidth(BUSSTOP_STROKE)

	renderCircle(gc, Vertex{0, 0}, BUSSTOP_RADIUS)
}

func RenderHome(gc *draw2dimg.GraphicContext, v Vertex) {
	gc.SetFillColor(HOME_FILL)
	gc.SetLineWidth(0)

	renderCircle(gc, v, HOME_RADIUS)
}

func RenderEnvironment(gc *draw2dimg.GraphicContext) {
	// add other environment

	// BUS-Line
	gc.SetStrokeColor(color.Black)
	gc.MoveTo(0, 0)
	gc.SetLineWidth(10)
	gc.LineTo(0, float64(HEIGHT))
	gc.FillStroke()
	// Bottom-Line
	gc.SetStrokeColor(color.Black)
	gc.MoveTo(0, 0)
	gc.SetLineWidth(10)
	gc.LineTo(float64(WIDTH), 0)
	gc.FillStroke()

	RenderBusStop(gc)
}

func renderCircle(gc *draw2dimg.GraphicContext, center Vertex, radius float64) {
	for vi := 0.0; vi < float64(CIRCLE_VERTICES); vi++ {
		x := center.X + math.Cos(vi*2.0*math.Pi/float64(CIRCLE_VERTICES))*radius
		y := center.Y + math.Sin(vi*2.0*math.Pi/float64(CIRCLE_VERTICES))*radius
		if vi == 0 {
			gc.MoveTo(x, y)
		} else {
			gc.LineTo(x, y)
		}
	}
	gc.Close()
	gc.FillStroke()
}
