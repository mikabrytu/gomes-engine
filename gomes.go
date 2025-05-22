package gomesengine

import (
	"fmt"

	"github.com/mikabrytu/gomes-engine/audio"
	"github.com/mikabrytu/gomes-engine/dependencies"
	"github.com/mikabrytu/gomes-engine/events"
	"github.com/mikabrytu/gomes-engine/input"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/physics"
	"github.com/mikabrytu/gomes-engine/render"
	"github.com/veandco/go-sdl2/sdl" // This is just for the centered window position flag. Games will not read this
)

func HiGomes() {
	fmt.Println("Hi Gomes!")
}

func Init(Title string, ScreenWidth, ScreenHeight int32) {
	dependencies.Init()
	events.Init()
	lifecycle.Init()

	setupScreen(Title, ScreenWidth, ScreenHeight)
	setupInput()
	setupPhysics()
	setupAudio()
}

func Run() {
	defer audio.StopSoundtrack(true)
	defer audio.ClearBuffer()
	defer dependencies.Quit()

	lifecycle.Run()
}

func setupScreen(title string, width, height int32) {
	specs := render.ScreenSpecs{
		Title:  title,
		Posx:   sdl.WINDOWPOS_CENTERED, // TODO: Make flags default positions
		Posy:   sdl.WINDOWPOS_CENTERED,
		Width:  width,
		Height: height,
	}
	render.CreateScreen(specs)
	lifecycle.RegisterRender(lifecycle.GameObject{
		Update:  render.Render,
		Destroy: render.Destroy,
	})
}

func setupInput() {
	lifecycle.RegisterInput(lifecycle.GameObject{
		Update: input.ListenToInput,
	})
}

func setupPhysics() {
	physics.Init()
}

func setupAudio() {
	audio.Init()
}
