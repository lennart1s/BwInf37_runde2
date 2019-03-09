package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
)

func main() {
	img := image.NewRGBA(image.Rect(0, 0, 400, 400))
	draw.Draw(img, img.Bounds(), &image.Uniform{color.RGBA{220, 220, 220, 255}}, image.ZP, draw.Src)

	gc := draw2dimg.NewGraphicContext(img)
	gc.SetFillColor(color.RGBA{100, 100, 100, 255})
	gc.SetStrokeColor(color.RGBA{50, 50, 50, 255})
	gc.SetLineWidth(3)

	gc.MoveTo(10, 10)
	gc.LineTo(100, 50)
	gc.QuadCurveTo(100, 10, 10, 10)
	gc.Close()
	gc.FillStroke()

	f, _ := os.Create("save.png")
	png.Encode(f, img)
}
