package main

import (
	"fmt"
	"os"

	"github.com/suseldev/kitty-icebreaker/internal/firecracker"
)

func main() {
	if len(os.Args) != 2 || os.Args[1] != "run" {
		fmt.Fprintf(os.Stderr, "usage: %s run\n(only run is supported for now)\n", os.Args[0])
		os.Exit(1)
	}
	if err := firecracker.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "kittyice: %v\n", err)
		os.Exit(1)
	}
}
