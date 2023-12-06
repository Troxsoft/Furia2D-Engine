package main

import (
	"fmt"

	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var obj *e.GameObject
var obj2 *e.GameObject

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(700, 400), func(ge *e.GameEvent) {
		obj, _ = e.CreateGameObject("you", e.SHAPE_RECTANGLE, e.NewSize(30, 30), e.NewPosition(30, 30))
		obj = obj.Instance(e.GetCurrentScene(), nil)

		obj2, _ = e.CreateGameObject("youe", e.SHAPE_RECTANGLE, e.NewSize(30, 30), e.NewPosition(30, 30))
		obj2.SetColor(e.BROWN)
		obj2 = obj2.Instance(e.GetCurrentScene(), nil)

		//fmt.Println(obj.F)
		e.NewLinealPath(e.NewPos(0, 50), e.NewPos(500, 50)).TeleportToPath(obj)
		e.NewLinealPath(e.NewPos(500, 100), e.NewPos(0, 100)).TeleportToPath(obj2)
	},
		func(ge *e.GameEvent) {

			e.NewLinealPath(e.NewPos(0, 50), e.NewPos(500, 50)).Travel(2, obj)
			e.NewLinealPath(e.NewPos(500, 100), e.NewPos(0, 100)).Travel(-2, obj2)
			fmt.Println(e.NewLinealPath(e.NewPos(0, 50), e.NewPos(500, 50)).IsDone(obj))
			fmt.Println(e.NewLinealPath(e.NewPos(500, 100), e.NewPos(0, 100)).IsDone(obj2))
		})
}
