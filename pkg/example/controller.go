package example

import rl "github.com/gen2brain/raylib-go/raylib"

type input struct {
	left  bool
	right bool
	up    bool
	down  bool
}

func getInput() input {
	i := input{
		left:  rl.IsKeyDown(rl.KeyLeft),
		right: rl.IsKeyDown(rl.KeyRight),
		up:    rl.IsKeyDown(rl.KeyUp),
		down:  rl.IsKeyDown(rl.KeyDown),
	}
	return i
}

type Controller struct {
	Circle *Circle
}

func (p *Controller) Update() {
	in := getInput()
	move := rl.Vector3{}
	hasInput := false
	if in.left {
		move.X = -1
		hasInput = true
	}
	if in.right {
		move.X = 1
		hasInput = true
	}
	if in.up {
		move.Z = -1
		hasInput = true
	}
	if in.down {
		move.Z = 1
		hasInput = true
	}
	if hasInput {
		p.Circle.Move(move)
	}
}
