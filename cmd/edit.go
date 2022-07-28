package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/thaffenden/notes/internal/config"
)

// NewCmdEdit creates the new command "edit" used to open your editor to edit existing notes.
func NewCmdEdit() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := config.FromContext(ccmd.Context())
			if err != nil {
				return err
			}

			fmt.Println(conf)
			// cmd := exec.Command(conf.Editor, conf.Directory)
			// if err := cmd.Run(); err != nil {
			// 	return err
			// }

			return nil
		},
		Short: "open your notes directory in your specified editor",
		Use:   "edit",
	}

	return cmd
}
