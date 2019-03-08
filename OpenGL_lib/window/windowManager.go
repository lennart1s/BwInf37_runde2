package window

import (
	"github.com/go-gl/glfw/v3.2/glfw"
)

var window *glfw.Window

func Create(width int, height int, fullscreen bool, title string, resizable bool) {
	check(glfw.Init())

	if resizable {
		glfw.WindowHint(glfw.Resizable, glfw.True)
	} else {
		glfw.WindowHint(glfw.Resizable, glfw.False)
	}
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	var err error
	window, err = glfw.CreateWindow(width, height, title, nil, nil)
	check(err)

	if fullscreen {
		monitor := glfw.GetPrimaryMonitor()
		vmode := monitor.GetVideoMode()
		window.SetMonitor(monitor, 0, 0, vmode.Width, vmode.Height, 144)
	}
	window.MakeContextCurrent()
}

func Update() {
	window.SwapBuffers()
	glfw.PollEvents()
	//TODO: ectl. fps-cap, manuell
}

func CloseRequested() bool {
	isRequested := window.ShouldClose()

	if isRequested {
		Close()
	}

	return isRequested
}

func Close() {
	glfw.Terminate()
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
