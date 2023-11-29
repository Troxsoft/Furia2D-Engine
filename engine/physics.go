package engine

import "fmt"

func _newGameObjectUnregister(pos Position, siz Size) *GameObject {
	return &GameObject{
		name: fmt.Sprintf("%s_'un_register_______________________________'", len(gameObjects)),
	}
}

/*
physics type: Colliders and collisions[ON]
*/
func (g *GameObject) MoveTo(npos Position) bool {

	coll2 := _newGameObjectUnregister(npos, NewSize(g.size.W, g.size.H))
	if h, _ := coll2.collision.OnCollision(); h != nil {

		return false
	} else {
		g.SetPosition(npos)
		return true
	}

}
