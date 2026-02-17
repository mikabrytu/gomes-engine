package render

import (
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/utils"
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
	text     string
	position math.Vector2
	rect     utils.RectSpecs
	color    Color
	screen   math.Vector2
	update   bool
	enabled  bool
}

func NewFont(specs FontSpecs, screenSize math.Vector2) *Font {
	font := &Font{
		screen:  screenSize,
		update:  false,
		enabled: true,
	}

	var err error
	font.instance, err = ttf.OpenFont(specs.Path, specs.Size)
	if err != nil {
		panic(err)
	}

	return font
}

func (f *Font) Init(text string, color Color, position math.Vector2) {
	f.text = text
	f.color = color
	f.position = position

	RegisterFont(f)
}

func (f *Font) Enable() {
	f.enabled = true
}

func (f *Font) Disable() {
	f.enabled = false
}

func (f *Font) IsEnable() bool {
	return f.enabled
}

func (f *Font) UpdateText(text string) {
	f.Reset()
	f.Init(text, f.color, f.position)
}

func (f *Font) GetText() string {
	return f.text
}

func (f *Font) UpdateColor(color Color) {
	f.color = color
	f.update = true
}

func (f *Font) GetColor() Color {
	return f.color
}

func (f *Font) UpdatePosition(position math.Vector2) {
	f.position = position
	f.update = true
}

func (f *Font) GetPosition() math.Vector2 {
	return f.position
}

func (f *Font) GetSize() math.Vector2 {
	return math.Vector2{
		f.rect.Width,
		f.rect.Height,
	}
}

func (f *Font) AlignText(anchor Anchor, offset math.Vector2) {
	switch anchor {
	case TopLeft:
		f.rect.PosX = 0 + offset.X
		f.rect.PosY = 0 + offset.Y
	case TopCenter:
		f.rect.PosX = ((f.screen.X / 2) - (f.rect.Width / 2)) + offset.X
		f.rect.PosY = 0 + offset.Y
	case TopRight:
		f.rect.PosX = f.screen.X - (f.rect.Width + offset.X)
		f.rect.PosY = 0 + offset.Y
	case MiddleLeft:
		f.rect.PosX = 0 + offset.X
		f.rect.PosY = ((f.screen.Y / 2) - (f.rect.Height / 2)) + offset.Y
	case MiddleCenter:
		f.rect.PosX = ((f.screen.X / 2) - (f.rect.Width / 2)) + offset.X
		f.rect.PosY = ((f.screen.Y / 2) - (f.rect.Height / 2)) + offset.Y
	case MiddleRight:
		f.rect.PosX = f.screen.X - (f.rect.Width + offset.X)
		f.rect.PosY = ((f.screen.Y / 2) - (f.rect.Height / 2)) + offset.Y
	case BottomLeft:
		f.rect.PosX = 0 + offset.X
		f.rect.PosY = f.screen.Y - (f.rect.Height + offset.Y)
	case BottomCenter:
		f.rect.PosX = ((f.screen.X / 2) - (f.rect.Width / 2)) + offset.X
		f.rect.PosY = f.screen.Y - (f.rect.Height + offset.Y)
	case BottomRight:
		f.rect.PosX = f.screen.X - (f.rect.Width + offset.X)
		f.rect.PosY = f.screen.Y - (f.rect.Height + offset.Y)
	}

	f.UpdatePosition(math.Vector2{
		X: f.rect.PosX,
		Y: f.rect.PosY,
	})
}

func (f *Font) Reset() {
	ClearFont(f)
}

func (f *Font) Clear() {
	f.instance.Close()
	ClearFont(f)
}
