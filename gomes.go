package gomesengine

import (
	"fmt"
	"time"

	"github.com/mikabrytu/gomes-engine/screen"
)

var specs screen.Specs
var running bool

func HiGomes() {
	fmt.Println("Hi Gomes!")
}

func Init() {
	specs = screen.Specs{
		Title:  "Gomes Engine",
		Posx:   0,
		Posy:   0,
		Width:  800,
		Height: 600,
	}
	screen.CreateScreen(specs)
}

func Run() {
	running = true

	go Destroy()

	for running {
		screen.Render()
	}

	screen.Destroy()
}

func Destroy() {
	time.Sleep(5 * time.Second)
	running = false
}
