package engine

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Shape string

const (
	SHAPE_RECTANGLE Shape = "rectangle"
	SHAPE_IMAGE     Shape = "image"
)

func ConvertColor(color Color) rl.Color {
	return rl.NewColor(color.R, color.G, color.B, color.A)
}

func IsValidShape(shape Shape) bool {

	if shape != "rectangle" && shape != "image" {
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
var instancesGameObjects []*GameObject
var (
	ErrShapeInvalid = errors.New("the shape is invalid")
)

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
func GetGameObjectsInTheGroup(gr string) []*GameObject {
	groupsd := []*GameObject{}
	for i := 0; i < len(instancesGameObjects); i++ {
		if instancesGameObjects[i].IsInGroup(gr) {
			groupsd = append(groupsd, instancesGameObjects[i])
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
func (g *GameObject) SetPosition2(x, y int32) {
	g.position.X = x
	g.position.Y = y
}
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
func (g *GameObject) SetImage(img *FuriImage) {
	if g.shape == SHAPE_IMAGE {

		g.vars["image"] = img
	} else {
		panic("the image not is valid")
	}
}
func (g *GameObject) AddToGroup(groupName string) {
	g.groups[groupName] = true
}

func InstanceGameObject(name string, params any) *GameObject {
	i := gameObjects[name]
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
		id:       len(instancesGameObjects) + 1,
		groups:   i.groups,
	}
	k.collision = NewCollisionRectagle(k)
	k.Execute("start", params)
	instancesGameObjects = append(instancesGameObjects, k)
	return k
}

func (g *GameObject) SetFunction(nameFunction string, function func(*GameObject, any)) {
	g.funcs[nameFunction] = function
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
func (g *GameObject) Draw() {
	if IsRunning() == false {
		panic("the rufi and raylib not working")
	}
	if IsValidShape(g.shape) == false {
		panic("the shape is invalid !!!!!!!!!!!!!!!!!!!!!!")
	} else {
		if g.hide == false {

			if g.shape == SHAPE_RECTANGLE {
				rl.DrawRectangle(g.position.X, g.position.Y, int32(g.size.W), int32(g.size.H), ConvertColor(g.GetVar("color").(Color)))
			} else if g.shape == SHAPE_IMAGE {
				if g.GetVar("image") != nil {
					g.GetVar("image").(*FuriImage).DrawImage(g.Position(), g.GetVar("color").(Color), g.size)
				}
			}
		}
	}
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
			return j, nil
		}
	}
}
