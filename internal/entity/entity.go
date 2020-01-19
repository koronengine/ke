package entity

type Entity struct {
	controllable bool
	static       bool
	pos          Position
	collision    Collision
}

func New(p Position) Entity {
	return Entity{
		pos: p,
	}
}

// Pos returns the Position of the entity
func (e *Entity) Pos() Position {
	return e.pos
}

// SetPos sets the Position of the entity
func (e *Entity) SetPos(p Position) {
	e.pos = p
}

// Collision returns the collision function
func (e *Entity) Collision() Collision {
	return e.collision
}

// SetCollision sets the collision function
func (e *Entity) SetCollision(c Collision) {
	e.collision = c
}

// Controllable returns if the Entity is human controllable or not
func (e *Entity) Controllable() bool {
	return e.controllable
}

// SetControllable sets if the Entity is human controllable or not
func (e *Entity) SetControllable(c bool) {
	e.controllable = c
}

// Static returns if the entity is static or not
func (e *Entity) Static() bool {
	return e.static
}

// SetStatic sets the Entity to be static or not
func (e *Entity) SetStatic(s bool) {
	e.static = s
}
