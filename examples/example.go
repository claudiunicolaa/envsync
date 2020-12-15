package main

import (
	"fmt"
	"github.com/claudiunicolaa/envsync"
)

func main() {
	// because environment files aren't into the root directory we need to add "examples/"
	_, err := envsync.EnvSync("examples/.env", "examples/.env.example")

	if err != nil {
		fmt.Println(err)
		return
	}
	// ...
}
