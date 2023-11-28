package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var (
	playerG  *e.GameObject
	enemigoG *e.GameObject
	aliadoG  *e.GameObject
)

func main() {
	e.InitGame("controlador basico :)", e.NewSize(600, 400), func(eg *e.GameEvent) {
		//init
		playerG, _ = e.CreateGameObject("player", e.SHAPE_IMAGE, e.NewSize(200, 200), e.NewPosition(30, 30))
		enemigoG, _ = e.CreateGameObject("enemigo", e.SHAPE_RECTANGLE, e.NewSize(30, 30), e.NewPosition(100, 100))
		aliadoG, _ = e.CreateGameObject("aliado", e.SHAPE_RECTANGLE, e.NewSize(50, 50), e.NewPosition(40, 80))
		//playerG.Hide()
		aliadoG.SetColor(e.NewColor2(0, 0, 255))
		enemigoG.SetColor(e.NewColor2(255, 0, 0))
		playerG.SetImage(e.NewImage("sapo.jpg"))
		enemigoG.AddToGroup("ene")
		playerG.SetStart(func(g *e.GameObject, a any) {
			g.SetPosition(e.NewPosition(0, a.(int32)))
		})
		playerG.SetUpdate(func(g *e.GameObject, a *e.GameObjectEvent) {
			if collisi, _ := a.OnCollisionInTheGroup("ene"); collisi != nil {
				g.SetColor3(255, 0, 0)
			} else {
				g.SetColor3(255, 255, 255)
				if a.IsMouseDown(a.MouseButtonRight) {
					g.SetColor3(255, 0, 0)
				} else {
					g.SetColor3(255, 255, 255)
				}
			}
			if eg.IsKeyDown(eg.KeyLeft) {
				g.SetPosition2(g.Position().X-2, g.Position().Y)
			}
			if eg.IsKeyDown(eg.KeyRight) {
				g.SetPosition2(g.Position().X+2, g.Position().Y)
			}
			if eg.IsKeyDown(eg.KeyUp) {
				g.SetPosition2(g.Position().X, g.Position().Y-2)
			}
			if eg.IsKeyDown(eg.KeyDown) {
				g.SetPosition2(g.Position().X, g.Position().Y+2)
			}
		})
		e.InstanceGameObject("enemigo", nil)
		e.InstanceGameObject("player", int32(30))
		e.InstanceGameObject("player", int32(300))
		e.InstanceGameObject("aliado", nil)
	}, func(eg *e.GameEvent) {
		//update

	})
}
