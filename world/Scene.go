package world

import (
	"GoTracing/light"
	obj "GoTracing/object"
	"container/list"
)

type Scene struct {
	ViewPoint Camera
	VPlane    *ViewPlane
	Height    int
	Width     int
	Sphere    *obj.Sphere
	ObjList   *list.List
	Light     light.Light
}
