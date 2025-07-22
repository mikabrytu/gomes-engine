package lifecycle

import (
	"container/list"
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

type GameObject struct {
	Id      int
	Start   func()
	Update  func()
	Physics func()
	Render  func()
	Destroy func()
	started bool
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
	smoothing = 0.5
}

func Register(o *GameObject) *GameObject {
	o.Id = idCounter
	o.started = false
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

func Stop(o *GameObject) {
	var next *list.Element
	for e := objects.Front(); e != nil; e = next {
		next = e.Next()
		item := e.Value.(*GameObject)

		if o.Id == item.Id {
			if item.Destroy != nil {
				item.Destroy()
			}

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
	o := &GameObject{Id: id}
	Stop(o)
}

func StopInput() {
	if inputLayer.Destroy != nil {
		inputLayer.Destroy()
	}

	inputLayer = GameObject{}
}

func StopRender() {
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
	for running {
		/* Start() */
		if inputLayer.Start != nil && !inputLayer.started {
			inputLayer.Start()
			inputLayer.started = true
		}

		for e := objects.Front(); e != nil; e = e.Next() {
			item := e.Value.(*GameObject)
			if item.Start != nil && !item.started {
				item.Start()
				item.started = true
			}
		}

		if renderLayer.Start != nil && !renderLayer.started {
			renderLayer.Start()
			renderLayer.started = true
		}
		/* Start() */

		/* FPS Calculation */
		now := time.Now()
		delta := now.Sub(prevTime).Seconds()
		prevTime = now

		current := 1 / delta
		fps = fps*smoothing + current*(1-smoothing)
		/* FPS Calculation */

		/* Update() */
		if inputLayer.Update != nil {
			inputLayer.Update()
		}

		for e := objects.Front(); e != nil; e = e.Next() {
			item := e.Value.(*GameObject)
			if item.Update != nil {
				item.Update()
			}
		}
		/* Update() */

		/* Physics() */
		for e := objects.Front(); e != nil; e = e.Next() {
			item := e.Value.(*GameObject)
			if item.Physics != nil {
				item.Physics()
			}
		}
		/* Physics() */

		/* Render() */
		for e := objects.Front(); e != nil; e = e.Next() {
			item := e.Value.(*GameObject)
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

func registerSpecial(o GameObject, message string) GameObject {
	if isGameObjectNil(o) {
		m := fmt.Sprintf("Trying to register a nil loopable as %v", message)
		panic(m)
	}

	o.started = false
	return o
}

func isGameObjectNil(o GameObject) bool {
	return o.Start == nil && o.Update == nil && o.Destroy == nil
}
