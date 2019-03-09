package files

import (
	"io/ioutil"
	"strings"
)

func LoadText(path string) string {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(data)
}

func LoadLines(path string) []string {
	text := LoadText(path)
	var lines []string
	for _, lineA := range strings.Split(text, "\r\n") {
		for _, lineB := range strings.Split(lineA, "\n") {
			lines = append(lines, lineB)
		}
	}
	return lines
}
