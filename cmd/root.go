package cmd

import (
	"github.com/dhanushsr/clockin/clockin"
	"github.com/spf13/cobra"
)

func RootCommand(c *clockin.Config) *cobra.Command {
	var root = &cobra.Command{
		Use:   "clockin",
		Short: "ClockIn provides a simple CLI for tracking work progress.",
		Long:  `A fast and flexible CLI implemented in GO to help track work progress and time spent.`,
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}

	root.AddCommand(ConfigCommand(c))
	root.AddCommand(AddCommand())
	return root
}
