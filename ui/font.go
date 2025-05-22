package ui

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type FontSpecs struct {
	Name string
	Path string
	Size int
}

type Anchor int

const (
	TopLeft Anchor = iota
	TopRight
	TopCenter
	MiddleLeft
	MiddleRight
	MiddleCenter
	BottomLeft
	BottomRight
	BottomCenter
)

var font *ttf.Font
var surface *sdl.Surface
var texture *sdl.Texture
var copy render.CopySpecs

func LoadFont(specs FontSpecs) {
	var err error
	font, err = ttf.OpenFont(specs.Path, specs.Size)
	if err != nil {
		panic(err)
	}
}

func RenderText(text string, color render.Color, position math.Vector2) {
	var err error
	csdl := sdl.Color{
		R: color.R,
		G: color.G,
		B: color.B,
		A: color.A,
	}

	surface, err = font.RenderUTF8Blended(text, csdl)
	if err != nil {
		panic(err)
	}

	texture, err = render.GetRenderer().CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}

	copy = render.CopySpecs{
		Texture: texture,
		Rect: sdl.Rect{
			X: int32(position.X),
			Y: int32(position.Y),
			W: int32(surface.W),
			H: int32(surface.H),
		},
	}

	lifecycle.Register(lifecycle.GameObject{
		Render: func() {
			render.RenderCopy(copy)
		},
	})
}

func AlignText(anchor Anchor, screen math.Vector2) {
	copy.Rect.X = int32(screen.X)/2 - (copy.Rect.W / 2)
}

func ClearFont() {
	texture.Destroy()
	surface.Free()
	//font.Close() // TODO: This is causing a crash when closing the game
}
