package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Sound struct {
	path  string
	sound rl.Sound
}

func PlaySound(path string) {
	NewSound(path).Play()
}
func NewSound(path string) *Sound {
	return &Sound{
		path:  path,
		sound: rl.LoadSound(path),
	}
}
func (s *Sound) Stop() {
	rl.StopSound(s.sound)
}
func (s *Sound) Play() {
	rl.PlaySound(s.sound)
}
func (s *Sound) RlSound() rl.Sound {
	return s.sound
}
