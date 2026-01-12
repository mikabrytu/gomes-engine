package render

import (
	"github.com/mikabrytu/gomes-engine/debug"
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
	texture *sdl.Texture
	rect    *sdl.Rect
}

var window *sdl.Window
var renderer *sdl.Renderer
var copies []CopySpecs
var backgroundColor Color

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

	copies = make([]CopySpecs, 0)
	backgroundColor = Black
}

func Render() {
	// TODO: Check a better place for this
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			if debug.IsEnabled() {
				println("Quit")
			}

			lifecycle.StopRender()
			return
		}
	}

	renderer.SetDrawColor(
		backgroundColor.R,
		backgroundColor.G,
		backgroundColor.B,
		backgroundColor.A,
	)
	renderer.Clear()

	for _, copy := range copies {
		renderer.Copy(
			copy.texture,
			nil,
			copy.rect,
		)
	}

	renderer.Present()
}

func RegisterTexture(texture *sdl.Texture, rect utils.RectSpecs) {
	for _, c := range copies {
		if c.texture == texture {
			println("Texture already registered in copy list")
			return
		}
	}

	copy := CopySpecs{
		texture: texture,
		rect: &sdl.Rect{
			X: int32(rect.PosX),
			Y: int32(rect.PosY),
			W: int32(rect.Width),
			H: int32(rect.Height),
		},
	}
	copies = append(copies, copy)
}

func GetRenderer() *sdl.Renderer {
	return renderer
}

func SetBackgroundColor(color Color) {
	backgroundColor = color
}

func Destroy() {
	defer renderer.Destroy()
	defer window.Destroy()

	lifecycle.Kill()
}
