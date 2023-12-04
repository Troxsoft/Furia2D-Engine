package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var obj *e.GameObject

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {
		obj, _ = e.CreateGameObject("you", e.SHAPE_RECTANGLE, e.NewSize(30, 30), e.NewPosition(30, 30))
		obj.Instance(e.GetCurrentScene(), nil)

		//fmt.Println(obj.F)
	},
		func(ge *e.GameEvent) {

		})
}
