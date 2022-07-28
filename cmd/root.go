package cmd

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/thaffenden/notes/cmd/new"
	"github.com/thaffenden/notes/internal/config"
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
	ctx := context.WithValue(context.Background(), "config", conf)
	return rootCmd.ExecuteContext(ctx)
}

func init() {
	rootCmd.AddCommand(new.NewCmd())
	rootCmd.AddCommand(NewCmdEdit())
}
