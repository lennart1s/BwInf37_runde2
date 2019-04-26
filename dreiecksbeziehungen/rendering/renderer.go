package rendering

import (
	"BwInf37_runde2/dreiecksbeziehungen/lib"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"math"
	"os"

	"github.com/llgcode/draw2d/draw2dimg"
)

func RenderTriangles(triangles []*lib.Triangle, outputPath string) {
	w, h := getSize(triangles)
	Img := image.NewRGBA(image.Rect(0, 0, w+20, h+20))
	draw.Draw(Img, Img.Bounds(), &image.Uniform{color.RGBA{200, 200, 200, 255}}, image.ZP, draw.Src)
	Gc := draw2dimg.NewGraphicContext(Img)
	Gc.Scale(1, -1)
	Gc.Translate(10, -(float64(h) + 10))
	//Gc.Translate(150, 150)

	Gc.SetFillColor(color.RGBA{110, 110, 110, 255})
	Gc.SetStrokeColor(color.RGBA{40, 40, 40, 255})
	Gc.SetLineWidth(2)
	for _, t := range triangles {
		Gc.MoveTo(t.A.X, t.A.Y)
		Gc.LineTo(t.B.X, t.B.Y)
		Gc.LineTo(t.C.X, t.C.Y)
		Gc.Close()
		Gc.FillStroke()
	}

	CircleSmallestAngle(Gc, triangles)

	f, _ := os.Create(outputPath)
	png.Encode(f, Img)
}

func getSize(triangles []*lib.Triangle) (int, int) {
	var w, h float64
	for _, t := range triangles {
		for _, c := range t.Corners() {
			if c.X > w {
				w = c.X
			}
			if c.Y > h {
				h = c.Y
			}
		}
	}

	return int(w), int(h)
}

func CircleSmallestAngle(gc *draw2dimg.GraphicContext, triangles []*lib.Triangle) {
	for _, t := range triangles {
		angle := 0.0
		aIndex := -1
		for ai, a := range t.Angles() {
			if aIndex == -1 || a < angle {
				angle = a
				aIndex = ai
			}
		}

		gc.SetStrokeColor(color.RGBA{220, 40, 40, 255})
		renderCircle(gc, t.Corners()[aIndex].X, t.Corners()[aIndex].Y, 10)
	}
}

func renderCircle(gc *draw2dimg.GraphicContext, cx float64, cy float64, radius float64) {
	for vi := 0.0; vi < float64(16); vi++ {
		x := cx + math.Cos(vi*2.0*math.Pi/float64(16))*radius
		y := cy + math.Sin(vi*2.0*math.Pi/float64(16))*radius
		if vi == 0 {
			gc.MoveTo(x, y)
		} else {
			gc.LineTo(x, y)
		}
	}
	gc.Close()
	gc.Stroke()
}
