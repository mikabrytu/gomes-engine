package main

import (
	"fmt"
	"math/rand"
	"time"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/render"
	savesystem "github.com/mikabrytu/gomes-engine/systems/save"
	"github.com/mikabrytu/gomes-engine/ui"
	"github.com/mikabrytu/gomes-engine/utils"

	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
)

var SCREEN_SIZE = math.Vector2{
	X: 800,
	Y: 600,
}

func main() {
	gomesengine.HiGomes()
	gomesengine.Init("Save System", int32(SCREEN_SIZE.X), int32(SCREEN_SIZE.Y))

	lifecycle.SetSmoothStep(0.9)

	skip()

	gomesengine.Run()
}

type Score struct {
	Score     int       `json:"score"`
	Timestamp time.Time `json:"timestamp"`
}

func skip() {
	rect := utils.RectSpecs{
		PosX:   (SCREEN_SIZE.X / 2) - 50,
		PosY:   (SCREEN_SIZE.Y / 2) - 50,
		Width:  100,
		Height: 100,
	}

	var box *lifecycle.GameObject
	box = lifecycle.Register(&lifecycle.GameObject{
		Start: func() {
			events.Subscribe(events.INPUT_MOUSE_CLICK_DOWN, func(params ...any) error {
				lifecycle.Disable(box)
				fmt.Printf("Total Objects Registered: %v\n", lifecycle.GetTotalObjects())
				return nil
			})

			events.Subscribe(events.INPUT_MOUSE_CLICK_UP, func(params ...any) error {
				lifecycle.Enable(box)
				fmt.Printf("Total Objects Registered: %v\n", lifecycle.GetTotalObjects())
				return nil
			})
		},
		Render: func() {
			render.DrawSimpleShapes(rect, render.White)
		},
	})
}

func instantiate() {
	specs := ui.FontSpecs{
		Name: "Temp",
		Path: "test/assets/font/freesansbold.ttf",
		Size: 24,
	}
	fPos := math.Vector2{X: 0, Y: 16}

	font := ui.NewFont(specs, SCREEN_SIZE)
	font.Init("Test", render.White, fPos)
	font.AlignText(ui.TopCenter, fPos)

	fColors := []render.Color{render.Red, render.Green, render.Blue}

	events.Subscribe(events.INPUT_MOUSE_CLICK_DOWN, func(params ...any) error {
		pos := math.Vector2{
			X: params[0].([]any)[0].([]any)[0].(int),
			Y: params[0].([]any)[0].([]any)[1].(int),
		}

		fmt.Printf("%v\n", pos)

		rect := utils.RectSpecs{
			Width:  50,
			Height: 50,
			PosX:   pos.X,
			PosY:   pos.Y,
		}

		lifecycle.Register(&lifecycle.GameObject{
			Start: func() {
				println("Cube started!")
			},
			Render: func() {
				render.DrawSimpleShapes(rect, render.White)
			},
		})

		c := fColors[rand.Intn(len(fColors))]
		font.UpdateColor(c)

		return nil
	})

	events.Subscribe(events.INPUT_KEYBOARD_PRESSED_F, func(params ...any) error {
		println("F")
		return nil
	})

	events.Subscribe(events.INPUT_KEYBOARD_PRESSING_F, func(params ...any) error {
		println("Pressing F")
		return nil
	})
}

func save() {
	path := "test/assets/json/highscore.json"
	timestamp := time.Now()
	data := Score{
		Score:     200,
		Timestamp: timestamp,
	}

	savesystem.Save(data, path)
}

func load() {
	var score Score
	path := "test/assets/json/highscore.json"
	err := savesystem.Load(path, &score)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Score of %v saved at %v", score.Score, score.Timestamp)
}
