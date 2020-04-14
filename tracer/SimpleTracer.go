package tracer

import (
	geo "GoTracing/geometry"
	"GoTracing/light"
	"GoTracing/material"
	obj "GoTracing/object"
	"container/list"
	"image/color"
)

type SimpleTracer struct{}

// tracing the ray and return the shade info of this ray
func (t SimpleTracer) Tracing(objList list.List, light light.Light, ray geo.Ray) material.ShadeRec {
	var min float64 = -1.0
	var isHit bool = false
	var hitPoint geo.Point3D
	var hitObject obj.Object = nil

	// if the ray hit multi objects, so return the nearest one
	for i := objList.Front(); i != nil; i = i.Next() {
		obj := i.Value.(obj.Object)
		currentHit, currentHitPoint := obj.Hit(ray)
		distance := ray.Endpoint.Sub(currentHitPoint).Length()
		if currentHit {
			isHit = true
			if min == -1.0 || (distance < min && distance > 0) {
				min = distance
				hitPoint = currentHitPoint
				hitObject = obj
			}
		}
	}

	shadeRec := material.ShadeRec{}

	if isHit {
		shadeRec.IsHit = true
		shadeRec.Material = hitObject.GetMaterial()
		shadeRec.HitPoint = hitPoint
		shadeRec.Normal = hitObject.NormalVector(hitPoint)
		shadeRec.Ray = ray
		shadeRec.Light = light
		shadeRec.ObjPosition = hitObject.GetPosition()
		shadeRec.ObjX = hitObject.GetLocalX()
		shadeRec.ObjY = hitObject.GetLocalY()
		shadeRec.ObjZ = hitObject.GetLocalZ()
	} else {
		shadeRec.IsHit = false
	}
	return shadeRec
}

// get the color of ray
func (t SimpleTracer) GetColor(shadeRec material.ShadeRec, objList list.List, light light.Light) color.RGBA {
	if shadeRec.IsHit {
		localNormal := shadeRec.Normal
		shadeRec.VOut = shadeRec.Ray.Direction.Opposite()

		// temporary light direction
		lightIn := light.GetDirection(shadeRec.HitPoint).Normalize()
		shadeRec.VIn = lightIn

		lcoalRay := geo.Ray{
			Endpoint:  shadeRec.HitPoint,
			Direction: lightIn,
		}

		// simple shadow, if the ray from object hit point to light hit some other objects, then the point is in shadow
		lightShadeRec := t.Tracing(objList, light, lcoalRay)

		// simple diffuse
		rayDirect := shadeRec.Ray.Direction.Normalize()
		diffuseIn := geo.Vector3D{
			X: rayDirect.X - 2*rayDirect.Dot(localNormal)*localNormal.X,
			Y: rayDirect.Y - 2*rayDirect.Dot(localNormal)*localNormal.Y,
			Z: rayDirect.Z - 2*rayDirect.Dot(localNormal)*localNormal.Z,
		}
		diffuseRay := geo.Ray{
			Endpoint:  shadeRec.HitPoint,
			Direction: diffuseIn,
		}
		diffuseShadeRec := t.Tracing(objList, light, diffuseRay)
		diffuseShadeRec.VIn = light.GetDirection(diffuseShadeRec.HitPoint).Normalize()
		diffuseShadeRec.VOut = diffuseIn
		var diffuseColor color.RGBA
		if diffuseShadeRec.IsHit {
			diffuseColor = diffuseShadeRec.Material.Shade(diffuseShadeRec, diffuseShadeRec.IsHit, BACKGROUND)
		} else {
			diffuseColor = BACKGROUND
		}

		return shadeRec.Material.Shade(shadeRec, !lightShadeRec.IsHit, diffuseColor)
	} else {
		return BACKGROUND
	}
}

func NewSimpleTracer() Tracer {
	return SimpleTracer{}
}

func (t SimpleTracer) Tracing2(objList list.List, shadeRec *material.ShadeRec) color.RGBA {
	var min float64 = -1.0
	var isHit bool = false
	var hitPoint geo.Point3D
	var hitObject obj.Object = nil
	ray := shadeRec.Ray
	light := shadeRec.Light
	// if the ray hit multi objects, so return the nearest one
	for i := objList.Front(); i != nil; i = i.Next() {
		obj := i.Value.(obj.Object)
		currentHit, currentHitPoint := obj.Hit(ray)
		distance := ray.Endpoint.Sub(currentHitPoint).Length()
		if currentHit {
			isHit = true
			if min == -1.0 || (distance < min && distance > 0) {
				min = distance
				hitPoint = currentHitPoint
				hitObject = obj
			}
		}
	}

	var color color.RGBA
	if isHit {
		shadeRec.IsHit = true
		shadeRec.Material = hitObject.GetMaterial()
		shadeRec.HitPoint = hitPoint
		shadeRec.Normal = hitObject.NormalVector(hitPoint)
		shadeRec.ObjPosition = hitObject.GetPosition()
		shadeRec.ObjX = hitObject.GetLocalX()
		shadeRec.ObjY = hitObject.GetLocalY()
		shadeRec.ObjZ = hitObject.GetLocalZ()
		shadeRec.VOut = shadeRec.Ray.Direction.Opposite()
		lightIn := light.GetDirection(shadeRec.HitPoint).Normalize()
		shadeRec.VIn = lightIn
		lightShadeRec := material.ShadeRec{
			IsHit: false,
			Light: shadeRec.Light,
			Ray: geo.Ray{
				Endpoint:  shadeRec.HitPoint,
				Direction: lightIn,
			},
		}
		t.Tracing2(objList, &lightShadeRec)

		color = shadeRec.Material.Shade(*shadeRec, !lightShadeRec.IsHit, BACKGROUND)
	} else {
		shadeRec.IsHit = false
		color = BACKGROUND
	}

	return color
}

func (t SimpleTracer) GetMaxDepth() uint {
	return 1
}
