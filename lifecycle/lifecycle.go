package lifecycle

import (
	"container/list"
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Loopable struct {
	Id      int
	Init    func()
	Update  func()
	Destroy func()
}

var idCounter = 0
var running bool = false
var loopables *list.List
var first Loopable
var last Loopable

// Public API

func Init() {
	loopables = list.New()
	first = Loopable{}
	last = Loopable{}
	idCounter = 0
	running = true
}

func Register(l Loopable) Loopable {
	l.Id = idCounter
	idCounter++

	_ = loopables.PushFront(l)

	return l
}

func RegisterFirst(l Loopable) {
	first = registerSpecial(l, "First")
}

func RegisterLast(l Loopable) {
	last = registerSpecial(l, "Last")
}

func StopById(id int) {
	l := Loopable{Id: id}
	Stop(l)
}

func Stop(l Loopable) {
	var next *list.Element
	for e := loopables.Front(); e != nil; e = next {
		next = e.Next()
		item := Loopable(e.Value.(Loopable))

		if l.Id == item.Id {
			item.Destroy()
			loopables.Remove(e)
			break
		}
	}

	if loopables.Len() == 0 {
		fmt.Println("There's no more loopables on the list. Quitting application")
		running = false
	}
}

func Kill() {
	loopables.Init()
	running = false
}

func Run() {
	if running {
		if first.Init != nil {
			first.Init()
		}

		for e := loopables.Front(); e != nil; e = e.Next() {
			item := Loopable(e.Value.(Loopable))
			if item.Init != nil {
				item.Init()
			}
		}

		if last.Init != nil {
			last.Init()
		}
	}

	for running {
		if first.Update != nil {
			first.Update()
		}

		for e := loopables.Front(); e != nil; e = e.Next() {
			item := Loopable(e.Value.(Loopable))
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

func registerSpecial(l Loopable, message string) Loopable {
	if isLoopableNil(l) {
		m := fmt.Sprintf("Trying to register a nil loopable as %v", message)
		panic(m)
	}

	return l
}

func isLoopableNil(l Loopable) bool {
	return l.Init == nil && l.Update == nil && l.Destroy == nil
}
