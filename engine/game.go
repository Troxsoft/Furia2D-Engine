package engine

import (
	"errors"
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/tawesoft/golib/v2/dialog"
)

var isRunning = false

var (
	ErrGameNotRunning = errors.New("the game not running")
)

func IsRunning() bool {
	return isRunning
}

var colorBackColor Color = NewColor(50, 50, 80, 255)

func SetBackgroundColor(color_ Color) {
	colorBackColor = color_
}
func SetCamera(cam *Camera) {
	camera = cam
}
func GetCamera() *Camera {
	return camera
}
func GetMousePosition() Position {
	return NewPosition(rl.GetMouseX(), rl.GetMouseY())
}
func _update() {

	for i := 0; i < len(instancesGameObjects); i++ {
		//fmt.Println(rl.IsK)
		instancesGameObjects[i].Execute("update", NewGameObjectEvent(instancesGameObjects[i]))
		instancesGameObjects[i].Draw()

	}

}

func InitGame(title string, size Size, start func(*GameEvent), update func(*GameEvent)) {
	isRunning = true
	err := dialog.Init()
	if err != nil {
		fmt.Println(err)
	}
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(int32(size.W), int32(size.H), title)

	rl.SetTargetFPS(60)
	start(NewGameEvent())

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		if camera != nil {

			rl.ClearBackground(ConvertColor(colorBackColor))
			//rl.DrawText("ano", 30, 100, 29, rl.Brown)
			for i := 0; i < len(ui_colorZone); i++ {
				//fmt.Println(rl.IsK)
				ui_colorZone[i].Draw()

			}
			for i := 0; i < len(ui_text); i++ {
				//fmt.Println(rl.IsK)
				ui_text[i].Draw()
				//fmt.Println(ui_text[i])

			}
			for i := 0; i < len(ui_button); i++ {
				//fmt.Println(rl.IsK)
				ui_button[i].Draw()

			}
			rl.BeginMode2D(camera.camera)
			_update()

		} else {
			rl.ClearBackground(ConvertColor(colorBackColor))
			for i := 0; i < len(ui_colorZone); i++ {
				//fmt.Println(rl.IsK)
				ui_colorZone[i].Draw()

			}
			_update()
			for i := 0; i < len(ui_text); i++ {
				//fmt.Println(rl.IsK)
				ui_text[i].Draw()
				//fmt.Println(ui_text[i])

			}
			for i := 0; i < len(ui_button); i++ {
				//fmt.Println(rl.IsK)
				ui_button[i].Draw()

			}
		}
		update(NewGameEvent())

		//rl.Camera
		//fmt.Println(rl.GetKeyPressed())

		//rl.GetCharPressed()
		//rl.DrawRectangle(30, 30, 30, 30, rl.Black)
		rl.EndMode2D()
		rl.EndDrawing()
	}

	rl.CloseWindow()

	isRunning = false
}
