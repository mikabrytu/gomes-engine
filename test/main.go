package main

import (
	"fmt"

	gomesengine "github.com/mikabrytu/gomes-engine"

	"github.com/mikabrytu/gomes-engine/debug"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/ui"
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

	press_count := 0
	ui_text := "Spaces pressed:"
	font_specs := ui.FontSpecs{
		Name: "Font",
		Path: "test/assets/font/freesansbold.ttf",
		Size: 32,
	}
	font := ui.NewFont(font_specs, SCREEN_SIZE)
	font.Init(ui_text, render.White, math.Vector2{X: 0, Y: 0})
	font.AlignText(ui.TopCenter, math.Vector2{X: 0, Y: 16})

	// Event Listeners

	events.Subscribe(events.Input, events.INPUT_KEYBOARD_PRESSED_ESCAPE, func(data any) {
		lifecycle.Kill()
	})

	events.Subscribe(events.Input, events.INPUT_MOUSE_CLICK_DOWN, func(data any) {
		click := data.(events.InputMouseClickDownEvent)

		if click.Index.Left == 1 {
			sprite.Disable()
			font.Disable()
			lifecycle.Disable(game_object)
		}

		if click.Index.Right == 1 {
			sprite.Enable()
			font.Enable()
			lifecycle.Enable(game_object)
		}
	})

	events.Subscribe(events.Input, events.INPUT_KEYBOARD_PRESSED_SPACE, func(data any) {
		press_count += 1
		message := ui_text + " " + fmt.Sprint(press_count)
		font.UpdateText(message)
	})

	gomesengine.Run()
}
