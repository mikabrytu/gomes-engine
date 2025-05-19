package lifecycle

import (
	"container/list"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type GameObject struct {
	Id      int
	Start   func()
	Update  func()
	Destroy func()
}

var idCounter = 0
var running bool = false
var objects *list.List
var first GameObject
var last GameObject

// Public API

func Init() {
	objects = list.New()
	first = GameObject{}
	last = GameObject{}
	idCounter = 0
	running = true
}

func Register(o GameObject) GameObject {
	o.Id = idCounter
	idCounter++

	_ = objects.PushFront(o)

	return o
}

func RegisterFirst(o GameObject) {
	first = registerSpecial(o, "First")
}

func RegisterLast(o GameObject) {
	last = registerSpecial(o, "Last")
}

func Stop(o GameObject) {
	var next *list.Element
	for e := objects.Front(); e != nil; e = next {
		next = e.Next()
		item := GameObject(e.Value.(GameObject))

		if o.Id == item.Id {
			item.Destroy()
			objects.Remove(e)
			break
		}
	}

	if objects.Len() == 0 {
		fmt.Println("There's no more loopables on the list. Quitting application")
		running = false
	}
}

func StopById(id int) {
	l := GameObject{Id: id}
	Stop(l)
}

func StopFirst() {
	if first.Destroy != nil {
		first.Destroy()
	}

	first = GameObject{}
}

func StopLast() {
	if last.Destroy != nil {
		last.Destroy()
	}

	last = GameObject{}
}

func Kill() {
	objects.Init()
	first = GameObject{}
	last = GameObject{}
	running = false
}

func Run() {
	if running {
		if first.Start != nil {
			first.Start()
		}

		for e := objects.Front(); e != nil; e = e.Next() {
			item := GameObject(e.Value.(GameObject))
			if item.Start != nil {
				item.Start()
			}
		}

		if last.Start != nil {
			last.Start()
		}
	}

	for running {
		if first.Update != nil {
			first.Update()
		}

		for e := objects.Front(); e != nil; e = e.Next() {
			item := GameObject(e.Value.(GameObject))
			if item.Update != nil {
				item.Update()
			}
		}

		if last.Update != nil {
			last.Update()
		}

		sdl.Delay(33)
	}
}

// Private Implementation

func registerSpecial(o GameObject, message string) GameObject {
	if isGameObjectNil(o) {
		m := fmt.Sprintf("Trying to register a nil loopable as %v", message)
		panic(m)
	}

	return o
}

func isGameObjectNil(o GameObject) bool {
	return o.Start == nil && o.Update == nil && o.Destroy == nil
}
