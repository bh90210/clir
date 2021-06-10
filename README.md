### clir fork allowing for -s, --short.

### Example

```go
package main

import (
   "fmt"
   "log"

	"github.com/leaanthony/clir"
)

func main() {

	// Create new cli
	cli := clir.NewCli("Flags", "A simple example", "v0.0.1")

	// Name
	name := "Anonymous"
	cli.StringFlag("name", "short, "Your name", &name)

	// Define action for the command
	cli.Action(func() error {
		fmt.Printf("Hello %s!\n", name)
		return nil
	})

	if err := cli.Run(); err != nil {
		fmt.Printf("Error encountered: %v\n", err)
	}

}
```

#### Generated Help

```shell
$ flags --help
Flags v0.0.1 - A simple example

Flags:

  -help
        Get help on the 'flags' command.
  -name string
        Your name
```

#### Documentation

The main documentation may be found [here](https://clir.leaanthony.com).
