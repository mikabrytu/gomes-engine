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
	fontspecs := render.FontSpecs{
		Name: "Font",
		Path: "test/assets/font/freesansbold.ttf",
		Size: 32,
	}

	sprite := render.NewSprite(
		"Green",
		"test/assets/img/alien.png",
		rect,
		render.Red,
	)
	font := render.NewFont(fontspecs, SCREEN_SIZE)

	anchor := render.TopLeft
	offset := math.Vector2{X: 16, Y: 16}
	lifecycle.Register(&lifecycle.GameObject{
		Start: func() {
			sprite.Init()
			font.Init("Texto!", render.White, math.Vector2{X: 0, Y: 0})
		},
		Update: func() {
			rect.PosX += 1
			sprite.UpdateRect(rect)

			font.AlignText(anchor, offset)
		},
		Render: func() {
			render.DrawRect(utils.RectSpecs{PosX: 0, PosY: 128, Width: 256, Height: 32}, render.Green)
		},
		Destroy: func() {
			sprite.Clear()
			font.Clear()
		},
	})

	events.Subscribe(events.Input, events.INPUT_KEYBOARD_PRESSED_SPACE, func(data any) {
		sprite.UpdateImage("test/assets/img/mario.png", render.Blue)

		font.UpdateText("More Text to render...")
		font.UpdateColor(render.Magenta)
		anchor = render.TopCenter
		offset.X = 0
	})

	events.Subscribe(events.Input, events.INPUT_MOUSE_CLICK, func(data any) {
		if sprite.IsEnable() {
			sprite.Disable()
			font.Disable()
		} else {
			sprite.Enable()
			font.Enable()
		}
	})

	gomesengine.Run()
}
