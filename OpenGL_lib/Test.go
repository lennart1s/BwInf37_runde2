package main

import (
	"io"
	"log"
	"os"
	"runtime"

	"github.com/go-gl/gl/v2.1/gl"

	"BwInf37_runde2/OpenGL_lib/graphics"
	"BwInf37_runde2/OpenGL_lib/window"
)

var (
	windowWidht  = 800
	windowHeight = 600
	title        = "OpenGL-Window"
	fullscreen   = false
)

func main() {
	initLog()
	log.Println("Starting up..")
	runtime.LockOSThread()

	window.Create(windowWidht, windowHeight, fullscreen, title, false)
	log.Println("Initialized window.")

	initOpenGL()
	log.Println("Initialized OpenGL.")

	/*prog := gl.CreateProgram()
	gl.LinkProgram(prog)
	return prog*/
	graphics.NewShaderProgram("./res/GuiShader.glsl")

	log.Println("Starting window-loop...")
	for !window.CloseRequested() {

		window.Update()
	}
	log.Println("Window closed.")
	log.Print("Stopping program... \n\n\n")
}

func initOpenGL() {
	check(gl.Init())
	log.Println("OpenGL version:", gl.GoStr(gl.GetString(gl.VERSION)))
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
