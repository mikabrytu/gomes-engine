package events

import (
	"github.com/Papiermond/eventbus"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/utils"
)

type InputMouseClickEvent struct {
	Position math.Vector2
	Index    utils.MouseButtonIndex
}
type InputMouseClickUpEvent struct {
	Position math.Vector2
	Index    utils.MouseButtonIndex
}
type InputMouseClickDownEvent struct {
	Position math.Vector2
	Index    utils.MouseButtonIndex
}
type InputMouseMoveEvent struct {
	Position math.Vector2
}
type InputMouseWheelEvent struct{}
type InputKeyboardPressedAEvent struct{}
type InputKeyboardPressedBEvent struct{}
type InputKeyboardPressedCEvent struct{}
type InputKeyboardPressedDEvent struct{}
type InputKeyboardPressedEEvent struct{}
type InputKeyboardPressedFEvent struct{}
type InputKeyboardPressedGEvent struct{}
type InputKeyboardPressedHEvent struct{}
type InputKeyboardPressedIEvent struct{}
type InputKeyboardPressedJEvent struct{}
type InputKeyboardPressedKEvent struct{}
type InputKeyboardPressedLEvent struct{}
type InputKeyboardPressedMEvent struct{}
type InputKeyboardPressedNEvent struct{}
type InputKeyboardPressedOEvent struct{}
type InputKeyboardPressedPEvent struct{}
type InputKeyboardPressedQEvent struct{}
type InputKeyboardPressedREvent struct{}
type InputKeyboardPressedSEvent struct{}
type InputKeyboardPressedTEvent struct{}
type InputKeyboardPressedUEvent struct{}
type InputKeyboardPressedVEvent struct{}
type InputKeyboardPressedWEvent struct{}
type InputKeyboardPressedXEvent struct{}
type InputKeyboardPressedYEvent struct{}
type InputKeyboardPressedZEvent struct{}
type InputKeyboardPressed0Event struct{}
type InputKeyboardPressed1Event struct{}
type InputKeyboardPressed2Event struct{}
type InputKeyboardPressed3Event struct{}
type InputKeyboardPressed4Event struct{}
type InputKeyboardPressed5Event struct{}
type InputKeyboardPressed6Event struct{}
type InputKeyboardPressed7Event struct{}
type InputKeyboardPressed8Event struct{}
type InputKeyboardPressed9Event struct{}
type InputKeyboardPressedSpaceEvent struct{}
type InputKeyboardPressedEnterEvent struct{}
type InputKeyboardPressedEscapeEvent struct{}
type InputKeyboardPressedTabEvent struct{}
type InputKeyboardPressedBackspaceEvent struct{}
type InputKeyboardPressedShiftEvent struct{}
type InputKeyboardPressedCtrlEvent struct{}
type InputKeyboardPressedAltEvent struct{}
type InputKeyboardPressedCapslockEvent struct{}
type InputKeyboardPressedLeftEvent struct{}
type InputKeyboardPressedRightEvent struct{}
type InputKeyboardPressedUpEvent struct{}
type InputKeyboardPressedDownEvent struct{}
type InputKeyboardPressedDeleteEvent struct{}
type InputKeyboardPressedHomeEvent struct{}
type InputKeyboardPressedEndEvent struct{}
type InputKeyboardPressedPageupEvent struct{}
type InputKeyboardPressedPagedownEvent struct{}
type InputKeyboardPressedInsertEvent struct{}
type InputKeyboardPressedPrintscreenEvent struct{}
type InputKeyboardPressedScrolllockEvent struct{}
type InputKeyboardPressedPauseEvent struct{}
type InputKeyboardPressedNumlockEvent struct{}
type InputKeyboardPressedMinusEvent struct{}
type InputKeyboardPressedEqualEvent struct{}
type InputKeyboardPressedLeftbracketEvent struct{}
type InputKeyboardPressedRightbracketEvent struct{}
type InputKeyboardPressedBackslashEvent struct{}
type InputKeyboardPressedSemicolonEvent struct{}
type InputKeyboardPressedCommaEvent struct{}
type InputKeyboardPressedPeriodEvent struct{}
type InputKeyboardPressedSlashEvent struct{}
type InputKeyboardPressingAEvent struct{}
type InputKeyboardPressingBEvent struct{}
type InputKeyboardPressingCEvent struct{}
type InputKeyboardPressingDEvent struct{}
type InputKeyboardPressingEEvent struct{}
type InputKeyboardPressingFEvent struct{}
type InputKeyboardPressingGEvent struct{}
type InputKeyboardPressingHEvent struct{}
type InputKeyboardPressingIEvent struct{}
type InputKeyboardPressingJEvent struct{}
type InputKeyboardPressingKEvent struct{}
type InputKeyboardPressingLEvent struct{}
type InputKeyboardPressingMEvent struct{}
type InputKeyboardPressingNEvent struct{}
type InputKeyboardPressingOEvent struct{}
type InputKeyboardPressingPEvent struct{}
type InputKeyboardPressingQEvent struct{}
type InputKeyboardPressingREvent struct{}
type InputKeyboardPressingSEvent struct{}
type InputKeyboardPressingTEvent struct{}
type InputKeyboardPressingUEvent struct{}
type InputKeyboardPressingVEvent struct{}
type InputKeyboardPressingWEvent struct{}
type InputKeyboardPressingXEvent struct{}
type InputKeyboardPressingYEvent struct{}
type InputKeyboardPressingZEvent struct{}
type InputKeyboardPressing0Event struct{}
type InputKeyboardPressing1Event struct{}
type InputKeyboardPressing2Event struct{}
type InputKeyboardPressing3Event struct{}
type InputKeyboardPressing4Event struct{}
type InputKeyboardPressing5Event struct{}
type InputKeyboardPressing6Event struct{}
type InputKeyboardPressing7Event struct{}
type InputKeyboardPressing8Event struct{}
type InputKeyboardPressing9Event struct{}
type InputKeyboardPressingSpaceEvent struct{}
type InputKeyboardPressingEnterEvent struct{}
type InputKeyboardPressingEscapeEvent struct{}
type InputKeyboardPressingTabEvent struct{}
type InputKeyboardPressingBackspaceEvent struct{}
type InputKeyboardPressingShiftEvent struct{}
type InputKeyboardPressingCtrlEvent struct{}
type InputKeyboardPressingAltEvent struct{}
type InputKeyboardPressingCapslockEvent struct{}
type InputKeyboardPressingLeftEvent struct{}
type InputKeyboardPressingRightEvent struct{}
type InputKeyboardPressingUpEvent struct{}
type InputKeyboardPressingDownEvent struct{}
type InputKeyboardPressingDeleteEvent struct{}
type InputKeyboardPressingHomeEvent struct{}
type InputKeyboardPressingEndEvent struct{}
type InputKeyboardPressingPageupEvent struct{}
type InputKeyboardPressingPagedownEvent struct{}
type InputKeyboardPressingInsertEvent struct{}
type InputKeyboardPressingPrintscreenEvent struct{}
type InputKeyboardPressingScrolllockEvent struct{}
type InputKeyboardPressingPauseEvent struct{}
type InputKeyboardPressingNumlockEvent struct{}
type InputKeyboardPressingMinusEvent struct{}
type InputKeyboardPressingEqualEvent struct{}
type InputKeyboardPressingLeftbracketEvent struct{}
type InputKeyboardPressingRightbracketEvent struct{}
type InputKeyboardPressingBackslashEvent struct{}
type InputKeyboardPressingSemicolonEvent struct{}
type InputKeyboardPressingCommaEvent struct{}
type InputKeyboardPressingPeriodEvent struct{}
type InputKeyboardPressingSlashEvent struct{}
type InputKeyboardReleasedAEvent struct{}
type InputKeyboardReleasedBEvent struct{}
type InputKeyboardReleasedCEvent struct{}
type InputKeyboardReleasedDEvent struct{}
type InputKeyboardReleasedEEvent struct{}
type InputKeyboardReleasedFEvent struct{}
type InputKeyboardReleasedGEvent struct{}
type InputKeyboardReleasedHEvent struct{}
type InputKeyboardReleasedIEvent struct{}
type InputKeyboardReleasedJEvent struct{}
type InputKeyboardReleasedKEvent struct{}
type InputKeyboardReleasedLEvent struct{}
type InputKeyboardReleasedMEvent struct{}
type InputKeyboardReleasedNEvent struct{}
type InputKeyboardReleasedOEvent struct{}
type InputKeyboardReleasedPEvent struct{}
type InputKeyboardReleasedQEvent struct{}
type InputKeyboardReleasedREvent struct{}
type InputKeyboardReleasedSEvent struct{}
type InputKeyboardReleasedTEvent struct{}
type InputKeyboardReleasedUEvent struct{}
type InputKeyboardReleasedVEvent struct{}
type InputKeyboardReleasedWEvent struct{}
type InputKeyboardReleasedXEvent struct{}
type InputKeyboardReleasedYEvent struct{}
type InputKeyboardReleasedZEvent struct{}
type InputKeyboardReleased0Event struct{}
type InputKeyboardReleased1Event struct{}
type InputKeyboardReleased2Event struct{}
type InputKeyboardReleased3Event struct{}
type InputKeyboardReleased4Event struct{}
type InputKeyboardReleased5Event struct{}
type InputKeyboardReleased6Event struct{}
type InputKeyboardReleased7Event struct{}
type InputKeyboardReleased8Event struct{}
type InputKeyboardReleased9Event struct{}
type InputKeyboardReleasedSpaceEvent struct{}
type InputKeyboardReleasedEnterEvent struct{}
type InputKeyboardReleasedEscapeEvent struct{}
type InputKeyboardReleasedTabEvent struct{}
type InputKeyboardReleasedBackspaceEvent struct{}
type InputKeyboardReleasedShiftEvent struct{}
type InputKeyboardReleasedCtrlEvent struct{}
type InputKeyboardReleasedAltEvent struct{}
type InputKeyboardReleasedCapslockEvent struct{}
type InputKeyboardReleasedLeftEvent struct{}
type InputKeyboardReleasedRightEvent struct{}
type InputKeyboardReleasedUpEvent struct{}
type InputKeyboardReleasedDownEvent struct{}
type InputKeyboardReleasedDeleteEvent struct{}
type InputKeyboardReleasedHomeEvent struct{}
type InputKeyboardReleasedEndEvent struct{}
type InputKeyboardReleasedPageupEvent struct{}
type InputKeyboardReleasedPagedownEvent struct{}
type InputKeyboardReleasedInsertEvent struct{}
type InputKeyboardReleasedPrintscreenEvent struct{}
type InputKeyboardReleasedScrolllockEvent struct{}
type InputKeyboardReleasedPauseEvent struct{}
type InputKeyboardReleasedNumlockEvent struct{}
type InputKeyboardReleasedMinusEvent struct{}
type InputKeyboardReleasedEqualEvent struct{}
type InputKeyboardReleasedLeftbracketEvent struct{}
type InputKeyboardReleasedRightbracketEvent struct{}
type InputKeyboardReleasedBackslashEvent struct{}
type InputKeyboardReleasedSemicolonEvent struct{}
type InputKeyboardReleasedCommaEvent struct{}
type InputKeyboardReleasedPeriodEvent struct{}
type InputKeyboardReleasedSlashEvent struct{}

func (e InputMouseClickEvent) GetType() eventbus.EventType { return INPUT_MOUSE_CLICK }

func (e InputMouseClickUpEvent) GetType() eventbus.EventType { return INPUT_MOUSE_CLICK_UP }

func (e InputMouseClickDownEvent) GetType() eventbus.EventType { return INPUT_MOUSE_CLICK_DOWN }

func (e InputMouseMoveEvent) GetType() eventbus.EventType { return INPUT_MOUSE_MOVE }

func (e InputMouseWheelEvent) GetType() eventbus.EventType { return INPUT_MOUSE_WHEEL }

func (e InputKeyboardPressedAEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_A }

func (e InputKeyboardPressedBEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_B }

func (e InputKeyboardPressedCEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_C }

func (e InputKeyboardPressedDEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_D }

func (e InputKeyboardPressedEEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_E }

func (e InputKeyboardPressedFEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_F }

func (e InputKeyboardPressedGEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_G }

func (e InputKeyboardPressedHEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_H }

func (e InputKeyboardPressedIEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_I }

func (e InputKeyboardPressedJEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_J }

func (e InputKeyboardPressedKEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_K }

func (e InputKeyboardPressedLEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_L }

func (e InputKeyboardPressedMEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_M }

func (e InputKeyboardPressedNEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_N }

func (e InputKeyboardPressedOEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_O }

func (e InputKeyboardPressedPEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_P }

func (e InputKeyboardPressedQEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_Q }

func (e InputKeyboardPressedREvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_R }

func (e InputKeyboardPressedSEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_S }

func (e InputKeyboardPressedTEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_T }

func (e InputKeyboardPressedUEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_U }

func (e InputKeyboardPressedVEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_V }

func (e InputKeyboardPressedWEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_W }

func (e InputKeyboardPressedXEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_X }

func (e InputKeyboardPressedYEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_Y }

func (e InputKeyboardPressedZEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_Z }

func (e InputKeyboardPressed0Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_0 }

func (e InputKeyboardPressed1Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_1 }

func (e InputKeyboardPressed2Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_2 }

func (e InputKeyboardPressed3Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_3 }

func (e InputKeyboardPressed4Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_4 }

func (e InputKeyboardPressed5Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_5 }

func (e InputKeyboardPressed6Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_6 }

func (e InputKeyboardPressed7Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_7 }

func (e InputKeyboardPressed8Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_8 }

func (e InputKeyboardPressed9Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_9 }

func (e InputKeyboardPressedSpaceEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_SPACE
}

func (e InputKeyboardPressedEnterEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_ENTER
}

func (e InputKeyboardPressedEscapeEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_ESCAPE
}

func (e InputKeyboardPressedTabEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_TAB }

func (e InputKeyboardPressedBackspaceEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_BACKSPACE
}

func (e InputKeyboardPressedShiftEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_SHIFT
}

func (e InputKeyboardPressedCtrlEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_CTRL
}

func (e InputKeyboardPressedAltEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_ALT }

func (e InputKeyboardPressedCapslockEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_CAPSLOCK
}

func (e InputKeyboardPressedLeftEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_LEFT
}

func (e InputKeyboardPressedRightEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_RIGHT
}

func (e InputKeyboardPressedUpEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_UP }

func (e InputKeyboardPressedDownEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_DOWN
}

func (e InputKeyboardPressedDeleteEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_DELETE
}

func (e InputKeyboardPressedHomeEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_HOME
}

func (e InputKeyboardPressedEndEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSED_END }

func (e InputKeyboardPressedPageupEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_PAGEUP
}

func (e InputKeyboardPressedPagedownEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_PAGEDOWN
}

func (e InputKeyboardPressedInsertEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_INSERT
}

func (e InputKeyboardPressedPrintscreenEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_PRINTSCREEN
}

func (e InputKeyboardPressedScrolllockEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_SCROLLLOCK
}

func (e InputKeyboardPressedPauseEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_PAUSE
}

func (e InputKeyboardPressedNumlockEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_NUMLOCK
}

func (e InputKeyboardPressedMinusEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_MINUS
}

func (e InputKeyboardPressedEqualEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_EQUAL
}

func (e InputKeyboardPressedLeftbracketEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_LEFTBRACKET
}

func (e InputKeyboardPressedRightbracketEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_RIGHTBRACKET
}

func (e InputKeyboardPressedBackslashEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_BACKSLASH
}

func (e InputKeyboardPressedSemicolonEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_SEMICOLON
}

func (e InputKeyboardPressedCommaEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_COMMA
}

func (e InputKeyboardPressedPeriodEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_PERIOD
}

func (e InputKeyboardPressedSlashEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSED_SLASH
}

func (e InputKeyboardPressingAEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_A }

func (e InputKeyboardPressingBEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_B }

func (e InputKeyboardPressingCEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_C }

func (e InputKeyboardPressingDEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_D }

func (e InputKeyboardPressingEEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_E }

func (e InputKeyboardPressingFEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_F }

func (e InputKeyboardPressingGEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_G }

func (e InputKeyboardPressingHEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_H }

func (e InputKeyboardPressingIEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_I }

func (e InputKeyboardPressingJEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_J }

func (e InputKeyboardPressingKEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_K }

func (e InputKeyboardPressingLEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_L }

func (e InputKeyboardPressingMEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_M }

func (e InputKeyboardPressingNEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_N }

func (e InputKeyboardPressingOEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_O }

func (e InputKeyboardPressingPEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_P }

func (e InputKeyboardPressingQEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_Q }

func (e InputKeyboardPressingREvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_R }

func (e InputKeyboardPressingSEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_S }

func (e InputKeyboardPressingTEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_T }

func (e InputKeyboardPressingUEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_U }

func (e InputKeyboardPressingVEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_V }

func (e InputKeyboardPressingWEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_W }

func (e InputKeyboardPressingXEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_X }

func (e InputKeyboardPressingYEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_Y }

func (e InputKeyboardPressingZEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_Z }

func (e InputKeyboardPressing0Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_0 }

func (e InputKeyboardPressing1Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_1 }

func (e InputKeyboardPressing2Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_2 }

func (e InputKeyboardPressing3Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_3 }

func (e InputKeyboardPressing4Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_4 }

func (e InputKeyboardPressing5Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_5 }

func (e InputKeyboardPressing6Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_6 }

func (e InputKeyboardPressing7Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_7 }

func (e InputKeyboardPressing8Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_8 }

func (e InputKeyboardPressing9Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_9 }

func (e InputKeyboardPressingSpaceEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_SPACE
}

func (e InputKeyboardPressingEnterEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_ENTER
}

func (e InputKeyboardPressingEscapeEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_ESCAPE
}

func (e InputKeyboardPressingTabEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_TAB
}

func (e InputKeyboardPressingBackspaceEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_BACKSPACE
}

func (e InputKeyboardPressingShiftEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_SHIFT
}

func (e InputKeyboardPressingCtrlEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_CTRL
}

func (e InputKeyboardPressingAltEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_ALT
}

func (e InputKeyboardPressingCapslockEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_CAPSLOCK
}

func (e InputKeyboardPressingLeftEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_LEFT
}

func (e InputKeyboardPressingRightEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_RIGHT
}

func (e InputKeyboardPressingUpEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_PRESSING_UP }

func (e InputKeyboardPressingDownEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_DOWN
}

func (e InputKeyboardPressingDeleteEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_DELETE
}

func (e InputKeyboardPressingHomeEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_HOME
}

func (e InputKeyboardPressingEndEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_END
}

func (e InputKeyboardPressingPageupEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_PAGEUP
}

func (e InputKeyboardPressingPagedownEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_PAGEDOWN
}

func (e InputKeyboardPressingInsertEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_INSERT
}

func (e InputKeyboardPressingPrintscreenEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_PRINTSCREEN
}

func (e InputKeyboardPressingScrolllockEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_SCROLLLOCK
}

func (e InputKeyboardPressingPauseEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_PAUSE
}

func (e InputKeyboardPressingNumlockEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_NUMLOCK
}

func (e InputKeyboardPressingMinusEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_MINUS
}

func (e InputKeyboardPressingEqualEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_EQUAL
}

func (e InputKeyboardPressingLeftbracketEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_LEFTBRACKET
}

func (e InputKeyboardPressingRightbracketEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_RIGHTBRACKET
}

func (e InputKeyboardPressingBackslashEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_BACKSLASH
}

func (e InputKeyboardPressingSemicolonEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_SEMICOLON
}

func (e InputKeyboardPressingCommaEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_COMMA
}

func (e InputKeyboardPressingPeriodEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_PERIOD
}

func (e InputKeyboardPressingSlashEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_PRESSING_SLASH
}

func (e InputKeyboardReleasedAEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_A }

func (e InputKeyboardReleasedBEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_B }

func (e InputKeyboardReleasedCEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_C }

func (e InputKeyboardReleasedDEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_D }

func (e InputKeyboardReleasedEEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_E }

func (e InputKeyboardReleasedFEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_F }

func (e InputKeyboardReleasedGEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_G }

func (e InputKeyboardReleasedHEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_H }

func (e InputKeyboardReleasedIEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_I }

func (e InputKeyboardReleasedJEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_J }

func (e InputKeyboardReleasedKEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_K }

func (e InputKeyboardReleasedLEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_L }

func (e InputKeyboardReleasedMEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_M }

func (e InputKeyboardReleasedNEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_N }

func (e InputKeyboardReleasedOEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_O }

func (e InputKeyboardReleasedPEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_P }

func (e InputKeyboardReleasedQEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_Q }

func (e InputKeyboardReleasedREvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_R }

func (e InputKeyboardReleasedSEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_S }

func (e InputKeyboardReleasedTEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_T }

func (e InputKeyboardReleasedUEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_U }

func (e InputKeyboardReleasedVEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_V }

func (e InputKeyboardReleasedWEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_W }

func (e InputKeyboardReleasedXEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_X }

func (e InputKeyboardReleasedYEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_Y }

func (e InputKeyboardReleasedZEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_Z }

func (e InputKeyboardReleased0Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_0 }

func (e InputKeyboardReleased1Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_1 }

func (e InputKeyboardReleased2Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_2 }

func (e InputKeyboardReleased3Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_3 }

func (e InputKeyboardReleased4Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_4 }

func (e InputKeyboardReleased5Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_5 }

func (e InputKeyboardReleased6Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_6 }

func (e InputKeyboardReleased7Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_7 }

func (e InputKeyboardReleased8Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_8 }

func (e InputKeyboardReleased9Event) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_9 }

func (e InputKeyboardReleasedSpaceEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_SPACE
}

func (e InputKeyboardReleasedEnterEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_ENTER
}

func (e InputKeyboardReleasedEscapeEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_ESCAPE
}

func (e InputKeyboardReleasedTabEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_TAB
}

func (e InputKeyboardReleasedBackspaceEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_BACKSPACE
}

func (e InputKeyboardReleasedShiftEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_SHIFT
}

func (e InputKeyboardReleasedCtrlEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_CTRL
}

func (e InputKeyboardReleasedAltEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_ALT
}

func (e InputKeyboardReleasedCapslockEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_CAPSLOCK
}

func (e InputKeyboardReleasedLeftEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_LEFT
}

func (e InputKeyboardReleasedRightEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_RIGHT
}

func (e InputKeyboardReleasedUpEvent) GetType() eventbus.EventType { return INPUT_KEYBOARD_RELEASED_UP }

func (e InputKeyboardReleasedDownEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_DOWN
}

func (e InputKeyboardReleasedDeleteEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_DELETE
}

func (e InputKeyboardReleasedHomeEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_HOME
}

func (e InputKeyboardReleasedEndEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_END
}

func (e InputKeyboardReleasedPageupEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_PAGEUP
}

func (e InputKeyboardReleasedPagedownEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_PAGEDOWN
}

func (e InputKeyboardReleasedInsertEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_INSERT
}

func (e InputKeyboardReleasedPrintscreenEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_PRINTSCREEN
}

func (e InputKeyboardReleasedScrolllockEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_SCROLLLOCK
}

func (e InputKeyboardReleasedPauseEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_PAUSE
}

func (e InputKeyboardReleasedNumlockEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_NUMLOCK
}

func (e InputKeyboardReleasedMinusEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_MINUS
}

func (e InputKeyboardReleasedEqualEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_EQUAL
}

func (e InputKeyboardReleasedLeftbracketEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_LEFTBRACKET
}

func (e InputKeyboardReleasedRightbracketEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_RIGHTBRACKET
}

func (e InputKeyboardReleasedBackslashEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_BACKSLASH
}

func (e InputKeyboardReleasedSemicolonEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_SEMICOLON
}

func (e InputKeyboardReleasedCommaEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_COMMA
}

func (e InputKeyboardReleasedPeriodEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_PERIOD
}

func (e InputKeyboardReleasedSlashEvent) GetType() eventbus.EventType {
	return INPUT_KEYBOARD_RELEASED_SLASH
}
