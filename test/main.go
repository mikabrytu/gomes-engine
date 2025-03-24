package main

import (
	"fmt"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
)

func main() {
	gomesengine.HiGomes()
	gomesengine.Init("Genius", 430, 430)

	genius()

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
		},
		Update: func() {
			render.DrawSimpleShapes(rect, color)
		},
	})
}
