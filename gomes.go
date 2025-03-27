package gomesengine

import (
	"fmt"

	"github.com/mikabrytu/gomes-engine/dependencies"
	"github.com/mikabrytu/gomes-engine/input"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/veandco/go-sdl2/sdl"
)

func HiGomes() {
	fmt.Println("Hi Gomes!")
}

func Init(Title string, ScreenWidth, ScreenHeight int32) {
	dependencies.Init()
	lifecycle.Init()

	setupScreen(Title, ScreenWidth, ScreenHeight)
	setupInput()
}

func Run() {
	defer dependencies.Quit()

	lifecycle.Run()
}

func setupScreen(title string, width, height int32) {
	specs := render.ScreenSpecs{
		Title:  title,
		Posx:   sdl.WINDOWPOS_CENTERED,
		Posy:   sdl.WINDOWPOS_CENTERED,
		Width:  width,
		Height: height,
	}
	render.CreateScreen(specs)
	lifecycle.RegisterLast(lifecycle.Loopable{
		Update:  render.Render,
		Destroy: render.Destroy,
	})
}

func setupInput() {
	lifecycle.RegisterFirst(lifecycle.Loopable{
		Update: input.ListenToInput,
	})
}
