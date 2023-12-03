package engine

var (
	BLACK  = NewColor2(0, 0, 0)
	WHITE  = NewColor2(255, 255, 255)
	RED    = NewColor2(255, 0, 0)
	GREEN  = NewColor2(0, 255, 0)
	BLUE   = NewColor2(0, 0, 255)
	YELLOW = NewColor2(255, 255, 0)
	PURPLE = NewColor2(87, 35, 100)
	ORANGE = NewColor2(255, 128, 0)
	BROWN  = NewColor2(138, 64, 0)
)

type Color struct {
	R uint8
	G uint8
	B uint8
	A uint8
}

func NewColor2(r, g, b uint8) Color {
	return Color{R: r, G: g, B: b, A: 255}
}

func NewColor(r, g, b, a uint8) Color {
	return Color{R: r, G: g, B: b, A: a}
}
