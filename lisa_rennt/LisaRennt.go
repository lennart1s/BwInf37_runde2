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
	"strconv"

	"github.com/llgcode/draw2d/draw2dimg"
)

var (
	Img *image.RGBA
	Gc  *image.RGBA

	Obstacles []*lib.Polygon
	Home      lib.Vertex
)

func main() {
	args := os.Args[1:]
	Obstacles, Home = file.Read(args[0])

	g := graph.Create(Home, Obstacles)
	//rendering.RenderGraph(Gc, &g)

	finish := graph.BellmanFord(&g)

	dist := 0.0
	n := finish
	description := "  " + getNodeInfo(n) + "\n"
	for n.ShortestParent != nil {
		description = "  " + getNodeInfo(n.ShortestParent) + "\n" + description
		dist += math.Sqrt(math.Pow(n.X-n.ShortestParent.X, 2) + math.Pow(n.Y-n.ShortestParent.Y, 2))
		n = n.ShortestParent
	}
	neededTime := dist / (15 / 3.6)
	busTime := finish.Y / (30 / 3.6)
	startDelay := busTime - neededTime
	startTime := millisToTime(7.5*60*60*1000 + int(startDelay*1000))

	fmt.Printf("Datei:\t\t%v\n\n", args[0])
	fmt.Printf("Startzeit:\t%v\n", startTime)
	fmt.Printf("Zielzeit:\t%v\n", millisToTime(7.5*60*60*1000+int(busTime)*1000))
	fmt.Printf("y-Koordinate:\t%vm\n", int(finish.Y+0.5))
	fmt.Printf("Dauer:\t\t%v\n", toMinLength(int(neededTime/60), 2)+":"+toMinLength(int(neededTime)%60, 2))
	fmt.Printf("Länge:\t\t%vm\n", int(dist+0.5))
	fmt.Printf("x- und y-Koordinaten der Eckpunkte:\n%v", description)

	Img := image.NewRGBA(image.Rect(0, 0, rendering.WIDTH, rendering.HEIGHT))
	draw.Draw(Img, Img.Bounds(), &image.Uniform{rendering.BACKGROUND}, image.ZP, draw.Src)
	Gc := draw2dimg.NewGraphicContext(Img)
	Gc.Scale(1, -1)
	Gc.Translate(rendering.BUSSTOP_RADIUS*1.5, -(float64(rendering.HEIGHT) - rendering.BUSSTOP_RADIUS*1.5))
	rendering.RenderEnvironment(Gc)
	rendering.RenderHome(Gc, Home)
	rendering.RenderObstacles(Gc, Obstacles...)
	rendering.RenderPath(Gc, &g, finish)

	f, _ := os.Create("./lisa_rennt_path.png")
	png.Encode(f, Img)
}

func getNodeInfo(n *lib.Node) string {
	info := n.Info["ID"] + " "
	info += "x: " + ftoa(n.X) + " y: " + ftoa(n.Y)

	return info
}

func testData(g *lib.Graph, f *lib.Node) {
	dist := 0.0
	n := f
	for n.ShortestParent != nil {
		dist += math.Sqrt(math.Pow(n.X-n.ShortestParent.X, 2) + math.Pow(n.Y-n.ShortestParent.Y, 2))
		n = n.ShortestParent
	}
	fmt.Println("Distance:", dist, "m")

	neededTime := dist / (15 / 3.6)
	fmt.Println("Needed time:", neededTime, "s")

	fmt.Println("Bus-reach-Y:", f.Y, "m")

	busTime := f.Y / (30 / 3.6)
	fmt.Println("Bus time:", busTime, "s")

	startDelay := busTime - neededTime
	fmt.Println("Start delay:", startDelay, "s")

	startTime := millisToTime(7.5*60*60*1000 + int(startDelay*1000))
	fmt.Println("Start time:", startTime)
}

func millisToTime(millis int) string {
	hours := millis / (60 * 60 * 1000)
	millis -= hours * 60 * 60 * 1000
	mins := millis / (60 * 1000)
	millis -= mins * (60 * 1000)
	secs := millis / 1000
	millis -= secs * 1000
	if millis >= 500 {
		secs += 1
		if secs >= 60 {
			mins += 1
			secs -= 60
			if mins >= 60 {
				hours += 1
				mins -= 60
			}
		}
	}

	time := toMinLength(hours, 2) + ":" + toMinLength(mins, 2) + ":" + toMinLength(secs, 2) // + "." + toMinLength(millis, 3)

	return time
}

func toMinLength(x int, l int) string {
	var s string = strconv.Itoa(x)
	for len(s) < l {
		s = "0" + s
	}
	return s
}

func ftoa(f float64) string {
	s := strconv.Itoa(int(f + 0.5))
	return s
}
