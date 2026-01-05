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

	path := "test/assets/img/alien.png"
	rect := utils.RectSpecs{
		PosX:   (SCREEN_SIZE.X / 2) - 32,
		PosY:   (SCREEN_SIZE.Y / 2) - 32,
		Width:  64,
		Height: 64,
	}
	sprite := render.NewSprite("Sprite", path, rect, render.White)
	game_object := lifecycle.Register(&lifecycle.GameObject{
		Start: func() {
			sprite.Init()
		},
		Render: func() {
			render.DrawRect(rect, render.White)
		},
	})

	// Event Listeners

	events.Subscribe(events.Input, events.INPUT_KEYBOARD_PRESSED_ESCAPE, func(data any) {
		lifecycle.Kill()
	})

	events.Subscribe(events.Input, events.INPUT_MOUSE_CLICK_DOWN, func(data any) {
		click := data.(events.InputMouseClickDownEvent)

		if click.Index.Left == 1 {
			sprite.Disable()
			lifecycle.Disable(game_object)
		}

		if click.Index.Right == 1 {
			sprite.Enable()
			lifecycle.Enable(game_object)
		}
	})

	gomesengine.Run()
}
