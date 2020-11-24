package cmd

import (
	"GoTracing/pkg/config"
	"GoTracing/pkg/world"
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

	conf, err := config.NewConfiguration(configFile)
	if err != nil {
		log.Fatal("Unable to load configuration file", configFile)
		return
	}

	checkoutOutput(&conf, outputPath)

	camera, err := world.NewCamera(conf.Camera)
	if err != nil {
		log.Fatal("Build camera failed.", err)
		return
	}

	tracer := config.GenerateTracer(conf)
	if tracer == nil {
		log.Fatal("Build tracer failed.")
		return
	}

	scene := world.Scene{
		ObjList:   list.New(),
		ViewPoint: camera,
		Tracer:    tracer,
	}

	world.Build(&scene, conf)
	world.Render(&scene, conf)
}

func checkoutOutput(config *config.Configuration, output string) {
	if output != "" {
		config.Main.Output = output
		log.Println("[Using give output path \"" + outputPath + "\" to override the path in configuration.]")
	}

}
