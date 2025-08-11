package main

import (
	gomesengine "github.com/mikabrytu/gomes-engine"

	"github.com/mikabrytu/gomes-engine/debug"
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
	gomesengine.Init("Version 1.4", int32(SCREEN_SIZE.X), int32(SCREEN_SIZE.Y))
	debug.EnableDebug()
	lifecycle.SetSmoothStep(0.9)

	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_ESCAPE, func(params ...any) error {
		lifecycle.Kill()
		return nil
	})

	tint()

	gomesengine.Run()
}

func tint() {
	lifecycle.Register(&lifecycle.GameObject{
		Start: func() {
			specs := utils.RectSpecs{
				PosX:   0,
				PosY:   0,
				Width:  128,
				Height: 128,
			}

			s1 := render.NewSprite("sprite1", "test/assets/img/square.png")
			s1.Init(specs, render.Transparent)

			specs.PosX = 128

			s2 := render.NewSprite("sprite2", "test/assets/img/square.png")
			s2.Init(specs, render.Yellow)

			events.Subscribe(events.INPUT_MOUSE_CLICK, func(params ...any) error {
				println("Clicked")

				specs.PosY = 128
				s2.UpdateRect(specs)
				s2.UpdateColor(render.Red)

				return nil
			})
		},
	})
}
