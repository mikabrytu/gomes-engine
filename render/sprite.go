package render

import (
	"github.com/mikabrytu/gomes-engine/utils"
)

type Sprite struct {
	name    string
	path    string
	rect    utils.RectSpecs
	color   Color
	enabled bool
	update  bool
}

func NewSprite(name string, path string, rect utils.RectSpecs, color Color) *Sprite {
	sprite := &Sprite{
		name:    name,
		path:    path,
		rect:    rect,
		color:   color,
		enabled: true,
		update:  false,
	}

	return sprite
}

func (s *Sprite) Init() {
	RegisterSprite(s)
}

func (s *Sprite) Clear() {
	ClearSprite(s)
}

func (s *Sprite) Enable() {
	s.enabled = true
}

func (s *Sprite) Disable() {
	s.enabled = false
}

func (s *Sprite) IsEnable() bool {
	return s.enabled
}

func (s *Sprite) UpdateImage(path string, color Color) {
	s.path = path
	s.color = color

	s.Clear()
	s.Init()
}

func (s *Sprite) UpdateColor(color Color) {
	s.color = color
	s.update = true
}

func (s *Sprite) GetColor() Color {
	return s.color
}

func (s *Sprite) UpdateRect(rect utils.RectSpecs) {
	s.rect = rect
	s.update = true
}

func (s *Sprite) GetRect() utils.RectSpecs {
	return s.rect
}
