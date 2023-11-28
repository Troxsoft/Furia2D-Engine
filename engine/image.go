package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type FuriaImage struct {
	path_         string
	raylibTexture rl.Texture2D
	Ignore        bool
}

func (fi *FuriaImage) DrawImage(pos Position, color Color, size Size) {
	if fi.Ignore == false {
		fi.raylibTexture.Width = int32(size.W)
		fi.raylibTexture.Height = int32(size.H)
		rl.DrawTexture(fi.raylibTexture, pos.X, pos.Y, ConvertColor(color))
	}
}
func (fi *FuriaImage) Destroy() {
	rl.UnloadTexture(fi.raylibTexture)
}
func NewImage(path string) *FuriaImage {
	return &FuriaImage{
		path_:         path,
		raylibTexture: rl.LoadTexture(path),
		Ignore:        false,
	}
}
