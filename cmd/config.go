package cmd

import (
	"github.com/dhanushsr/clockin/clockin"
	"github.com/spf13/cobra"
)

func ConfigCommand(c *clockin.Config) *cobra.Command {
	var edit bool
	config := &cobra.Command{
		Use:   "config",
		Short: "View/Edit config related to clockIn",
		Run: func(cmd *cobra.Command, args []string) {
			if edit {
				c.ReadAndSave()
			} else {
				c.Print()
			}
		},
	}
	config.Flags().BoolVarP(&edit, "edit", "e", false, "")
	return config
}
