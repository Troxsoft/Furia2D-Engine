MIT LICENCE


# Furia2D-Engine
A game engine written using Golang 100% using the powerful Raylib library
                  AND VERY EASY



![Raylib](https://github.com/raysan5/raylib)
![Raylib-go](https://github.com/gen2brain/raylib-go)

### Examples
- Hello rectangle 
```go
package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var obj *e.GameObject

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {
		obj, _ = e.CreateGameObject("you object", e.SHAPE_RECTANGLE, e.NewSize(30, 30), e.NewPosition(30, 30))
		obj.Instance(e.GetCurrentScene(),nil)

		//fmt.Println(obj.F)
	},
		func(ge *e.GameEvent) {

		})
}

```
- Instances/reuse code/inheritance ???
```go
package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var obj *e.GameObject

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {
		obj, _ = e.CreateGameObject("you", e.SHAPE_RECTANGLE, e.NewSize(30, 30), e.NewPosition(30, 30))
		obj.SetStart(func(g *e.GameObject, a any) {
			g.SetPosition(a.(e.Position))
			g.SetColor(e.GREEN)
		})
		obj.Instance(e.GetCurrentScene(), e.NewPos(20, 20))
		obj.Instance(e.GetCurrentScene(), e.NewPos(60, 60))
		//fmt.Println(obj.F)
	},
		func(ge *e.GameEvent) {

		})
}

```
- more examples in'examples'
#### FURIA2D is easy or very easy ???
