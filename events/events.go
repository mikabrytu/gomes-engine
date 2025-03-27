package events

import (
	"fmt"

	"github.com/AlexanderGrom/go-event"
)

var dispatcher event.Dispatcher

func Init() {
	dispatcher = event.New()
}

func Subscribe(name string, callback func(params ...any) error) {
	if name == "" {
		panic("Event Name is empty")
	}

	if callback == nil {
		m := fmt.Sprintf("Event Callback for %s is nil", name)
		panic(m)
	}

	dispatcher.On(name, func(args ...any) error {
		callback(args)
		return nil
	})
}

func Emit(name string, params ...any) {
	dispatcher.Go(name, params)
}
