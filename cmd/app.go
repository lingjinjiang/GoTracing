package cmd

import (
	"GoTracing/config"
	"GoTracing/light"
	"GoTracing/world"

	geo "GoTracing/geometry"
	"container/list"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	outputPath string
	configFile string
)

func NewGoTracingCommand() *cobra.Command {

	cmd := &cobra.Command{
		Run: run,
	}

	flags := cmd.PersistentFlags()

	flags.StringVarP(&outputPath, "out", "o", "", "the image output path")
	flags.StringVarP(&configFile, "config", "c", "", "configuration file")

	return cmd
}

func run(cmd *cobra.Command, args []string) {

	config, err := config.NewConfiguration(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}

	checkoutOutput(config, outputPath)

	vp := world.ViewPlane{
		Width:   config.Width,
		Height:  config.Height,
		Samples: config.Sample,
	}

	camera := world.Camera{
		U:        geo.Vector3D{X: 1.0, Y: 0.0, Z: 0.0}.Normalize(),
		V:        geo.Vector3D{X: 0.0, Y: 2.0, Z: -1.0}.Normalize(),
		W:        geo.Vector3D{X: 0.0, Y: 1.0, Z: 2.0}.Normalize(),
		Distance: 900.0,
		Position: geo.Point3D{X: 0, Y: 500, Z: 1000},
	}

	scene := world.Scene{
		ObjList:   list.New(),
		VPlane:    &vp,
		ViewPoint: camera,
		Light: light.SimplePointLight{
			Position: geo.Point3D{
				X: 10000000,
				Y: 10000000,
				Z: 10000000,
			},
		},
	}

	world.Build(&scene, *config)
	world.Render(&scene, *config)
}

func checkoutOutput(config *config.Configuration, output string) {
	if output != "" {
		config.Output = output
		fmt.Println("[Using give output path \"" + outputPath + "\" to override the path in configuration.]")
	}

}
