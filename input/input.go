package input

import "github.com/veandco/go-sdl2/sdl"

func ListenToInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		println("Listening for Inputs...")

		switch event.(type) {
		case *sdl.MouseButtonEvent:
			println("Mouse button pressed")
			break
		}
	}
}
