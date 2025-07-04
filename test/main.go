package main

import (
	"fmt"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/ui"
	"github.com/mikabrytu/gomes-engine/utils"
)

var SCREEN_SIZE = math.Vector2{
	X: 800,
	Y: 600,
}

func main() {
	gomesengine.HiGomes()
	gomesengine.Init("Gong", int32(SCREEN_SIZE.X), int32(SCREEN_SIZE.Y))
	lifecycle.SetSmoothStep(0.9)

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
		PosY:   (SCREEN_SIZE.Y / 2) - (ph / 2),
		Width:  pw,
		Height: ph,
	}
	palletRight := palletLeft
	palletRight.PosX = (SCREEN_SIZE.X - (pw + off))

	drawPallets(palletLeft, render.White, "palletLeft")
	drawPallets(palletRight, render.White, "palletRight")
}

func prepareBall(pw int) {
	ball := utils.RectSpecs{
		PosX:   (SCREEN_SIZE.X / 2) - (pw / 2),
		PosY:   (SCREEN_SIZE.Y / 2) - (pw / 2),
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
			if (ball.PosX + ball.Width) > SCREEN_SIZE.X {
				direction = -1
			}
			if ball.PosX < 0 {
				direction = 1
			}

			ball.PosX += speed * direction
		},
		Render: func() {
			render.DrawSimpleShapes(ball, render.White)

			fps := lifecycle.ShowFPS()
			fmt.Printf("FPS: %v\n", int(fps))
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
	offset := math.Vector2{X: -64, Y: 10}

	p1Score := ui.NewFont(font, SCREEN_SIZE)
	p1Score.Init("Player 1", render.White, position)
	p1Score.AlignText(ui.BottomCenter, offset)

	position.Y = SCREEN_SIZE.Y - 10
	p2Score := ui.NewFont(font, SCREEN_SIZE)
	p2Score.Init("Player 2", render.White, position)
	p2Score.AlignText(ui.TopCenter, offset)

	count := 0
	events.Subscribe(events.INPUT_MOUSE_CLICK, func(params ...any) error {
		count++
		text := "Player " + fmt.Sprint(count)
		p1Score.UpdateText(text)

		fmt.Printf("Updated Count: %v\n", count)

		return nil
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
