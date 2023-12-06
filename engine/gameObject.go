package engine

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Shape string

const (
	SHAPE_RECTANGLE       Shape = "rectangle"
	SHAPE_IMAGE           Shape = "image"
	SHAPE_ANIMATION_IMAGE Shape = "animationImage"
)

func ConvertColor(color Color) rl.Color {
	return rl.NewColor(color.R, color.G, color.B, color.A)
}

func IsValidShape(shape Shape) bool {

	if shape != "rectangle" && shape != "image" && shape != "animationImage" {
		return false
	}
	return true
}

type GameObject struct {
	name      string
	size      Size
	position  Position
	shape     Shape
	funcs     map[string]func(*GameObject, any)
	hide      bool
	vars      map[string]any
	collision CollisionRectangle
	id        int
	groups    map[string]bool
}

var gameObjects map[string]*GameObject = make(map[string]*GameObject)
var (
	ErrShapeInvalid = errors.New("the shape is invalid")
)

func (g *GameObject) Vars() map[string]any {
	//hhhhhhhhhhhhhhhhhhhhhhhhhhhh
	newV := make(map[string]any)
	for k, v := range g.vars {
		newV[k] = v
	}
	return newV

}
func (g *GameObject) Id() int {
	return g.id
}
func (g *GameObject) CollisionRect() CollisionRectangle {
	return g.collision
}
func (g *GameObject) Name() string {
	return g.name
}
func (g *GameObject) Shape() Shape {
	return g.shape
}

func (g *GameObject) Size() Size {
	return g.size
}
func (g *GameObject) Hide() {
	g.hide = true
}
func (g *GameObject) Show() {
	g.hide = false
}
func (g *GameObject) SetSize(siz Size) {
	g.size = siz
}
func (g *GameObject) SetSize2(w, h uint) {
	g.size.W = w
	g.size.H = h
}

//	func (g *GameObject) PrintGameObject() {
//		fmt.Printf(
//			`
//		-Game object Name: '%s' SHAPE: '%s'
//
// -Position: X = %f | Y = %f
// -Color: R = %s | G = %s | B = %s | A(alpha) = %s
// `, g.Name(), g.Shape(), g.Pos().X, g.Pos().Y, g.Color().R, g.Color().G, g.Color().B, g.Color().A)
// }
func GetGameObjectsInTheGroup(gr string) []*GameObject {
	groupsd := []*GameObject{}
	for i := 0; i < len(scene_32445.instancesGameObjects); i++ {
		if scene_32445.instancesGameObjects[i].IsInGroup(gr) {
			groupsd = append(groupsd, scene_32445.instancesGameObjects[i])
		}
	}
	return groupsd
}
func (g *GameObject) DeleteGroup(gr string) {
	g.groups[gr] = false
}
func (g *GameObject) IsInGroup(gr string) bool {
	v, ok := g.groups[gr]
	if ok == true {

		if v == true {
			return true
		} else {
			return false
		}
	}
	return false
}

func (g *GameObject) Position() Position {
	return g.position
}
func (g *GameObject) Pos() Position {
	return g.Position()
}

/*
physics type: Colliders and collisions[OFF]
*/
func (g *GameObject) SetPos(pos Position) {
	g.SetPosition(pos)
}

/*
physics type: Colliders and collisions[OFF]
*/
func (g *GameObject) SetPos2(x, y float64) {
	g.SetPosition2(x, y)
}

/*
physics type: Colliders and collisions[OFF]
*/
func (g *GameObject) SetPosition2(x, y float64) {
	g.position.X = x
	g.position.Y = y
}

/*
physics type: Colliders and collisions[OFF]
*/
func (g *GameObject) SetPosition(pos Position) {
	g.position = pos
}

func (g *GameObject) GetVar(name string) any {
	return g.vars[name]
}
func (g *GameObject) SetVar(name string, value any) {
	g.vars[name] = value
}

func (g *GameObject) SetUpdate(f func(*GameObject, *GameObjectEvent)) {
	g.funcs["update"] = func(g *GameObject, a any) {
		f(g, a.(*GameObjectEvent))
	}
}

func (g *GameObject) SetStart(f func(*GameObject, any)) {
	g.funcs["start"] = f
}
func (g *GameObject) SetImage(img *FuriaImage) {
	if g.shape == SHAPE_IMAGE {

		g.vars["image"] = img
	} else {
		panic("the image not is valid")
	}
}
func (g *GameObject) Image() *FuriaImage {
	if g.shape == SHAPE_IMAGE {
		switch any(g.GetVar("image")).(type) {
		case *FuriaImage:
			return g.GetVar("image").(*FuriaImage)
			break
		}

	} else {
		panic("the image not is valid")
	}
	return nil
}

func (g *GameObject) AddToGroup(groupName string) {
	g.groups[groupName] = true
}
func (g *GameObject) Groups() []string {
	keys := []string{}
	values := []bool{}
	arr := []string{}
	for k, v := range g.groups {
		keys = append(keys, k)
		values = append(values, v)
	}
	for i, v := range values {
		if v == true {
			arr = append(arr, keys[i])
		}
	}
	return arr

}
func (g *GameObject) Instance(scene *Scene, params any) *GameObject {
	i := g
	vars__ := make(map[string]any)
	for k, v := range i.vars {
		vars__[k] = v
	}
	funcs__ := make(map[string]func(*GameObject, any))
	for k, v := range i.funcs {
		funcs__[k] = v
	}
	k := &GameObject{
		name:     i.name,
		size:     i.size,
		position: i.position,
		shape:    i.shape,
		funcs:    funcs__,
		hide:     i.hide,
		vars:     vars__,
		id:       len(scene.instancesGameObjects) + 1,
		groups:   i.groups,
	}
	k.collision = NewCollisionRectagle(k)
	k.Execute("start", params)
	scene.instancesGameObjects = append(scene.instancesGameObjects, k)
	return k
}

func (g *GameObject) SetFunction(nameFunction string, function func(*GameObject, any)) {
	g.funcs[nameFunction] = function
}

func (g *GameObject) Functions() map[string]func(*GameObject, any) {
	newV := make(map[string]func(*GameObject, any))
	for k, v := range g.funcs {
		newV[k] = v
	}
	return newV
}
func (g *GameObject) Execute(nameFunction string, params any) {
	g.funcs[nameFunction](g, params)
}

func (g *GameObject) SetColor(color Color) {
	g.SetVar("color", color)
}
func (g3 *GameObject) SetColor3(r, g, b uint8) {
	g3.SetVar("color", NewColor2(r, g, b))
}
func (g3 *GameObject) SetColor2(r, g, b, a uint8) {
	g3.SetVar("color", NewColor(r, g, b, a))
}
func (g *GameObject) Color() Color {
	return g.GetVar("color").(Color)
}
func (g *GameObject) Draw() {
	if IsRunning() == false {
		panic("the rufi and raylib not working")
	}
	if IsValidShape(g.shape) == false {
		panic("the shape is invalid !!!!!!!!!!!!!!!!!!!!!!")
	} else {
		if g.hide == false {

			if g.shape == SHAPE_RECTANGLE {
				rl.DrawRectangle(int32(g.position.X), int32(g.position.Y), int32(g.size.W), int32(g.size.H), ConvertColor(g.GetVar("color").(Color)))
			} else if g.shape == SHAPE_IMAGE {
				if g.GetVar("image") != nil {
					g.GetVar("image").(*FuriaImage).DrawImage(g.Position(), g.GetVar("color").(Color), g.size)
				}
			} else if g.shape == SHAPE_ANIMATION_IMAGE {
				if g.GetVar("animationImage") != nil {
					g.GetVar("animationImage").(*AnimationImage).Draw(g.GetVar("color").(Color), g.Position(), g.Size())
				}
			} else {

				panic("shape invalid in the drawing object function :(")
			}
		}
	}
}

func (g *GameObject) SetAnimationImage(ani *AnimationImage) {
	if g.shape != SHAPE_ANIMATION_IMAGE {
		panic("the shape not is 'animationImage' :(")
	} else {
		g.SetVar("animationImage", ani)
	}
}
func (g *GameObject) AnimationImage() *AnimationImage {
	if g.shape != SHAPE_ANIMATION_IMAGE {
		panic("the shape not is 'animationImage' :(")
	}
	switch any(g.GetVar("animationImage")).(type) {
	case *AnimationImage:
		return g.GetVar("animationImage").(*AnimationImage)
		break
	}
	return nil
}
func CreateGameObject(name string, shape Shape, size Size, position Position) (*GameObject, error) {
	if IsRunning() == false {
		return nil, ErrGameNotRunning
	} else {
		if IsValidShape(shape) == false {
			return nil, ErrShapeInvalid
		} else {
			keys := make([]string, 0, len(gameObjects))
			for k := range gameObjects {
				keys = append(keys, k)
			}
			j := &GameObject{
				name:     name,
				size:     size,
				position: position,
				shape:    shape,
				funcs:    make(map[string]func(*GameObject, any)),
				hide:     false,
				vars:     make(map[string]any),
				groups:   make(map[string]bool),
			}
			j.collision = NewCollisionRectagle(j)
			gameObjects[name] = j
			//gameObjects[name].collision = NewCollisionRectagle(&gameObjects[name])
			j.SetFunction("start", func(g *GameObject, a any) {

			})
			j.SetFunction("update", func(g *GameObject, a any) {

			})

			j.SetVar("color", NewColor2(0, 0, 0))

			if shape == SHAPE_IMAGE {
				j.SetVar("image", nil)
				j.SetVar("color", NewColor2(255, 255, 255))

			}
			if shape == SHAPE_ANIMATION_IMAGE {
				j.SetVar("animationImage", nil)
				j.SetVar("color", NewColor2(255, 255, 255))
			}
			return j, nil
		}
	}
}
