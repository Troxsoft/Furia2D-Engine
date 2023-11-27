MIT LICENCE


# Furia2D-Engine
Un motor de videojuego escrito usando Golang 100% usando la poderosa libreria Raylib
                  Y MUY FACIL
### Ejemplos
- mover personaje con un enemigo
```go
package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

func main() {
	e.InitGame("controlador basico :)", e.NewSize(600, 400), func() {
		//init
		e.CreateGameObject("player", e.SHAPE_IMAGE, e.NewSize(50, 50), e.NewPosition(30, 30))
		e.CreateGameObject("enemigo", e.SHAPE_RECTANGLE, e.NewSize(30, 30), e.NewPosition(100, 100))
		e.CreateGameObject("aliado", e.SHAPE_RECTANGLE, e.NewSize(50, 50), e.NewPosition(40, 80))
		e.SetColorGameObject("aliado", e.NewColor2(0, 0, 255))
		e.SetColorGameObject("enemigo", e.NewColor2(255, 0, 0))
		e.SetImageGameObject("player", e.NewImage("sapo.jpg"))
		e.AddToGroupGameObject("enemigo", "ene")
		e.SetStartGameObject("player", func(g *e.GameObject, a any) {
			g.SetPosition(e.NewPosition(0, a.(int32)))
		})
		e.SetUpdateGameObject("player", func(g *e.GameObject, a *e.Event) {
			if collisi, _ := a.OnCollision(); collisi != nil {
				if collisi.IsInGroup("ene") {
					g.SetColor3(255, 0, 0)
				} else {
					g.SetColor3(255, 255, 255)
				}
			} else {
				g.SetColor3(255, 255, 255)
			}

			//a.OnCollisionTo()
			if a.IsKeyDown(e.KeyLeft) {
				g.SetPosition2(g.Position().X-2, g.Position().Y)
			}
			if a.IsKeyDown(e.KeyRight) {
				g.SetPosition2(g.Position().X+2, g.Position().Y)
			}
			if a.IsKeyDown(e.KeyUp) {
				g.SetPosition2(g.Position().X, g.Position().Y-2)
			}
			if a.IsKeyDown(e.KeyDown) {
				g.SetPosition2(g.Position().X, g.Position().Y+2)
			}
		})
		e.InstanceGameObject("enemigo", nil)
		e.InstanceGameObject("player", int32(30))
		e.InstanceGameObject("aliado", nil)

	}, func() {
		//update

	})

}
```
    
