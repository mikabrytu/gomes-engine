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

func AddListener(layer EventLayer, key string, callback func(data any)) {
	switch layer {
	case Audio:
		callBusSub(audioBus, EventType(key), callback)
	case Game:
		callBusSub(gameBus, EventType(key), callback)
	case Input:
		callBusSub(inputBus, EventType(key), callback)
	case Lifecycle:
		callBusSub(lifecycleBus, EventType(key), callback)
	case Physics:
		callBusSub(physicsBus, EventType(key), callback)
	case Render:
		callBusSub(renderBus, EventType(key), callback)
	}
}

func RemoveListener(layer EventLayer, key string) {
	switch layer {
	case Audio:
		callBusUnsub(audioBus, EventType(key))
	case Game:
		callBusUnsub(gameBus, EventType(key))
	case Input:
		callBusUnsub(inputBus, EventType(key))
	case Lifecycle:
		callBusUnsub(lifecycleBus, EventType(key))
	case Physics:
		callBusUnsub(physicsBus, EventType(key))
	case Render:
		callBusUnsub(renderBus, EventType(key))
	}
}

func callBusSub(bus EventBus, eventType EventType, callback func(data any)) {
	bus.subscribe(eventType, func(e Event) {
		callback(e)
	})
}

func callBusUnsub(bus EventBus, eventType EventType) {
	bus.unsubscribe(eventType)
}

func publishAsync(bus EventBus, event Event) {
	go func() {
		bus.publish(event)
	}()
}
