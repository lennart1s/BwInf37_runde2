package graphics

import (
	"BwInf37_runde2/OpenGL_lib/files"

	"github.com/go-gl/gl/v2.1/gl"
)

type ShaderProgram struct {
	ID               uint32
	UniformLocations map[string]Uniform
}

type Uniform struct {
	Name     uint8
	Location int32
	Type     uint32
}

func NewShaderProgram() *ShaderProgram {
	shaderProgram := ShaderProgram{ID: gl.CreateProgram()} // hier id hohlen
	files.LoadShader("./res/GuiShader.glsl", gl.VERTEX_SHADER)

	/*gl.AttachShader(shaderProgram.ID, 0)
	gl.AttachShader(shaderProgram.ID, 1)

	gl.LinkProgram(shaderProgram.ID)

	gl.ValidateProgram(shaderProgram.ID)

	var uniformCount int32
	gl.GetProgramiv(shaderProgram.ID, gl.ACTIVE_UNIFORMS, &uniformCount)
	for i := int32(0); i < uniformCount; i++ {
		u := Uniform{Location: i}
		gl.GetActiveUniform(shaderProgram.ID, uint32(i), 64, nil, nil, &u.Type, &u.Name)
		strName := string(u.Name)
		shaderProgram.UniformLocations[strName] = u
	}*/

	return &shaderProgram
}
