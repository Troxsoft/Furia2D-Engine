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
func GetMousePosition() Position {
	return NewPosition(rl.GetMouseX(), rl.GetMouseY())
}
func _update() {
	for i := 0; i < len(ui_colorZone); i++ {
		//fmt.Println(rl.IsK)
		ui_colorZone[i].Draw()

	}
	for i := 0; i < len(instancesGameObjects); i++ {
		//fmt.Println(rl.IsK)
		instancesGameObjects[i].Execute("update", NewGameObjectEvent(instancesGameObjects[i]))
		instancesGameObjects[i].Draw()

	}
	for i := 0; i < len(ui_text); i++ {
		//fmt.Println(rl.IsK)
		ui_text[i].Draw()

	}
	for i := 0; i < len(ui_button); i++ {
		//fmt.Println(rl.IsK)
		ui_button[i].Draw()

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
		rl.ClearBackground(ConvertColor(colorBackColor))
		//fmt.Println(rl.GetKeyPressed())
		update(NewGameEvent())
		_update()
		//rl.GetCharPressed()
		rl.EndDrawing()
	}

	rl.CloseWindow()

	isRunning = false
}
