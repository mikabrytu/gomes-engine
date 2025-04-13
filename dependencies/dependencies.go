package dependencies

import (
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

func Init() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	if err := mix.Init(mix.INIT_OGG); err != nil {
		panic(err)
	}

	if err := mix.OpenAudio(
		mix.DEFAULT_FREQUENCY,
		mix.DEFAULT_FORMAT,
		mix.DEFAULT_CHANNELS,
		mix.DEFAULT_CHUNKSIZE,
	); err != nil {
		panic(err)
	}
}

func Quit() {
	mix.CloseAudio()
	mix.Quit()
	sdl.Quit()
}
