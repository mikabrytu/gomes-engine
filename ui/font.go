package ui

import (
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

type Alignment int

const (
	TopLeft Alignment = iota
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
var specs render.TextureSpecs

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

	specs = render.TextureSpecs{
		Texture:  texture,
		Position: position,
		Size: math.Vector2{
			X: int(surface.W),
			Y: int(surface.H),
		},
	}

	render.RenderTexture(specs)
}

func ClearFont() {
	texture.Destroy()
	surface.Free()
	font.Close()
}
