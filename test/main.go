package main

import (
	"fmt"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/audio"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
)

func main() {
	gomesengine.HiGomes()
	gomesengine.Init("Genius", 430, 430)

	genius()
	keyboard()

	gomesengine.Run()
}

func genius() {
	var size int32 = 200
	var offset int32 = 10

	rect := render.RectSpecs{
		PosX:   offset,
		PosY:   offset,
		Width:  size,
		Height: size,
	}
	redRect := rect
	greenRect := rect
	blueRect := rect
	yellowRect := rect

	greenRect.PosX += size + offset
	blueRect.PosY += size + offset
	yellowRect.PosX += size + offset
	yellowRect.PosY += size + offset

	drawRect(redRect, render.Red, "RED")
	drawRect(greenRect, render.Green, "GREEN")
	drawRect(blueRect, render.Blue, "BLUE")
	drawRect(yellowRect, render.Yellow, "YELLOW")
}

func drawRect(rect render.RectSpecs, color render.Color, message string) {
	lifecycle.Register(lifecycle.Loopable{
		Init: func() {
			fmt.Printf("Initializing %v\n", message)
			clicks(rect, message)
		},
		Update: func() {
			render.DrawSimpleShapes(rect, color)
		},
	})
}

func clicks(rect render.RectSpecs, name string) {
	events.Subscribe(events.INPUT_MOUSE_CLICK_DOWN, func(params ...any) error {
		positions := params[0].([]any)
		posx := positions[0].([]any)[0].(int32)
		posy := positions[0].([]any)[1].(int32)

		if posx >= rect.PosX && posx <= (rect.PosX+rect.Width) && posy >= rect.PosY && posy <= (rect.PosY+rect.Height) {
			fmt.Printf("Clicked on %v\n", name)
			audio.Play("test/assets/audio/Go.ogg")
		}

		return nil
	})
}

func keyboard() {
	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_W, func(params ...any) error {
		println("Button pressed!")
		return nil
	})

	events.Subscribe(events.INPUT_KEYBOARD_RELEASED_W, func(params ...any) error {
		println("Button released!")
		return nil
	})
}
