package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var player *e.GameObject
var piso *e.GameObject
var camara *e.Camera
var obstaculos *e.GameObject
var cada int = 0
var overScene *e.Scene = e.NewScene()

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {

		e.NewUiText(overScene, "perdistes", e.NewPos(29, 20), 30)

		player, _ = e.CreateGameObject("player", e.SHAPE_RECTANGLE, e.NewSize(30, 30), e.NewPosition(30, 30))
		player.SetColor(e.BLUE)
		piso, _ = e.CreateGameObject("piso", e.SHAPE_RECTANGLE, e.NewSize(900, 50), e.NewPosition(0, 350))
		camara = e.NewCamera(e.NewPosition(0, 0), player.Position())
		obstaculos, _ = e.CreateGameObject("obstaculos", e.SHAPE_RECTANGLE, e.NewSize(40, 40), e.NewPosition(0, 0))
		piso.SetColor(e.GREEN)
		obstaculos.SetColor(e.NewColor2(255, 0, 0))
		obstaculos.SetStart(func(g *e.GameObject, a any) {
			g.SetPosition(a.(e.Position))
		})
		obstaculos.SetUpdate(func(g *e.GameObject, goe *e.GameObjectEvent) {

		})
		piso.SetUpdate(func(cl *e.GameObject, goe *e.GameObjectEvent) {
			cl.SetSize(e.NewSize(cl.Size().W+5, cl.Size().H))
		})
		player.SetUpdate(func(cl *e.GameObject, goe *e.GameObjectEvent) {
			cl.MoveTo(e.NewPosition(cl.Position().X, cl.Position().Y+5))
			if gh := cl.MoveTo(e.NewPosition(cl.Position().X+2, cl.Position().Y)); gh != nil {
				e.SetScene(overScene)
			}
		})
		e.GetCurrentScene().SetCamera(camara)
		player = player.Instance(e.GetCurrentScene(), nil)
		piso = piso.Instance(e.GetCurrentScene(), nil)
		cada = int(player.Position().X + 10)
	},
		func(ge *e.GameEvent) {
			//e.SetScene(overScene)
			camara.SetTarget(e.NewPosition(player.Position().X, piso.Position().Y-200))
			if cada > 150 {
				obstaculos.Instance(e.GetCurrentScene(), e.NewPosition(player.Position().X+20+camara.Target().X, 320))
				cada = 0
			}
			cada += 1
			suelo2 := e.RayCast(player.Id(), player.Position(), e.DOWN, 35, 1)
			if ge.IsKeyPressed(ge.KeyUp) && suelo2 != nil {
				player.MoveTo(e.NewPosition(player.Position().X+20, player.Position().Y-200))
			}
		})
}
