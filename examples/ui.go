package main

import (
	"fmt"

	e "github.com/Troxsoft/Furia2D-Engine/engine"
)

var text *e.UiText
var contador int = 0
var btn *e.UiButton

func main() {
	e.InitGame("Ejemplos de Gui en Furia2D-Engine :)", e.NewSize(500, 400), func(ge *e.GameEvent) {

		e.ShowMessageBoxAlert("alerta", "empieza la diversion")
		e.ShowMessageBoxError(":(", "mala")
		e.ShowMessageBoxInfo(":)", "fue una broma")
		text = e.NewUiText(e.GetCurrentScene(), fmt.Sprint(contador), e.NewPosition(50, 50), 20)
		k := e.NewUiColorZone(e.GetCurrentScene(), e.NewColor2(0, 255, 0), e.NewPosition(0, 0), e.NewSize(500, 400))
		k.SetRoundness(100)

		k.SetSegments(30)
		btn = e.NewUiButton(e.GetCurrentScene(), "hola como ", e.NewColor2(30, 30, 70), e.NewPosition(273, 200), e.NewSize(100, 30))

		//btn.SetPosition(e.NewPosition(200, 200)) //fmt.Println(obj.F)
		//e.NewUiText(fmt.Sprint(contador), e.NewPosition(50, 50), 20)
	},
		func(ge *e.GameEvent) {

			if btn.IsPressed() {
				e.ShowMessageBoxInfo("turbio", "me lo as presionado")
				btn.Text.SetText("aña")
			}
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
