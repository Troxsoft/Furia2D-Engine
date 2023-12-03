package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var player *e.GameObject
var piso *e.GameObject
var camara *e.Camera
var obstaculos *e.GameObject
var cada int = 0

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {
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

			if em := goe.OnCollisionInTheGroup("o"); em != nil {
				e.NewUiText("perdistes aña", e.NewPosition(100, 100), 50)
			}
		})
		piso.SetUpdate(func(cl *e.GameObject, goe *e.GameObjectEvent) {
			cl.SetSize(e.NewSize(cl.Size().W+5, cl.Size().H))
		})
		player.SetUpdate(func(cl *e.GameObject, goe *e.GameObjectEvent) {
			cl.MoveTo(e.NewPosition(cl.Position().X, cl.Position().Y+5))
			if gh := cl.MoveTo(e.NewPosition(cl.Position().X+2, cl.Position().Y)); gh != nil {
				enzo := e.NewUiText("perdistes aña", e.NewPosition(0, 0), 20)
				enzo.SetColor3(255, 0, 0)
			}
		})
		e.SetCamera(camara)
		player = player.Instance(nil)
		piso = piso.Instance(nil)
		cada = int(player.Position().X + 10)
	},
		func(ge *e.GameEvent) {
			camara.SetTarget(e.NewPosition(player.Position().X, piso.Position().Y-200))
			if cada > 150 {
				obstaculos.Instance(e.NewPosition(player.Position().X+20+camara.Target().X, 320))
				cada = 0
			}
			cada += 1
			if ge.IsKeyDown(ge.KeyUp) {
				player.MoveTo(e.NewPosition(player.Position().X+2, player.Position().Y-50))
			}
		})
}
