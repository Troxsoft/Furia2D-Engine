package engine

type CollisionArea struct {
	position    Position
	size        Size
	gob         *GameObject
	sc          *Scene
	onCollision func(*CollisionArea, *GameObject, *GameObjectEvent)
	debugColor  *Color
}

func (p *CollisionArea) Scene() *Scene {
	return p.sc
}
func (p *CollisionArea) Size() Size {
	return p.size
}
func (p *CollisionArea) Position() Position {
	return p.position
}
func NewCollisionArea(scene *Scene, pos Position, siz Size) *CollisionArea {
	p := &CollisionArea{
		position: pos,
		size:     siz,
		gob:      _newGameObjectUnregister(pos, siz, SHAPE_RECTANGLE, len(scene.GameObjectsObjects())+1),
		sc:       scene,
		onCollision: func(coll *CollisionArea, g *GameObject, goe *GameObjectEvent) {

		},
		debugColor: nil,
	}
	scene.areas = append(scene.areas, p)
	return p
}
func (p *CollisionArea) SetDebugColor(color Color) {
	p.debugColor = &color
}
func (p *CollisionArea) QuitDebugColor() {
	p.debugColor = nil
}
func (p *CollisionArea) SetOnCollision(c func(*CollisionArea, *GameObject, *GameObjectEvent)) {
	p.onCollision = c
}
