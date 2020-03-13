package world

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"

	geo "GoTracing/geometry"
	"GoTracing/material"
	obj "GoTracing/object"
	"GoTracing/tracer"
)

const INT_MAX = ^int(0)

func Render(s *Scene) {
	fmt.Println("Rendering ...")
	file, err := os.Create("/home/ling/test.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	vp := s.VPlane
	img := image.NewRGBA(image.Rect(0, 0, int(vp.Width), int(vp.Height)))

	// sample
	numSamples := vp.Samples
	n := int(math.Sqrt(float64(numSamples)))

	for x := 0.0; x < vp.Width; x++ {
		for y := 0.0; y < vp.Height; y++ {
			r := 0
			g := 0
			b := 0
			for p := 0; p < n; p++ {
				for q := 0; q < n; q++ {
					lcoalX := float64(x) - 0.5*float64(vp.Width) + (float64(p)+0.5)/float64(n)
					lcoalY := float64(vp.Height)*0.5 - float64(y) + (float64(q)+0.5)/float64(n)
					ray := s.ViewPoint.GetRay(lcoalX, lcoalY)

					shadeRec := tracer.Tracing(*s.ObjList, *ray)
					localColor := tracer.GetColor(shadeRec, *s.ObjList, s.Light)

					r += int(localColor.R)
					g += int(localColor.G)
					b += int(localColor.B)
				}
			}
			color := color.RGBA{
				R: uint8(r / numSamples),
				G: uint8(g / numSamples),
				B: uint8(b / numSamples),
				A: 255,
			}

			img.Set(int(x), int(y), color)
		}
	}
	jpeg.Encode(file, img, nil)
}

func Build(s *Scene, sceneFile string) {
	// @Todo load configuration file
	fmt.Println("load configuration from ", sceneFile)

	// Add object to scene
	sphere1 := obj.NewSphere(250, 40, 250, 75)
	materail1 := material.SpecularPhong{
		Ks:    0.6,
		Kd:    0.75,
		Cd:    0.5,
		Color: color.RGBA{255, 0, 255, 255},
	}
	sphere1.SetMaterial(materail1)

	sphere2 := obj.NewSphere(250, 40, -250, 75)
	materail2 := material.SpecularPhong{
		Ks:    0.6,
		Kd:    0.75,
		Cd:    0.5,
		Color: color.RGBA{0, 255, 255, 255},
	}
	sphere2.SetMaterial(materail2)

	sphere3 := obj.NewSphere(-250, 40, 250, 75)
	materail3 := material.SpecularPhong{
		Ks:    0.6,
		Kd:    0.75,
		Cd:    0.5,
		Color: color.RGBA{255, 255, 0, 255},
	}
	sphere3.SetMaterial(materail3)

	sphere4 := obj.NewSphere(-250, 40, -250, 75)
	materail4 := material.SpecularPhong{
		Ks:    0.6,
		Kd:    0.75,
		Cd:    0.5,
		Color: color.RGBA{128, 128, 128, 255},
	}
	sphere4.SetMaterial(materail4)

	sphere5 := obj.NewSphere(0, 40, 0, 75)
	materail5 := material.SpecularPhong{
		Ks:    0.1,
		Kd:    0.9,
		Cd:    0.9,
		Color: color.RGBA{255, 255, 255, 255},
	}
	sphere5.SetMaterial(materail5)

	s.ObjList.PushBack(sphere1)
	s.ObjList.PushBack(sphere2)
	s.ObjList.PushBack(sphere3)
	s.ObjList.PushBack(sphere4)
	s.ObjList.PushBack(sphere5)

	plane1 := obj.Plane{
		Normal: geo.Vector3D{
			X: 0,
			Y: 1,
			Z: 0,
		},
		Position: geo.Point3D{
			X: 0,
			Y: -200,
			Z: 0,
		},
	}
	planeMaterial := material.SV_Matte{
		Color1: color.RGBA{255, 255, 255, 255},
		Color2: color.RGBA{0, 0, 0, 255},
		Size:   100,
	}
	// planeMaterial := brdf.Matte {
	// 	Color: color.RGBA {255, 255, 255, 255},
	// }
	plane1.SetMaterial(planeMaterial)

	s.ObjList.PushBack(plane1)
}
