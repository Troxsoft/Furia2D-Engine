package main

import e "github.com/Troxsoft/Furia2D-Engine/engine"

var animacion *e.GameObject

func main() {
	e.InitGame("animaciones", e.NewSize(500, 500), func(ge *e.GameEvent) {
		animacion, _ = e.CreateGameObject("ani", e.SHAPE_ANIMATION_IMAGE, e.NewSize(500, 500), e.NewPosition(0, 0))
		animacion.SetAnimationImage(e.NewAnimationImage([]*e.FuriaImage{
			e.NewImage("sapo.jpg"),
			e.NewImage("sapo.jpg"),
			e.NewImage("sapo.jpg"),
			e.NewImage("sapo2.jpg"),
			e.NewImage("sapo2.jpg"),
			e.NewImage("sapo2.jpg"),
			e.NewImage("petro.jpg"),
			e.NewImage("petro.jpg"),
			e.NewImage("petro.jpg"),
		}, 200))
		e.InstanceGameObject("ani", nil)
	}, func(ge *e.GameEvent) {

	})
}
