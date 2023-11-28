package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type AnimationImage struct {
	images []*FuriaImage
	fps    int

	stop  bool
	index int
}

func NewAnimationImage(images []*FuriaImage, fps int) *AnimationImage {

	return &AnimationImage{
		images: images,
		fps:    fps * 2,

		index: 0,
		stop:  false,
	}
}
func (ai *AnimationImage) Draw(color Color, pos Position, size Size) {
	if ai.stop {
		ai.images[ai.index].DrawImage(pos, color, size)

	} else {
		index := (int(rl.GetTime()) * ai.fps) % int(len(ai.images))
		ai.index = index
		ai.images[ai.index].DrawImage(pos, color, size)
	}

}

func (ai *AnimationImage) Stop() {
	ai.stop = true
}

func (ai *AnimationImage) Run() {
	ai.stop = false
}

/*
fps:5
significa que la animacion se dibujara 5 veces en un segundo

*/
