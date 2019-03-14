package main

import "BwInf37_runde2/lisa_rennt/file"

func main2() {
	polygons, _ := file.Read("./lisa_rennt/examples/cw_cv.txt")

	println(polygons[0].IsClockwise())

	polygons[0].ArrangeClockwise()

	println(polygons[0].IsClockwise())
}
