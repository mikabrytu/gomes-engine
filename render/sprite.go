package render

import (
	"github.com/mikabrytu/gomes-engine/utils"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type Sprite struct {
	name    string
	path    string
	rect    utils.RectSpecs
	texture *sdl.Texture
	enabled bool
}

func NewSprite(name string, path string, rect utils.RectSpecs) *Sprite {
	sprite := &Sprite{
		name:    name,
		path:    path,
		rect:    rect,
		enabled: true,
	}

	return sprite
}

func (s *Sprite) Init() {
	surface, err := img.Load(s.path)
	if err != nil {
		panic(err)
	}

	s.texture, err = renderer.CreateTextureFromSurface(surface)
	if err != nil {
		panic(err)
	}
	surface.Free()

	RegisterTexture(s.texture, s.rect)
}

// func (s *Sprite) Enable() {
// 	s.enabled = true
// }

// func (s *Sprite) Disable() {
// 	s.enabled = false
// }

// func (s *Sprite) IsEnabled() bool {
// 	return s.enabled
// }

func (s *Sprite) Clear() {
	err := s.texture.Destroy()
	if err != nil {
		panic(err)
	}
}
