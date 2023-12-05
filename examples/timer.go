package main

import (
	"fmt"

	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var timer *e.Timer

func main() {
	e.InitGame("welcome to Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {
		timer = e.NewTimer(2)
		timer.Start()

		//fmt.Println(obj.F)
	},
		func(ge *e.GameEvent) {
			if timer.IsDone() {
				timer.Reinitiate()
				fmt.Println("hello")
			}
		})
}
