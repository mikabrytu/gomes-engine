package events

import (
	"github.com/Papiermond/eventbus"
)

type EventLayer int

const (
	Audio EventLayer = iota
	Game
	Input
	Lifecycle
	Physics
	Render
)

var audioBus eventbus.EventBus
var gameBus eventbus.EventBus
var inputBus eventbus.EventBus
var lifecycleBus eventbus.EventBus
var physicsBus eventbus.EventBus
var renderBus eventbus.EventBus

func Init() {
	audioBus = eventbus.New()
	gameBus = eventbus.New()
	inputBus = eventbus.New()
	lifecycleBus = eventbus.New()
	physicsBus = eventbus.New()
	renderBus = eventbus.New()
}

func Emit(layer EventLayer, event any) {
	go func() {
		inputBus.Publish(event.(eventbus.Event))
	}()
}

func Subscribe(layer EventLayer, key string, callback func(data any)) {
	inputBus.Subscribe(eventbus.EventType(key), func(e eventbus.Event) {
		callback(e)
	})
}
