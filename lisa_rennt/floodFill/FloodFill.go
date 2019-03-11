package floodFill

import (
	"BwInf37_runde2/lisa_rennt/rendering"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/llgcode/draw2d"
)

var (
	fillColor = color.RGBA{10, 10, 180, 255}
)

func Calculate(img *image.RGBA, m draw2d.Matrix, home rendering.Vertex) {
	var toExplore []rendering.Vertex
	paint(img, m, home)
	toExplore = append(toExplore, getSourrounding(img, home, m)...)

	for {
		var new []rendering.Vertex
		for _, v := range toExplore {
			if img.At(int(v.X), int(v.Y)) != fillColor {
				paint(img, m, v)
				new = append(new, getSourrounding(img, v, m)...)
			}
		}

		// Way Found
		for _, v := range new {
			if v.X == 0 {
				break
			}
		}

		// No Way
		toExplore = toExplore[:0]
		toExplore = append(toExplore, new...)
		if len(toExplore) == 0 {
			break
		}
		println(len(toExplore))

		f, _ := os.Create("./lisa_rennt/auto.png")
		png.Encode(f, img)
	}
}

func getSourrounding(img *image.RGBA, v rendering.Vertex, m draw2d.Matrix) []rendering.Vertex {
	v.X, v.Y = m.TransformPoint(v.X, v.Y)
	var s []rendering.Vertex
	for x := -1.0; x <= 1.0; x++ {
		for y := -1.0; y <= 1.0; y++ {
			if (x == 0 && y != 0) || (x != 0 && y == 0) {
				o := rendering.Vertex{v.X + x, v.Y + y}
				if !outOfBounds(o) && img.At(int(o.X), int(o.Y)) != fillColor &&
					img.At(int(o.X), int(o.Y)) != rendering.OBSTACLE_OUTLINE {
					s = append(s, o)
				}
			}
		}
	}

	return s
}

func outOfBounds(v rendering.Vertex) bool {
	if v.X < 0 || v.Y < 0 {
		return true
	} else if v.X >= float64(rendering.WIDTH) || v.Y >= float64(rendering.HEIGHT) {
		return true
	}
	return false
}

func paint(img *image.RGBA, m draw2d.Matrix, v rendering.Vertex) {
	x, y := m.TransformPoint(v.X, v.Y)
	img.Set(int(x), int(y), fillColor)
}
