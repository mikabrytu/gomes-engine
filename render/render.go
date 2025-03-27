package render

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/veandco/go-sdl2/sdl"
)

type ScreenSpecs struct {
	Title  string
	Posx   int32
	Posy   int32
	Width  int32
	Height int32
}

var window *sdl.Window
var renderer *sdl.Renderer

func CreateScreen(s ScreenSpecs) {
	var err error

	window, err = sdl.CreateWindow(s.Title, s.Posx, s.Posy, s.Width, s.Height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}
}

func Render() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			lifecycle.StopById(0)
			return
		}
	}

	renderer.Present()
	renderer.SetDrawColor(Black.R, Black.G, Black.B, Black.A)
	renderer.Clear()
}

func DrawSimpleShapes(shape RectSpecs, color Color) {
	rect := sdl.Rect{
		X: shape.PosX,
		Y: shape.PosY,
		W: shape.Width,
		H: shape.Height,
	}

	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	renderer.DrawRect(&rect)
	renderer.FillRect(&rect)
}

func Destroy() {
	defer window.Destroy()
	defer renderer.Destroy()

	lifecycle.Kill()
}
