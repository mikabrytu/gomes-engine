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

type Font struct {
	instance *ttf.Font
	surface  *sdl.Surface
	texture  *sdl.Texture
	copy     render.CopySpecs
}

func NewFont(specs FontSpecs) *Font {
	font := &Font{}

	var err error
	font.instance, err = ttf.OpenFont(specs.Path, specs.Size)
	if err != nil {
		panic(err)
	}

	return font
}

func (f *Font) RenderText(text string, color render.Color, position math.Vector2) {
	var err error
	csdl := sdl.Color{
		R: color.R,
		G: color.G,
		B: color.B,
		A: color.A,
	}

	f.surface, err = f.instance.RenderUTF8Blended(text, csdl)
	if err != nil {
		panic(err)
	}

	f.texture, err = render.GetRenderer().CreateTextureFromSurface(f.surface)
	if err != nil {
		panic(err)
	}

	f.copy = render.CopySpecs{
		Texture: f.texture,
		Rect: sdl.Rect{
			X: int32(position.X),
			Y: int32(position.Y),
			W: int32(f.surface.W),
			H: int32(f.surface.H),
		},
	}

	lifecycle.Register(lifecycle.GameObject{
		Render: func() {
			render.RenderCopy(f.copy)
		},
		Destroy: func() {
			f.ClearFont()
		},
	})
}

func (f *Font) AlignText(anchor Anchor, screen math.Vector2) {
	f.copy.Rect.X = int32(screen.X)/2 - (f.copy.Rect.W / 2)
}

func (f *Font) ClearFont() {
	f.texture.Destroy()
	f.surface.Free()
	//font.Close() // TODO: This is causing a crash when closing the game
}
