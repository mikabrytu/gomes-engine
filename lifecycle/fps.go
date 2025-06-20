package lifecycle

import "time"

var prevTime time.Time
var fps float64
var smoothing float64

func SetSmoothStep(step float64) {
	smoothing = step
}

func ShowFPS() float64 {
	return fps
}
