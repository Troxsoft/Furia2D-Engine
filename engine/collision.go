package engine

import rl "github.com/gen2brain/raylib-go/raylib"

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
func (cr *CollisionRectangle) OnCollisionInTheGroup(group string) (*GameObject, bool) {
	if h, _ := cr.OnCollision(); h != nil {
		if h.IsInGroup(group) {
			return h, true
		}

	}
	return nil, false
}

func (cr *CollisionRectangle) OnCollision() (*GameObject, bool) {
	for i := 0; i < len(instancesGameObjects); i++ {

		if cr.OnCollisionTo(instancesGameObjects[i].CollisionRect()) {
			return instancesGameObjects[i], true
		}
	}
	return nil, false

}

func (cr *CollisionRectangle) OnCollisionTo(other CollisionRectangle) bool {
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
		rl.NewRectangle(float32(cr.gameObject.position.X), float32(cr.gameObject.position.Y), float32(cr.gameObject.size.W), float32(cr.gameObject.size.H)),
		rl.NewRectangle(float32(other.gameObject.position.X), float32(other.gameObject.position.Y), float32(other.gameObject.size.W), float32(other.gameObject.size.H)))
}
