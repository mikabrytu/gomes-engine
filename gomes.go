package gomesengine

import (
	"fmt"

	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
)

func HiGomes() {
	fmt.Println("Hi Gomes!")
}

func Init() {
	lifecycle.Init()

	specs := render.ScreenSpecs{
		Title:  "Gomes Engine",
		Posx:   0,
		Posy:   0,
		Width:  800,
		Height: 600,
	}
	render.CreateScreen(specs)
	lifecycle.Register(lifecycle.Loopable{
		Update:  render.Render,
		Destroy: render.Destroy,
	})
}

func Run() {
	lifecycle.Run()
}
