package world

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
	"sync"
	"time"

	"GoTracing/config"
	geo "GoTracing/geometry"
	"GoTracing/material"
	obj "GoTracing/object"
	"GoTracing/tracer"
)

const INT_MAX = ^int(0)

var (
	wg       sync.WaitGroup
	threadCh chan int
)

func Render(s *Scene, config config.Configuration) {
	fmt.Println("Rendering ..." + config.Output)
	file, err := os.Create(config.Output)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	vp := s.VPlane
	img := image.NewRGBA(image.Rect(0, 0, int(vp.Width), int(vp.Height)))
	threadCh = make(chan int, config.RenderThreads)

	wg.Add(int(vp.Width) * int(vp.Height))
	beginTime := time.Now()
	for x := 0.0; x < vp.Width; x++ {
		for y := 0.0; y < vp.Height; y++ {
			threadCh <- 1
			go Tracing(x, y, vp, s, img)
		}
	}
	wg.Wait()
	finishTime := time.Now()
	fmt.Println("Using", finishTime.Second()-beginTime.Second(), "seconds")
	jpeg.Encode(file, img, nil)
}

func Build(s *Scene, config config.Configuration) {
	// @Todo load configuration file
	//fmt.Println("load configuration from ", sceneFile)

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

	plane1 := obj.Rect{
		Position: geo.Point3D{
			X: 250,
			Y: -100,
			Z: 250,
		},
		Length: 300,
		LVector: geo.Vector3D{
			X: 1,
			Y: 0,
			Z: 0,
		},
		Width: 300,
		WVector: geo.Vector3D{
			X: 0,
			Y: 0,
			Z: 1,
		},
	}

	plane2 := obj.Rect{
		Position: geo.Point3D{
			X: -250,
			Y: -100,
			Z: 250,
		},
		Length: 300,
		LVector: geo.Vector3D{
			X: 1,
			Y: 0,
			Z: 0,
		},
		Width: 300,
		WVector: geo.Vector3D{
			X: 0,
			Y: 0,
			Z: 1,
		},
	}

	plane3 := obj.Rect{
		Position: geo.Point3D{
			X: -250,
			Y: -100,
			Z: -250,
		},
		Length: 300,
		LVector: geo.Vector3D{
			X: 1,
			Y: 0,
			Z: 0,
		},
		Width: 300,
		WVector: geo.Vector3D{
			X: 0,
			Y: 0,
			Z: 1,
		},
	}

	plane4 := obj.Rect{
		Position: geo.Point3D{
			X: 250,
			Y: -100,
			Z: -250,
		},
		Length: 300,
		LVector: geo.Vector3D{
			X: 1,
			Y: 0,
			Z: 0,
		},
		Width: 300,
		WVector: geo.Vector3D{
			X: 0,
			Y: 0,
			Z: 1,
		},
	}

	bottom := obj.Rect{
		Position: geo.Point3D{
			X: 0,
			Y: -150,
			Z: 0,
		},
		Length: 1200,
		LVector: geo.Vector3D{
			X: 1,
			Y: 0,
			Z: 0,
		},
		Width: 1000,
		WVector: geo.Vector3D{
			X: 0,
			Y: 0,
			Z: 1,
		},
	}

	planeMaterial1 := material.SV_Matte{
		Color1: color.RGBA{225, 100, 255, 255},
		Color2: color.RGBA{100, 255, 160, 255},
		Size:   50,
	}

	planeMaterial2 := material.SV_Matte{
		Color1: color.RGBA{100, 200, 255, 255},
		Color2: color.RGBA{225, 240, 120, 255},
		Size:   50,
	}

	bottomMaterial := material.SV_Matte{
		Color1: color.RGBA{255, 255, 255, 255},
		Color2: color.RGBA{0, 0, 0, 255},
		Size:   75,
	}

	plane1.SetMaterial(planeMaterial1)
	plane2.SetMaterial(planeMaterial2)
	plane3.SetMaterial(planeMaterial1)
	plane4.SetMaterial(planeMaterial2)
	bottom.SetMaterial(bottomMaterial)

	s.ObjList.PushBack(plane1)
	s.ObjList.PushBack(plane2)
	s.ObjList.PushBack(plane3)
	s.ObjList.PushBack(plane4)
	s.ObjList.PushBack(bottom)
}

func Tracing(x float64, y float64, vp *ViewPlane, s *Scene, img *image.RGBA) {
	defer wg.Done()
	numSamples := vp.Samples
	n := int(math.Sqrt(float64(numSamples)))
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
	<-threadCh
}
