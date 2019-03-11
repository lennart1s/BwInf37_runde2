package main

import (
	"BwInf37_runde2/lisa_rennt/file"
	"BwInf37_runde2/lisa_rennt/graph"
	"BwInf37_runde2/lisa_rennt/lib"
	"BwInf37_runde2/lisa_rennt/rendering"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
)

var (
	Img *image.RGBA
	Gc  *image.RGBA

	Obstacles []*lib.Polygon
	Home      lib.Vertex
)

func main() {
	// Setup
	Img := image.NewRGBA(image.Rect(0, 0, rendering.WIDTH, rendering.HEIGHT))
	draw.Draw(Img, Img.Bounds(), &image.Uniform{rendering.BACKGROUND}, image.ZP, draw.Src)
	Gc := draw2dimg.NewGraphicContext(Img)
	Gc.Scale(1, -1)
	Gc.Translate(rendering.BUSSTOP_RADIUS*1.5, -(float64(rendering.HEIGHT) - rendering.BUSSTOP_RADIUS*1.5))

	// Read-File
	Obstacles, Home = file.Read("./lisa_rennt/examples/lisarennt5.txt")

	rendering.RenderEnvironment(Gc)
	rendering.RenderHome(Gc, Home)
	rendering.RenderObstacles(Gc, Obstacles...)

	g := graph.Create(Home, Obstacles)

	f, _ := os.Create("./lisa_rennt/save.png")
	png.Encode(f, Img)
}
