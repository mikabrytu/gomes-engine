package gomesengine

import (
	"fmt"

	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/screen"
)

var specs screen.Specs

func HiGomes() {
	fmt.Println("Hi Gomes!")
}

func Init() {
	lifecycle.Init()

	specs = screen.Specs{
		Title:  "Gomes Engine",
		Posx:   0,
		Posy:   0,
		Width:  800,
		Height: 600,
	}
	screen.CreateScreen(specs)
	lifecycle.Register(lifecycle.Loopable{
		Update:  screen.Render,
		Destroy: screen.Destroy,
	})
}

func Run() {
	lifecycle.Run()
}
