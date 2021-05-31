package main

import (
	"fmt"

	"github.com/leaanthony/clir"
)

func main() {

	// Create new cli
	cli := clir.NewCli("Flags", "An example of using flags", "v0.0.1")

	// Name
	var name string
	cli.StringFlag("name", "Your name", &name)

	// Age
	var age int
	cli.IntFlag("age", "Your age", &age)

	// Awesome
	var awesome bool
	cli.BoolFlag("awesome", "a", "Are you awesome?", &awesome)
	var lol bool
	cli.BoolFlag("lol", "l", "Are you LOL?", &lol)
	var test bool
	cli.BoolFlag("test", "", "Are you TEST?", &test)
	// Define action for the command
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

		if lol {
			fmt.Println("lol")
		}

		return nil
	})

	// Run!
	if err := cli.Run(); err != nil {
		fmt.Printf("Error: %s\n", err)
	}

}
