package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var player *e.GameObject
var piso *e.GameObject

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {
		player, _ = e.CreateGameObject("player", e.SHAPE_RECTANGLE, e.NewSize(30, 30), e.NewPosition(30, 30))
		player.Instance(nil)
		piso, _ = e.CreateGameObject("piso", e.SHAPE_RECTANGLE, e.NewSize(900, 50), e.NewPosition(0, 350))
		piso.SetUpdate(func(pl *e.GameObject, goe *e.GameObjectEvent) {
			//pl.MoveTo(e.ClampPos())

		})
		piso.Instance(nil)
		//fmt.Println(obj.F)
	},
		func(ge *e.GameEvent) {

		})
}
