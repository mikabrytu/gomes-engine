package audio

import (
	"container/list"

	"github.com/veandco/go-sdl2/mix"
)

var buffer *list.List
var soundtrack *mix.Music

func Init() {
	if buffer == nil {
		buffer = list.New()
	}
}

func PlaySFX(path string) {
	checkBuffer()

	s, err := mix.LoadWAV(path)
	if err != nil {
		panic(err)
	}

	s.Play(-1, 0)

	if buffer == nil {
		panic("Audio buffer not initialized")
	}

	buffer.PushBack(s)
}

func PlaySoundtrack(path string) {
	var err error

	soundtrack, err = mix.LoadMUS(path)
	if err != nil {
		panic(err)
	}

	soundtrack.Play(0)
}

func PauseSoundtrack() {
	mix.PauseMusic()
}

func ResumeSoundtrack() {
	mix.ResumeMusic()
}

func IsSoundtrackPlaying() bool {
	return mix.PlayingMusic() && !mix.PausedMusic()
}

func StopSoundtrack(finish bool) {
	if !finish {
		mix.HaltMusic()
	}

	soundtrack.Free()
	soundtrack = nil
}

func ClearBuffer() {
	for e := buffer.Front(); e != nil; e = e.Next() {
		s := e.Value.(*mix.Chunk)
		s.Free()
		s = nil
	}

	buffer = list.New()
	println("Audio buffer cleared")
}

func checkBuffer() {
	if buffer.Len() > 16 {
		ClearBuffer()
	}
}
