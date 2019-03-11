package main

import (
	"BwInf37_runde2/lisa_rennt/file"
	"BwInf37_runde2/lisa_rennt/lib"
	"BwInf37_runde2/lisa_rennt/rendering"
	"fmt"
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
	l1 := lib.Line{lib.Vertex{-5.88, -0.58}, lib.Vertex{0.6, 1.76}}
	l2 := lib.Line{lib.Vertex{-7.96, -0.18}, lib.Vertex{0.5, -3.5}}
	fmt.Println(lib.LineSegementIntersection(l1, l2))

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

	f, _ := os.Create("./lisa_rennt/save.png")
	png.Encode(f, Img)
}
