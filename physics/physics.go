package physics

import (
	"container/list"
	"fmt"

	"github.com/mikabrytu/gomes-engine/utils"
)

type RigidBody struct {
	Name string
	Rect *utils.RectSpecs
}

var bodies *list.List

func Init() {
	bodies = list.New()
}

func RegisterBody(b *utils.RectSpecs, name string) {
	bodies.PushFront(RigidBody{
		Name: name,
		Rect: b,
	})

	fmt.Printf("Registered body: %s\n", name)
}

func GetBodies() *list.List {
	return bodies
}
