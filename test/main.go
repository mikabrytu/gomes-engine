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
	gomesengine.Init("Version 1.4 - DEVELOP", int32(SCREEN_SIZE.X), int32(SCREEN_SIZE.Y))
	debug.EnableDebug()

	events.AddListener(events.Input, events.INPUT_KEYBOARD_PRESSED_ESCAPE, func(data any) {
		lifecycle.Kill()
	})

	rect := utils.RectSpecs{
		PosX:   (SCREEN_SIZE.X / 2) - 64,
		PosY:   (SCREEN_SIZE.Y / 2) - 64,
		Width:  128,
		Height: 128,
	}

	sprite_1 := render.NewSprite("sprite-1", "test/assets/img/mario.png", rect, render.White, 1)
	sprite_1.Init()
	sprite_1.Enable()

	rect_2 := rect
	rect_2.PosX += 64
	rect_2.PosY += 64

	sprite_2 := render.NewSprite("sprite-2", "test/assets/img/alien2.jpg", rect_2, render.White, 3)
	sprite_2.Init()
	sprite_2.Enable()

	rect_3 := rect_2
	rect_3.PosX += 64
	rect_3.PosY += 64

	sprite_3 := render.NewSprite("sprite-3", "test/assets/img/alien.png", rect_3, render.White, 2)
	sprite_3.Init()
	sprite_3.Enable()

	events.AddListener(events.Input, events.INPUT_KEYBOARD_PRESSED_ENTER, func(data any) {
		sprite_1.Clear()
		sprite_2.Clear()
		sprite_3.Clear()
	})

	gomesengine.Run()
}
