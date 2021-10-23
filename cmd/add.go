package cmd

import (
	"fmt"

	"github.com/dhanushsr/clockin/clockin"
	"github.com/spf13/cobra"
)

var projectName string
var moduleName string

const PROJECT_KEY string = "project"
const MODULE_KEY string = "module"

func AddCommand(c *clockin.Config) *cobra.Command {
	var add = &cobra.Command{
		Use:   "add",
		Short: "Add a new project/sub-project to track.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(projectName) == 0 {
				fmt.Printf("Invalid Usage.\n\n")
				_ = cmd.Help()
			} else {
				clockin.AddProject(c, projectName, moduleName)
			}
		},
	}
	add.Flags().StringVarP(&projectName, PROJECT_KEY, "p", "", "Project to be added")
	add.Flags().StringVarP(&moduleName, MODULE_KEY, "m", "", "Module to be added")
	add.MarkFlagRequired(PROJECT_KEY)

	return add
}
