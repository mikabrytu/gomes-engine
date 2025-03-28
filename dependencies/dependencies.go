package dependencies

import "github.com/veandco/go-sdl2/sdl"

func Init() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}

func Quit() {
	sdl.Quit()
}
