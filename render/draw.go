package render

import (
	"github.com/mikabrytu/gomes-engine/utils"
	"github.com/veandco/go-sdl2/sdl"
)

func DrawRect(shape utils.RectSpecs, color Color) {
	rect := sdl.Rect{
		X: int32(shape.PosX),
		Y: int32(shape.PosY),
		W: int32(shape.Width),
		H: int32(shape.Height),
	}

	renderer.SetDrawColor(color.R, color.G, color.B, color.A)
	renderer.DrawRect(&rect)
	renderer.FillRect(&rect)
}
