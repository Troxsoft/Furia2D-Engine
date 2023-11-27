package engine

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
