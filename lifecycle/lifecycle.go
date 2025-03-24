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

func Init() {
	loopables = list.New()
	idCounter = 0
	running = true
}

func Register(l Loopable) Loopable {
	l.Id = idCounter
	idCounter++

	_ = loopables.PushFront(l)

	return l
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
		for e := loopables.Front(); e != nil; e = e.Next() {
			item := Loopable(e.Value.(Loopable))
			if item.Init != nil {
				item.Init()
			}
		}
	}

	for running {
		for e := loopables.Front(); e != nil; e = e.Next() {
			item := Loopable(e.Value.(Loopable))
			if item.Update != nil {
				item.Update()
			}
		}

		sdl.Delay(33)
	}
}
