package main

import (
	"fmt"

	"github.com/leaanthony/clir"
)

func main() {

	// Vars
	var name string
	var age int
	var awesome bool

	// Create new cli
	cli := clir.NewCli("Flags", "An example of using flags", "v0.0.1")

	// Define flags
	cli.StringFlag("name", "n", "Your name", &name).
		IntFlag("age", "a", "Your age", &age).
		BoolFlag("awesome", "", "Are you awesome?", &awesome)

	// Define action
	cli.Action(func() error {

		if name == "" {
			name = "Anonymous"
		}
		fmt.Printf("Hello %s!\n", name)

		if age > 0 {
			fmt.Printf("You are %d years old!\n", age)
		}

		if awesome {
			fmt.Println("You are awesome!")
		}

		return nil
	})

	cli.Run()

}
