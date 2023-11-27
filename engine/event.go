package engine

import rl "github.com/gen2brain/raylib-go/raylib"

const (

	// KeyNull is used for no key pressed
	KeyNull = 0

	// Keyboard Function Keys
	KeySpace        = 32
	KeyEscape       = 256
	KeyEnter        = 257
	KeyTab          = 258
	KeyBackspace    = 259
	KeyInsert       = 260
	KeyDelete       = 261
	KeyRight        = 262
	KeyLeft         = 263
	KeyDown         = 264
	KeyUp           = 265
	KeyPageUp       = 266
	KeyPageDown     = 267
	KeyHome         = 268
	KeyEnd          = 269
	KeyCapsLock     = 280
	KeyScrollLock   = 281
	KeyNumLock      = 282
	KeyPrintScreen  = 283
	KeyPause        = 284
	KeyF1           = 290
	KeyF2           = 291
	KeyF3           = 292
	KeyF4           = 293
	KeyF5           = 294
	KeyF6           = 295
	KeyF7           = 296
	KeyF8           = 297
	KeyF9           = 298
	KeyF10          = 299
	KeyF11          = 300
	KeyF12          = 301
	KeyLeftShift    = 340
	KeyLeftControl  = 341
	KeyLeftAlt      = 342
	KeyLeftSuper    = 343
	KeyRightShift   = 344
	KeyRightControl = 345
	KeyRightAlt     = 346
	KeyRightSuper   = 347
	KeyKbMenu       = 348
	KeyLeftBracket  = 91
	KeyBackSlash    = 92
	KeyRightBracket = 93
	KeyGrave        = 96

	// Keyboard Number Pad Keys
	KeyKp0        = 320
	KeyKp1        = 321
	KeyKp2        = 322
	KeyKp3        = 323
	KeyKp4        = 324
	KeyKp5        = 325
	KeyKp6        = 326
	KeyKp7        = 327
	KeyKp8        = 328
	KeyKp9        = 329
	KeyKpDecimal  = 330
	KeyKpDivide   = 331
	KeyKpMultiply = 332
	KeyKpSubtract = 333
	KeyKpAdd      = 334
	KeyKpEnter    = 335
	KeyKpEqual    = 336

	// Keyboard Alpha Numeric Keys
	KeyApostrophe = 39
	KeyComma      = 44
	KeyMinus      = 45
	KeyPeriod     = 46
	KeySlash      = 47
	KeyZero       = 48
	KeyOne        = 49
	KeyTwo        = 50
	KeyThree      = 51
	KeyFour       = 52
	KeyFive       = 53
	KeySix        = 54
	KeySeven      = 55
	KeyEight      = 56
	KeyNine       = 57
	KeySemicolon  = 59
	KeyEqual      = 61
	KeyA          = 65
	KeyB          = 66
	KeyC          = 67
	KeyD          = 68
	KeyE          = 69
	KeyF          = 70
	KeyG          = 71
	KeyH          = 72
	KeyI          = 73
	KeyJ          = 74
	KeyK          = 75
	KeyL          = 76
	KeyM          = 77
	KeyN          = 78
	KeyO          = 79
	KeyP          = 80
	KeyQ          = 81
	KeyR          = 82
	KeyS          = 83
	KeyT          = 84
	KeyU          = 85
	KeyV          = 86
	KeyW          = 87
	KeyX          = 88
	KeyY          = 89
	KeyZ          = 90

	// Android keys
	KeyBack       = 4
	KeyMenu       = 82
	KeyVolumeUp   = 24
	KeyVolumeDown = 25

	MouseLeftButton   = MouseButtonLeft
	MouseRightButton  = MouseButtonRight
	MouseMiddleButton = MouseButtonMiddle
)

// Mouse Buttons
const (
	MouseButtonLeft = iota
	MouseButtonRight
	MouseButtonMiddle
	MouseButtonSide
	MouseButtonExtra
	MouseButtonForward
	MouseButtonBack
)

type Event struct {
	obj *GameObject
}

func (k *Event) OnCollisionTo(other *GameObject) bool {
	return k.obj.collision.OnCollisionTo(other.collision)
}
func (k *Event) OnCollision() (*GameObject, bool) {
	return k.obj.collision.OnCollision()
}
func (k *Event) IsMouseReleased(mouse int32) bool {

	return rl.IsMouseButtonReleased(mouse)
}
func (k *Event) IsMouseUp(mouse int32) bool {
	return rl.IsMouseButtonUp(mouse)
}
func (k *Event) IsMousePressed(mouse int32) bool {
	return rl.IsMouseButtonPressed(mouse)
}
func (k *Event) IsMouseDown(mouse int32) bool {
	return rl.IsMouseButtonPressed(mouse)
}

func (k *Event) IsKeyReleased(key int32) bool {

	return rl.IsKeyReleased(key)
}
func (k *Event) IsKeyUp(key int32) bool {
	return rl.IsKeyUp(key)
}
func (k *Event) IsKeyPressed(key int32) bool {
	return rl.IsKeyPressed(key)
}
func (k *Event) IsKeyDown(key int32) bool {
	return rl.IsKeyDown(key)
}
