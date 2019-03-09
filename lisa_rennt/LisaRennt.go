package main

import (
	"BwInf37_runde2/lisa_rennt/file"
	"BwInf37_runde2/lisa_rennt/rendering"
	"image"
	"image/draw"
	"image/png"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
)

var (
	Obstacles []*rendering.Polygon
	Home      rendering.Vertex
)

func main() {
	// Setup
	img := image.NewRGBA(image.Rect(0, 0, rendering.WIDTH, rendering.HEIGHT))
	draw.Draw(img, img.Bounds(), &image.Uniform{rendering.BACKGROUND}, image.ZP, draw.Src)
	gc := draw2dimg.NewGraphicContext(img)
	gc.Scale(1, -1)
	gc.Translate(rendering.BUSSTOP_RADIUS*1.5, -(float64(rendering.HEIGHT) - rendering.BUSSTOP_RADIUS*1.5))

	// Read-File
	Obstacles, Home = file.Read("./lisa_rennt/examples/lisarennt5.txt")

	rendering.RenderEnvironment(gc)
	rendering.RenderHome(gc, Home)
	rendering.RenderObstacles(gc, Obstacles...)

	f, _ := os.Create("./lisa_rennt/save.png")
	png.Encode(f, img)
}
