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

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, 800, 600, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	loopable = lifecycle.Loopable{Update: update}
	lifecycle.Run(loopable)
}

func update() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quit")
			lifecycle.Stop()
			break
		}
	}

	sdl.Delay(33)
}
