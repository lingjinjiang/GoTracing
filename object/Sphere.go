package object

import (
	geo "GoTracing/geometry"
	"GoTracing/material"
	"errors"
	"log"
	"strconv"

	"math"
)

type Sphere struct {
	center   geo.Point3D
	radius   float64
	material material.Material
	localX   geo.Vector3D
	localY   geo.Vector3D
	localZ   geo.Vector3D
}

// if the endpoint of the ray is (a, b, c) and direction is (dx, dy, dz)
// so it can be confirmed that (x-a)/dx = (y-b)/dy = (z-c)/dz = t,
// so if we get the result t, we can get the hit point.
// and as (x-a)^2 + (y-b)^2 + (z-c)^2 = distance^2 which distance is the hit point between endpoint and dx^2 + dy^2 + dz^2 = 1.
// so the t^2 = distance^2
func (s Sphere) Hit(ray geo.Ray) (bool, geo.Point3D) {
	p := ray.Endpoint
	v := ray.Direction
	radius := s.radius
	center := s.center
	b := 2 * ((p.X-center.X)*v.X + (p.Y-center.Y)*v.Y + (p.Z-center.Z)*v.Z)
	c := (p.X-center.X)*(p.X-center.X) + (p.Y-center.Y)*(p.Y-center.Y) + (p.Z-center.Z)*(p.Z-center.Z) - radius*radius
	delta := b*b - 4*c
	var isHit bool
	var hitPoint geo.Point3D
	if delta >= 0.0 {
		isHit = true
		resultA := -(b + math.Sqrt(delta)) / 2.0
		resultB := -(b - math.Sqrt(delta)) / 2.0
		var result float64
		if resultA*resultA < resultB*resultB {
			result = resultA
		} else {
			result = resultB
		}

		if result < 0.00001 {
			isHit = false
		}

		hitPoint = geo.Point3D{
			X: result*v.X + p.X,
			Y: result*v.Y + p.Y,
			Z: result*v.Z + p.Z,
		}
	} else {
		isHit = false
		hitPoint = geo.Point3D{}
	}

	return isHit, hitPoint
}

func (s Sphere) NormalVector(point geo.Point3D) geo.Vector3D {
	return s.Center().Sub(point).Normalize()
}

func NewSphere(material material.Material, args map[string]string) (Object, error) {
	sphere := Sphere{}
	if r, err := strconv.ParseFloat(args["radius"], 64); err == nil {
		if r <= 0.0 {
			return nil, errors.New("The radius should be a positive value: " + args["radius"])
		}
		sphere.SetRadius(r)
	} else {
		log.Fatal("The radius is illegal: ", args["radius"])
		return nil, err
	}
	if position, err := geo.ParsePoint(args["center"]); err == nil {
		sphere.SetCenter(*position)
	} else {
		log.Fatal("The postion is illegal: ", args["center"])
		return nil, err
	}

	if localX, err := geo.ParseNormalVector(args["localX"]); err == nil {
		sphere.localX = *localX
	} else {
		log.Fatal("The length is illegal: ", args["localX"])
		return nil, err
	}

	if localY, err := geo.ParseNormalVector(args["localY"]); err == nil {
		sphere.localY = *localY
	} else {
		log.Fatal("The length is illegal: ", args["localY"])
		return nil, err
	}

	if localZ, err := geo.ParseNormalVector(args["localZ"]); err == nil {
		sphere.localZ = *localZ
	} else {
		log.Fatal("The length is illegal: ", args["localZ"])
		return nil, err
	}

	// check local coordinate
	if sphere.localX.Dot(sphere.localY) != 0 || sphere.localX.Dot(sphere.localZ) != 0 || sphere.localY.Dot(sphere.localZ) != 0 {
		return nil, errors.New("The local cooridnate is invalied.")
	}

	sphere.SetMaterial(material)

	return sphere, nil
}

func (s Sphere) Center() geo.Point3D {
	return s.center
}

func (s *Sphere) SetCenter(center geo.Point3D) {
	s.center = center
}

func (s Sphere) GetRadius() float64 {
	return s.radius
}

func (s *Sphere) SetMaterial(material material.Material) {
	s.material = material
}

func (s Sphere) GetMaterial() material.Material {
	return s.material
}

func (s Sphere) GetPosition() geo.Point3D {
	return geo.Point3D{}
}

func (s Sphere) GetLocalX() geo.Vector3D {
	return s.localX
}

func (s Sphere) GetLocalY() geo.Vector3D {
	return s.localY
}

func (s Sphere) GetLocalZ() geo.Vector3D {
	return s.localZ
}

func (s *Sphere) SetRadius(radius float64) {
	s.radius = radius
}
