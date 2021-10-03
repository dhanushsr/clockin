package main

import (
	"fmt"
	"os"

	"github.com/dhanushsr/clockin/clockin"
	"github.com/dhanushsr/clockin/cmd"
	"github.com/dhanushsr/clockin/utils"
)

func main() {
	config, err := clockin.LoadOrInitialiseConfig()
	if err != nil {
		utils.PrintError(err)
	}
	if err := cmd.RootCommand(config).Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
