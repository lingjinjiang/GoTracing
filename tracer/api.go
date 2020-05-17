package tracer

import (
	"GoTracing/material"
	"container/list"
	"image/color"
)

var BACKGROUND color.RGBA = color.RGBA{20, 20, 20, 255}

type Tracer interface {
	Tracing(objList list.List, shadeRec *material.ShadeRec) color.RGBA
	GetMaxDepth() uint
}
