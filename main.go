package main

import (
	geo "GoTracing/geometry"
	"GoTracing/world"
	"container/list"
)

func main() {
	vp := world.ViewPlane {
		Width: 1280,
		Height: 720,
		Samples: 16,
	}

	camera := world.Camera {
		U: geo.Vector3D { X: 1.0, Y: 0.0, Z: 0.0}.Normalize(),
		V: geo.Vector3D { X: 0.0, Y: 2.0, Z: -1.0}.Normalize(),
		W: geo.Vector3D { X: 0.0, Y: 1.0, Z: 2.0}.Normalize(),
		Distance: 900.0,
		Position: geo.Point3D { X: 0, Y: 500, Z: 1000},
	}

	scene := world.Scene {
		ObjList: list.New(),
		VPlane: &vp,
		ViewPoint: camera,
	}

	world.Build(&scene, "word.yaml")
	world.Render(&scene)
}