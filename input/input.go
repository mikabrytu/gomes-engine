package input

import (
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/veandco/go-sdl2/sdl"
)

func ListenToInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.MouseButtonEvent:
			events.Emit(events.INPUT_MOUSE_CLICK)
		case *sdl.QuitEvent:
			println("Quit")
			lifecycle.StopFirst()
			return
		}
	}
}
