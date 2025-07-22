package main

import (
	gomesengine "github.com/mikabrytu/gomes-engine"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

var SCREEN_SIZE = math.Vector2{
	X: 800,
	Y: 600,
}

func main() {
	gomesengine.HiGomes()
	gomesengine.Init("Save System", int32(SCREEN_SIZE.X), int32(SCREEN_SIZE.Y))

	lifecycle.SetSmoothStep(0.9)
	draw()

	gomesengine.Run()
}

func draw() {
	specs := utils.RectSpecs{
		PosX:   0,
		PosY:   0,
		Width:  72,
		Height: 64,
	}

	sprite := render.NewSprite("alien", "test/assets/img/alien.png")
	sprite.Init(specs)

	lifecycle.Register(&lifecycle.GameObject{
		Start: func() {
			events.Subscribe(events.INPUT_MOUSE_CLICK, func(params ...any) error {
				sprite.UpdateImage("test/assets/img/alien2.jpg")
				return nil
			})
		},
		Update: func() {
			rect := sprite.GetRect()
			rect.PosX += 1

			sprite.UpdateRect(rect)
		},
		Destroy: func() {
			sprite.ClearSprite()
		},
	})
}
