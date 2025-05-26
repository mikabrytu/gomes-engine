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
	text     string
	position math.Vector2
	copy     render.CopySpecs
	color    render.Color
	screen   math.Vector2
	update   bool
}

// Public API

func NewFont(specs FontSpecs, screenSize math.Vector2) *Font {
	font := &Font{
		screen: screenSize,
		update: false,
	}

	var err error
	font.instance, err = ttf.OpenFont(specs.Path, specs.Size)
	if err != nil {
		panic(err)
	}

	return font
}

func (f *Font) Init(text string, color render.Color, position math.Vector2) {
	f.text = text
	f.color = color
	f.position = position
	f.prepareRender()

	lifecycle.Register(lifecycle.GameObject{
		Update: func() {
			if f.update {
				f.update = false
				f.texture.Destroy()
				f.surface.Free()
				f.prepareRender()
			}
		},
		Render: func() {
			render.RenderCopy(f.copy)
		},
		Destroy: func() {
			f.ClearFont()
		},
	})
}

func (f *Font) UpdateText(text string) {
	f.text = text
	f.update = true
}

func (f *Font) AlignText(anchor Anchor, offset math.Vector2) {
	switch anchor {
	case TopLeft:
		f.copy.Rect.X = int32(0 + offset.X)
		f.copy.Rect.Y = int32(0 + offset.Y)
	case TopCenter:
		f.copy.Rect.X = ((int32(f.screen.X) / 2) - (f.surface.W / 2)) + int32(offset.X)
		f.copy.Rect.Y = int32(0 + offset.Y)
	case TopRight:
		f.copy.Rect.X = int32(f.screen.X) - (f.surface.W + int32(offset.X))
		f.copy.Rect.Y = int32(0 + offset.Y)
	case MiddleLeft:
		f.copy.Rect.X = int32(0 + offset.X)
		f.copy.Rect.Y = ((int32(f.screen.Y) / 2) - (f.surface.H / 2)) + int32(offset.Y)
	case MiddleCenter:
		f.copy.Rect.X = ((int32(f.screen.X) / 2) - (f.surface.W / 2)) + int32(offset.X)
		f.copy.Rect.Y = ((int32(f.screen.Y) / 2) - (f.surface.H / 2)) + int32(offset.Y)
	case MiddleRight:
		f.copy.Rect.X = int32(f.screen.X) - (f.surface.W + int32(offset.X))
		f.copy.Rect.Y = ((int32(f.screen.Y) / 2) - (f.surface.H / 2)) + int32(offset.Y)
	case BottomLeft:
		f.copy.Rect.X = int32(0 + offset.X)
		f.copy.Rect.Y = int32(f.screen.Y) - (f.surface.H + int32(offset.Y))
	case BottomCenter:
		f.copy.Rect.X = ((int32(f.screen.X) / 2) - (f.surface.W / 2)) + int32(offset.X)
		f.copy.Rect.Y = int32(f.screen.Y) - (f.surface.H + int32(offset.Y))
	case BottomRight:
		f.copy.Rect.X = int32(f.screen.X) - (f.surface.W + int32(offset.X))
		f.copy.Rect.Y = int32(f.screen.Y) - (f.surface.H + int32(offset.Y))
	}

	f.position = math.Vector2{
		X: int(f.copy.Rect.X),
		Y: int(f.copy.Rect.Y),
	}
}

func (f *Font) ClearFont() {
	f.texture.Destroy()
	f.surface.Free()
	//font.Close() // TODO: This is causing a crash when closing the game
}

// Private Implementation

func (f *Font) prepareRender() {
	var err error
	csdl := colorToSDL(f.color)

	f.surface, err = f.instance.RenderUTF8Blended(f.text, csdl)
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
			X: int32(f.position.X),
			Y: int32(f.position.Y),
			W: int32(f.surface.W),
			H: int32(f.surface.H),
		},
	}
}

func colorToSDL(color render.Color) sdl.Color {
	return sdl.Color{
		R: color.R,
		G: color.G,
		B: color.B,
		A: color.A,
	}
}
