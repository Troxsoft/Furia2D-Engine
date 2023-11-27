package engine

type Position struct {
	X int32
	Y int32
}

type Size struct {
	W uint
	H uint
}

func NewSize(w, h uint) Size {
	return Size{
		W: w,
		H: h,
	}
}
func NewPosition(x, y int32) Position {
	return Position{
		X: x,
		Y: y,
	}
}
