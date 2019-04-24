package files

import (
	"BwInf37_runde2/dreiecksbeziehungen/lib"
	"io/ioutil"
	"strconv"
	"strings"
)

func Load(path string) []*lib.Triangle {
	var triangles []*lib.Triangle

	lines := loadLines(path)
	for _, line := range lines[1:] {
		triangles = append(triangles, triangleFromLine(line))
	}

	return triangles
}

func triangleFromLine(line string) *lib.Triangle {
	parts := strings.Split(line, " ")[1:]
	a := vectorFromStrings(parts[0], parts[1])
	b := vectorFromStrings(parts[2], parts[3])
	c := vectorFromStrings(parts[4], parts[5])

	triangle := &lib.Triangle{A: &a, B: &b, C: &c}

	alpha := lib.AddVector(&b, lib.NegateVector(&a)).GetAngle(lib.AddVector(&c, lib.NegateVector(&a)))
	beta := lib.AddVector(&a, lib.NegateVector(&b)).GetAngle(lib.AddVector(&c, lib.NegateVector(&b)))
	gamma := lib.AddVector(&a, lib.NegateVector(&c)).GetAngle(lib.AddVector(&b, lib.NegateVector(&c)))

	triangle.Alpha = alpha
	triangle.Beta = beta
	triangle.Gamma = gamma

	triangle.ArrangeAntiClockwise()

	return triangle
}

func vectorFromStrings(xStr string, yStr string) lib.Vector {
	x, _ := strconv.ParseFloat(xStr, 64)
	y, _ := strconv.ParseFloat(yStr, 64)
	return lib.Vector{x, y}
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
