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
	for i := 0; i < len(musics_update); i++ {

		rl.UpdateMusicStream(musics_update[i].music)
	}
	b49 := *scene_32445
	for i := 0; i < len(b49.instancesGameObjects); i++ {
		//fmt.Println(rl.IsK)
		b49.instancesGameObjects[i].Execute("update", NewGameObjectEvent(b49.instancesGameObjects[i]))
		b49.instancesGameObjects[i].Draw()

	}

}

var scene_32445 *Scene = NewScene()

func InitGame(title string, size Size, start func(*GameEvent), update func(*GameEvent)) {
	isRunning = true
	err := dialog.Init()
	if err != nil {
		fmt.Println(err)
	}
	rl.SetConfigFlags(rl.FlagVsyncHint)
	rl.InitWindow(int32(size.W), int32(size.H), title)
	rl.InitAudioDevice()
	rl.SetTargetFPS(60)
	start(NewGameEvent())

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		scene_32445.Draw()
		update(NewGameEvent())
		rl.EndMode2D()
		rl.EndDrawing()
	}

	rl.CloseWindow()

	isRunning = false
}
