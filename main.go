package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
	example "github.com/seanoneillcode/physics/pkg/example"
	"github.com/seanoneillcode/physics/pkg/physics"
)

func main() {
	rl.InitWindow(1280, 720, "physics")
	rl.SetTargetFPS(60)

	blockModel := rl.LoadModel("pkg/example/block.gltf")
	blockPosition := rl.NewVector3(2, 0, 0)
	var blockScale float32 = 1
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

		ray := rl.Ray{
			Position:  circle.GetPos(),
			Direction: rl.Vector3Normalize(rl.Vector3Subtract(blockPosition, circle.GetPos())),
		}
		for _, m := range blockModel.GetMeshes() {

			collision := rl.GetRayCollisionMesh(ray, m, blockModel.Transform)
			if collision.Hit {
				if collision.Distance < 0 {
					fmt.Printf("hit mesh\n")
					fmt.Printf("normal: %v\n", collision.Normal)
				}
			}
		}

		// camera.Update(circle.GetPos(), 0)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		rl.BeginMode3D(camera.GetCamera())

		rl.DrawCubeWires(rl.Vector3{}, 20, 20, 20, rl.White)
		circle.Draw()
		rl.DrawModel(blockModel, blockPosition, blockScale, rl.White)

		rl.EndMode3D()
		rl.EndDrawing()
	}

	rl.CloseWindow()
}
