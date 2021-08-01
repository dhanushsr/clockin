package main

import (
	"fmt"
	"os"

	"github.com/dhanushsr/clockin/cmd"
)

func main() {
	if err := cmd.RootCommand().Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
