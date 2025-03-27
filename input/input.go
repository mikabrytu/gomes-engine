package input

import (
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/veandco/go-sdl2/sdl"
)

func ListenToInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		println("Listening for Inputs...")

		switch event.(type) {
		case *sdl.MouseButtonEvent:
			println("Mouse button pressed")
		case *sdl.QuitEvent:
			println("Quit")
			lifecycle.StopFirst()
			return
		}
	}
}
