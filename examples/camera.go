package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var obj *e.GameObject

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {
		obj, _ = e.CreateGameObject("you", e.SHAPE_RECTANGLE, e.NewSize(100, 100), e.NewPosition(0, 0))
		e.NewCamera(e.NewPosition(0, 0), e.NewPosition(0, 0))
		h, _ := e.CreateGameObject("j", e.SHAPE_RECTANGLE, e.NewSize(20, 20), e.NewPosition(240, 100))
		h.Instance(e.GetCurrentScene(), nil)
		obj = obj.Instance(e.GetCurrentScene(), nil)
		e.GetCurrentScene().SetCamera(e.NewCamera(e.NewPosition(0, 0), e.NewPosition(0, 0)))

	},
		func(ge *e.GameEvent) {
			e.GetCurrentScene().Camera().SetTarget(e.NewPosition(obj.Position().X-250, obj.Position().Y-200))
			if ge.IsKeyDown(ge.KeyRight) {
				obj.MoveTo(e.NewPosition(obj.Position().X+5, obj.Position().Y))
			}
			if ge.IsKeyDown(ge.KeyDown) {
				obj.MoveTo(e.NewPosition(obj.Position().X, obj.Position().Y+5))
			}
			if ge.IsKeyDown(ge.KeyUp) {
				obj.MoveTo(e.NewPosition(obj.Position().X, obj.Position().Y-5))
			}
			if ge.IsKeyDown(ge.KeyLeft) {
				obj.MoveTo(e.NewPosition(obj.Position().X-5, obj.Position().Y))
			}
		})
}
