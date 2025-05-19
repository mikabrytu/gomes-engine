package main

import (
	"fmt"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
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

	palletLeft := utils.RectSpecs{
		PosX:   off,
		PosY:   (SCREEN_HEIGHT / 2) - (ph / 2),
		Width:  pw,
		Height: ph,
	}
	palletRight := palletLeft
	palletRight.PosX = (SCREEN_WIDTH - (pw + off))

	drawPallets(palletLeft, render.White, "palletLeft")
	drawPallets(palletRight, render.White, "palletRight")

	ball := utils.RectSpecs{
		PosX:   (SCREEN_WIDTH / 2) - (pw / 2),
		PosY:   (SCREEN_HEIGHT / 2) - (pw / 2),
		Width:  pw,
		Height: pw,
	}
	var direction int = 1
	var speed int = 5

	lifecycle.Register(lifecycle.GameObject{
		Start: func() {
			physics.RegisterBody(&ball, "ball")
		},
		Update: func() {
			bodies := physics.GetBodies()

			for e := bodies.Front(); e != nil; e = e.Next() {
				item := physics.RigidBody(e.Value.(physics.RigidBody))

				if item.Name != "ball" {
					continue
				}

				fmt.Printf("Body %s at position: (%v, %v)\n", item.Name, item.Rect.PosX, item.Rect.PosY)
			}
		},
		Physics: func() {
			if (ball.PosX + ball.Width) > SCREEN_WIDTH {
				direction = -1
			}
			if ball.PosX < 0 {
				direction = 1
			}

			ball.PosX += speed * direction
		},
		Render: func() {
			render.DrawSimpleShapes(ball, render.White)
		},
	})
}

func drawPallets(rect utils.RectSpecs, color render.Color, name string) {
	lifecycle.Register(lifecycle.GameObject{
		Start: func() {
			physics.RegisterBody(&rect, name)
		},
		Render: func() {
			render.DrawSimpleShapes(rect, color)
		},
	})
}
