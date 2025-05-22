package render

import (
	"container/list"

	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
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

type TextureSpecs struct {
	Texture  *sdl.Texture
	Position math.Vector2
	Size     math.Vector2
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
			lifecycle.StopLast()
			return
		}
	}

	if (renderCopies != nil) && (renderCopies.Len() > 0) {
		for e := renderCopies.Front(); e != nil; e = e.Next() {
			specs := e.Value.(TextureSpecs)
			rect := sdl.Rect{
				X: int32(specs.Position.X),
				Y: int32(specs.Position.Y),
				W: int32(specs.Size.X),
				H: int32(specs.Size.Y),
			}
			renderer.Copy(specs.Texture, nil, &rect)
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

func RenderTexture(texture TextureSpecs) {
	if renderCopies == nil {
		renderCopies = list.New()
	}

	renderCopies.PushBack(texture)
}

func GetRenderer() *sdl.Renderer {
	return renderer
}

func Destroy() {
	defer window.Destroy()
	defer renderer.Destroy()

	lifecycle.Kill()
}
