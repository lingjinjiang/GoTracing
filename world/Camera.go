package world

import (
	geo "GoTracing/geometry"
)

type Camera struct {
	U        geo.Vector3D // local x direction
	V        geo.Vector3D // local y direction
	W        geo.Vector3D // local z direction
	Distance float64      // distance to view plane
	Position geo.Point3D  // position of camera
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
