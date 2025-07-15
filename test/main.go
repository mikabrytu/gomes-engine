package main

import (
	"fmt"
	"time"

	gomesengine "github.com/mikabrytu/gomes-engine"
	savesystem "github.com/mikabrytu/gomes-engine/systems/save"

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

	//save()
	//load()

	gomesengine.Run()
}

type Score struct {
	Score     int       `json:"score"`
	Timestamp time.Time `json:"timestamp"`
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
