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
	//palletRight.PosY -= 50

	drawPallets(palletLeft, render.White, "palletLeft")
	drawPallets(palletRight, render.White, "palletRight")
}

func prepareBall(pw int) {
	ball := utils.RectSpecs{
		PosX:   (SCREEN_SIZE.X / 2) - (pw / 2),
		PosY:   (SCREEN_SIZE.Y / 2) - (pw / 2) + 1,
		Width:  pw,
		Height: pw,
	}
	var speed int = 5
	var body physics.RigidBody

	lifecycle.Register(lifecycle.GameObject{
		Start: func() {
			body = physics.RegisterBody(&ball, "ball")
			body.Axis.X = 1
			physics.EnableDynamicCollision(&body)
		},
		Physics: func() {
			ball.PosX += speed * body.Axis.X
			ball.PosY += speed * body.Axis.Y

			physics.ResolveDynamicCollisions(&body, false, false)

			if body.Rect.PosY < 0 {
				body.Axis.Y = 1
			}

			if (body.Rect.PosY + body.Rect.Height) > SCREEN_SIZE.Y {
				body.Axis.Y = -1
			}
		},
		Render: func() {
			render.DrawSimpleShapes(ball, render.White)

			// fps := lifecycle.ShowFPS()
			// fmt.Printf("FPS: %v\n", int(fps))
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
	var body physics.RigidBody
	speed := 5

	lifecycle.Register(lifecycle.GameObject{
		Start: func() {
			body = physics.RegisterBody(&rect, name)

			events.Subscribe(events.INPUT_KEYBOARD_PRESSED_W, func(params ...any) error {
				body.Axis.Y = -1
				return nil
			})

			events.Subscribe(events.INPUT_KEYBOARD_PRESSED_S, func(params ...any) error {
				body.Axis.Y = 1
				return nil
			})

			events.Subscribe(events.INPUT_KEYBOARD_RELEASED_W, func(params ...any) error {
				body.Axis.Y = 0
				return nil
			})

			events.Subscribe(events.INPUT_KEYBOARD_RELEASED_S, func(params ...any) error {
				body.Axis.Y = 0
				return nil
			})
		},
		Physics: func() {
			body.Rect.PosY += body.Axis.Y * speed
		},
		Render: func() {
			render.DrawSimpleShapes(rect, color)
		},
	})
}
