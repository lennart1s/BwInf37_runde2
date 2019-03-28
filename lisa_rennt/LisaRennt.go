package main

import (
	"BwInf37_runde2/lisa_rennt/file"
	graph "BwInf37_runde2/lisa_rennt/graph_solution"
	"BwInf37_runde2/lisa_rennt/lib"
	"BwInf37_runde2/lisa_rennt/rendering"
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"math"
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
	Obstacles, Home = file.Read("./lisa_rennt/examples/lisarennt3.txt")

	rendering.RenderEnvironment(Gc)
	rendering.RenderHome(Gc, Home)
	rendering.RenderObstacles(Gc, Obstacles...)

	g := graph.Create(Home, Obstacles)
	rendering.RenderGraph(Gc, &g)

	finish := graph.Dijkstra(&g)
	rendering.RenderPath(Gc, &g, finish)

	testData(&g, finish)

	f, _ := os.Create("./lisa_rennt/save.png")
	png.Encode(f, Img)
}

func testData(g *lib.Graph, f *lib.Node) {
	dist := 0.0
	n := f
	for n != nil {
		dist += math.Sqrt(math.Pow(f.X-f.ShortestParent.X, 2) + math.Pow(f.Y-f.ShortestParent.Y, 2))
		n = n.ShortestParent
	}
	fmt.Println("Distance:", dist, "m")

	neededTime := dist / (15 / 3.6)
	fmt.Println("Needed time:", neededTime, "s")

	fmt.Println("Bus-reach-Y:", f.Y, "m")

	busTime := f.Y / (30 / 3.6)
	fmt.Println("Bus time:", busTime, "s")

	startDelay := neededTime - busTime
	fmt.Println("Start delay:", startDelay, "s")
}
