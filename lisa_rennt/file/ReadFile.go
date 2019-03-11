package file

import (
	"BwInf37_runde2/lisa_rennt/lib"
	"io/ioutil"
	"strconv"
	"strings"
)

func Read(path string) ([]*lib.Polygon, lib.Vertex) {
	var obstacles []*lib.Polygon
	var home lib.Vertex

	lines := loadLines(path)
	numPolygons, err := strconv.Atoi(lines[0])
	for _, line := range lines[1 : 1+numPolygons] {
		obstacles = append(obstacles, lib.PolygonFromLine(line))
	}

	homeParts := strings.Split(lines[len(lines)-1], " ")
	home.X, err = strconv.ParseFloat(homeParts[0], 64)
	home.Y, err = strconv.ParseFloat(homeParts[1], 64)

	if err != nil {
		panic("Invalid data in file: " + path)
		return nil, lib.Vertex{}
	}

	return obstacles, home
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
