package render

import (
	"github.com/mikabrytu/gomes-engine/utils"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	name    string
	path    string
	texture *sdl.Texture
	copy    CopySpecs
	rect    utils.RectSpecs
	color   Color
}

func NewSprite(name string, path string) *Sprite {
	sprite := &Sprite{
		name: name,
		path: path,
	}

	return sprite
}

// It prepares the necessary render dependencies and register the texture in the render loop.
// Color is an optional tint rgb value that can be multiplied to the texture.
// Use any transparent color to ignore this step
func (s *Sprite) Init(specs utils.RectSpecs, color Color) {
	s.rect = specs
	s.color = color

	s.newTexture()
	s.newRect()

	AddToRenderer(&s.copy)
}

func (s *Sprite) UpdateRect(rect utils.RectSpecs) {
	s.rect = rect
	s.newRect()
}

func (s *Sprite) UpdateColor(color Color) {
	s.color = color
	s.newRect()
}

func (s *Sprite) UpdateImage(path string) {
	s.path = path
	s.newTexture()
}

func (s *Sprite) GetRect() utils.RectSpecs {
	return s.rect
}

func (s *Sprite) ClearSprite() {
	s.texture.Destroy()
}

func (s *Sprite) newTexture() {
	var err error
	s.texture, err = img.LoadTexture(renderer, s.path)
	if err != nil {
		panic(err)
	}
}

func (s *Sprite) newRect() {
	s.copy = CopySpecs{
		Texture: s.texture,
		Rect: sdl.Rect{
			X: int32(s.rect.PosX),
			Y: int32(s.rect.PosY),
			W: int32(s.rect.Width),
			H: int32(s.rect.Height),
		},
		Color:  s.color,
		Update: true,
	}
}
