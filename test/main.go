package main

import (
	"fmt"

	gomesengine "github.com/mikabrytu/gomes-engine"

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
	//debug.EnableDebug()
	lifecycle.SetSmoothStep(0.9)

	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_ESCAPE, func(params ...any) error {
		lifecycle.Kill()
		return nil
	})

	pool()

	gomesengine.Run()
}

func pool() {
	var matrix [][]*lifecycle.GameObject = make([][]*lifecycle.GameObject, 10)

	for i := range 10 {
		matrix[i] = make([]*lifecycle.GameObject, 10)

		for j := range 10 {
			rect := utils.RectSpecs{
				PosX:   i * (32 + 4),
				PosY:   j * (32 + 4),
				Width:  32,
				Height: 32,
			}

			matrix[i][j] = lifecycle.Register(&lifecycle.GameObject{
				Start: func() {
					name := fmt.Sprintf("sprite-%v-%v", i, j)
					s := render.NewSprite(name, "test/assets/img/alien.png")
					s.Init(rect, render.Transparent)

					events.Subscribe(events.INPUT_MOUSE_CLICK_DOWN, func(params ...any) error {
						x := params[0].([]any)[0].([]any)[0].(int)
						y := params[0].([]any)[0].([]any)[1].(int)
						//fmt.Printf("[%v, %v]\n", x, y)

						if (x > rect.PosX && x < (rect.PosX+rect.Width)) && (y > rect.PosY && y < (rect.PosY+rect.Height)) {
							lifecycle.Disable(matrix[i][j])
						}

						return nil
					})
				},
			})
		}
	}

	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_SPACE, func(params ...any) error {
		for i := range 10 {
			for j := range 10 {
				if !matrix[i][j].IsEnable {
					lifecycle.Enable(matrix[i][j])
				}
			}
		}

		return nil
	})
}
