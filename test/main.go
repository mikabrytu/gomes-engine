package main

import (
	"container/list"
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

	save()
	load()

	gomesengine.Run()
}

type Score struct {
	Score     int       `json:"score"`
	Timestamp time.Time `json:"timestamp"`
}

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func save() {
	timestamp := time.Now()
	data := Score{
		Score:     100,
		Timestamp: timestamp,
	}

	savesystem.Save(data)
}

func load() {
	var users Users
	path := "test/assets/json/users.json"
	err := savesystem.Load(path, &users)

	if err != nil {
		panic(err)
	}

	names := list.New()
	for i := 0; i < len(users.Users); i++ {
		names.PushBack(users.Users[i].Name)
	}

	for e := names.Front(); e != nil; e = e.Next() {
		name := e.Value.(string)
		println(name)
	}
}
