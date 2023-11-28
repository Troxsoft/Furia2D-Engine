package engine

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
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
func _update() {
	for i := 0; i < len(instancesGameObjects); i++ {
		//fmt.Println(rl.IsK)
		instancesGameObjects[i].Execute("update", NewGameObjectEvent(instancesGameObjects[i]))
		instancesGameObjects[i].Draw()

	}
}

func InitGame(title string, size Size, start func(*GameEvent), update func(*GameEvent)) {
	isRunning = true

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
