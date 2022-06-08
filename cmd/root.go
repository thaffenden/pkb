package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thaffenden/notes/pkg/config"
)

var rootCmd = &cobra.Command{
	RunE: func(ccmd *cobra.Command, args []string) error {
		return nil
	},
	Short: "manage notes in markdown files",
	Use:   "notes",
}

// Execute executes the root command.
func Execute(conf config.Config) error {
	return rootCmd.Execute()
}
