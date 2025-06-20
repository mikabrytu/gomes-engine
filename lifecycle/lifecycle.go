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
	Physics func()
	Render  func()
	Destroy func()
}

var idCounter = 0
var running bool = false
var objects *list.List
var inputLayer GameObject
var renderLayer GameObject

// Public API

func Init() {
	objects = list.New()
	inputLayer = GameObject{}
	renderLayer = GameObject{}
	idCounter = 0
	running = true
}

func Register(o GameObject) GameObject {
	o.Id = idCounter
	idCounter++

	_ = objects.PushFront(o)

	return o
}

func RegisterInput(o GameObject) {
	inputLayer = registerSpecial(o, "First")
}

func RegisterRender(o GameObject) {
	renderLayer = registerSpecial(o, "Last")
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
	if inputLayer.Destroy != nil {
		inputLayer.Destroy()
	}

	inputLayer = GameObject{}
}

func StopLast() {
	if renderLayer.Destroy != nil {
		renderLayer.Destroy()
	}

	renderLayer = GameObject{}
}

func Kill() {
	objects.Init()
	inputLayer = GameObject{}
	renderLayer = GameObject{}
	running = false
}

func Run() {
	if running {
		if inputLayer.Start != nil {
			inputLayer.Start()
		}

		for e := objects.Front(); e != nil; e = e.Next() {
			item := GameObject(e.Value.(GameObject))
			if item.Start != nil {
				item.Start()
			}
		}

		if renderLayer.Start != nil {
			renderLayer.Start()
		}
	}

	for running {
		if inputLayer.Update != nil {
			inputLayer.Update()
		}

		for e := objects.Front(); e != nil; e = e.Next() {
			item := GameObject(e.Value.(GameObject))
			if item.Update != nil {
				item.Update()
			}
		}

		for e := objects.Front(); e != nil; e = e.Next() {
			item := GameObject(e.Value.(GameObject))
			if item.Physics != nil {
				item.Physics()
			}
		}

		for e := objects.Front(); e != nil; e = e.Next() {
			item := GameObject(e.Value.(GameObject))
			if item.Render != nil {
				item.Render()
			}
		}

		if renderLayer.Update != nil {
			renderLayer.Update()
		}

		sdl.Delay(15)
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
