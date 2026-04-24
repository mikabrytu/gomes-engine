package events

import (
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

func (e InputMouseClickEvent) GetType() EventType { return INPUT_MOUSE_CLICK }

func (e InputMouseClickUpEvent) GetType() EventType { return INPUT_MOUSE_CLICK_UP }

func (e InputMouseClickDownEvent) GetType() EventType { return INPUT_MOUSE_CLICK_DOWN }

func (e InputMouseMoveEvent) GetType() EventType { return INPUT_MOUSE_MOVE }

func (e InputMouseWheelEvent) GetType() EventType { return INPUT_MOUSE_WHEEL }

func (e InputKeyboardPressedAEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_A }

func (e InputKeyboardPressedBEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_B }

func (e InputKeyboardPressedCEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_C }

func (e InputKeyboardPressedDEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_D }

func (e InputKeyboardPressedEEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_E }

func (e InputKeyboardPressedFEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_F }

func (e InputKeyboardPressedGEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_G }

func (e InputKeyboardPressedHEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_H }

func (e InputKeyboardPressedIEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_I }

func (e InputKeyboardPressedJEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_J }

func (e InputKeyboardPressedKEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_K }

func (e InputKeyboardPressedLEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_L }

func (e InputKeyboardPressedMEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_M }

func (e InputKeyboardPressedNEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_N }

func (e InputKeyboardPressedOEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_O }

func (e InputKeyboardPressedPEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_P }

func (e InputKeyboardPressedQEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_Q }

func (e InputKeyboardPressedREvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_R }

func (e InputKeyboardPressedSEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_S }

func (e InputKeyboardPressedTEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_T }

func (e InputKeyboardPressedUEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_U }

func (e InputKeyboardPressedVEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_V }

func (e InputKeyboardPressedWEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_W }

func (e InputKeyboardPressedXEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_X }

func (e InputKeyboardPressedYEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_Y }

func (e InputKeyboardPressedZEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_Z }

func (e InputKeyboardPressed0Event) GetType() EventType { return INPUT_KEYBOARD_PRESSED_0 }

func (e InputKeyboardPressed1Event) GetType() EventType { return INPUT_KEYBOARD_PRESSED_1 }

func (e InputKeyboardPressed2Event) GetType() EventType { return INPUT_KEYBOARD_PRESSED_2 }

func (e InputKeyboardPressed3Event) GetType() EventType { return INPUT_KEYBOARD_PRESSED_3 }

func (e InputKeyboardPressed4Event) GetType() EventType { return INPUT_KEYBOARD_PRESSED_4 }

func (e InputKeyboardPressed5Event) GetType() EventType { return INPUT_KEYBOARD_PRESSED_5 }

func (e InputKeyboardPressed6Event) GetType() EventType { return INPUT_KEYBOARD_PRESSED_6 }

func (e InputKeyboardPressed7Event) GetType() EventType { return INPUT_KEYBOARD_PRESSED_7 }

func (e InputKeyboardPressed8Event) GetType() EventType { return INPUT_KEYBOARD_PRESSED_8 }

func (e InputKeyboardPressed9Event) GetType() EventType { return INPUT_KEYBOARD_PRESSED_9 }

func (e InputKeyboardPressedSpaceEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_SPACE
}

func (e InputKeyboardPressedEnterEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_ENTER
}

func (e InputKeyboardPressedEscapeEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_ESCAPE
}

func (e InputKeyboardPressedTabEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_TAB }

func (e InputKeyboardPressedBackspaceEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_BACKSPACE
}

func (e InputKeyboardPressedShiftEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_SHIFT
}

func (e InputKeyboardPressedCtrlEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_CTRL
}

func (e InputKeyboardPressedAltEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_ALT }

func (e InputKeyboardPressedCapslockEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_CAPSLOCK
}

func (e InputKeyboardPressedLeftEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_LEFT
}

func (e InputKeyboardPressedRightEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_RIGHT
}

func (e InputKeyboardPressedUpEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_UP }

func (e InputKeyboardPressedDownEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_DOWN
}

func (e InputKeyboardPressedDeleteEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_DELETE
}

func (e InputKeyboardPressedHomeEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_HOME
}

func (e InputKeyboardPressedEndEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSED_END }

func (e InputKeyboardPressedPageupEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_PAGEUP
}

func (e InputKeyboardPressedPagedownEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_PAGEDOWN
}

func (e InputKeyboardPressedInsertEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_INSERT
}

func (e InputKeyboardPressedPrintscreenEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_PRINTSCREEN
}

func (e InputKeyboardPressedScrolllockEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_SCROLLLOCK
}

func (e InputKeyboardPressedPauseEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_PAUSE
}

func (e InputKeyboardPressedNumlockEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_NUMLOCK
}

func (e InputKeyboardPressedMinusEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_MINUS
}

func (e InputKeyboardPressedEqualEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_EQUAL
}

func (e InputKeyboardPressedLeftbracketEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_LEFTBRACKET
}

func (e InputKeyboardPressedRightbracketEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_RIGHTBRACKET
}

func (e InputKeyboardPressedBackslashEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_BACKSLASH
}

func (e InputKeyboardPressedSemicolonEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_SEMICOLON
}

func (e InputKeyboardPressedCommaEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_COMMA
}

func (e InputKeyboardPressedPeriodEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_PERIOD
}

func (e InputKeyboardPressedSlashEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSED_SLASH
}

func (e InputKeyboardPressingAEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_A }

func (e InputKeyboardPressingBEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_B }

func (e InputKeyboardPressingCEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_C }

func (e InputKeyboardPressingDEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_D }

func (e InputKeyboardPressingEEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_E }

func (e InputKeyboardPressingFEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_F }

func (e InputKeyboardPressingGEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_G }

func (e InputKeyboardPressingHEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_H }

func (e InputKeyboardPressingIEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_I }

func (e InputKeyboardPressingJEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_J }

func (e InputKeyboardPressingKEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_K }

func (e InputKeyboardPressingLEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_L }

func (e InputKeyboardPressingMEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_M }

func (e InputKeyboardPressingNEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_N }

func (e InputKeyboardPressingOEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_O }

func (e InputKeyboardPressingPEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_P }

func (e InputKeyboardPressingQEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_Q }

func (e InputKeyboardPressingREvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_R }

func (e InputKeyboardPressingSEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_S }

func (e InputKeyboardPressingTEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_T }

func (e InputKeyboardPressingUEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_U }

func (e InputKeyboardPressingVEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_V }

func (e InputKeyboardPressingWEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_W }

func (e InputKeyboardPressingXEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_X }

func (e InputKeyboardPressingYEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_Y }

func (e InputKeyboardPressingZEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_Z }

func (e InputKeyboardPressing0Event) GetType() EventType { return INPUT_KEYBOARD_PRESSING_0 }

func (e InputKeyboardPressing1Event) GetType() EventType { return INPUT_KEYBOARD_PRESSING_1 }

func (e InputKeyboardPressing2Event) GetType() EventType { return INPUT_KEYBOARD_PRESSING_2 }

func (e InputKeyboardPressing3Event) GetType() EventType { return INPUT_KEYBOARD_PRESSING_3 }

func (e InputKeyboardPressing4Event) GetType() EventType { return INPUT_KEYBOARD_PRESSING_4 }

func (e InputKeyboardPressing5Event) GetType() EventType { return INPUT_KEYBOARD_PRESSING_5 }

func (e InputKeyboardPressing6Event) GetType() EventType { return INPUT_KEYBOARD_PRESSING_6 }

func (e InputKeyboardPressing7Event) GetType() EventType { return INPUT_KEYBOARD_PRESSING_7 }

func (e InputKeyboardPressing8Event) GetType() EventType { return INPUT_KEYBOARD_PRESSING_8 }

func (e InputKeyboardPressing9Event) GetType() EventType { return INPUT_KEYBOARD_PRESSING_9 }

func (e InputKeyboardPressingSpaceEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_SPACE
}

func (e InputKeyboardPressingEnterEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_ENTER
}

func (e InputKeyboardPressingEscapeEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_ESCAPE
}

func (e InputKeyboardPressingTabEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_TAB
}

func (e InputKeyboardPressingBackspaceEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_BACKSPACE
}

func (e InputKeyboardPressingShiftEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_SHIFT
}

func (e InputKeyboardPressingCtrlEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_CTRL
}

func (e InputKeyboardPressingAltEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_ALT
}

func (e InputKeyboardPressingCapslockEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_CAPSLOCK
}

func (e InputKeyboardPressingLeftEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_LEFT
}

func (e InputKeyboardPressingRightEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_RIGHT
}

func (e InputKeyboardPressingUpEvent) GetType() EventType { return INPUT_KEYBOARD_PRESSING_UP }

func (e InputKeyboardPressingDownEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_DOWN
}

func (e InputKeyboardPressingDeleteEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_DELETE
}

func (e InputKeyboardPressingHomeEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_HOME
}

func (e InputKeyboardPressingEndEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_END
}

func (e InputKeyboardPressingPageupEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_PAGEUP
}

func (e InputKeyboardPressingPagedownEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_PAGEDOWN
}

func (e InputKeyboardPressingInsertEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_INSERT
}

func (e InputKeyboardPressingPrintscreenEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_PRINTSCREEN
}

func (e InputKeyboardPressingScrolllockEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_SCROLLLOCK
}

func (e InputKeyboardPressingPauseEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_PAUSE
}

func (e InputKeyboardPressingNumlockEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_NUMLOCK
}

func (e InputKeyboardPressingMinusEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_MINUS
}

func (e InputKeyboardPressingEqualEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_EQUAL
}

func (e InputKeyboardPressingLeftbracketEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_LEFTBRACKET
}

func (e InputKeyboardPressingRightbracketEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_RIGHTBRACKET
}

func (e InputKeyboardPressingBackslashEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_BACKSLASH
}

func (e InputKeyboardPressingSemicolonEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_SEMICOLON
}

func (e InputKeyboardPressingCommaEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_COMMA
}

func (e InputKeyboardPressingPeriodEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_PERIOD
}

func (e InputKeyboardPressingSlashEvent) GetType() EventType {
	return INPUT_KEYBOARD_PRESSING_SLASH
}

func (e InputKeyboardReleasedAEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_A }

func (e InputKeyboardReleasedBEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_B }

func (e InputKeyboardReleasedCEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_C }

func (e InputKeyboardReleasedDEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_D }

func (e InputKeyboardReleasedEEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_E }

func (e InputKeyboardReleasedFEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_F }

func (e InputKeyboardReleasedGEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_G }

func (e InputKeyboardReleasedHEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_H }

func (e InputKeyboardReleasedIEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_I }

func (e InputKeyboardReleasedJEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_J }

func (e InputKeyboardReleasedKEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_K }

func (e InputKeyboardReleasedLEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_L }

func (e InputKeyboardReleasedMEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_M }

func (e InputKeyboardReleasedNEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_N }

func (e InputKeyboardReleasedOEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_O }

func (e InputKeyboardReleasedPEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_P }

func (e InputKeyboardReleasedQEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_Q }

func (e InputKeyboardReleasedREvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_R }

func (e InputKeyboardReleasedSEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_S }

func (e InputKeyboardReleasedTEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_T }

func (e InputKeyboardReleasedUEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_U }

func (e InputKeyboardReleasedVEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_V }

func (e InputKeyboardReleasedWEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_W }

func (e InputKeyboardReleasedXEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_X }

func (e InputKeyboardReleasedYEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_Y }

func (e InputKeyboardReleasedZEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_Z }

func (e InputKeyboardReleased0Event) GetType() EventType { return INPUT_KEYBOARD_RELEASED_0 }

func (e InputKeyboardReleased1Event) GetType() EventType { return INPUT_KEYBOARD_RELEASED_1 }

func (e InputKeyboardReleased2Event) GetType() EventType { return INPUT_KEYBOARD_RELEASED_2 }

func (e InputKeyboardReleased3Event) GetType() EventType { return INPUT_KEYBOARD_RELEASED_3 }

func (e InputKeyboardReleased4Event) GetType() EventType { return INPUT_KEYBOARD_RELEASED_4 }

func (e InputKeyboardReleased5Event) GetType() EventType { return INPUT_KEYBOARD_RELEASED_5 }

func (e InputKeyboardReleased6Event) GetType() EventType { return INPUT_KEYBOARD_RELEASED_6 }

func (e InputKeyboardReleased7Event) GetType() EventType { return INPUT_KEYBOARD_RELEASED_7 }

func (e InputKeyboardReleased8Event) GetType() EventType { return INPUT_KEYBOARD_RELEASED_8 }

func (e InputKeyboardReleased9Event) GetType() EventType { return INPUT_KEYBOARD_RELEASED_9 }

func (e InputKeyboardReleasedSpaceEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_SPACE
}

func (e InputKeyboardReleasedEnterEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_ENTER
}

func (e InputKeyboardReleasedEscapeEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_ESCAPE
}

func (e InputKeyboardReleasedTabEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_TAB
}

func (e InputKeyboardReleasedBackspaceEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_BACKSPACE
}

func (e InputKeyboardReleasedShiftEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_SHIFT
}

func (e InputKeyboardReleasedCtrlEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_CTRL
}

func (e InputKeyboardReleasedAltEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_ALT
}

func (e InputKeyboardReleasedCapslockEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_CAPSLOCK
}

func (e InputKeyboardReleasedLeftEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_LEFT
}

func (e InputKeyboardReleasedRightEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_RIGHT
}

func (e InputKeyboardReleasedUpEvent) GetType() EventType { return INPUT_KEYBOARD_RELEASED_UP }

func (e InputKeyboardReleasedDownEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_DOWN
}

func (e InputKeyboardReleasedDeleteEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_DELETE
}

func (e InputKeyboardReleasedHomeEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_HOME
}

func (e InputKeyboardReleasedEndEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_END
}

func (e InputKeyboardReleasedPageupEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_PAGEUP
}

func (e InputKeyboardReleasedPagedownEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_PAGEDOWN
}

func (e InputKeyboardReleasedInsertEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_INSERT
}

func (e InputKeyboardReleasedPrintscreenEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_PRINTSCREEN
}

func (e InputKeyboardReleasedScrolllockEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_SCROLLLOCK
}

func (e InputKeyboardReleasedPauseEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_PAUSE
}

func (e InputKeyboardReleasedNumlockEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_NUMLOCK
}

func (e InputKeyboardReleasedMinusEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_MINUS
}

func (e InputKeyboardReleasedEqualEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_EQUAL
}

func (e InputKeyboardReleasedLeftbracketEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_LEFTBRACKET
}

func (e InputKeyboardReleasedRightbracketEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_RIGHTBRACKET
}

func (e InputKeyboardReleasedBackslashEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_BACKSLASH
}

func (e InputKeyboardReleasedSemicolonEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_SEMICOLON
}

func (e InputKeyboardReleasedCommaEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_COMMA
}

func (e InputKeyboardReleasedPeriodEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_PERIOD
}

func (e InputKeyboardReleasedSlashEvent) GetType() EventType {
	return INPUT_KEYBOARD_RELEASED_SLASH
}
