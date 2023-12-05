package main

import (
	"fmt"

	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var obj *e.GameObject
var obj23 *e.GameObject

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(500, 600), func(ge *e.GameEvent) {
		obj, _ = e.CreateGameObject("you", e.SHAPE_RECTANGLE, e.NewSize(30, 30), e.NewPosition(30, 35))
		//player input controller
		obj23, _ = e.CreateGameObject("youj", e.SHAPE_RECTANGLE, e.NewSize(33, 33), e.NewPosition(400, 26))
		obj26787979, _ := e.CreateGameObject("youjj", e.SHAPE_RECTANGLE, e.NewSize(33, 33), e.NewPosition(0, 26))
		obj26787979.Instance(e.GetCurrentScene(), nil)

		ob_abajo_mutante, _ := e.CreateGameObject("youjjj", e.SHAPE_RECTANGLE, e.NewSize(33, 33), e.NewPosition(150, 300))
		ob_abajo_mutante.Instance(e.GetCurrentScene(), nil)

		ob_abajo_mutante_mmamamamam, _ := e.CreateGameObject("youjjjj", e.SHAPE_RECTANGLE, e.NewSize(33, 33), e.NewPosition(150, 0))
		ob_abajo_mutante_mmamamamam.Instance(e.GetCurrentScene(), nil)

		obj.SetUpdate(func(g *e.GameObject, goe *e.GameObjectEvent) {
			//fmt.Println(g.Position())
			g.MoveTo(e.ClampPos(g.Position(), e.NewPos(500, 400), e.NewPos(0, 0)))
			if ge.IsKeyDown(ge.KeyRight) {
				g.MoveTo(e.NewPos(g.Position().X+5, g.Position().Y))
			}
			if ge.IsKeyDown(ge.KeyLeft) {
				g.MoveTo(e.NewPos(g.Position().X-5, g.Position().Y))
			}
			//fmt.Println(e.GetCurrentScene().GameObjectsObjects()[0].Id(), " ", e.GetCurrentScene().GameObjectsObjects()[1].Id())
			if jk := goe.RayCast(g.Position(), e.RIGHT, 200, 1); jk != nil {
				fmt.Println("caliente 2")
				//fmt.Println(jk)
			}
			e.DrawRayCast(g.Position(), e.RIGHT, 200, e.GREEN)

			if jk := goe.RayCast(g.Position(), e.LEFT, 100, 1); jk != nil {
				fmt.Println("caliente 1")
				//fmt.Println(jk)
			}
			e.DrawRayCast(g.Position(), e.LEFT, 100, e.GREEN)

			if jk := goe.RayCast(g.Position(), e.DOWN, 400, 1); jk != nil {
				fmt.Println("caliente 3")
				//fmt.Println(jk)
			}
			e.DrawRayCast(g.Position(), e.UP, 400, e.GREEN)

			if jk := goe.RayCast(g.Position(), e.UP, 400, 1); jk != nil {
				fmt.Println("caliente 4")
				//fmt.Println(jk)
			}
			e.DrawRayCast(g.Position(), e.DOWN, 400, e.GREEN)
		})
		obj.Instance(e.GetCurrentScene(), nil)
		obj23.Instance(e.GetCurrentScene(), nil)

		//fmt.Println(obj.F)
	},
		func(ge *e.GameEvent) {

		})
}
