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

func GetBodyByName(name string) RigidBody {
	if (bodies == nil) || (bodies.Len() == 0) {
		panic("No bodies registered")
	}

	for e := bodies.Front(); e != nil; e = e.Next() {
		item := e.Value.(RigidBody)

		if item.Name == name {
			return item
		}
	}

	return RigidBody{
		Name: "nil",
	}
}

func GetBodyByRect(rect *utils.RectSpecs) RigidBody {
	if (bodies == nil) || (bodies.Len() == 0) {
		panic("No bodies registered")
	}

	for e := bodies.Front(); e != nil; e = e.Next() {
		item := e.Value.(RigidBody)

		if item.Rect == rect {
			return item
		}
	}

	return RigidBody{
		Name: "nil",
	}
}

func CheckCollision(collider RigidBody) RigidBody {
	if (bodies == nil) || (bodies.Len() == 0) {
		panic("No bodies registered")
	}

	result := RigidBody{
		Name: "nil",
	}

	for e := bodies.Front(); e != nil; e = e.Next() {
		target := e.Value.(RigidBody)

		if target.Name == collider.Name { // Change to pointer check
			continue
		}

		cLeft := collider.Rect.PosX
		cTop := collider.Rect.PosY
		cRight := collider.Rect.PosX + collider.Rect.Width
		cBottom := collider.Rect.PosY + collider.Rect.Height

		tLeft := target.Rect.PosX
		tTop := target.Rect.PosY
		tRight := target.Rect.PosX + target.Rect.Width
		tBottom := target.Rect.PosY + target.Rect.Height

		if cBottom < tTop || cTop > tBottom || cRight < tLeft || cLeft > tRight {
			continue
		}

		result = target
	}

	return result
}
