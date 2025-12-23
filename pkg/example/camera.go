package example

import rl "github.com/gen2brain/raylib-go/raylib"

const cameraOrbitSpeed = 0.05

type CustomCamera struct {
	camera   rl.Camera3D
	target   rl.Vector3
	distance float32
	rotation float32
	height   float32

	targetRotation float32
}

func NewCustomCamera(isFlat bool) *CustomCamera {
	camera := rl.Camera3D{}
	camera.Target = rl.NewVector3(0, 0, 0)
	camera.Up = rl.NewVector3(0.0, 1.0, 0.0)
	var distance float32
	if isFlat {
		camera.Fovy = 45.0
		camera.Projection = rl.CameraOrthographic
		camera.Position = rl.NewVector3(0, 1, 0)
		distance = 0.01
	} else {
		camera.Fovy = 45.0
		camera.Projection = rl.CameraPerspective
		camera.Position = rl.NewVector3(0, 1, 0)
		distance = 6
	}

	return &CustomCamera{
		camera:   camera,
		target:   rl.Vector3{},
		distance: distance,
		rotation: 0,
		height:   4,
	}
}

func (c *CustomCamera) Update(target rl.Vector3, targetRotation float32) {
	c.target = target
	c.rotation = rl.Lerp(c.rotation, targetRotation, cameraOrbitSpeed)
	var rotation = rl.MatrixRotate(rl.GetCameraUp(&c.camera), float32(c.rotation))
	view := rl.Vector3Transform(rl.Vector3{X: 0, Y: c.height, Z: c.distance}, rotation)
	c.camera.Position = rl.Vector3Add(c.target, view)
	c.camera.Target = target
}

func (c *CustomCamera) GetCamera() rl.Camera {
	return c.camera
}
