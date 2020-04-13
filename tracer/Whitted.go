package tracer

import (
	geo "GoTracing/geometry"
	"GoTracing/light"
	"GoTracing/material"
	"container/list"
	"image/color"
)

type Whitted struct {
	MaxDepth int
}

func (t Whitted) Tracing(objList list.List, light light.Light, ray geo.Ray) material.ShadeRec {
	return material.ShadeRec{}
}

func (t Whitted) GetColor(shadeRec material.ShadeRec, objList list.List, light light.Light) color.RGBA {
	return color.RGBA{}
}

func NewWhitted() Tracer {
	return Whitted{}
}

func (t Whitted) Tracing2(objList list.List, shadeRec *material.ShadeRec) color.RGBA {
	return BACKGROUND
}
