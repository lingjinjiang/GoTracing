package main

import (
	"GoTracing/cmd"
	"log"
)

func main() {
	cmd := cmd.NewGoTracingCommand()
	err := cmd.Execute()
	if err != nil {
		log.Fatal("Some error happend", err)
	}
}
