package main

import (
	gomesengine "github.com/mikabrytu/gomes-engine"
	"github.com/mikabrytu/gomes-engine/screen"
	"github.com/veandco/go-sdl2/sdl"
)

func main() {
	gomesengine.HiGomes()

	specs := screen.Specs{
		Title:  "Gomes Demo",
		Posx:   sdl.WINDOWPOS_UNDEFINED,
		Posy:   sdl.WINDOWPOS_UNDEFINED,
		Width:  800,
		Height: 600,
	}
	screen.CreateScreen(specs)
}
