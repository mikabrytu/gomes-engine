package math

type Vector2 struct {
	X int
	Y int
}

func ClampInt(n int) int {
	if n < -1 {
		return -1
	}

	if n > 1 {
		return 1
	}

	return n
}
