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
	switch layer {
	case Audio:
		publishAsync(audioBus, event.(eventbus.Event))
	case Game:
		publishAsync(gameBus, event.(eventbus.Event))
	case Input:
		publishAsync(inputBus, event.(eventbus.Event))
	case Lifecycle:
		publishAsync(lifecycleBus, event.(eventbus.Event))
	case Physics:
		publishAsync(physicsBus, event.(eventbus.Event))
	case Render:
		publishAsync(renderBus, event.(eventbus.Event))
	}
}

func Subscribe(layer EventLayer, key string, callback func(data any)) {
	switch layer {
	case Audio:
		subscribe(audioBus, eventbus.EventType(key), callback)
	case Game:
		subscribe(gameBus, eventbus.EventType(key), callback)
	case Input:
		subscribe(inputBus, eventbus.EventType(key), callback)
	case Lifecycle:
		subscribe(lifecycleBus, eventbus.EventType(key), callback)
	case Physics:
		subscribe(physicsBus, eventbus.EventType(key), callback)
	case Render:
		subscribe(renderBus, eventbus.EventType(key), callback)
	}
}

func publishAsync(bus eventbus.EventBus, event eventbus.Event) {
	go func() {
		bus.Publish(event)
	}()
}

func subscribe(bus eventbus.EventBus, eventType eventbus.EventType, callback func(data any)) {
	bus.Subscribe(eventType, func(e eventbus.Event) {
		callback(e)
	})
}
