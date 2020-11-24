package world

import (
	"GoTracing/pkg/light"
	obj "GoTracing/pkg/object"
	"GoTracing/pkg/tracer"
	"container/list"
)

type Scene struct {
	ViewPoint Camera
	Height    int
	Width     int
	Sphere    *obj.Sphere
	ObjList   *list.List
	Light     light.Light
	Tracer    tracer.Tracer
}
