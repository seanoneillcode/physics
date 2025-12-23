package physics

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const speedAmount = 0.1

type Object struct {
	Pos      rl.Vector3
	Velocity rl.Vector3
	Radius   float32
}

func (o *Object) Move(move rl.Vector3) {
	o.Velocity.X = move.X * speedAmount
	o.Velocity.Y = move.Y * speedAmount
	o.Velocity.Z = move.Z * speedAmount
}

type Engine struct {
	objects []*Object
}

func NewEngine() *Engine {
	return &Engine{
		objects: []*Object{},
	}
}

func (e *Engine) Update() {
	for _, obj := range e.objects {
		obj.Pos = rl.Vector3Add(obj.Pos, obj.Velocity)
		// collide with walls
		if obj.Pos.X+obj.Radius > 10 {
			obj.Velocity.X = -obj.Velocity.X
		}
		if obj.Pos.X-obj.Radius < -10 {
			obj.Velocity.X = -obj.Velocity.X
		}
		if obj.Pos.Z+obj.Radius > 10 {
			obj.Velocity.Z = -obj.Velocity.Z
		}
		if obj.Pos.Z-obj.Radius < -10 {
			obj.Velocity.Z = -obj.Velocity.Z
		}
	}
}

func (e *Engine) GetObject() *Object {
	newObj := &Object{
		Pos:      rl.NewVector3(0, 0, 0),
		Velocity: rl.NewVector3(0, 0, 0),
		Radius:   1.0,
	}
	e.objects = append(e.objects, newObj)
	return newObj
}

func (e *Engine) RemoveObject(obj *Object) {
	newObjects := []*Object{}
	for _, o := range e.objects {
		if o != obj {
			newObjects = append(newObjects, o)
		}
	}
	e.objects = newObjects
}
