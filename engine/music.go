package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Music struct {
	path  string
	music rl.Music
}

var (
	musics_update []*Music
)

func NewMusic(path string) *Music {
	//mus := rl.LoadAudioStream(path)
	p := &Music{
		path:  path,
		music: rl.LoadMusicStream(path),
	}
	musics_update = append(musics_update, p)
	return p
}
func (m *Music) Play() {
	rl.PlayMusicStream(m.music)
}
func (m *Music) Stop() {
	rl.StopMusicStream(m.music)
}
func (m *Music) RlMusic() rl.Music {
	return m.music
}
func (m *Music) MusicPath() string {
	return m.path
}
