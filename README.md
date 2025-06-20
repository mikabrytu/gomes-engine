# Gomes Engine
**Gomes Engine** is a simple 2D game engine focused on prototype and quick iteration, using GO as the language and SDL for rendering behind the scenes.

You need [GO installed](https://go.dev/doc/install) to use this engine.  
If you need help with GO, check the [official documentation](https://go.dev/doc/)

## How To Install
You can install Gomes by creating a new go project and importing the current version

```
go mod init my_project
go get github.com/mikabrytu/gomes-engine
```

You can test if everything is working by running this code in your _main.go_ file

```go
package main

import (
	gomesengine "github.com/mikabrytu/gomes-engine"
)

func main() {
	gomesengine.HiGomes()
}
```

### How To Use
Gomes idea is to be fast and simple, so the idea is to render your idea and gameplay prototype as quick as possible. But you can still use it to full games, just keep in mind this is a simple engine that are still being developed.  
Here's a suggestion on how to organize your project to do that.

In your `main()` function you just need to call 2 functions to show the screen and start the game loop
```go
package main

import (
	gomesengine "github.com/mikabrytu/gomes-engine"
)

func main() {
  gomesengine.Init("MyGame", 800, 600)
	gomesengine.Run()
}
```

The `Run()` function starts the loop so after that nothing will be called until the game stops, so you need to start you code between `Init()` and `Run()`, like this:

```go
func main() {
  gomesengine.Init("MyGame", 800, 600)
  myGame()
	gomesengine.Run()
}

func myGame() {
  print("Hello Game!")
}
```

You can find examples here:
- [Genius](https://github.com/mikabrytu/gomes-genius)
- [Gong](https://github.com/mikabrytu/gong)

#### GameObject
**GameObject** is the data that will be checked every frame to run logic and render code, similar to concepts found in [Unity](https://docs.unity3d.com/ScriptReference/GameObject.html), for example  
To insert new objects in the gameloop, you need to register them in the lifecycle by calling:

```
lifecycle.Register(lifecycle.GameObject{
  Start:   func() {
    print("Object Instantiated!")
  },
})
```

Because GO doesn't have inheritance like other common languages such as Java or C#, you can create a file that holds a struct as templates and create new pointers so you can easily track and manage your GameObject instances. 
Here is an example:

_box.go_
```go
package objects

import (
  "github.com/mikabrytu/gomes-engine/lifecycle"
  "github.com/mikabrytu/gomes-engine/render"
  "github.com/mikabrytu/gomes-engine/utils"
)

type Box struct {
  rect utils.RectSpecs
  color render.Color
}

func NewBox(rect utils.RectSpecs, color render.Color) *Box{
  box := &Box {
    rect: rect,
    color: color,
  }

  lifecycle.Register(lifecycle.GameObject{
    Render:  ball.render,
  })

  return box
}

func (b *Box) render() {
  render.DrawSimpleShapes(b.rect, b.color)
}
```

_main.go_
```go
func myGame() {
  // Instantiating a new ball
  ball = objects.NewBall(rect, render.White)
}
```

### Layer Reference
In Gomes you can find specific functions and tools on **layers**. Each layer is responsible for an area of the game.

#### Lifecycle
Lifecycle is responsible for the game loop and registering game objects that should be updated every frame.
This layer maintain a list of GameObject that run functions like, `Start()` and `Update()`.
All GameObject functions:

- `Start()` -> Runs on the first frame the object is added to the game loop
- `Update()` -> Runs every frame
- `Physics()` -> Runs every frame
- `Render()` -> Runs every frame
- `Destroy()` -> Runs once the frame before the object will be removed from the loop list

This is the order of execution of a frame in Gomes.

#### Render
Layer responsible for render the object.
Gomes are currently using [SDL2](https://www.libsdl.org) as the rendering library but no SDL knowledge is required to use. 
Keep in mind thougn that any limitation found on SDL2 will be present on Gomes as well

#### Physics
Layer responsible for physics simulations. 
Is recommended that all calculations regarding movement, collision detection or resolution should happen on the `Physics()` function of the GameObject

#### Input
Layer that listen to inputs and fire events for the user to listen to, like button `INPUT_MOUSE_CLICK`

#### Events
Gomes uses events to communicate between layers and is currently using the [go-event](github.com/AlexanderGrom/go-event) library for that.
There's already some events registered but you can create your own like this:

```go
package main

import "github.com/mikabrytu/gomes-engine/events"

var id string = "EVENT_ID"

func main() {
  events.Emit(id)
  events.Subscribe(id, func(params ...any) error {
    print("Event Fired!")
    return nil
  })
}
```

#### Audio
Audio can store data for sfx and soundtrack and play/resume then as you like.

#### UI
Layer responsible for setting up fonts and menus.
It's important to note that UI doesn't render anything. It just compiles the data and send to the render layer to present on the screen
