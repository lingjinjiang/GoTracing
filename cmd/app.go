package cmd

import (
	"GoTracing/light"
	"GoTracing/world"
	"os"
	"os/user"

	geo "GoTracing/geometry"
	"container/list"
	"fmt"

	"github.com/spf13/cobra"
)

var (
	outputPath    string
	configFile    string
	defaultOutput string
)

func NewGoTracingCommand() *cobra.Command {

	currentUser, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defaultOutput = currentUser.HomeDir + "/" + "GoTracing.jpg"

	cmd := &cobra.Command{
		Run: run,
	}

	flags := cmd.PersistentFlags()

	flags.StringVarP(&outputPath, "out", "o", "", "the image output path")
	flags.StringVarP(&configFile, "config", "c", "", "configuration file")

	return cmd
}

func run(cmd *cobra.Command, args []string) {

	checkConfigurationFile(configFile)

	checkoutOutput(outputPath)

	vp := world.ViewPlane{
		Width:   1280,
		Height:  720,
		Samples: 16,
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

	config := world.Build(&scene, configFile, outputPath)
	world.Render(&scene, config)
}

func checkoutOutput(output string) {
	if output == "" {
		outputPath = defaultOutput
		fmt.Println("[Using default output path \"" + outputPath + "\".]")
	}
}

func checkConfigurationFile(configFile string) {
	fmt.Println("[Configuration file is useless at current time. It will be using in the future.]")
}
