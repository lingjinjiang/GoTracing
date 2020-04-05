package cmd

import (
	"GoTracing/config"
	"GoTracing/world"
	"log"

	"container/list"

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
		log.Fatal("Unable to load configuration file", configFile)
		return
	}

	checkoutOutput(&config, outputPath)

	camera, err := world.NewCamera(config.Camera)
	if err != nil {
		log.Fatal("Build camera failed.", err)
		return
	}

	scene := world.Scene{
		ObjList:   list.New(),
		ViewPoint: camera,
	}

	world.Build(&scene, config)
	world.Render(&scene, config)
}

func checkoutOutput(config *config.Configuration, output string) {
	if output != "" {
		config.Main.Output = output
		log.Println("[Using give output path \"" + outputPath + "\" to override the path in configuration.]")
	}

}
