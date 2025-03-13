package screen

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/veandco/go-sdl2/sdl"
)

type Specs struct {
	Title  string
	Posx   int32
	Posy   int32
	Width  int32
	Height int32
}

var loopable lifecycle.Loopable

func CreateScreen(s Specs) {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()

	window, err := sdl.CreateWindow(s.Title, s.Posx, s.Posy, s.Width, s.Height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	loopable = lifecycle.Loopable{Update: update}
	lifecycle.Register(loopable)
	lifecycle.Run()
}

func update() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			lifecycle.Stop(loopable)
			break
		}
	}

	sdl.Delay(33)
}
