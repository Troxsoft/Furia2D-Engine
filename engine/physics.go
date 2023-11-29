package engine

import "fmt"

func _newGameObjectUnregister(pos Position, siz Size, shape Shape, id int) *GameObject {
	j := &GameObject{
		name:     fmt.Sprintf("%s_'un_register_______________________________'", len(gameObjects)),
		size:     siz,
		position: pos,
		shape:    shape,
		id:       id,
	}
	j.collision = NewCollisionRectagle(j)
	return j
}

/*
physics type: Colliders and collisions[ON]
*/
func (g *GameObject) MoveTo(npos Position) (*GameObject, bool) {

	coll2 := _newGameObjectUnregister(npos, NewSize(g.size.W, g.size.H), g.shape, g.id)
	if h, _ := coll2.collision.OnCollision(); h != nil {

		return h, false
	} else {
		g.SetPosition(npos)
		return nil, true
	}

}
