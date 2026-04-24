package events

import (
	"sync"
)

type EventType string

type Event interface {
	GetType() EventType
}

type Callback func(Event)

type EventListener struct {
	Id       uint64
	Callback Callback
}

type EventBus interface {
	subscribe(eventType EventType, callback Callback) *EventListener
	unsubscribe(eventType EventType, id uint64)
	publish(event Event)
	count(eventType EventType) int
}

type eventBusImpl struct {
	listeners map[EventType][]EventListener
	mutex     sync.Mutex
	nextId    uint64
}

func newBus() EventBus {
	return &eventBusImpl{
		listeners: make(map[EventType][]EventListener),
	}
}

func (bus *eventBusImpl) subscribe(eventType EventType, callback Callback) *EventListener {
	bus.mutex.Lock()
	defer bus.mutex.Unlock()

	bus.nextId++

	listener := EventListener{
		Id:       bus.nextId,
		Callback: callback,
	}
	bus.listeners[eventType] = append(bus.listeners[eventType], listener)

	return &listener
}

func (bus *eventBusImpl) unsubscribe(eventType EventType, id uint64) {
	bus.mutex.Unlock()
	defer bus.mutex.Lock()

	for i, l := range bus.listeners[eventType] {
		if l.Id == id {
			bus.listeners[eventType] = append(bus.listeners[eventType][:i], bus.listeners[eventType][i+1:]...)
			break
		}
	}
}

func (bus *eventBusImpl) publish(event Event) {
	bus.mutex.Lock()
	defer bus.mutex.Unlock()

	listeners, ok := bus.listeners[event.GetType()]

	if ok {
		for _, listener := range listeners {
			listener.Callback(event)
		}
	}
}

func (bus *eventBusImpl) count(eventType EventType) int {
	return len(bus.listeners[eventType])
}
