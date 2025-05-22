package main

import (
	"fmt"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/ui"
	"github.com/mikabrytu/gomes-engine/utils"
)

const SCREEN_WIDTH = 800
const SCREEN_HEIGHT = 600

func main() {
	gomesengine.HiGomes()
	gomesengine.Init("Gong", SCREEN_WIDTH, SCREEN_HEIGHT)

	gong()

	gomesengine.Run()
}

func gong() {
	const pw int = 50
	const ph int = 200
	const off int = 10

	preparePallets(off, pw, ph)
	prepareBall(pw)
	prepareText()
}

func preparePallets(off, pw, ph int) {
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
}

func prepareBall(pw int) {
	ball := utils.RectSpecs{
		PosX:   (SCREEN_WIDTH / 2) - (pw / 2),
		PosY:   (SCREEN_HEIGHT / 2) - (pw / 2),
		Width:  pw,
		Height: pw,
	}
	var direction int = -1
	var speed int = 5

	lifecycle.Register(lifecycle.GameObject{
		Start: func() {
			physics.RegisterBody(&ball, "ball")
		},
		Physics: func() {
			body := physics.GetBodyByName("ball")

			if body.Name != "nil" {
				collider := physics.CheckCollision(body)
				if collider.Name != "nil" {
					if collider.Rect.PosX > ball.PosX {
						direction = -1
					}

					if collider.Rect.PosX < ball.PosX {
						direction = 1
					}
				}
			} else {
				fmt.Printf("Ball is not a rigidbody")
			}

			// Continue movement until end of screen
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

func prepareText() {
	font := ui.FontSpecs{
		Name: "Sans",
		Path: "test/assets/font/freesansbold.ttf",
		Size: 24,
	}

	position := math.Vector2{X: 10, Y: 10}

	ui.LoadFont(font)
	ui.RenderText("Gong", render.Blue, position)
	ui.AlignText(ui.TopCenter, math.Vector2{X: SCREEN_WIDTH, Y: SCREEN_HEIGHT})
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
