package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Camera struct {
	camera   rl.Camera2D
	target   Position
	zoom     float32
	rotation float32
	offset   Position
}

func ConvertPosition(pos Position) rl.Vector2 {
	return rl.NewVector2(float32(pos.X), float32(pos.Y))
}
func NewCamera(offset Position, target Position) *Camera {

	return &Camera{
		camera:   rl.NewCamera2D(ConvertPosition(offset), ConvertPosition(target), 0, 1),
		target:   target,
		zoom:     1,
		rotation: 0,
		offset:   offset,
	}
}

func (c *Camera) RlCamera() rl.Camera2D {
	return c.camera
}
func (c *Camera) Offset() Position {

	return c.offset
}
func (c *Camera) SetOffset(o Position) {
	c.offset = o
	c.camera.Offset = ConvertPosition(o)
}
func (c *Camera) SetRotation(r float32) {
	c.rotation = r
	c.camera.Rotation = r
}
func (c *Camera) Rotation() float32 {
	return c.rotation
}
func (c *Camera) SetZoom(nZoom float32) {
	c.zoom = nZoom
	c.camera.Zoom = nZoom

}
func (c *Camera) Zoom() float32 {
	return c.zoom
}
func (c *Camera) SetTarget(target Position) {
	c.target = target
	c.camera.Target = ConvertPosition(target)
	//rl.move
}
func (c *Camera) Target() Position {
	return c.target
}

//const Camera = CreateGameObject("camera", SHAPE_RECTANGLE, rl.GetWindo)
