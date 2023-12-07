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
	areas     []*CollisionArea
}

func (s *Scene) IsCurrent() bool {
	return s == scene_32445
}
func (s *Scene) Buttons() []*UiButton {
	return s.ui_button
}
func (s *Scene) ColorZones() []*UiColorZone {
	return s.ui_colorZone
}
func (s *Scene) Texts() []*UiText {
	return s.ui_text
}
func (s *Scene) GameObjectsObjects() []*GameObject {
	return s.instancesGameObjects
}

func (s *Scene) GameObjectsInTheGroup(gr string) []*GameObject {
	return GetGameObjectsInTheGroup(gr)
}
func (s *Scene) FistGameObjectsInTheGroup(gr string) *GameObject {
	return GetGameObjectsInTheGroup(gr)[0]
}
func (s *Scene) LastGameObjectsInTheGroup(gr string) *GameObject {
	return GetGameObjectsInTheGroup(gr)[len(GetGameObjectsInTheGroup(gr))-1]
}
func (s *Scene) FistGameObject() *GameObject {
	return s.instancesGameObjects[0]
}
func (s *Scene) LastGameObject() *GameObject {
	return s.instancesGameObjects[len(s.instancesGameObjects)-1]
}

func (s *Scene) GameObjectWithName(name string) []*GameObject {
	v := []*GameObject{}
	for i := 0; i < len(s.instancesGameObjects); i++ {
		if s.instancesGameObjects[i].Name() == name {
			v = append(v, s.instancesGameObjects[i])
		}
	}
	return v
}
func (s *Scene) LastGameObjectWithName(name string) *GameObject {
	v := []*GameObject{}
	for i := 0; i < len(s.instancesGameObjects); i++ {
		if s.instancesGameObjects[i].Name() == name {
			v = append(v, s.instancesGameObjects[i])
		}
	}
	return v[len(v)-1]
}
func (s *Scene) FirstGameObjectWithName(name string) *GameObject {

	for i := 0; i < len(s.instancesGameObjects); i++ {
		if s.instancesGameObjects[i].Name() == name {
			return s.instancesGameObjects[i]
		}
	}
	return nil

}

func (s *Scene) Camera() *Camera {
	return s.camera
}
func (s *Scene) SetCamera(c *Camera) {
	s.camera = c
}
func (s *Scene) CollisionsAreas() []*CollisionArea {
	return s.areas
}
func NewScene() *Scene {
	return &Scene{
		instancesGameObjects: []*GameObject{},
		ui_text:              []*UiText{},
		ui_colorZone:         []*UiColorZone{},
		ui_button:            []*UiButton{},
		camera:               nil,
		areas:                []*CollisionArea{},
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
		for i := 0; i < len(s.areas); i++ {
			p5 := s.areas[i].gob.collision.OnCollision()
			var xbsx *GameObjectEvent = nil
			if p5 != nil {
				xbsx = NewGameObjectEvent(p5)
			}
			s.areas[i].onCollision(s.areas[i], p5, xbsx)
			if s.areas[i].debugColor != nil {
				rl.DrawRectangle(int32(s.areas[i].position.X), int32(s.areas[i].position.Y), int32(s.areas[i].size.W), int32(s.areas[i].size.H), ConvertColor(*s.areas[i].debugColor))
			}
		}
		_update()

	} else {
		rl.ClearBackground(ConvertColor(colorBackColor))

		for i := 0; i < len(s.ui_colorZone); i++ {
			//fmt.Println(rl.IsK)
			s.ui_colorZone[i].Draw()

		}
		for i := 0; i < len(s.areas); i++ {
			p5 := s.areas[i].gob.collision.OnCollision()
			var xbsx *GameObjectEvent = nil
			if p5 != nil {
				xbsx = NewGameObjectEvent(p5)
			}
			s.areas[i].onCollision(s.areas[i], p5, xbsx)
			if s.areas[i].debugColor != nil {
				rl.DrawRectangle(int32(s.areas[i].position.X), int32(s.areas[i].position.Y), int32(s.areas[i].size.W), int32(s.areas[i].size.H), ConvertColor(*s.areas[i].debugColor))
			}
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
