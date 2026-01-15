package render

import (
	"github.com/mikabrytu/gomes-engine/debug"
	"github.com/mikabrytu/gomes-engine/lifecycle"
	"github.com/mikabrytu/gomes-engine/utils"
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

type CopySprite struct {
	sprite  *Sprite
	texture *sdl.Texture
	rect    *sdl.Rect
	color   Color
}

type CopyFont struct {
	font    *Font
	texture *sdl.Texture
	rect    *sdl.Rect
}

var window *sdl.Window
var renderer *sdl.Renderer

var spriteCreateQueue []*Sprite
var spriteDestroyQueue []*Sprite
var fontCreateQueue []*Font
var fontDestroyQueue []*Font
var spriteCopies []CopySprite
var fontCopies []CopyFont

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

	spriteCreateQueue = make([]*Sprite, 0)
	spriteDestroyQueue = make([]*Sprite, 0)
	fontCreateQueue = make([]*Font, 0)
	fontDestroyQueue = make([]*Font, 0)
	spriteCopies = make([]CopySprite, 0)
	fontCopies = make([]CopyFont, 0)

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

	createSprites()
	createFonts()

	for _, copy := range spriteCopies {
		if copy.sprite.enabled == false {
			continue
		}

		if copy.sprite.update {
			copy.sprite.update = false
			copy.rect = &sdl.Rect{
				X: int32(copy.sprite.rect.PosX),
				Y: int32(copy.sprite.rect.PosY),
				W: int32(copy.sprite.rect.Width),
				H: int32(copy.sprite.rect.Height),
			}
		}

		renderer.Copy(
			copy.texture,
			nil,
			copy.rect,
		)
	}

	for _, copy := range fontCopies {
		if copy.font.enabled == false {
			continue
		}

		if copy.font.update {
			copy.font.update = false
			copy.texture.SetColorMod(
				copy.font.color.R,
				copy.font.color.G,
				copy.font.color.B,
			)

			copy.rect = &sdl.Rect{
				X: int32(copy.font.rect.PosX),
				Y: int32(copy.font.rect.PosY),
				W: int32(copy.font.rect.Width),
				H: int32(copy.font.rect.Height),
			}
		}

		renderer.Copy(
			copy.texture,
			nil,
			copy.rect,
		)
	}

	renderer.Present()

	destroySprites()
	destroyFonts()
}

func RegisterSprite(sprite *Sprite) {
	spriteCreateQueue = append(spriteCreateQueue, sprite)
}

func RegisterFont(font *Font) {
	fontCreateQueue = append(fontCreateQueue, font)
}

func ClearSprite(sprite *Sprite) {
	spriteDestroyQueue = append(spriteDestroyQueue, sprite)
}

func ClearFont(font *Font) {
	fontDestroyQueue = append(fontDestroyQueue, font)
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

func createSprites() {
	if len(spriteCreateQueue) == 0 {
		return
	}

	for _, sprite := range spriteCreateQueue {
		texture, err := img.LoadTexture(renderer, sprite.path)
		if err != nil {
			panic(err)
		}

		copy := CopySprite{
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
		spriteCopies = append(spriteCopies, copy)
	}

	spriteCreateQueue = spriteCreateQueue[:0]
}

func createFonts() {
	if len(fontCreateQueue) == 0 {
		return
	}

	for _, font := range fontCreateQueue {
		csdl := sdl.Color{
			R: font.color.R,
			G: font.color.G,
			B: font.color.B,
			A: font.color.A,
		}

		surface, err := font.instance.RenderUTF8Blended(font.text, csdl)
		if err != nil {
			panic(err)
		}

		texture, err := renderer.CreateTextureFromSurface(surface)
		if err != nil {
			panic(err)
		}

		copy := CopyFont{
			font:    font,
			texture: texture,
			rect: &sdl.Rect{
				X: int32(font.position.X),
				Y: int32(font.position.Y),
				W: int32(surface.W),
				H: int32(surface.H),
			},
		}

		font.rect = utils.RectSpecs{
			PosX:   int(copy.rect.X),
			PosY:   int(copy.rect.Y),
			Width:  int(copy.rect.W),
			Height: int(copy.rect.H),
		}
		fontCopies = append(fontCopies, copy)

		surface.Free()
	}

	fontCreateQueue = fontCreateQueue[:0]
}

func destroySprites() {
	if len(spriteDestroyQueue) == 0 {
		return
	}

	for _, sprite := range spriteDestroyQueue {
		for i, copy := range spriteCopies {
			if copy.sprite != sprite {
				continue
			}

			err := copy.texture.Destroy()
			if err != nil {
				panic(err)
			}

			spriteCopies = append(spriteCopies[:i], spriteCopies[i+1:]...)
			break
		}
	}

	spriteDestroyQueue = spriteDestroyQueue[:0]
}

func destroyFonts() {
	if len(fontDestroyQueue) == 0 {
		return
	}

	for _, font := range fontDestroyQueue {
		for i, copy := range fontCopies {
			if copy.font != font {
				continue
			}

			err := copy.texture.Destroy()
			if err != nil {
				panic(err)
			}

			fontCopies = append(fontCopies[:i], fontCopies[i+1:]...)
			break
		}
	}

	fontDestroyQueue = fontDestroyQueue[:0]
}
