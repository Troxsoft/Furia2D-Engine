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
