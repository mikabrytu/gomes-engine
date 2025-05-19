package main

import (
	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
)

const SCREEN_WIDTH = 800
const SCREEN_HEIGHT = 600

func main() {
	gomesengine.HiGomes()
	gomesengine.Init("Genius", SCREEN_WIDTH, SCREEN_HEIGHT)

	gong()

	gomesengine.Run()
}

func gong() {
	const pw int = 50
	const ph int = 200
	const off int = 10

	palletLeft := render.RectSpecs{
		PosX:   off,
		PosY:   (SCREEN_HEIGHT / 2) - (ph / 2),
		Width:  pw,
		Height: ph,
	}
	palletRight := palletLeft
	palletRight.PosX = (SCREEN_WIDTH - (pw + off))

	drawPallets(palletLeft, render.White)
	drawPallets(palletRight, render.White)

	ball := render.RectSpecs{
		PosX:   (SCREEN_WIDTH / 2) - (pw / 2),
		PosY:   (SCREEN_HEIGHT / 2) - (pw / 2),
		Width:  pw,
		Height: pw,
	}

	lifecycle.Register(lifecycle.GameObject{
		Update: func() {
			render.DrawSimpleShapes(ball, render.White)
		},
	})
}

func drawPallets(rect render.RectSpecs, color render.Color) {
	lifecycle.Register(lifecycle.GameObject{
		Update: func() {
			render.DrawSimpleShapes(rect, color)
		},
	})
}
