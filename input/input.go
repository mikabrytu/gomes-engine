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

var pressCount int = 0

func ListenToInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event := event.(type) {
		case *sdl.MouseButtonEvent:
			handleMouse(event)
		case *sdl.KeyboardEvent:
			handleKeyboard(event)
		case *sdl.QuitEvent:
			println("Quit")
			lifecycle.StopInput()
			return
		}
	}
}

func handleMouse(e *sdl.MouseButtonEvent) {
	events.Emit(events.INPUT_MOUSE_CLICK)

	switch e.State {
	case sdl.PRESSED:
		events.Emit(events.INPUT_MOUSE_CLICK_DOWN, int(e.X), int(e.Y))
	case sdl.RELEASED:
		events.Emit(events.INPUT_MOUSE_CLICK_UP, int(e.X), int(e.Y))
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
		{sdl.K_0, events.INPUT_KEYBOARD_PRESSED_0},
		{sdl.K_1, events.INPUT_KEYBOARD_PRESSED_1},
		{sdl.K_2, events.INPUT_KEYBOARD_PRESSED_2},
		{sdl.K_3, events.INPUT_KEYBOARD_PRESSED_3},
		{sdl.K_4, events.INPUT_KEYBOARD_PRESSED_4},
		{sdl.K_5, events.INPUT_KEYBOARD_PRESSED_5},
		{sdl.K_6, events.INPUT_KEYBOARD_PRESSED_6},
		{sdl.K_7, events.INPUT_KEYBOARD_PRESSED_7},
		{sdl.K_8, events.INPUT_KEYBOARD_PRESSED_8},
		{sdl.K_9, events.INPUT_KEYBOARD_PRESSED_9},
		{sdl.K_SPACE, events.INPUT_KEYBOARD_PRESSED_SPACE},
		{sdl.K_RETURN, events.INPUT_KEYBOARD_PRESSED_ENTER},
		{sdl.K_ESCAPE, events.INPUT_KEYBOARD_PRESSED_ESCAPE},
		{sdl.K_TAB, events.INPUT_KEYBOARD_PRESSED_TAB},
		{sdl.K_BACKSPACE, events.INPUT_KEYBOARD_PRESSED_BACKSPACE},
		{sdl.K_LSHIFT, events.INPUT_KEYBOARD_PRESSED_SHIFT},
		{sdl.K_RSHIFT, events.INPUT_KEYBOARD_PRESSED_SHIFT},
		{sdl.K_LCTRL, events.INPUT_KEYBOARD_PRESSED_CTRL},
		{sdl.K_RCTRL, events.INPUT_KEYBOARD_PRESSED_CTRL},
		{sdl.K_LALT, events.INPUT_KEYBOARD_PRESSED_ALT},
		{sdl.K_RALT, events.INPUT_KEYBOARD_PRESSED_ALT},
		{sdl.K_CAPSLOCK, events.INPUT_KEYBOARD_PRESSED_CAPSLOCK},
		{sdl.K_LEFT, events.INPUT_KEYBOARD_PRESSED_LEFT},
		{sdl.K_RIGHT, events.INPUT_KEYBOARD_PRESSED_RIGHT},
		{sdl.K_UP, events.INPUT_KEYBOARD_PRESSED_UP},
		{sdl.K_DOWN, events.INPUT_KEYBOARD_PRESSED_DOWN},
		{sdl.K_DELETE, events.INPUT_KEYBOARD_PRESSED_DELETE},
		{sdl.K_HOME, events.INPUT_KEYBOARD_PRESSED_HOME},
		{sdl.K_END, events.INPUT_KEYBOARD_PRESSED_END},
		{sdl.K_PAGEUP, events.INPUT_KEYBOARD_PRESSED_PAGEUP},
		{sdl.K_PAGEDOWN, events.INPUT_KEYBOARD_PRESSED_PAGEDOWN},
		{sdl.K_INSERT, events.INPUT_KEYBOARD_PRESSED_INSERT},
		{sdl.K_PRINTSCREEN, events.INPUT_KEYBOARD_PRESSED_PRINTSCREEN},
		{sdl.K_SCROLLLOCK, events.INPUT_KEYBOARD_PRESSED_SCROLLLOCK},
		{sdl.K_PAUSE, events.INPUT_KEYBOARD_PRESSED_PAUSE},
		{sdl.K_NUMLOCKCLEAR, events.INPUT_KEYBOARD_PRESSED_NUMLOCK},
		{sdl.K_MINUS, events.INPUT_KEYBOARD_PRESSED_MINUS},
		{sdl.K_EQUALS, events.INPUT_KEYBOARD_PRESSED_EQUAL},
		{sdl.K_LEFTBRACKET, events.INPUT_KEYBOARD_PRESSED_LEFTBRACKET},
		{sdl.K_RIGHTBRACKET, events.INPUT_KEYBOARD_PRESSED_RIGHTBRACKET},
		{sdl.K_BACKSLASH, events.INPUT_KEYBOARD_PRESSED_BACKSLASH},
		{sdl.K_SEMICOLON, events.INPUT_KEYBOARD_PRESSED_SEMICOLON},
		{sdl.K_COMMA, events.INPUT_KEYBOARD_PRESSED_COMMA},
		{sdl.K_PERIOD, events.INPUT_KEYBOARD_PRESSED_PERIOD},
		{sdl.K_SLASH, events.INPUT_KEYBOARD_PRESSED_SLASH},
	}

	var pressingEvents []KeyboardEvent = []KeyboardEvent{
		{sdl.K_a, events.INPUT_KEYBOARD_PRESSING_A},
		{sdl.K_b, events.INPUT_KEYBOARD_PRESSING_B},
		{sdl.K_c, events.INPUT_KEYBOARD_PRESSING_C},
		{sdl.K_d, events.INPUT_KEYBOARD_PRESSING_D},
		{sdl.K_e, events.INPUT_KEYBOARD_PRESSING_E},
		{sdl.K_f, events.INPUT_KEYBOARD_PRESSING_F},
		{sdl.K_g, events.INPUT_KEYBOARD_PRESSING_G},
		{sdl.K_h, events.INPUT_KEYBOARD_PRESSING_H},
		{sdl.K_i, events.INPUT_KEYBOARD_PRESSING_I},
		{sdl.K_j, events.INPUT_KEYBOARD_PRESSING_J},
		{sdl.K_k, events.INPUT_KEYBOARD_PRESSING_K},
		{sdl.K_l, events.INPUT_KEYBOARD_PRESSING_L},
		{sdl.K_m, events.INPUT_KEYBOARD_PRESSING_M},
		{sdl.K_n, events.INPUT_KEYBOARD_PRESSING_N},
		{sdl.K_o, events.INPUT_KEYBOARD_PRESSING_O},
		{sdl.K_p, events.INPUT_KEYBOARD_PRESSING_P},
		{sdl.K_q, events.INPUT_KEYBOARD_PRESSING_Q},
		{sdl.K_r, events.INPUT_KEYBOARD_PRESSING_R},
		{sdl.K_s, events.INPUT_KEYBOARD_PRESSING_S},
		{sdl.K_t, events.INPUT_KEYBOARD_PRESSING_T},
		{sdl.K_u, events.INPUT_KEYBOARD_PRESSING_U},
		{sdl.K_v, events.INPUT_KEYBOARD_PRESSING_V},
		{sdl.K_w, events.INPUT_KEYBOARD_PRESSING_W},
		{sdl.K_x, events.INPUT_KEYBOARD_PRESSING_X},
		{sdl.K_y, events.INPUT_KEYBOARD_PRESSING_Y},
		{sdl.K_z, events.INPUT_KEYBOARD_PRESSING_Z},
		{sdl.K_0, events.INPUT_KEYBOARD_PRESSING_0},
		{sdl.K_1, events.INPUT_KEYBOARD_PRESSING_1},
		{sdl.K_2, events.INPUT_KEYBOARD_PRESSING_2},
		{sdl.K_3, events.INPUT_KEYBOARD_PRESSING_3},
		{sdl.K_4, events.INPUT_KEYBOARD_PRESSING_4},
		{sdl.K_5, events.INPUT_KEYBOARD_PRESSING_5},
		{sdl.K_6, events.INPUT_KEYBOARD_PRESSING_6},
		{sdl.K_7, events.INPUT_KEYBOARD_PRESSING_7},
		{sdl.K_8, events.INPUT_KEYBOARD_PRESSING_8},
		{sdl.K_9, events.INPUT_KEYBOARD_PRESSING_9},
		{sdl.K_SPACE, events.INPUT_KEYBOARD_PRESSING_SPACE},
		{sdl.K_RETURN, events.INPUT_KEYBOARD_PRESSING_ENTER},
		{sdl.K_ESCAPE, events.INPUT_KEYBOARD_PRESSING_ESCAPE},
		{sdl.K_TAB, events.INPUT_KEYBOARD_PRESSING_TAB},
		{sdl.K_BACKSPACE, events.INPUT_KEYBOARD_PRESSING_BACKSPACE},
		{sdl.K_LSHIFT, events.INPUT_KEYBOARD_PRESSING_SHIFT},
		{sdl.K_RSHIFT, events.INPUT_KEYBOARD_PRESSING_SHIFT},
		{sdl.K_LCTRL, events.INPUT_KEYBOARD_PRESSING_CTRL},
		{sdl.K_RCTRL, events.INPUT_KEYBOARD_PRESSING_CTRL},
		{sdl.K_LALT, events.INPUT_KEYBOARD_PRESSING_ALT},
		{sdl.K_RALT, events.INPUT_KEYBOARD_PRESSING_ALT},
		{sdl.K_CAPSLOCK, events.INPUT_KEYBOARD_PRESSING_CAPSLOCK},
		{sdl.K_LEFT, events.INPUT_KEYBOARD_PRESSING_LEFT},
		{sdl.K_RIGHT, events.INPUT_KEYBOARD_PRESSING_RIGHT},
		{sdl.K_UP, events.INPUT_KEYBOARD_PRESSING_UP},
		{sdl.K_DOWN, events.INPUT_KEYBOARD_PRESSING_DOWN},
		{sdl.K_DELETE, events.INPUT_KEYBOARD_PRESSING_DELETE},
		{sdl.K_HOME, events.INPUT_KEYBOARD_PRESSING_HOME},
		{sdl.K_END, events.INPUT_KEYBOARD_PRESSING_END},
		{sdl.K_PAGEUP, events.INPUT_KEYBOARD_PRESSING_PAGEUP},
		{sdl.K_PAGEDOWN, events.INPUT_KEYBOARD_PRESSING_PAGEDOWN},
		{sdl.K_INSERT, events.INPUT_KEYBOARD_PRESSING_INSERT},
		{sdl.K_PRINTSCREEN, events.INPUT_KEYBOARD_PRESSING_PRINTSCREEN},
		{sdl.K_SCROLLLOCK, events.INPUT_KEYBOARD_PRESSING_SCROLLLOCK},
		{sdl.K_PAUSE, events.INPUT_KEYBOARD_PRESSING_PAUSE},
		{sdl.K_NUMLOCKCLEAR, events.INPUT_KEYBOARD_PRESSING_NUMLOCK},
		{sdl.K_MINUS, events.INPUT_KEYBOARD_PRESSING_MINUS},
		{sdl.K_EQUALS, events.INPUT_KEYBOARD_PRESSING_EQUAL},
		{sdl.K_LEFTBRACKET, events.INPUT_KEYBOARD_PRESSING_LEFTBRACKET},
		{sdl.K_RIGHTBRACKET, events.INPUT_KEYBOARD_PRESSING_RIGHTBRACKET},
		{sdl.K_BACKSLASH, events.INPUT_KEYBOARD_PRESSING_BACKSLASH},
		{sdl.K_SEMICOLON, events.INPUT_KEYBOARD_PRESSING_SEMICOLON},
		{sdl.K_COMMA, events.INPUT_KEYBOARD_PRESSING_COMMA},
		{sdl.K_PERIOD, events.INPUT_KEYBOARD_PRESSING_PERIOD},
		{sdl.K_SLASH, events.INPUT_KEYBOARD_PRESSING_SLASH},
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
		{sdl.K_0, events.INPUT_KEYBOARD_RELEASED_0},
		{sdl.K_1, events.INPUT_KEYBOARD_RELEASED_1},
		{sdl.K_2, events.INPUT_KEYBOARD_RELEASED_2},
		{sdl.K_3, events.INPUT_KEYBOARD_RELEASED_3},
		{sdl.K_4, events.INPUT_KEYBOARD_RELEASED_4},
		{sdl.K_5, events.INPUT_KEYBOARD_RELEASED_5},
		{sdl.K_6, events.INPUT_KEYBOARD_RELEASED_6},
		{sdl.K_7, events.INPUT_KEYBOARD_RELEASED_7},
		{sdl.K_8, events.INPUT_KEYBOARD_RELEASED_8},
		{sdl.K_9, events.INPUT_KEYBOARD_RELEASED_9},
		{sdl.K_SPACE, events.INPUT_KEYBOARD_RELEASED_SPACE},
		{sdl.K_RETURN, events.INPUT_KEYBOARD_RELEASED_ENTER},
		{sdl.K_ESCAPE, events.INPUT_KEYBOARD_RELEASED_ESCAPE},
		{sdl.K_TAB, events.INPUT_KEYBOARD_RELEASED_TAB},
		{sdl.K_BACKSPACE, events.INPUT_KEYBOARD_RELEASED_BACKSPACE},
		{sdl.K_LSHIFT, events.INPUT_KEYBOARD_RELEASED_SHIFT},
		{sdl.K_RSHIFT, events.INPUT_KEYBOARD_RELEASED_SHIFT},
		{sdl.K_LCTRL, events.INPUT_KEYBOARD_RELEASED_CTRL},
		{sdl.K_RCTRL, events.INPUT_KEYBOARD_RELEASED_CTRL},
		{sdl.K_LALT, events.INPUT_KEYBOARD_RELEASED_ALT},
		{sdl.K_RALT, events.INPUT_KEYBOARD_RELEASED_ALT},
		{sdl.K_CAPSLOCK, events.INPUT_KEYBOARD_RELEASED_CAPSLOCK},
		{sdl.K_LEFT, events.INPUT_KEYBOARD_RELEASED_LEFT},
		{sdl.K_RIGHT, events.INPUT_KEYBOARD_RELEASED_RIGHT},
		{sdl.K_UP, events.INPUT_KEYBOARD_RELEASED_UP},
		{sdl.K_DOWN, events.INPUT_KEYBOARD_RELEASED_DOWN},
		{sdl.K_DELETE, events.INPUT_KEYBOARD_RELEASED_DELETE},
		{sdl.K_HOME, events.INPUT_KEYBOARD_RELEASED_HOME},
		{sdl.K_END, events.INPUT_KEYBOARD_RELEASED_END},
		{sdl.K_PAGEUP, events.INPUT_KEYBOARD_RELEASED_PAGEUP},
		{sdl.K_PAGEDOWN, events.INPUT_KEYBOARD_RELEASED_PAGEDOWN},
		{sdl.K_INSERT, events.INPUT_KEYBOARD_RELEASED_INSERT},
		{sdl.K_PRINTSCREEN, events.INPUT_KEYBOARD_RELEASED_PRINTSCREEN},
		{sdl.K_SCROLLLOCK, events.INPUT_KEYBOARD_RELEASED_SCROLLLOCK},
		{sdl.K_PAUSE, events.INPUT_KEYBOARD_RELEASED_PAUSE},
		{sdl.K_NUMLOCKCLEAR, events.INPUT_KEYBOARD_RELEASED_NUMLOCK},
		{sdl.K_MINUS, events.INPUT_KEYBOARD_RELEASED_MINUS},
		{sdl.K_EQUALS, events.INPUT_KEYBOARD_RELEASED_EQUAL},
		{sdl.K_LEFTBRACKET, events.INPUT_KEYBOARD_RELEASED_LEFTBRACKET},
		{sdl.K_RIGHTBRACKET, events.INPUT_KEYBOARD_RELEASED_RIGHTBRACKET},
		{sdl.K_BACKSLASH, events.INPUT_KEYBOARD_RELEASED_BACKSLASH},
		{sdl.K_SEMICOLON, events.INPUT_KEYBOARD_RELEASED_SEMICOLON},
		{sdl.K_COMMA, events.INPUT_KEYBOARD_RELEASED_COMMA},
		{sdl.K_PERIOD, events.INPUT_KEYBOARD_RELEASED_PERIOD},
		{sdl.K_SLASH, events.INPUT_KEYBOARD_RELEASED_SLASH},
	}

	switch e.State {
	case sdl.PRESSED:
		pressCount++

		if pressCount > 1 {
			emit(pressingEvents, e.Keysym.Sym)
		} else {
			emit(pressedEvents, e.Keysym.Sym)
		}
	case sdl.RELEASED:
		println("Released")
		pressCount = 0
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
