package input

import (
	"fmt"

	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardEvent struct {
	key  sdl.Keycode
	name string
}

func ListenToInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.MouseButtonEvent:
			handleMouse(event.(*sdl.MouseButtonEvent))
		case *sdl.KeyboardEvent:
			handleKeyboard(event.(*sdl.KeyboardEvent))
		case *sdl.QuitEvent:
			println("Quit")
			lifecycle.StopFirst()
			return
		}
	}
}

func handleMouse(e *sdl.MouseButtonEvent) {
	events.Emit(events.INPUT_MOUSE_CLICK)

	switch e.State {
	case sdl.PRESSED:
		events.Emit(events.INPUT_MOUSE_CLICK_DOWN, e.X, e.Y)
	case sdl.RELEASED:
		events.Emit(events.INPUT_MOUSE_CLICK_UP, e.X, e.Y)
	}
}

func handleKeyboard(e *sdl.KeyboardEvent) {
	var pressedEvents []KeyboardEvent = []KeyboardEvent{
		{sdl.K_a, events.INPUT_KEYBOARD_PRESSED_A},
		{sdl.K_b, events.INPUT_KEYBOARD_PRESSED_B},
		{sdl.K_c, events.INPUT_KEYBOARD_PRESSED_C},
		{sdl.K_d, events.INPUT_KEYBOARD_PRESSED_D},
		{sdl.K_e, events.INPUT_KEYBOARD_PRESSED_E},
		{sdl.K_f, events.INPUT_KEYBOARD_PRESSED_F},
		{sdl.K_g, events.INPUT_KEYBOARD_PRESSED_G},
		{sdl.K_h, events.INPUT_KEYBOARD_PRESSED_H},
		{sdl.K_i, events.INPUT_KEYBOARD_PRESSED_I},
		{sdl.K_j, events.INPUT_KEYBOARD_PRESSED_J},
		{sdl.K_k, events.INPUT_KEYBOARD_PRESSED_K},
		{sdl.K_l, events.INPUT_KEYBOARD_PRESSED_L},
		{sdl.K_m, events.INPUT_KEYBOARD_PRESSED_M},
		{sdl.K_n, events.INPUT_KEYBOARD_PRESSED_N},
		{sdl.K_o, events.INPUT_KEYBOARD_PRESSED_O},
		{sdl.K_p, events.INPUT_KEYBOARD_PRESSED_P},
		{sdl.K_q, events.INPUT_KEYBOARD_PRESSED_Q},
		{sdl.K_r, events.INPUT_KEYBOARD_PRESSED_R},
		{sdl.K_s, events.INPUT_KEYBOARD_PRESSED_S},
		{sdl.K_t, events.INPUT_KEYBOARD_PRESSED_T},
		{sdl.K_u, events.INPUT_KEYBOARD_PRESSED_U},
		{sdl.K_v, events.INPUT_KEYBOARD_PRESSED_V},
		{sdl.K_w, events.INPUT_KEYBOARD_PRESSED_W},
		{sdl.K_x, events.INPUT_KEYBOARD_PRESSED_X},
		{sdl.K_y, events.INPUT_KEYBOARD_PRESSED_Y},
		{sdl.K_z, events.INPUT_KEYBOARD_PRESSED_Z},
	}

	var releasedEvents []KeyboardEvent = []KeyboardEvent{
		{sdl.K_a, events.INPUT_KEYBOARD_RELEASED_A},
		{sdl.K_b, events.INPUT_KEYBOARD_RELEASED_B},
		{sdl.K_c, events.INPUT_KEYBOARD_RELEASED_C},
		{sdl.K_d, events.INPUT_KEYBOARD_RELEASED_D},
		{sdl.K_e, events.INPUT_KEYBOARD_RELEASED_E},
		{sdl.K_f, events.INPUT_KEYBOARD_RELEASED_F},
		{sdl.K_g, events.INPUT_KEYBOARD_RELEASED_G},
		{sdl.K_h, events.INPUT_KEYBOARD_RELEASED_H},
		{sdl.K_i, events.INPUT_KEYBOARD_RELEASED_I},
		{sdl.K_j, events.INPUT_KEYBOARD_RELEASED_J},
		{sdl.K_k, events.INPUT_KEYBOARD_RELEASED_K},
		{sdl.K_l, events.INPUT_KEYBOARD_RELEASED_L},
		{sdl.K_m, events.INPUT_KEYBOARD_RELEASED_M},
		{sdl.K_n, events.INPUT_KEYBOARD_RELEASED_N},
		{sdl.K_o, events.INPUT_KEYBOARD_RELEASED_O},
		{sdl.K_p, events.INPUT_KEYBOARD_RELEASED_P},
		{sdl.K_q, events.INPUT_KEYBOARD_RELEASED_Q},
		{sdl.K_r, events.INPUT_KEYBOARD_RELEASED_R},
		{sdl.K_s, events.INPUT_KEYBOARD_RELEASED_S},
		{sdl.K_t, events.INPUT_KEYBOARD_RELEASED_T},
		{sdl.K_u, events.INPUT_KEYBOARD_RELEASED_U},
		{sdl.K_v, events.INPUT_KEYBOARD_RELEASED_V},
		{sdl.K_w, events.INPUT_KEYBOARD_RELEASED_W},
		{sdl.K_x, events.INPUT_KEYBOARD_RELEASED_X},
		{sdl.K_y, events.INPUT_KEYBOARD_RELEASED_Y},
		{sdl.K_z, events.INPUT_KEYBOARD_RELEASED_Z},
	}

	switch e.State {
	case sdl.PRESSED:
		emit(pressedEvents, e.Keysym.Sym)
	case sdl.RELEASED:
		emit(releasedEvents, e.Keysym.Sym)
	}
}

func emit(list []KeyboardEvent, code sdl.Keycode) {
	e := findEventByKeycode(list, code)
	if e.name == "" {
		m := fmt.Sprintf("%vcouldn't find event for key %v. Event not sent", utils.ERROR_PREFIX, code)
		println(m)
	}

	events.Emit(e.name)
}

func findEventByKeycode(list []KeyboardEvent, filter sdl.Keycode) KeyboardEvent {
	for _, k := range list {
		if k.key == filter {
			return k
		}
	}

	return KeyboardEvent{}
}
