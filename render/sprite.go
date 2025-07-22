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
}

func NewSprite(name string, path string) *Sprite {
	sprite := &Sprite{
		name: name,
		path: path,
	}

	return sprite
}

func (s *Sprite) Init(specs utils.RectSpecs) {
	var err error
	s.texture, err = img.LoadTexture(renderer, s.path)
	if err != nil {
		panic(err)
	}

	s.copy = CopySpecs{
		Texture: s.texture,
		Rect: sdl.Rect{
			X: int32(specs.PosX),
			Y: int32(specs.PosY),
			W: int32(specs.Width),
			H: int32(specs.Height),
		},
	}
}

func (s *Sprite) Register() {
	AddToRenderer(&s.copy)
}

func (s *Sprite) ClearSprite() {
	s.texture.Destroy()
}
