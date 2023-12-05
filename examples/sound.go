package main

import (
	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {

		e.NewSound("mario.mp3").Play()
		//fmt.Println(obj.F)
	},
		func(ge *e.GameEvent) {

		})
}
