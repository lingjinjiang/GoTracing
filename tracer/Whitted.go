package tracer

import (
	geo "GoTracing/geometry"
	"GoTracing/light"
	"GoTracing/material"
	"container/list"
)

type Whitted struct {
	MaxDepth int
}

func (t Whitted) Tracing(objList list.List, light light.Light, ray geo.Ray) material.ShadeRec {
	return material.ShadeRec{}
}
