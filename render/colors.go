package render

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

var Red = Color{R: 255, G: 0, B: 0, A: 255}
var Green = Color{R: 0, G: 255, B: 0, A: 255}
var Blue = Color{R: 0, G: 0, B: 255, A: 255}
var Black = Color{R: 0, G: 0, B: 0, A: 255}
var White = Color{R: 255, G: 255, B: 255, A: 255}
var Yellow = Color{R: 255, G: 255, B: 0, A: 255}
var Cyan = Color{R: 0, G: 255, B: 255, A: 255}
var Magenta = Color{R: 255, G: 0, B: 255, A: 255}
var Gray = Color{R: 128, G: 128, B: 128, A: 255}
var LightGray = Color{R: 192, G: 192, B: 192, A: 255}
var DarkGray = Color{R: 64, G: 64, B: 64, A: 255}
var Orange = Color{R: 255, G: 200, B: 0, A: 255}
var Brown = Color{R: 128, G: 64, B: 0, A: 255}
var Pink = Color{R: 255, G: 175, B: 175, A: 255}
var Purple = Color{R: 128, G: 0, B: 128, A: 255}
var Transparent = Color{R: 0, G: 0, B: 0, A: 0}
