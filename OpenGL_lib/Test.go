package main

import (
	"io"
	"log"
	"os"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

var (
	windowWidht  = 800
	windowHeight = 600
	title        = "OpenGl-Window"
	fullscreen   = false
)

func main() {
	initLog()
	log.Println("Starting up..")
	runtime.LockOSThread()

	window := initWindow()
	log.Println("Initialized window.")
	defer glfw.Terminate()

	initOpenGL()
	log.Println("Initialized OpenGL.")

	log.Println("Starting window-loop...")
	for !window.ShouldClose() {

		window.SwapBuffers()
		glfw.PollEvents()
	}
	log.Println("Window closed.")
	log.Print("Stopping program... \n\n\n")
}

func initWindow() *glfw.Window {
	check(glfw.Init())

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(windowWidht, windowHeight, title, nil, nil)
	check(err)
	if fullscreen {
		monitor := glfw.GetPrimaryMonitor()
		vmode := monitor.GetVideoMode()
		window.SetMonitor(monitor, 0, 0, vmode.Width, vmode.Height, 40)
	}
	window.MakeContextCurrent()

	log.Println("Window-size:", windowWidht, "x", windowHeight, " Fullscreen:", fullscreen)

	return window
}

func initOpenGL() uint32 {
	check(gl.Init())
	log.Println("OpenGL version:", gl.GoStr(gl.GetString(gl.VERSION)))

	prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog
}

func initLog() {
	f, err := os.OpenFile("./log", os.O_RDWR|os.O_CREATE /*|os.O_APPEND*/, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
