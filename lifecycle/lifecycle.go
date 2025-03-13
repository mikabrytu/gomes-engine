package lifecycle

import (
	"container/list"
	"fmt"
)

type Loopable struct {
	Id     int
	Update func()
}

var idCounter = 0
var running bool = false
var loopables *list.List

func Init() {
	loopables = list.New()
	running = true
}

func Register(l Loopable) Loopable {
	idCounter++
	l.Id = idCounter
	_ = loopables.PushBack(l)

	return l
}

func Stop(l Loopable) {
	var next *list.Element
	for e := loopables.Front(); e != nil; e = next {
		next = e.Next()
		item := Loopable(e.Value.(Loopable))

		if l.Id == item.Id {
			loopables.Remove(e)
			break
		}
	}

	if loopables.Len() == 0 {
		fmt.Println("There's no more loopables on the list. Quitting application")
		running = false
	}
}

func Run() {
	for running {
		for e := loopables.Front(); e != nil; e = e.Next() {
			item := Loopable(e.Value.(Loopable))
			item.Update()
		}
	}
}
