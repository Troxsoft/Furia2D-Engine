package engine

import rl "github.com/gen2brain/raylib-go/raylib"

// type Key struct
type GameObjectEvent struct {
	obj *GameObject

	MouseLeftButton   int32
	MouseRightButton  int32
	MouseMiddleButton int32

	// Mouse Buttons

	MouseButtonLeft    int32
	MouseButtonRight   int32
	MouseButtonMiddle  int32
	MouseButtonSide    int32
	MouseButtonExtra   int32
	MouseButtonForward int32
	MouseButtonBack    int32
}

func (k *GameObjectEvent) OnCollisionInTheGroup(group string) *GameObject {
	return k.obj.collision.OnCollisionInTheGroup(group)
}
func (k *GameObjectEvent) OnCollisionTo(other *GameObject) bool {
	return k.obj.collision.OnCollisionTo(other.collision)
}
func (k *GameObjectEvent) OnCollision() *GameObject {
	return k.obj.collision.OnCollision()
}
func (k *GameObjectEvent) OnCollisionPos(pos Position) *GameObject {
	return k.obj.collision.OnCollisionPos(pos)
}

func (k *GameObjectEvent) OnCollisionInPosition(other *GameObject, pos Position) bool {
	return k.obj.collision.OnCollisionInPosition(other.collision, pos)
}

func (k *GameObjectEvent) IsMouseReleased(mouse int32) bool {
	if rl.IsMouseButtonReleased(mouse) {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()
		a := rl.NewRectangle(float32(mouseX), float32(mouseY), float32(2), float32(2))
		b := rl.NewRectangle(float32(k.obj.Position().X), float32(k.obj.position.Y), float32(k.obj.size.W), float32(k.obj.size.H))
		if rl.CheckCollisionRecs(a, b) {
			return true
		}
	}

	return false
}
func (k *GameObjectEvent) RayCast(initPos Position, dirreccion Position, long uint32, jump uint16) *GameObject {
	return RayCast(k.obj.id, initPos, dirreccion, long, jump)
}
func (k *GameObjectEvent) RayCastDefault(initPos Position, dirreccion Position, long uint32) *GameObject {
	return RayCast(k.obj.id, initPos, dirreccion, long, 1)
}
func (k *GameObjectEvent) IsMouseUp(mouse int32) bool {
	if rl.IsMouseButtonUp(mouse) {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()
		a := rl.NewRectangle(float32(mouseX), float32(mouseY), float32(2), float32(2))
		b := rl.NewRectangle(float32(k.obj.Position().X), float32(k.obj.position.Y), float32(k.obj.size.W), float32(k.obj.size.H))
		if rl.CheckCollisionRecs(a, b) {
			return true
		}
	}

	return false
}
func (k *GameObjectEvent) IsMousePressed(mouse int32) bool {
	if rl.IsMouseButtonPressed(mouse) {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()
		a := rl.NewRectangle(float32(mouseX), float32(mouseY), float32(2), float32(2))
		b := rl.NewRectangle(float32(k.obj.Position().X), float32(k.obj.position.Y), float32(k.obj.size.W), float32(k.obj.size.H))
		if rl.CheckCollisionRecs(a, b) {
			return true
		}
	}

	return false
}
func (k *GameObjectEvent) IsMouseDown(mouse int32) bool {
	if rl.IsMouseButtonDown(mouse) {
		mouseX := rl.GetMouseX()
		mouseY := rl.GetMouseY()
		a := rl.NewRectangle(float32(mouseX), float32(mouseY), float32(2), float32(2))
		b := rl.NewRectangle(float32(k.obj.Position().X), float32(k.obj.position.Y), float32(k.obj.size.W), float32(k.obj.size.H))
		if rl.CheckCollisionRecs(a, b) {
			return true
		}
	}

	return false
}

func NewGameObjectEvent(obj *GameObject) *GameObjectEvent {
	return &GameObjectEvent{
		obj:               obj,
		MouseLeftButton:   0,
		MouseRightButton:  1,
		MouseMiddleButton: 2,

		// Mouse Buttons

		MouseButtonLeft:    0,
		MouseButtonRight:   1,
		MouseButtonMiddle:  2,
		MouseButtonSide:    3,
		MouseButtonExtra:   4,
		MouseButtonForward: 5,
		MouseButtonBack:    6,
	}
}

// @copiar@ 10 20 sugerencia macro para Quick-lenguaje
// @ir@
type GameEvent struct {
	KeySpace        int32
	KeyEscape       int32
	KeyEnter        int32
	KeyTab          int32
	KeyBackspace    int32
	KeyInsert       int32
	KeyDelete       int32
	KeyRight        int32
	KeyLeft         int32
	KeyDown         int32
	KeyUp           int32
	KeyPageUp       int32
	KeyPageDown     int32
	KeyHome         int32
	KeyEnd          int32
	KeyCapsLock     int32
	KeyScrollLock   int32
	KeyNumLock      int32
	KeyPrintScreen  int32
	KeyPause        int32
	KeyF1           int32
	KeyF2           int32
	KeyF3           int32
	KeyF4           int32
	KeyF5           int32
	KeyF6           int32
	KeyF7           int32
	KeyF8           int32
	KeyF9           int32
	KeyF10          int32
	KeyF11          int32
	KeyF12          int32
	KeyLeftShift    int32
	KeyLeftControl  int32
	KeyLeftAlt      int32
	KeyLeftSuper    int32
	KeyRightShift   int32
	KeyRightControl int32
	KeyRightAlt     int32
	KeyRightSuper   int32
	KeyKbMenu       int32
	KeyLeftBracket  int32
	KeyBackSlash    int32
	KeyRightBracket int32
	KeyGrave        int32

	// Keyboard Number Pad Keys
	KeyKp0        int32
	KeyKp1        int32
	KeyKp2        int32
	KeyKp3        int32
	KeyKp4        int32
	KeyKp5        int32
	KeyKp6        int32
	KeyKp7        int32
	KeyKp8        int32
	KeyKp9        int32
	KeyKpDecimal  int32
	KeyKpDivide   int32
	KeyKpMultiply int32
	KeyKpSubtract int32
	KeyKpAdd      int32
	KeyKpEnter    int32
	KeyKpEqual    int32

	// Keyboard Alpha Numeric Keys
	KeyApostrophe int32
	KeyComma      int32
	KeyMinus      int32
	KeyPeriod     int32
	KeySlash      int32
	KeyZero       int32
	KeyOne        int32
	KeyTwo        int32
	KeyThree      int32
	KeyFour       int32
	KeyFive       int32
	KeySix        int32
	KeySeven      int32
	KeyEight      int32
	KeyNine       int32
	KeySemicolon  int32
	KeyEqual      int32
	KeyA          int32
	KeyB          int32
	KeyC          int32
	KeyD          int32
	KeyE          int32
	KeyF          int32
	KeyG          int32
	KeyH          int32
	KeyI          int32
	KeyJ          int32
	KeyK          int32
	KeyL          int32
	KeyM          int32
	KeyN          int32
	KeyO          int32
	KeyP          int32
	KeyQ          int32
	KeyR          int32
	KeyS          int32
	KeyT          int32
	KeyU          int32
	KeyV          int32
	KeyW          int32
	KeyX          int32
	KeyY          int32
	KeyZ          int32
	// Android keys
	KeyBack           int32
	KeyMenu           int32
	KeyVolumeUp       int32
	KeyVolumeDown     int32
	MouseLeftButton   int32
	MouseRightButton  int32
	MouseMiddleButton int32
	// Mouse Buttons
	MouseButtonLeft    int32
	MouseButtonRight   int32
	MouseButtonMiddle  int32
	MouseButtonSide    int32
	MouseButtonExtra   int32
	MouseButtonForward int32
	MouseButtonBack    int32
}

func NewGameEvent() *GameEvent {
	return &GameEvent{
		KeySpace:        32,
		KeyEscape:       256,
		KeyEnter:        257,
		KeyTab:          258,
		KeyBackspace:    259,
		KeyInsert:       260,
		KeyDelete:       261,
		KeyRight:        262,
		KeyDown:         264,
		KeyLeft:         263,
		KeyUp:           265,
		KeyPageUp:       266,
		KeyPageDown:     267,
		KeyHome:         268,
		KeyEnd:          269,
		KeyCapsLock:     280,
		KeyScrollLock:   281,
		KeyNumLock:      282,
		KeyPrintScreen:  283,
		KeyPause:        284,
		KeyF1:           290,
		KeyF2:           291,
		KeyF3:           292,
		KeyF4:           293,
		KeyF5:           294,
		KeyF6:           295,
		KeyF7:           296,
		KeyF8:           297,
		KeyF9:           298,
		KeyF10:          299,
		KeyF11:          300,
		KeyF12:          301,
		KeyLeftShift:    340,
		KeyLeftControl:  341,
		KeyLeftAlt:      342,
		KeyLeftSuper:    343,
		KeyRightShift:   344,
		KeyRightControl: 345,
		KeyRightAlt:     346,
		KeyRightSuper:   347,
		KeyKbMenu:       348,
		KeyLeftBracket:  91,
		KeyBackSlash:    92,
		KeyRightBracket: 93,
		KeyGrave:        96,

		// Keyboard Number Pad Keys
		KeyKp0:        320,
		KeyKp1:        321,
		KeyKp2:        322,
		KeyKp3:        323,
		KeyKp4:        324,
		KeyKp5:        325,
		KeyKp6:        326,
		KeyKp7:        327,
		KeyKp8:        328,
		KeyKp9:        329,
		KeyKpDecimal:  330,
		KeyKpDivide:   331,
		KeyKpMultiply: 332,
		KeyKpSubtract: 333,
		KeyKpAdd:      334,
		KeyKpEnter:    335,
		KeyKpEqual:    336,

		// Keyboard Alpha Numeric Keys
		KeyApostrophe: 39,
		KeyComma:      44,
		KeyMinus:      45,
		KeyPeriod:     46,
		KeySlash:      47,
		KeyZero:       48,
		KeyOne:        49,
		KeyTwo:        50,
		KeyThree:      51,
		KeyFour:       52,
		KeyFive:       53,
		KeySix:        54,
		KeySeven:      55,
		KeyEight:      56,
		KeyNine:       57,
		KeySemicolon:  59,
		KeyEqual:      61,
		KeyA:          65,
		KeyB:          66,
		KeyC:          67,
		KeyD:          68,
		KeyE:          69,
		KeyF:          70,
		KeyG:          71,
		KeyH:          72,
		KeyI:          73,
		KeyJ:          74,
		KeyK:          75,
		KeyL:          76,
		KeyM:          77,
		KeyN:          78,
		KeyO:          79,
		KeyP:          80,
		KeyQ:          81,
		KeyR:          82,
		KeyS:          83,
		KeyT:          84,
		KeyU:          85,
		KeyV:          86,
		KeyW:          87,
		KeyX:          88,
		KeyY:          89,
		KeyZ:          90,

		// Android keys
		KeyBack:       4,
		KeyMenu:       82,
		KeyVolumeUp:   24,
		KeyVolumeDown: 25,

		MouseLeftButton:   0,
		MouseRightButton:  1,
		MouseMiddleButton: 2,

		// Mouse Buttons

		MouseButtonLeft:    0,
		MouseButtonRight:   1,
		MouseButtonMiddle:  2,
		MouseButtonSide:    3,
		MouseButtonExtra:   4,
		MouseButtonForward: 5,
		MouseButtonBack:    6,
	}
}
func (k *GameEvent) IsMouseReleased(mouse int32) bool {

	return rl.IsMouseButtonReleased(mouse)
}
func (k *GameEvent) IsMouseUp(mouse int32) bool {
	return rl.IsMouseButtonUp(mouse)
}
func (k *GameEvent) IsMousePressed(mouse int32) bool {
	return rl.IsMouseButtonPressed(mouse)
}
func (k *GameEvent) IsMouseDown(mouse int32) bool {
	return rl.IsMouseButtonPressed(mouse)
}

func (k *GameEvent) IsKeyReleased(key int32) bool {

	return rl.IsKeyReleased(key)
}
func (k *GameEvent) IsKeyUp(key int32) bool {
	return rl.IsKeyUp(key)
}
func (k *GameEvent) IsKeyPressed(key int32) bool {
	return rl.IsKeyPressed(key)
}
func (k *GameEvent) IsKeyDown(key int32) bool {
	return rl.IsKeyDown(key)
}
