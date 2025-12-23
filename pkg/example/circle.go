package example

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/seanoneillcode/physics/pkg/physics"
)

type Circle struct {
	obj *physics.Object
}

func NewCircle(e *physics.Engine) *Circle {
	return &Circle{
		obj: e.GetObject(),
	}
}

func (p *Circle) Update() {
	// p.obj.Update()
}

func (p *Circle) Draw() {
	rl.DrawSphere(p.obj.Pos, p.obj.Radius, rl.Red)
}

func (p *Circle) GetPos() rl.Vector3 {
	return p.obj.Pos
}

func (p *Circle) Move(move rl.Vector3) {
	p.obj.Move(move)
}
