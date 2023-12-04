package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Scene struct {
	instancesGameObjects []*GameObject
	ui_text              []*UiText
	//color zone
	ui_colorZone []*UiColorZone

	//button
	ui_button []*UiButton
	camera    *Camera
}

func (s *Scene) GameObjectsObjects() []*GameObject {
	return s.instancesGameObjects
}

func (s *Scene) Camera() *Camera {
	return s.camera
}
func (s *Scene) SetCamera(c *Camera) {
	s.camera = c
}

func NewScene() *Scene {
	return &Scene{
		instancesGameObjects: []*GameObject{},
		ui_text:              []*UiText{},
		ui_colorZone:         []*UiColorZone{},
		ui_button:            []*UiButton{},
		camera:               nil,
	}
}
func SetScene(s *Scene) {
	scene_32445 = s
}
func GetCurrentScene() *Scene {
	return scene_32445
}
func (s *Scene) Draw() {
	if s.camera != nil {
		rl.ClearBackground(ConvertColor(colorBackColor))
		//rl.DrawText("ano", 30, 100, 29, rl.Brown)
		for i := 0; i < len(s.ui_colorZone); i++ {
			//fmt.Println(rl.IsK)
			s.ui_colorZone[i].Draw()

		}
		for i := 0; i < len(s.ui_text); i++ {
			//fmt.Println(rl.IsK)
			s.ui_text[i].Draw()
			//fmt.Println(ui_text[i])

		}
		for i := 0; i < len(s.ui_button); i++ {
			//fmt.Println(rl.IsK)
			s.ui_button[i].Draw()

		}
		rl.BeginMode2D(s.camera.camera)
		_update()
	} else {
		rl.ClearBackground(ConvertColor(colorBackColor))
		for i := 0; i < len(s.ui_colorZone); i++ {
			//fmt.Println(rl.IsK)
			s.ui_colorZone[i].Draw()

		}
		_update()
		for i := 0; i < len(s.ui_text); i++ {
			//fmt.Println(rl.IsK)
			s.ui_text[i].Draw()
			//fmt.Println(ui_text[i])

		}
		for i := 0; i < len(s.ui_button); i++ {
			//fmt.Println(rl.IsK)
			s.ui_button[i].Draw()

		}
	}
}
