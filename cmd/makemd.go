package main

import (
	"fmt"
	"os"

	"github.com/detouri/makemd/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, "makemd: ", err)
		os.Exit(1)
	}
}
