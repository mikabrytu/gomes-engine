package main

import (
	"fmt"

	gomesengine "github.com/mikabrytu/gomes-engine"

	"github.com/mikabrytu/gomes-engine/debug"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
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

	events.Subscribe(events.Input, events.INPUT_MOUSE_CLICK, func(data any) {
		click := data.(events.InputMouseClickEvent)

		index := ""
		if click.Index.Left == 1 {
			index = "left"
		}
		if click.Index.Right == 1 {
			index = "right"
		}
		if click.Index.Middle == 1 {
			index = "middle"
		}

		message := fmt.Sprintf("Clicked at position {%d, %d} with button %v\n", click.Position.X, click.Position.Y, index)
		print(message)
	})

	gomesengine.Run()
}
