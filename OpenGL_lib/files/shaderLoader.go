package files

import (
	"errors"
	"log"
	"strings"

	"github.com/go-gl/gl/v2.1/gl"
)

const (
	END_FLAG             = "//#end"
	VERTEX_SHADER_FLAG   = "//#vertexShader"
	FRAGMENT_SHADER_FLAG = "//#fragmentShader"
)

func LoadShader(path string, shaderType uint32) uint32 {
	lines := LoadLines(path)
	var source []string
	inShaderSource := false
	for _, line := range lines {
		if inShaderSource && strings.HasPrefix(line, END_FLAG) {
			source = append(source, "\x00")
			break
		} else if inShaderSource {
			source = append(source, line)
		} else if (shaderType == gl.VERTEX_SHADER && strings.HasPrefix(line, VERTEX_SHADER_FLAG)) ||
			(shaderType == gl.FRAGMENT_SHADER && strings.HasPrefix(line, FRAGMENT_SHADER_FLAG)) {
			inShaderSource = true
		}
	}

	id := gl.CreateShader(shaderType)
	csource, free := gl.Strs(strings.Join(source, "\n"))
	gl.ShaderSource(id, 1, csource, nil)
	free()
	if shaderType == gl.VERTEX_SHADER {
		log.Println("Compiling vertex-shader:", path)
	} else if shaderType == gl.FRAGMENT_SHADER {
		log.Println("Compiling fragment-shader:", path)
	}
	err, info := compileShader(id)
	if err != nil {
		log.Println(err)
		log.Fatal(info)
	}

	return id
}

func compileShader(id uint32) (error, string) {
	gl.CompileShader(id)
	var status int32
	gl.GetShaderiv(id, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(id, gl.INFO_LOG_LENGTH, &logLength)
		infoLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(id, logLength, nil, gl.Str(infoLog))
		return errors.New("Could not compile shader"), infoLog
	}

	return nil, ""
}
