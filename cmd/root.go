package cmd

import (
	"github.com/dhanushsr/clockin/clockin"
	"github.com/dhanushsr/clockin/utils"
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
	root.AddCommand(AddCommand(c))
	root.AddCommand(PrintCommand(c))
	return root
}

func PrintCommand(c *clockin.Config) *cobra.Command {
	var print = &cobra.Command{
		Use:   "show",
		Short: "Print data stored.",
		Run: func(cmd *cobra.Command, args []string) {
			err := clockin.PrintAllProjects(c)
			if err != nil {
				utils.PrintError(err)
			}
		},
	}
	return print
}
