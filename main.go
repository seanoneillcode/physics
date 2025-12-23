package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	example "github.com/seanoneillcode/physics/pkg/example"
)

func main() {
	rl.InitWindow(1280, 720, "physics")
	rl.SetTargetFPS(60)

	camera := example.NewCustomCamera(true)

	circle := example.NewCircle()
	circle.Move(rl.NewVector3(1, 0, 2))
	camera.Update(circle.GetPos(), 0)

	controller := example.Controller{
		Circle: circle,
	}

	for !rl.WindowShouldClose() {

		controller.Update()
		circle.Update()

		// camera.Update(circle.GetPos(), 0)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.BeginMode3D(camera.GetCamera())

		circle.Draw()

		rl.EndMode3D()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
