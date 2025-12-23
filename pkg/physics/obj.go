package physics

import rl "github.com/gen2brain/raylib-go/raylib"

const speedAmount = 0.1

type Object struct {
	Pos      rl.Vector3
	Velocity rl.Vector3
	Radius   float32
}

func (o *Object) Update() {
	o.Pos = rl.Vector3Add(o.Pos, o.Velocity)
}

func (o *Object) Move(move rl.Vector3) {
	o.Velocity.X = move.X * speedAmount
	o.Velocity.Y = move.Y * speedAmount
	o.Velocity.Z = move.Z * speedAmount
}
