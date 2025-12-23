package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	example "github.com/seanoneillcode/physics/pkg/example"
	"github.com/seanoneillcode/physics/pkg/physics"
)

func main() {
	rl.InitWindow(1280, 720, "physics")
	rl.SetTargetFPS(60)

	camera := example.NewCustomCamera(true)
	engine := physics.NewEngine()

	circle := example.NewCircle(engine)
	circle.Move(rl.NewVector3(1, 0, 2))
	camera.Update(circle.GetPos(), 0)

	controller := example.Controller{
		Circle: circle,
	}

	for !rl.WindowShouldClose() {

		controller.Update()
		engine.Update()
		circle.Update()

		// camera.Update(circle.GetPos(), 0)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.BeginMode3D(camera.GetCamera())

		rl.DrawCubeWires(rl.Vector3{}, 20, 20, 20, rl.White)
		circle.Draw()

		rl.EndMode3D()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
