package graphics

import (
	"BwInf37_runde2/OpenGL_lib/files"
	"strings"

	"github.com/go-gl/gl/v2.1/gl"
)

type ShaderProgram struct {
	ID               uint32
	UniformLocations map[string]Uniform
}

type Uniform struct {
	Name     string
	Location uint32
	Type     uint32
}

func NewShaderProgram(path string) *ShaderProgram {
	shaderProgram := ShaderProgram{ID: gl.CreateProgram(),
		UniformLocations: make(map[string]Uniform)}

	gl.AttachShader(shaderProgram.ID, files.LoadShader(path, gl.VERTEX_SHADER))
	gl.AttachShader(shaderProgram.ID, files.LoadShader(path, gl.FRAGMENT_SHADER))

	gl.LinkProgram(shaderProgram.ID)
	gl.ValidateProgram(shaderProgram.ID)

	var uniformCount int32
	gl.GetProgramiv(shaderProgram.ID, gl.ACTIVE_UNIFORMS, &uniformCount)
	for i := uint32(0); i < uint32(uniformCount); i++ {
		u := Uniform{Location: i}
		var a, b int32
		u.Name = strings.Repeat("\x00", int(128+1))
		gl.GetActiveUniform(shaderProgram.ID, i, 128, &a, &b, &u.Type, gl.Str(u.Name))
		shaderProgram.UniformLocations[u.Name] = u
		//println(u.Type, gl.FLOAT_MAT4)
	}

	return &shaderProgram
}
