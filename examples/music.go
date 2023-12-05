package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var (
	musica *e.Music
)

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {
		e.PlaySound("mario.mp3") // EASY
		/*	DIFFICULT

			musica = e.NewMusic("mario.mp3")

			musica.Play()
		*/
	},
		func(ge *e.GameEvent) {
			if ge.IsKeyPressed(ge.KeyA) {
				musica.Stop()
			}
			if ge.IsKeyPressed(ge.KeyB) {
				musica.Play()
			}
		})
}
