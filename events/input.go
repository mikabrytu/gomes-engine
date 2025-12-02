package events

import (
	"github.com/Papiermond/eventbus"
	"github.com/mikabrytu/gomes-engine/math"
)

type InputMouseClickEvent struct {
	Position math.Vector2
}

func (e InputMouseClickEvent) GetType() eventbus.EventType {
	return INPUT_MOUSE_CLICK
}
