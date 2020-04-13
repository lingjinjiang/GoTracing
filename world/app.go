package world

import (
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
	"sync"
	"time"

	"github.com/nfnt/resize"

	"GoTracing/config"
	"GoTracing/light"
)

const INT_MAX = ^int(0)

var (
	wg       sync.WaitGroup
	threadCh chan int
)

func Render(s *Scene, config config.Configuration) {
	log.Println("Rendering", config.Main.Output, "...")
	file, err := os.Create(config.Main.Output)
	if err != nil {
		log.Fatal(err, "Unable to create image", config.Main.Output)
	}
	defer file.Close()

	vp := s.ViewPoint.VPlane
	img := image.NewRGBA(image.Rect(0, 0, int(vp.Width), int(vp.Height)))
	threadCh = make(chan int, config.Main.RenderThreads)

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
	log.Println("Using", finishTime.Second()-beginTime.Second(), "seconds")
	// using main config to build image, using view plane to tracing
	output := resize.Resize(uint(config.Main.Width), uint(config.Main.Height), img, resize.Lanczos3)
	jpeg.Encode(file, output, nil)
}

func Build(s *Scene, conf config.Configuration) {
	objlist := config.GenerateObjects(conf)
	if objlist == nil {
		return
	}

	s.ObjList = objlist

	lightList := config.GenerateLights(conf)
	if lightList == nil {
		return
	}
	s.Light = lightList.Back().Value.(light.Light)
}

func Tracing(x float64, y float64, vp ViewPlane, s *Scene, img *image.RGBA) {
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

			// shadeRec := material.ShadeRec{
			// 	Light: s.Light,
			// 	Ray:   *ray,
			// }
			// localColor := s.Tracer.Tracing2(*s.ObjList, &shadeRec)

			shadeRec := s.Tracer.Tracing(*s.ObjList, s.Light, *ray)
			localColor := s.Tracer.GetColor(shadeRec, *s.ObjList, s.Light)

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
