package render

import (
	"github.com/mikabrytu/gomes-engine/debug"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type ScreenSpecs struct {
	Title  string
	Posx   int32
	Posy   int32
	Width  int32
	Height int32
}

type CopySpecs struct {
	sprite  *Sprite
	texture *sdl.Texture
	rect    *sdl.Rect
	color   Color
}

var window *sdl.Window
var renderer *sdl.Renderer
var toCreate []*Sprite
var toDestroy []*Sprite
var copies []CopySpecs
var backgroundColor Color

func CreateScreen(s ScreenSpecs) {
	var err error

	window, err = sdl.CreateWindow(s.Title, s.Posx, s.Posy, s.Width, s.Height, sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		panic(err)
	}

	toCreate = make([]*Sprite, 0)
	toDestroy = make([]*Sprite, 0)
	copies = make([]CopySpecs, 0)

	backgroundColor = Black
}

func Render() {
	// TODO: Check a better place for this
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			if debug.IsEnabled() {
				println("Quit")
			}

			lifecycle.StopRender()
			return
		}
	}

	renderer.SetDrawColor(
		backgroundColor.R,
		backgroundColor.G,
		backgroundColor.B,
		backgroundColor.A,
	)
	renderer.Clear()

	createTextures()

	for _, copy := range copies {
		if copy.sprite.enabled == false {
			continue
		}

		renderer.Copy(
			copy.texture,
			nil,
			copy.rect,
		)
	}

	renderer.Present()

	destroyTextures()
}

func RegisterSprite(sprite *Sprite) {
	toCreate = append(toCreate, sprite)
}

func ClearSprite(sprite *Sprite) {
	toDestroy = append(toDestroy, sprite)
}

func GetRenderer() *sdl.Renderer {
	return renderer
}

func SetBackgroundColor(color Color) {
	backgroundColor = color
}

func Destroy() {
	defer renderer.Destroy()
	defer window.Destroy()

	lifecycle.Kill()
}

func createTextures() {
	if len(toCreate) == 0 {
		return
	}

	for _, sprite := range toCreate {
		texture, err := img.LoadTexture(renderer, sprite.path)
		if err != nil {
			panic(err)
		}

		copy := CopySpecs{
			sprite:  sprite,
			texture: texture,
			rect: &sdl.Rect{
				X: int32(sprite.rect.PosX),
				Y: int32(sprite.rect.PosY),
				W: int32(sprite.rect.Width),
				H: int32(sprite.rect.Height),
			},
			color: sprite.color,
		}
		copies = append(copies, copy)
	}

	toCreate = toCreate[:0]
}

func destroyTextures() {
	if len(toDestroy) == 0 {
		return
	}

	for _, sprite := range toDestroy {
		for i, copy := range copies {
			if copy.sprite != sprite {
				continue
			}

			err := copy.texture.Destroy()
			if err != nil {
				panic(err)
			}

			copies = append(copies[:i], copies[i+1:]...)
			break
		}
	}

	toDestroy = toDestroy[:0]
}
