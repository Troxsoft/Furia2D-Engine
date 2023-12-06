package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Position struct {
	X float64
	Y float64
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
func NewPosition(x, y float64) Position {
	return Position{
		X: x,
		Y: y,
	}
}
func NewPos(x, y float64) Position {
	return NewPosition(x, y)
}

var (
	DOWN  = NewPosition(0, -1) // physics -> yes,draw ->UP
	UP    = NewPosition(0, 1)  // physics -> yes,draw ->DOWN
	LEFT  = NewPosition(-1, 0) // physics -> yes,draw ->yes
	RIGHT = NewPosition(1, 0)  // physics -> yes,draw ->yes
)

func ClampPos(value, max, min Position) Position {
	newPos := value
	//     X
	if newPos.X > max.X {
		newPos.X = max.X
	}
	if newPos.X < min.X {
		newPos.X = min.X
	}
	//     Y
	if newPos.Y > max.Y {
		newPos.Y = max.Y
	}
	if newPos.Y < min.Y {
		newPos.Y = min.Y
	}
	return newPos
}

func Clamp(value, max, min int) int {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}
func Clampf64(value, max, min float64) float64 {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}

func Clampf32(value, max, min float32) float32 {
	if value > max {
		return max
	}
	if value < min {
		return min
	}
	return value
}

func (p Position) Normalize() Position {
	vec := rl.Vector2Normalize(rl.NewVector2(float32(p.X), float32(p.Y)))
	return NewPos(float64(vec.X), float64(vec.Y))
}
func (p Position) DistanceTo(pos Position) float32 {
	vec := rl.Vector2Distance(rl.NewVector2(float32(p.X), float32(p.Y)), rl.NewVector2(float32(pos.X), float32(pos.Y)))
	return vec
}
func (p Position) DistanceSqrTo(pos Position) float32 {
	vec := rl.Vector2DistanceSqr(rl.NewVector2(float32(p.X), float32(p.Y)), rl.NewVector2(float32(pos.X), float32(pos.Y)))
	return vec
}

// const maximoComunDivisorRecursivo = (a, b) => {
//     if (b === 0) return a;
//     return maximoComunDivisorRecursivo(b, a % b);
// };

func GreatestCommonDivisor(a, b int) int {
	if b == 0 {
		return a
	}
	return GreatestCommonDivisor(b, a%b)
}
