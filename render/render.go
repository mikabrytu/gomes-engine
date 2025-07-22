package render

import (
	"container/list"

	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/utils"
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

func DrawSimpleShapes(shape utils.RectSpecs, color Color) {
	rect := sdl.Rect{
		X: int32(shape.PosX),
		Y: int32(shape.PosY),
		W: int32(shape.Width),
		H: int32(shape.Height),
	}

	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	renderer.DrawRect(&rect)
	renderer.FillRect(&rect)
}

func RenderCopy(copy *CopySpecs) {
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
