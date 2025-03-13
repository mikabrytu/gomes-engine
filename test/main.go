package main

import (
	"fmt"
	"time"

	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/screen"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	gomesengine.HiGomes()

	//testScreen()
	//testLifecycle()
}

func testScreen() {
	lifecycle.Init()

	specs := screen.Specs{
		Title:  "Gomes Demo",
		Posx:   sdl.WINDOWPOS_UNDEFINED,
		Posy:   sdl.WINDOWPOS_UNDEFINED,
		Width:  800,
		Height: 600,
	}
	screen.CreateScreen(specs)
}

func testLifecycle() {
	l1 := lifecycle.Loopable{
		Update: func() {
			fmt.Println("Item 1 is in the Loop")
		},
	}

	l2 := lifecycle.Loopable{
		Update: func() {
			fmt.Println("Item 2 is in the Loop")
		},
	}

	lifecycle.Init()
	l1 = lifecycle.Register(l1)
	l2 = lifecycle.Register(l2)

	go lifecycle.Run()

	time.Sleep(2 * time.Second)
	lifecycle.Stop(l1)

	time.Sleep(2 * time.Second)
	lifecycle.Stop(l2)

	select {}
}
