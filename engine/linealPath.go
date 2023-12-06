package engine

type LinealPath struct {
	PointA Position
	PointB Position
}

func NewLinealPath(a, b Position) *LinealPath {

	return &LinealPath{
		PointA: a,
		PointB: b,
	}
}
func (lp *LinealPath) TeleportToPath(obj *GameObject) {
	obj.MoveTo(lp.PointA)

}
func (lp *LinealPath) IsDone(obj *GameObject) bool {
	return obj.Pos().X == lp.PointB.X && obj.Pos().Y == lp.PointB.Y

}
func (lp *LinealPath) Travel(speed float64, obj *GameObject) {

	if obj.Pos().X != lp.PointB.X {
		obj.MoveTo(NewPos(obj.position.X+float64(speed), obj.position.Y))
	}
	if obj.Pos().Y != lp.PointB.Y {
		obj.MoveTo(NewPos(obj.position.X, obj.position.Y+speed))
	}

}
