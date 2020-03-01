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
	obj "GoTracing/object"
	"GoTracing/brdf"
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
					lcoalX := float64(x) - 0.5 * float64(vp.Width) + (float64(p) + 0.5) / float64(n)
					lcoalY := float64(vp.Height) * 0.5 - float64(y) + (float64(q) + 0.5) / float64(n)
					ray := s.ViewPoint.GetRay(lcoalX, lcoalY)

					localColor := Tracing(*s, *ray)
					r += int(localColor.R)
					g += int(localColor.G)
					b += int(localColor.B)
				}
			}
			color := color.RGBA {
				R: uint8(r / numSamples),
				G: uint8(g / numSamples),
				B: uint8(b / numSamples),
				A: 255,
			}
			// ray := s.ViewPoint.GetRay(float64(x) - float64(vp.Width) / 2, float64(vp.Height) / 2 - float64(y))
			// color := Tracing(*s, *ray)
			img.Set(int(x), int(y), color)
		}
	}
	jpeg.Encode(file, img, nil)
}

func Tracing(s Scene, ray geo.Ray) color.RGBA {
	objList := s.ObjList
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
			if min == -1.0 || distance < min {
				min = distance
				hitPoint = currentHitPoint
				hitObject = obj
			}
		}
	}

	if isHit {
		vOut := ray.Direction.Opposite()
		vIn := geo.Vector3D {
			X: 1,
			Y: 1,
			Z: 1,
		}.Normalize()

		return hitObject.GetMaterial().Shade(vIn, vOut, hitObject.NormalVector(hitPoint), hitPoint)
	} else {
		return color.RGBA{20, 20, 20, 255}
	}
}

func Build(s *Scene, sceneFile string) {
	// @Todo load configuration file
	fmt.Println("load configuration from ", sceneFile)
	
	// Add object to scene
	sphere1 := obj.NewSphere(250, 40, 250, 75)
	materail1 := brdf.Phong {
		Ks: 0.6,
		Kd: 0.75,
		Cd: 0.5,
		Color: color.RGBA {255, 0, 255, 255},
	}
	sphere1.SetMaterial(materail1)

	sphere2 := obj.NewSphere(250, 40, -250, 75)
	materail2 := brdf.Phong {
		Ks: 0.6,
		Kd: 0.75,
		Cd: 0.5,
		Color: color.RGBA{0, 255, 255, 255},
	}
	sphere2.SetMaterial(materail2)

	sphere3 := obj.NewSphere(-250, 40, 250, 75)
	materail3 := brdf.Phong {
		Ks: 0.6,
		Kd: 0.75,
		Cd: 0.5,
		Color: color.RGBA{255, 255, 0, 255},
	}
	sphere3.SetMaterial(materail3)

	sphere4 := obj.NewSphere(-250, 40, -250, 75)
	materail4 := brdf.Phong {
		Ks: 0.6,
		Kd: 0.75,
		Cd: 0.5,
		Color: color.RGBA{128, 128, 128, 255},
	}
	sphere4.SetMaterial(materail4)
	
	s.ObjList.PushBack(sphere1)
	s.ObjList.PushBack(sphere2)
	s.ObjList.PushBack(sphere3)
	s.ObjList.PushBack(sphere4)

	plane1 := obj.Plane {
		Normal: geo.Vector3D {
			X: 0,
			Y: 1,
			Z: 0,
		},
		Position: geo.Point3D {
			X: 0,
			Y: -200,
			Z: 0,
		},
	}
	planeMaterial := brdf.SV_Matte {
		Color1: color.RGBA {255, 255, 255, 255},
		Color2: color.RGBA {0, 0, 0, 255},
		Size: 100,
	}
	// planeMaterial := brdf.Matte {
	// 	Color: color.RGBA {255, 255, 255, 255},
	// }
	plane1.SetMaterial(planeMaterial)

	s.ObjList.PushBack(plane1)
}