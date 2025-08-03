package debug

var enabled bool

func EnableDebug() {
	enabled = true
}

func DisableDebug() {
	enabled = false
}

func IsEnabled() bool {
	return enabled
}
