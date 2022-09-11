package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thaffenden/pkb/internal/config"
)

// CmdNew creates the new command "new" used to create new notes.
func CmdNew() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := config.FromContext(ccmd.Context())
			if err != nil {
				return err
			}

			fmt.Println("new command")

			if len(conf.Templates) == 1 {
				fmt.Printf("using template: %+v", conf.Templates[0])
			}

			// open survey picker to selet template type
			// TODO: SUB_TEMPLATES if node has sub templates prompt for them

			// get doc name (flag or prompt)

			// open new doc with selected template type
			return nil
		},
		Short: "create a new note",
		Use:   "new",
	}

	return cmd
}