package object

import (
	geo "GoTracing/geometry"
	"GoTracing/brdf"

	"math"
)

type Sphere struct {
	center geo.Point3D
	radius float64
	material brdf.BRDF
}

// if the endpoint of the ray is (a, b, c) and direction is (dx, dy, dz)
// so it can be coconfirmed that (x-a)/dx = (y-b)/dy = (z-c)/dz = t,
// so if we get the result t, we can get the hit point.
// and as (x-a)^2 + (y-b)^2 + (z-c)^2 = distance^2 which distance is the hit point between endpoint and dx^2 + dy^2 + dz^2 = 1.
// so the t^2 = distance^2
func (s Sphere) Hit(ray geo.Ray) (bool, geo.Point3D) {
	p := ray.Endpoint
	v := ray.Direction
	radius := s.radius
	center := s.center
	b := 2 * ((p.X - center.X) * v.X + (p.Y - center.Y) * v.Y + (p.Z -center.Z) * v.Z)
	c := (p.X - center.X) * (p.X - center.X) + (p.Y - center.Y) * (p.Y - center.Y) + (p.Z -center.Z) * (p.Z -center.Z) - radius * radius
	delta := b * b - 4 * c
	var isHit bool
	var hitPoint geo.Point3D
	if delta >= 0.0 {
		isHit = true
		resultA := - (b + math.Sqrt(delta)) / 2.0
		resultB := - (b - math.Sqrt(delta)) / 2.0
		var result float64
		if resultA * resultA < resultB * resultB {
			result = resultA
		} else {
			result = resultB
		}

		hitPoint = geo.Point3D {
			X: result * v.X + p.X,
			Y: result * v.Y + p.Y,
			Z: result * v.Z + p.Z,
		}
	} else {
		isHit = false
		hitPoint = geo.Point3D {}
	}

	return isHit, hitPoint
}

func (s Sphere) NormalVector(point geo.Point3D) geo.Vector3D {
	return s.Center().Sub(point).Normalize()
}

func NewSphere(x float64, y float64, z float64, r float64) *Sphere {
	sphere := &Sphere {
		center: geo.Point3D {
			X: x,
			Y: y,
			Z: z,
		},
		radius: r,
	}
	return sphere
}

func (s Sphere) Center() geo.Point3D {
	return s.center
}

func (s Sphere) Radius() float64 {
	return s.radius
}

func (s *Sphere) SetMaterial(material brdf.BRDF) {
	s.material = material
}

func (s Sphere) GetMaterial() brdf.BRDF {
	return s.material
}