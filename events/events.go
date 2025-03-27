package events

import (
	"fmt"

	"github.com/AlexanderGrom/go-event"
)

var dispatcher event.Dispatcher

func Init() {
	dispatcher = event.New()
}

func Subscribe(name string, callback func() error) {
	if name == "" {
		panic("Event Name is empty")
	}

	if callback == nil {
		m := fmt.Sprintf("Event Callback for %s is nil", name)
		panic(m)
	}

	dispatcher.On(name, callback)
}

func Emit(name string) {
	dispatcher.Go(name)
}
