package gomesengine

import (
	"fmt"

	"github.com/mikabrytu/gomes-engine/dependencies"
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

	specs := render.ScreenSpecs{
		Title:  Title,
		Posx:   sdl.WINDOWPOS_CENTERED,
		Posy:   sdl.WINDOWPOS_CENTERED,
		Width:  ScreenWidth,
		Height: ScreenHeight,
	}
	render.CreateScreen(specs)
	lifecycle.Register(lifecycle.Loopable{
		Update:  render.Render,
		Destroy: render.Destroy,
	})
}

func Run() {
	defer dependencies.Quit()

	lifecycle.Run()
}
