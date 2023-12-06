package engine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

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
func RayCast(idToIgnore int, initPos Position, dirreccion Position, long uint32, jump uint16) *GameObject {
	var i uint32
	ojb := _newGameObjectUnregister(initPos, NewSize(2, 2), SHAPE_RECTANGLE, idToIgnore)
	for i = 0; i < long; i += uint32(jump) {
		x := initPos.X + dirreccion.X + float64(i)
		if dirreccion.X < 0 {
			x = initPos.X + dirreccion.X - float64(i)
		}
		if dirreccion.X == 0 {
			x = initPos.X
		}

		y := initPos.Y + dirreccion.Y + float64(i)
		if dirreccion.Y < 0 {
			y = initPos.Y + dirreccion.Y + float64(i)
		}
		if dirreccion.Y == 0 {
			y = initPos.Y
		}
		vectorPos := NewPosition(x, y)
		//fmt.Println(vectorPos)
		if cp23 := ojb.collision.OnCollisionPos(vectorPos); cp23 != nil {
			return cp23
		}
	}
	return nil
}
func DrawRayCast(initPos Position, dirreccion Position, long uint32, color Color) {
	rl.DrawLine(int32(initPos.X), int32(initPos.Y), int32(initPos.X+(dirreccion.X*float64(long))), int32(initPos.Y+(dirreccion.Y*float64(long))), ConvertColor(color))
}

/*
physics type: Colliders and collisions[ON]
*/
func (g *GameObject) MoveTo(npos Position) *GameObject {

	coll2 := _newGameObjectUnregister(npos, NewSize(g.size.W, g.size.H), g.shape, g.id)
	if h := coll2.collision.OnCollision(); h != nil {

		return h
	} else {
		g.SetPosition(npos)
		return nil
	}

}
