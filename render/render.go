package render

import (
	"container/list"

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

type CopySpecs struct {
	Texture *sdl.Texture
	Rect    sdl.Rect
}

var window *sdl.Window
var renderer *sdl.Renderer
var renderCopies *list.List

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
			lifecycle.StopRender()
			return
		}
	}

	// The copies should render before the present/clear section, otherwise it won't render.
	if (renderCopies != nil) && (renderCopies.Len() > 0) {
		for e := renderCopies.Front(); e != nil; e = e.Next() {
			specs := e.Value.(*CopySpecs)
			renderer.Copy(specs.Texture, nil, &specs.Rect)
		}
	}

	renderer.Present()
	renderer.SetDrawColor(Black.R, Black.G, Black.B, Black.A)
	renderer.Clear()
}

func AddToRenderer(copy *CopySpecs) {
	if renderCopies == nil {
		renderCopies = list.New()
	}

	for e := renderCopies.Front(); e != nil; e = e.Next() {
		if e.Value.(*CopySpecs) == copy {
			return
		}
	}

	renderCopies.PushBack(copy)
}

func GetRenderer() *sdl.Renderer {
	return renderer
}

func Destroy() {
	renderCopies = list.New()

	defer window.Destroy()
	defer renderer.Destroy()

	lifecycle.Kill()
}
