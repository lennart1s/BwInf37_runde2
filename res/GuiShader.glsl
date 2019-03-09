//#vertexShader
#version 330

layout(location = 0) in vec2 position;

out vec2 textureCoords;

uniform mat4 transformationMatrix;

void main() {
  gl_Position = transformationMatrix * vec4(position, 0.0, 1.0);
  textureCoords = vec2((position.x+1.0)/2.0, 1 - (position.y+1.0)/2.0);
}
//#end

//#fragmentShader
#version 330

in vec2 textureCoords;

layout(location = 0) out vec4 color;

uniform sampler2D guiTexture;

void main() {
  color = texture(guiTexture, textureCoords);
}
//#end