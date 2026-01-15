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

const REPEAT_EVENT string = "REPEAT_EVENT"

func main() {
	gomesengine.HiGomes()
	gomesengine.Init("Version 1.4", int32(SCREEN_SIZE.X), int32(SCREEN_SIZE.Y))
	debug.EnableDebug()
	lifecycle.SetSmoothStep(0.9)
	render.SetBackgroundColor(render.Pink)

	events.Subscribe(events.Input, events.INPUT_KEYBOARD_PRESSED_ESCAPE, func(data any) {
		lifecycle.Kill()
	})

	rect := utils.RectSpecs{
		PosX:   0,
		PosY:   0,
		Width:  64,
		Height: 64,
	}
	sprite := render.NewSprite(
		"Green",
		"test/assets/img/alien.png",
		rect,
	)

	lifecycle.Register(&lifecycle.GameObject{
		Start: func() {
			sprite.Init()
		},
		Destroy: func() {
			sprite.Clear()
		},
	})

	events.Subscribe(events.Input, events.INPUT_KEYBOARD_PRESSED_SPACE, func(data any) {
		sprite.UpdateImage("test/assets/img/mario.png")
	})

	gomesengine.Run()
}
