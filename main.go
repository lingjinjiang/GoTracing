package main

import (
	"GoTracing/cmd"
	"fmt"
)

func main() {
	cmd := cmd.NewGoTracingCommand()
	err := cmd.Execute()
	if err != nil {
		fmt.Println(err)
	}
}
