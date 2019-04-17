package main

import (
	"BwInf37_runde2/dreiecksbeziehungen/files"
	"BwInf37_runde2/dreiecksbeziehungen/updown"
)

func main() {
	triangles := files.Load("./dreiecksbeziehungen/examples/dreiecke1.txt")

	updown.UpDown(triangles)
}
