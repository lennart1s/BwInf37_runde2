package main

import (
	"BwInf37_runde2/lisa_rennt/rendering"
	"image"
	"image/draw"
	"image/png"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/llgcode/draw2d/draw2dimg"
)

var (
	obstacles []*rendering.Polygon
	home      rendering.Vertex
)

func main() {
	// Setup
	img := image.NewRGBA(image.Rect(0, 0, rendering.WIDTH, rendering.HEIGHT))
	draw.Draw(img, img.Bounds(), &image.Uniform{rendering.BACKGROUND}, image.ZP, draw.Src)
	gc := draw2dimg.NewGraphicContext(img)
	gc.Scale(1, -1)
	gc.Translate(rendering.BUSSTOP_RADIUS*1.5, -(float64(rendering.HEIGHT) - rendering.BUSSTOP_RADIUS*1.5))

	// Read-File
	readFile("./lisa_rennt/examples/lisarennt2.txt")

	rendering.RenderEnvironment(gc)
	rendering.RenderHome(gc, home)
	rendering.RenderObstacles(gc, obstacles...)

	f, _ := os.Create("save.png")
	png.Encode(f, img)
}

func readFile(path string) {
	lines := loadLines(path)
	numPolygons, err := strconv.Atoi(lines[0])
	for _, line := range lines[1 : 1+numPolygons] {
		obstacles = append(obstacles, rendering.PolygonFromLine(line))
	}

	homeParts := strings.Split(lines[len(lines)-1], " ")
	home.X, err = strconv.ParseFloat(homeParts[0], 64)
	home.Y, err = strconv.ParseFloat(homeParts[1], 64)

	if err != nil {
		panic("Invalid data in file: " + path)
	}
}

func loadText(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func loadLines(path string) []string {
	text := loadText(path)
	var lines []string
	for _, lineA := range strings.Split(text, "\r\n") {
		for _, lineB := range strings.Split(lineA, "\n") {
			lines = append(lines, lineB)
		}
	}
	return lines
}
