package physics

import (
	"container/list"
	"fmt"

	"github.com/mikabrytu/gomes-engine/math"
	"github.com/mikabrytu/gomes-engine/utils"
)

type RigidBody struct {
	Name      string
	Rect      *utils.RectSpecs
	Axis      math.Vector2
	IsDynamic bool
}

var bodies *list.List

func Init() {
	bodies = list.New()
}

func RegisterBody(b *utils.RectSpecs, name string) RigidBody {
	body := RigidBody{
		Name:      name,
		Rect:      b,
		Axis:      math.Vector2{X: 0, Y: 0},
		IsDynamic: false,
	}

	bodies.PushFront(body)
	fmt.Printf("Registered body: %s\n", name)

	return body
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

func EnableDynamicCollision(body *RigidBody) {
	body.IsDynamic = true
}

func DisableDynamicCollision(body *RigidBody) {
	body.IsDynamic = false
}

func CheckCollision(collider *RigidBody) RigidBody {
	if (bodies == nil) || (bodies.Len() == 0) {
		panic("No bodies registered")
	}

	result := RigidBody{
		Name: "nil",
	}

	for e := bodies.Front(); e != nil; e = e.Next() {
		target := e.Value.(RigidBody)

		// TODO: Change to pointer check
		if target.Name == collider.Name {
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

func ResolveDynamicCollisions(body *RigidBody) {
	if !body.IsDynamic {
		fmt.Printf("Body %v is not set to Dynamic resolution.\n", body.Name)
		return
	}

	collision := CheckCollision(body)

	if collision.Name == "nil" {
		return
	}

	if collision.Rect.PosX < body.Rect.PosX {
		body.Axis.X = 1
	}

	if collision.Rect.PosX > body.Rect.PosX {
		body.Axis.X = -1
	}

	if collision.Rect.PosY < body.Rect.PosY {
		body.Axis.Y = 1
	}

	if collision.Rect.PosY > body.Rect.PosY {
		body.Axis.Y = -1
	}
}
