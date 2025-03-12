package lifecycle

type Loopable struct {
	Update func()
}

var running bool = false

func Run(l Loopable) {
	running = true

	for running {
		l.Update()
	}
}

func Stop() {
	running = false
}
