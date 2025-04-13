package audio

import (
	"container/list"
	"fmt"

	"github.com/veandco/go-sdl2/mix"
)

var buffer *list.List

func Init() {
	if buffer == nil {
		buffer = list.New()
	}
}

func Play(path string) {
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

func ClearBuffer() {
	for e := buffer.Front(); e != nil; e = e.Next() {
		s := e.Value.(*mix.Chunk)
		s.Free()
		s = nil

		fmt.Println("Buffer cleared a single chunk from memory")
	}

	buffer = list.New()
}
