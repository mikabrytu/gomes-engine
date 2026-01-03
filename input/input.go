package input

import (
	"fmt"

	"github.com/mikabrytu/gomes-engine/debug"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/utils"
	"github.com/veandco/go-sdl2/sdl"
)

type KeyboardEvent struct {
	key  sdl.Keycode
	name any
}

var pressCount int = 0

func ListenToInput() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event := event.(type) {
		case *sdl.MouseButtonEvent:
			handleMouseClick(event)
		case *sdl.MouseMotionEvent:
			handleMouseMotion(event)
		case *sdl.KeyboardEvent:
			handleKeyboard(event)
		case *sdl.QuitEvent:
			if debug.IsEnabled() {
				println("Quit")
			}

			lifecycle.StopInput()
			return
		}
	}
}

func handleMouseClick(e *sdl.MouseButtonEvent) {
	index := utils.MouseButtonIndex{}
	if e.Button == sdl.BUTTON_LEFT {
		index.Left = 1
	}
	if e.Button == sdl.BUTTON_RIGHT {
		index.Right = 1
	}
	if e.Button == sdl.BUTTON_MIDDLE {
		index.Middle = 1
	}

	events.Emit(events.Input, events.InputMouseClickEvent{
		Position: math.Vector2{
			X: int(e.X),
			Y: int(e.Y),
		},
		Index: index,
	})

	switch e.State {
	case sdl.PRESSED:
		event := events.InputMouseClickDownEvent{
			Position: math.Vector2{
				X: int(e.X),
				Y: int(e.Y),
			},
			Index: index,
		}

		events.Emit(events.Input, event)
	case sdl.RELEASED:
		event := events.InputMouseClickUpEvent{
			Position: math.Vector2{
				X: int(e.X),
				Y: int(e.Y),
			},
			Index: index,
		}

		events.Emit(events.Input, event)
	}
}

func handleMouseMotion(e *sdl.MouseMotionEvent) {
	events.Emit(events.Input, events.InputMouseMoveEvent{
		Position: math.Vector2{
			X: int(e.X),
			Y: int(e.Y),
		},
	})
}

func handleKeyboard(e *sdl.KeyboardEvent) {
	var pressedEvents []KeyboardEvent = []KeyboardEvent{
		{sdl.K_a, events.InputKeyboardPressedAEvent{}},
		{sdl.K_b, events.InputKeyboardPressedBEvent{}},
		{sdl.K_c, events.InputKeyboardPressedCEvent{}},
		{sdl.K_d, events.InputKeyboardPressedDEvent{}},
		{sdl.K_e, events.InputKeyboardPressedEEvent{}},
		{sdl.K_f, events.InputKeyboardPressedFEvent{}},
		{sdl.K_g, events.InputKeyboardPressedGEvent{}},
		{sdl.K_h, events.InputKeyboardPressedHEvent{}},
		{sdl.K_i, events.InputKeyboardPressedIEvent{}},
		{sdl.K_j, events.InputKeyboardPressedJEvent{}},
		{sdl.K_k, events.InputKeyboardPressedKEvent{}},
		{sdl.K_l, events.InputKeyboardPressedLEvent{}},
		{sdl.K_m, events.InputKeyboardPressedMEvent{}},
		{sdl.K_n, events.InputKeyboardPressedNEvent{}},
		{sdl.K_o, events.InputKeyboardPressedOEvent{}},
		{sdl.K_p, events.InputKeyboardPressedPEvent{}},
		{sdl.K_q, events.InputKeyboardPressedQEvent{}},
		{sdl.K_r, events.InputKeyboardPressedREvent{}},
		{sdl.K_s, events.InputKeyboardPressedSEvent{}},
		{sdl.K_t, events.InputKeyboardPressedTEvent{}},
		{sdl.K_u, events.InputKeyboardPressedUEvent{}},
		{sdl.K_v, events.InputKeyboardPressedVEvent{}},
		{sdl.K_w, events.InputKeyboardPressedWEvent{}},
		{sdl.K_x, events.InputKeyboardPressedXEvent{}},
		{sdl.K_y, events.InputKeyboardPressedYEvent{}},
		{sdl.K_z, events.InputKeyboardPressedZEvent{}},

		{sdl.K_0, events.InputKeyboardPressed0Event{}},
		{sdl.K_1, events.InputKeyboardPressed1Event{}},
		{sdl.K_2, events.InputKeyboardPressed2Event{}},
		{sdl.K_3, events.InputKeyboardPressed3Event{}},
		{sdl.K_4, events.InputKeyboardPressed4Event{}},
		{sdl.K_5, events.InputKeyboardPressed5Event{}},
		{sdl.K_6, events.InputKeyboardPressed6Event{}},
		{sdl.K_7, events.InputKeyboardPressed7Event{}},
		{sdl.K_8, events.InputKeyboardPressed8Event{}},
		{sdl.K_9, events.InputKeyboardPressed9Event{}},

		{sdl.K_SPACE, events.InputKeyboardPressedSpaceEvent{}},
		{sdl.K_RETURN, events.InputKeyboardPressedEnterEvent{}},
		{sdl.K_ESCAPE, events.InputKeyboardPressedEscapeEvent{}},
		{sdl.K_TAB, events.InputKeyboardPressedTabEvent{}},
		{sdl.K_BACKSPACE, events.InputKeyboardPressedBackspaceEvent{}},

		{sdl.K_LSHIFT, events.InputKeyboardPressedShiftEvent{}},
		{sdl.K_RSHIFT, events.InputKeyboardPressedShiftEvent{}},

		{sdl.K_LCTRL, events.InputKeyboardPressedCtrlEvent{}},
		{sdl.K_RCTRL, events.InputKeyboardPressedCtrlEvent{}},

		{sdl.K_LALT, events.InputKeyboardPressedAltEvent{}},
		{sdl.K_RALT, events.InputKeyboardPressedAltEvent{}},

		{sdl.K_CAPSLOCK, events.InputKeyboardPressedCapslockEvent{}},

		{sdl.K_LEFT, events.InputKeyboardPressedLeftEvent{}},
		{sdl.K_RIGHT, events.InputKeyboardPressedRightEvent{}},
		{sdl.K_UP, events.InputKeyboardPressedUpEvent{}},
		{sdl.K_DOWN, events.InputKeyboardPressedDownEvent{}},

		{sdl.K_DELETE, events.InputKeyboardPressedDeleteEvent{}},
		{sdl.K_HOME, events.InputKeyboardPressedHomeEvent{}},
		{sdl.K_END, events.InputKeyboardPressedEndEvent{}},
		{sdl.K_PAGEUP, events.InputKeyboardPressedPageupEvent{}},
		{sdl.K_PAGEDOWN, events.InputKeyboardPressedPagedownEvent{}},
		{sdl.K_INSERT, events.InputKeyboardPressedInsertEvent{}},

		{sdl.K_PRINTSCREEN, events.InputKeyboardPressedPrintscreenEvent{}},
		{sdl.K_SCROLLLOCK, events.InputKeyboardPressedScrolllockEvent{}},
		{sdl.K_PAUSE, events.InputKeyboardPressedPauseEvent{}},
		{sdl.K_NUMLOCKCLEAR, events.InputKeyboardPressedNumlockEvent{}},

		{sdl.K_MINUS, events.InputKeyboardPressedMinusEvent{}},
		{sdl.K_EQUALS, events.InputKeyboardPressedEqualEvent{}},
		{sdl.K_LEFTBRACKET, events.InputKeyboardPressedLeftbracketEvent{}},
		{sdl.K_RIGHTBRACKET, events.InputKeyboardPressedRightbracketEvent{}},
		{sdl.K_BACKSLASH, events.InputKeyboardPressedBackslashEvent{}},
		{sdl.K_SEMICOLON, events.InputKeyboardPressedSemicolonEvent{}},
		{sdl.K_COMMA, events.InputKeyboardPressedCommaEvent{}},
		{sdl.K_PERIOD, events.InputKeyboardPressedPeriodEvent{}},
		{sdl.K_SLASH, events.InputKeyboardPressedSlashEvent{}},
	}

	var pressingEvents []KeyboardEvent = []KeyboardEvent{
		{sdl.K_a, events.InputKeyboardPressingAEvent{}},
		{sdl.K_b, events.InputKeyboardPressingBEvent{}},
		{sdl.K_c, events.InputKeyboardPressingCEvent{}},
		{sdl.K_d, events.InputKeyboardPressingDEvent{}},
		{sdl.K_e, events.InputKeyboardPressingEEvent{}},
		{sdl.K_f, events.InputKeyboardPressingFEvent{}},
		{sdl.K_g, events.InputKeyboardPressingGEvent{}},
		{sdl.K_h, events.InputKeyboardPressingHEvent{}},
		{sdl.K_i, events.InputKeyboardPressingIEvent{}},
		{sdl.K_j, events.InputKeyboardPressingJEvent{}},
		{sdl.K_k, events.InputKeyboardPressingKEvent{}},
		{sdl.K_l, events.InputKeyboardPressingLEvent{}},
		{sdl.K_m, events.InputKeyboardPressingMEvent{}},
		{sdl.K_n, events.InputKeyboardPressingNEvent{}},
		{sdl.K_o, events.InputKeyboardPressingOEvent{}},
		{sdl.K_p, events.InputKeyboardPressingPEvent{}},
		{sdl.K_q, events.InputKeyboardPressingQEvent{}},
		{sdl.K_r, events.InputKeyboardPressingREvent{}},
		{sdl.K_s, events.InputKeyboardPressingSEvent{}},
		{sdl.K_t, events.InputKeyboardPressingTEvent{}},
		{sdl.K_u, events.InputKeyboardPressingUEvent{}},
		{sdl.K_v, events.InputKeyboardPressingVEvent{}},
		{sdl.K_w, events.InputKeyboardPressingWEvent{}},
		{sdl.K_x, events.InputKeyboardPressingXEvent{}},
		{sdl.K_y, events.InputKeyboardPressingYEvent{}},
		{sdl.K_z, events.InputKeyboardPressingZEvent{}},

		{sdl.K_0, events.InputKeyboardPressing0Event{}},
		{sdl.K_1, events.InputKeyboardPressing1Event{}},
		{sdl.K_2, events.InputKeyboardPressing2Event{}},
		{sdl.K_3, events.InputKeyboardPressing3Event{}},
		{sdl.K_4, events.InputKeyboardPressing4Event{}},
		{sdl.K_5, events.InputKeyboardPressing5Event{}},
		{sdl.K_6, events.InputKeyboardPressing6Event{}},
		{sdl.K_7, events.InputKeyboardPressing7Event{}},
		{sdl.K_8, events.InputKeyboardPressing8Event{}},
		{sdl.K_9, events.InputKeyboardPressing9Event{}},

		{sdl.K_SPACE, events.InputKeyboardPressingSpaceEvent{}},
		{sdl.K_RETURN, events.InputKeyboardPressingEnterEvent{}},
		{sdl.K_ESCAPE, events.InputKeyboardPressingEscapeEvent{}},
		{sdl.K_TAB, events.InputKeyboardPressingTabEvent{}},
		{sdl.K_BACKSPACE, events.InputKeyboardPressingBackspaceEvent{}},

		{sdl.K_LSHIFT, events.InputKeyboardPressingShiftEvent{}},
		{sdl.K_RSHIFT, events.InputKeyboardPressingShiftEvent{}},

		{sdl.K_LCTRL, events.InputKeyboardPressingCtrlEvent{}},
		{sdl.K_RCTRL, events.InputKeyboardPressingCtrlEvent{}},

		{sdl.K_LALT, events.InputKeyboardPressingAltEvent{}},
		{sdl.K_RALT, events.InputKeyboardPressingAltEvent{}},

		{sdl.K_CAPSLOCK, events.InputKeyboardPressingCapslockEvent{}},

		{sdl.K_LEFT, events.InputKeyboardPressingLeftEvent{}},
		{sdl.K_RIGHT, events.InputKeyboardPressingRightEvent{}},
		{sdl.K_UP, events.InputKeyboardPressingUpEvent{}},
		{sdl.K_DOWN, events.InputKeyboardPressingDownEvent{}},

		{sdl.K_DELETE, events.InputKeyboardPressingDeleteEvent{}},
		{sdl.K_HOME, events.InputKeyboardPressingHomeEvent{}},
		{sdl.K_END, events.InputKeyboardPressingEndEvent{}},
		{sdl.K_PAGEUP, events.InputKeyboardPressingPageupEvent{}},
		{sdl.K_PAGEDOWN, events.InputKeyboardPressingPagedownEvent{}},
		{sdl.K_INSERT, events.InputKeyboardPressingInsertEvent{}},

		{sdl.K_PRINTSCREEN, events.InputKeyboardPressingPrintscreenEvent{}},
		{sdl.K_SCROLLLOCK, events.InputKeyboardPressingScrolllockEvent{}},
		{sdl.K_PAUSE, events.InputKeyboardPressingPauseEvent{}},
		{sdl.K_NUMLOCKCLEAR, events.InputKeyboardPressingNumlockEvent{}},

		{sdl.K_MINUS, events.InputKeyboardPressingMinusEvent{}},
		{sdl.K_EQUALS, events.InputKeyboardPressingEqualEvent{}},
		{sdl.K_LEFTBRACKET, events.InputKeyboardPressingLeftbracketEvent{}},
		{sdl.K_RIGHTBRACKET, events.InputKeyboardPressingRightbracketEvent{}},
		{sdl.K_BACKSLASH, events.InputKeyboardPressingBackslashEvent{}},
		{sdl.K_SEMICOLON, events.InputKeyboardPressingSemicolonEvent{}},
		{sdl.K_COMMA, events.InputKeyboardPressingCommaEvent{}},
		{sdl.K_PERIOD, events.InputKeyboardPressingPeriodEvent{}},
		{sdl.K_SLASH, events.InputKeyboardPressingSlashEvent{}},
	}

	var releasedEvents []KeyboardEvent = []KeyboardEvent{
		{sdl.K_a, events.InputKeyboardReleasedAEvent{}},
		{sdl.K_b, events.InputKeyboardReleasedBEvent{}},
		{sdl.K_c, events.InputKeyboardReleasedCEvent{}},
		{sdl.K_d, events.InputKeyboardReleasedDEvent{}},
		{sdl.K_e, events.InputKeyboardReleasedEEvent{}},
		{sdl.K_f, events.InputKeyboardReleasedFEvent{}},
		{sdl.K_g, events.InputKeyboardReleasedGEvent{}},
		{sdl.K_h, events.InputKeyboardReleasedHEvent{}},
		{sdl.K_i, events.InputKeyboardReleasedIEvent{}},
		{sdl.K_j, events.InputKeyboardReleasedJEvent{}},
		{sdl.K_k, events.InputKeyboardReleasedKEvent{}},
		{sdl.K_l, events.InputKeyboardReleasedLEvent{}},
		{sdl.K_m, events.InputKeyboardReleasedMEvent{}},
		{sdl.K_n, events.InputKeyboardReleasedNEvent{}},
		{sdl.K_o, events.InputKeyboardReleasedOEvent{}},
		{sdl.K_p, events.InputKeyboardReleasedPEvent{}},
		{sdl.K_q, events.InputKeyboardReleasedQEvent{}},
		{sdl.K_r, events.InputKeyboardReleasedREvent{}},
		{sdl.K_s, events.InputKeyboardReleasedSEvent{}},
		{sdl.K_t, events.InputKeyboardReleasedTEvent{}},
		{sdl.K_u, events.InputKeyboardReleasedUEvent{}},
		{sdl.K_v, events.InputKeyboardReleasedVEvent{}},
		{sdl.K_w, events.InputKeyboardReleasedWEvent{}},
		{sdl.K_x, events.InputKeyboardReleasedXEvent{}},
		{sdl.K_y, events.InputKeyboardReleasedYEvent{}},
		{sdl.K_z, events.InputKeyboardReleasedZEvent{}},

		{sdl.K_0, events.InputKeyboardReleased0Event{}},
		{sdl.K_1, events.InputKeyboardReleased1Event{}},
		{sdl.K_2, events.InputKeyboardReleased2Event{}},
		{sdl.K_3, events.InputKeyboardReleased3Event{}},
		{sdl.K_4, events.InputKeyboardReleased4Event{}},
		{sdl.K_5, events.InputKeyboardReleased5Event{}},
		{sdl.K_6, events.InputKeyboardReleased6Event{}},
		{sdl.K_7, events.InputKeyboardReleased7Event{}},
		{sdl.K_8, events.InputKeyboardReleased8Event{}},
		{sdl.K_9, events.InputKeyboardReleased9Event{}},

		{sdl.K_SPACE, events.InputKeyboardReleasedSpaceEvent{}},
		{sdl.K_RETURN, events.InputKeyboardReleasedEnterEvent{}},
		{sdl.K_ESCAPE, events.InputKeyboardReleasedEscapeEvent{}},
		{sdl.K_TAB, events.InputKeyboardReleasedTabEvent{}},
		{sdl.K_BACKSPACE, events.InputKeyboardReleasedBackspaceEvent{}},

		{sdl.K_LSHIFT, events.InputKeyboardReleasedShiftEvent{}},
		{sdl.K_RSHIFT, events.InputKeyboardReleasedShiftEvent{}},

		{sdl.K_LCTRL, events.InputKeyboardReleasedCtrlEvent{}},
		{sdl.K_RCTRL, events.InputKeyboardReleasedCtrlEvent{}},

		{sdl.K_LALT, events.InputKeyboardReleasedAltEvent{}},
		{sdl.K_RALT, events.InputKeyboardReleasedAltEvent{}},

		{sdl.K_CAPSLOCK, events.InputKeyboardReleasedCapslockEvent{}},

		{sdl.K_LEFT, events.InputKeyboardReleasedLeftEvent{}},
		{sdl.K_RIGHT, events.InputKeyboardReleasedRightEvent{}},
		{sdl.K_UP, events.InputKeyboardReleasedUpEvent{}},
		{sdl.K_DOWN, events.InputKeyboardReleasedDownEvent{}},

		{sdl.K_DELETE, events.InputKeyboardReleasedDeleteEvent{}},
		{sdl.K_HOME, events.InputKeyboardReleasedHomeEvent{}},
		{sdl.K_END, events.InputKeyboardReleasedEndEvent{}},
		{sdl.K_PAGEUP, events.InputKeyboardReleasedPageupEvent{}},
		{sdl.K_PAGEDOWN, events.InputKeyboardReleasedPagedownEvent{}},
		{sdl.K_INSERT, events.InputKeyboardReleasedInsertEvent{}},

		{sdl.K_PRINTSCREEN, events.InputKeyboardReleasedPrintscreenEvent{}},
		{sdl.K_SCROLLLOCK, events.InputKeyboardReleasedScrolllockEvent{}},
		{sdl.K_PAUSE, events.InputKeyboardReleasedPauseEvent{}},
		{sdl.K_NUMLOCKCLEAR, events.InputKeyboardReleasedNumlockEvent{}},

		{sdl.K_MINUS, events.InputKeyboardReleasedMinusEvent{}},
		{sdl.K_EQUALS, events.InputKeyboardReleasedEqualEvent{}},
		{sdl.K_LEFTBRACKET, events.InputKeyboardReleasedLeftbracketEvent{}},
		{sdl.K_RIGHTBRACKET, events.InputKeyboardReleasedRightbracketEvent{}},
		{sdl.K_BACKSLASH, events.InputKeyboardReleasedBackslashEvent{}},
		{sdl.K_SEMICOLON, events.InputKeyboardReleasedSemicolonEvent{}},
		{sdl.K_COMMA, events.InputKeyboardReleasedCommaEvent{}},
		{sdl.K_PERIOD, events.InputKeyboardReleasedPeriodEvent{}},
		{sdl.K_SLASH, events.InputKeyboardReleasedSlashEvent{}},
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
		pressCount = 0
		emit(releasedEvents, e.Keysym.Sym)
	}
}

func emit(list []KeyboardEvent, code sdl.Keycode) {
	e := findEventByKeycode(list, code)
	if debug.IsEnabled() && e.name == "" {
		m := fmt.Sprintf("%vcouldn't find event for key %v. Event not sent", utils.ERROR_PREFIX, code)
		println(m)
	}

	events.Emit(events.Input, e.name)
}

func findEventByKeycode(list []KeyboardEvent, filter sdl.Keycode) KeyboardEvent {
	for _, k := range list {
		if k.key == filter {
			return k
		}
	}

	return KeyboardEvent{}
}
