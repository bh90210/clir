package main

import (
	"fmt"

	"github.com/leaanthony/clir"
)

func main() {

	// Create new cli
	cli := clir.NewCli("Basic", "A basic example", "v0.0.1")

	// Set long description
	cli.LongDescription("This app prints hello world")

	// Define action
	cli.Action(func() error {
		println("Hello World!")
		return nil
	})

	// Run!
	if err := cli.Run(); err != nil {
		fmt.Println(err)
	}

}
