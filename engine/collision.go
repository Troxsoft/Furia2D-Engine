package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type CollisionRectangle struct {
	gameObject *GameObject
	disabled   bool
}

/*
return a.X + a2.X > b.X &&

	a.X < b.X + b2.X &&
	a.Y + a2.Y > b.Y &&
	a.Y < b.Y + b2.Y
*/
func NewCollisionRectagle(g *GameObject) CollisionRectangle {
	return CollisionRectangle{
		gameObject: g,
		disabled:   false,
	}
}
func (cr *CollisionRectangle) OnCollisionInTheGroup(group string) *GameObject {
	if h := cr.OnCollision(); h != nil {
		if h.IsInGroup(group) {
			return h
		}

	}
	return nil
}

func (cr *CollisionRectangle) OnCollision() *GameObject {
	for i := 0; i < len(scene_32445.instancesGameObjects); i++ {

		if cr.OnCollisionTo(scene_32445.instancesGameObjects[i].CollisionRect()) {
			return scene_32445.instancesGameObjects[i]
		}
	}
	return nil

}

func (cr *CollisionRectangle) OnCollisionPos(pos Position) *GameObject {
	for i := 0; i < len(scene_32445.instancesGameObjects); i++ {

		if cr.OnCollisionInPosition(scene_32445.instancesGameObjects[i].CollisionRect(), pos) {
			return scene_32445.instancesGameObjects[i]
		}
	}
	return nil

}

func (cr *CollisionRectangle) OnCollisionInPosition(other CollisionRectangle, pos Position) bool {
	if cr.disabled {
		return false
	}
	if other.disabled {
		return false
	}
	if cr.gameObject.id == other.gameObject.id {
		return false
	}
	return rl.CheckCollisionRecs(
		rl.NewRectangle(float32(pos.X), float32(pos.Y), float32(cr.gameObject.size.W), float32(cr.gameObject.size.H)),
		rl.NewRectangle(float32(other.gameObject.position.X), float32(other.gameObject.position.Y), float32(other.gameObject.size.W), float32(other.gameObject.size.H)))
}

func (cr *CollisionRectangle) OnCollisionTo(other CollisionRectangle) bool {
	if cr.disabled {
		return false
	}
	if other.disabled {
		return false
	}
	if cr.gameObject.id == other.gameObject.id {
		//fmt.Println(cr.gameObject.id)
		return false
	}
	return rl.CheckCollisionRecs(
		rl.NewRectangle(float32(cr.gameObject.position.X), float32(cr.gameObject.position.Y), float32(cr.gameObject.size.W), float32(cr.gameObject.size.H)),
		rl.NewRectangle(float32(other.gameObject.position.X), float32(other.gameObject.position.Y), float32(other.gameObject.size.W), float32(other.gameObject.size.H)))
}
