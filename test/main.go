package main

import (
	"math/rand"

	gomesengine "github.com/mikabrytu/gomes-engine"

	"github.com/mikabrytu/gomes-engine/debug"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/mikabrytu/gomes-engine/utils"
)

var SCREEN_SIZE = math.Vector2{
	X: 800,
	Y: 600,
}

const REPEAT_EVENT string = "REPEAT_EVENT"

func main() {
	gomesengine.HiGomes()
	gomesengine.Init("Version 1.4 - DEVELOP", int32(SCREEN_SIZE.X), int32(SCREEN_SIZE.Y))
	debug.EnableDebug()

	events.AddListener(events.Input, events.INPUT_KEYBOARD_PRESSED_ESCAPE, func(data any) {
		lifecycle.Kill()
	})

	objs := make([]*obj, 0)
	events.AddListener(events.Input, events.INPUT_KEYBOARD_PRESSED_F, func(data any) {
		println()
		objs = append(objs, New())
	})

	events.AddListener(events.Input, events.INPUT_KEYBOARD_PRESSED_SPACE, func(data any) {
		index := rand.Intn(len(objs))
		o := objs[index]

		if o != nil {
			lifecycle.Stop(o.Instace)
			o = nil
		} else {
			println("Nil object. Try again")
		}
	})

	gomesengine.Run()
}

type obj struct {
	Instace  *lifecycle.GameObject
	Listener *events.EventListener
	rect     utils.RectSpecs
	name     string
	color    render.Color
}

func New() *obj {
	obj := &obj{
		rect: utils.RectSpecs{
			PosX:   rand.Intn(SCREEN_SIZE.X - 64),
			PosY:   rand.Intn(SCREEN_SIZE.Y - 64),
			Width:  64,
			Height: 64,
		},
	}

	obj.Listener = &events.EventListener{Id: 1001}
	obj.Instace = lifecycle.Register(&lifecycle.GameObject{
		Start: func() {
			obj.Listener = events.AddListener(events.Input, events.INPUT_MOUSE_CLICK_DOWN, func(data any) {
				click := data.(events.InputMouseClickDownEvent)

				if click.Position.X > obj.rect.PosX && click.Position.X < (obj.rect.PosX+obj.rect.Width) &&
					click.Position.Y > obj.rect.PosY && click.Position.Y < (obj.rect.PosY+obj.rect.Height) {
					println("clicked at", obj.name)
				}

			})
		},
		Destroy: func() {
			println("Removing Listener ", obj.Listener.Id)
			events.RemoveListener(events.Input, events.INPUT_MOUSE_CLICK_DOWN, obj.Listener.Id)
		},
		Render: func() {
			render.DrawRect(obj.rect, render.White)
		},
	})

	return obj
}

func (o *obj) GetListenerId() uint64 {
	return o.Listener.Id
}
