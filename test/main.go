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

	circle()

	gomesengine.Run()
}

func circle() {
	circle := utils.CircleSpecs{
		PosX:   SCREEN_SIZE.X / 2,
		PosY:   SCREEN_SIZE.Y / 2,
		Radius: 32,
	}

	lifecycle.Register(&lifecycle.GameObject{
		Update: func() {
			circle.PosX += 1
			circle.PosY += 1
		},
		Render: func() {
			render.DrawCircle(circle, render.White)
		},
	})
}
