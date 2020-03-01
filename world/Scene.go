package world

import (
	geo "GoTracing/geometry"
	obj "GoTracing/object"
	"container/list"
)

type Scene struct {
	ViewPoint Camera
	VPlane *ViewPlane
	Height int
	Width int
	Sphere *obj.Sphere
	ObjList *list.List
	Light *geo.Point3D
}