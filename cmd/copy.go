package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thaffenden/pkb/internal/config"
)

// CreateCopy creates the new command "copy" used to select a note to copy
// to your system clipboard.
func CreateCopy() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := config.FromContext(ccmd.Context())
			if err != nil {
				return err
			}

			fmt.Print(conf)
			return nil
		},
		Short: "select a note and copy it's content to your system clipboard",
		Use:   "copy",
	}

	return cmd
}
