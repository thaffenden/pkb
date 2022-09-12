package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thaffenden/pkb/internal/config"
	"github.com/thaffenden/pkb/internal/editor"
)

// CmdEdit creates the new command "edit" used to open your editor to edit existing notes.
func CmdEdit() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := config.FromContext(ccmd.Context())
			if err != nil {
				return err
			}

			err = editor.Open(conf.Editor, conf.Directory)
			if err != nil {
				return err
			}

			return nil
		},
		Short: "open your notes directory in your specified editor",
		Use:   "edit",
	}

	return cmd
}
