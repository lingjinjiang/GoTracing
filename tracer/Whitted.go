package tracer

import (
	geo "GoTracing/geometry"
	"GoTracing/material"
	obj "GoTracing/object"
	"container/list"
	"image/color"
)

type Whitted struct {
	maxDepth uint
}

func NewWhitted() Tracer {
	return Whitted{
		maxDepth: 2,
	}
}

func (t Whitted) Tracing(objList list.List, shadeRec *material.ShadeRec) color.RGBA {
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
	shadeRec.Depth -= 1

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
			Depth: 1,
		}

		t.Tracing(objList, &lightShadeRec)
		color = shadeRec.Material.Shade(*shadeRec, !lightShadeRec.IsHit, BACKGROUND)
	} else {
		shadeRec.IsHit = false
		color = BACKGROUND
	}

	return color
}

func (t Whitted) GetMaxDepth() uint {
	return t.maxDepth
}
