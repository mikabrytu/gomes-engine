package events

type EventLayer int

const (
	Audio EventLayer = iota
	Game
	Input
	Lifecycle
	Physics
	Render
)

var audioBus EventBus
var gameBus EventBus
var inputBus EventBus
var lifecycleBus EventBus
var physicsBus EventBus
var renderBus EventBus

func Init() {
	audioBus = newBus()
	gameBus = newBus()
	inputBus = newBus()
	lifecycleBus = newBus()
	physicsBus = newBus()
	renderBus = newBus()
}

func Emit(layer EventLayer, event any) {
	switch layer {
	case Audio:
		publishAsync(audioBus, event.(Event))
	case Game:
		publishAsync(gameBus, event.(Event))
	case Input:
		publishAsync(inputBus, event.(Event))
	case Lifecycle:
		publishAsync(lifecycleBus, event.(Event))
	case Physics:
		publishAsync(physicsBus, event.(Event))
	case Render:
		publishAsync(renderBus, event.(Event))
	}
}

func AddListener(layer EventLayer, key string, callback func(data any)) *EventListener {
	listener := &EventListener{}

	switch layer {
	case Audio:
		listener = callBusSub(audioBus, EventType(key), callback)
	case Game:
		listener = callBusSub(gameBus, EventType(key), callback)
	case Input:
		listener = callBusSub(inputBus, EventType(key), callback)
	case Lifecycle:
		listener = callBusSub(lifecycleBus, EventType(key), callback)
	case Physics:
		listener = callBusSub(physicsBus, EventType(key), callback)
	case Render:
		listener = callBusSub(renderBus, EventType(key), callback)
	}

	return listener
}

func RemoveListener(layer EventLayer, key string, id uint64) {
	switch layer {
	case Audio:
		callBusUnsub(audioBus, EventType(key), id)
	case Game:
		callBusUnsub(gameBus, EventType(key), id)
	case Input:
		callBusUnsub(inputBus, EventType(key), id)
	case Lifecycle:
		callBusUnsub(lifecycleBus, EventType(key), id)
	case Physics:
		callBusUnsub(physicsBus, EventType(key), id)
	case Render:
		callBusUnsub(renderBus, EventType(key), id)
	}
}

func GetListenerCount(layer EventLayer, key string) int {
	count := 0

	switch layer {
	case Audio:
		count = audioBus.count(EventType(key))
	case Game:
		count = gameBus.count(EventType(key))
	case Input:
		count = inputBus.count(EventType(key))
	case Lifecycle:
		count = lifecycleBus.count(EventType(key))
	case Physics:
		count = physicsBus.count(EventType(key))
	case Render:
		count = renderBus.count(EventType(key))
	}

	return count
}

func callBusSub(bus EventBus, eventType EventType, callback func(data any)) *EventListener {
	return bus.subscribe(eventType, func(e Event) {
		callback(e)
	})
}

func callBusUnsub(bus EventBus, eventType EventType, id uint64) {
	bus.unsubscribe(eventType, id)
}

func publishAsync(bus EventBus, event Event) {
	go func() {
		bus.publish(event)
	}()
}
