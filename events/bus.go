package events

import (
	"sync"
)

type EventType string

type Event interface {
	GetType() EventType
}

type EventListener func(Event)

type EventBus interface {
	subscribe(eventType EventType, listener EventListener)
	unsubscribe(eventType EventType)
	publish(event Event)
}

type eventBusImpl struct {
	listeners map[EventType][]EventListener
	mutex     sync.Mutex
}

func newBus() EventBus {
	return &eventBusImpl{
		listeners: make(map[EventType][]EventListener),
	}
}

func (bus *eventBusImpl) subscribe(eventType EventType, listener EventListener) {
	bus.mutex.Lock()
	defer bus.mutex.Unlock()

	bus.listeners[eventType] = append(bus.listeners[eventType], listener)
}

func (bus *eventBusImpl) unsubscribe(eventType EventType) {
	bus.mutex.Lock()
	defer bus.mutex.Unlock()

	delete(bus.listeners, eventType)
}

func (bus *eventBusImpl) publish(event Event) {
	bus.mutex.Lock()
	defer bus.mutex.Unlock()

	listeners, ok := bus.listeners[event.GetType()]

	if ok {
		for _, listener := range listeners {
			listener(event)
		}
	}
}
