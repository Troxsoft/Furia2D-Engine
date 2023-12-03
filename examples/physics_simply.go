package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var obj *e.GameObject
var obj2 *e.GameObject

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(600, 600), func(ge *e.GameEvent) {
		obj, _ = e.CreateGameObject("you", e.SHAPE_RECTANGLE, e.NewSize(30, 30), e.NewPosition(200, 30))
		obj2, _ = e.CreateGameObject("you not is enemy", e.SHAPE_RECTANGLE, e.NewSize(50, 50), e.NewPosition(300, 30))
		obj.SetUpdate(func(g *e.GameObject, goe *e.GameObjectEvent) {
			g.SetPosition(e.ClampPos(g.Position(), e.NewPosition(600, 600), e.NewPosition(0, 0)))
			if ge.IsKeyDown(ge.KeyRight) {
				objetoColision := g.MoveTo(e.NewPosition(g.Position().X+10, g.Position().Y))
				if objetoColision != nil {
					g.SetColor3(255, 0, 0)
				} else {
					g.SetColor3(0, 0, 0)
				}
			}
			if ge.IsKeyDown(ge.KeyLeft) {
				g.MoveTo(e.NewPosition(g.Position().X-10, g.Position().Y))
				g.SetColor3(0, 0, 0)
			}

		})
		obj.Instance(nil)
		obj2.Instance(nil)
		//fmt.Println(obj.F)
	},
		func(ge *e.GameEvent) {

		})
}
