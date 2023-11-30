package main

import (
	"fmt"

	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var text *e.UiText
var contador int = 0

func main() {
	e.InitGame("Ejemplos de Gui en Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {

		e.ShowMessageBoxAlert("alerta", "empieza la diversion")
		e.ShowMessageBoxError(":(", "mala")
		e.ShowMessageBoxInfo(":)", "fue una broma")
		text = e.NewUiText(fmt.Sprint(contador), e.NewPosition(50, 50), 20)
		k := e.NewColorZone(e.NewColor2(0, 255, 0), e.NewPosition(0, 0), e.NewSize(500, 400))
		k.SetRoundness(100)
		k.SetSegments(30)
		//fmt.Println(obj.F)
		//e.NewUiText(fmt.Sprint(contador), e.NewPosition(50, 50), 20)
	},
		func(ge *e.GameEvent) {
			contador += 1
			text.SetText(fmt.Sprint(contador))
			if contador > 400 {
				text.Hide()
			}
			if contador > 800 {
				text.Show()
			}
		})
}
