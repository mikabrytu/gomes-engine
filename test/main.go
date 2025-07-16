package main

import (
	"fmt"
	"time"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/render"
	savesystem "github.com/mikabrytu/gomes-engine/systems/save"
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

	instantiate()

	gomesengine.Run()
}

type Score struct {
	Score     int       `json:"score"`
	Timestamp time.Time `json:"timestamp"`
}

func instantiate() {
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
