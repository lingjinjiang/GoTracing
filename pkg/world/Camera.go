package world

import (
	"GoTracing/pkg/config"
	geo "GoTracing/pkg/geometry"
	"errors"
	"log"
	"strconv"
)

type Camera struct {
	U        geo.Vector3D // local x direction
	V        geo.Vector3D // local y direction
	W        geo.Vector3D // local z direction
	Distance float64      // distance to view plane
	Position geo.Point3D  // position of camera
	VPlane   ViewPlane
}

func (c Camera) GetRay(x float64, y float64) *geo.Ray {
	direction := geo.Vector3D{
		X: c.U.X*x + c.V.X*y - c.W.X*c.Distance,
		Y: c.U.Y*x + c.V.Y*y - c.W.Y*c.Distance,
		Z: c.U.Z*x + c.V.Z*y - c.W.Z*c.Distance,
	}

	return &geo.Ray{
		Endpoint:  c.Position,
		Direction: direction.Normalize(),
	}
}

func NewCamera(cameraInfo config.CameraInfo) (Camera, error) {
	camera := Camera{}

	if u, err := geo.ParseNormalVector(cameraInfo.U); err == nil {
		camera.U = *u
	} else {
		log.Fatal("The u vector is illegal: ", cameraInfo.U)
		return camera, err
	}

	if v, err := geo.ParseNormalVector(cameraInfo.V); err == nil {
		camera.V = *v
	} else {
		log.Fatal("The v vector is illegal: ", cameraInfo.V)
		return camera, err
	}

	if w, err := geo.ParseNormalVector(cameraInfo.W); err == nil {
		camera.W = *w
	} else {
		log.Fatal("The w vector is illegal: ", cameraInfo.W)
		return camera, err
	}

	// check local coordinate
	if camera.U.Dot(camera.V) != 0 || camera.U.Dot(camera.W) != 0 || camera.V.Dot(camera.W) != 0 {
		return camera, errors.New("The camera local cooridnate is invalied.")
	}

	if position, err := geo.ParsePoint(cameraInfo.Position); err == nil {
		camera.Position = *position
	} else {
		log.Fatal("The postion is illegal: ", cameraInfo.Position)
		return camera, err
	}

	if distance, err := strconv.ParseFloat(cameraInfo.Distance, 64); err == nil {
		if distance <= 0.0 {
			return camera, errors.New("The distance should be a positive value" + cameraInfo.Distance)
		}
		camera.Distance = distance
	} else {
		log.Fatal("The distance is illegal: ", cameraInfo.Distance)
		return camera, err
	}

	if vp, err := NewViewPlane(cameraInfo.VPlane); err == nil {
		camera.VPlane = vp
	} else {
		log.Fatal(err)
		e := errors.New("Can't build viewplane")
		return camera, e
	}

	return camera, nil
}
