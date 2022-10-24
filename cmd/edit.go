package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thaffenden/pkb/internal/config"
	"github.com/thaffenden/pkb/internal/editor"
	"github.com/thaffenden/pkb/internal/flags"
	"github.com/thaffenden/pkb/internal/prompt"
)

// CreateEdit creates the new command "edit" used to open your editor to edit existing notes.
func CreateEdit() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := config.Get()
			if err != nil {
				return err
			}

			if flags.Pick {
				file, err := prompt.SelectExistingFile(conf.Directory)
				if err != nil {
					return err
				}

				if err := editor.OpenFile(conf.Editor, conf.Directory, file); err != nil {
					return err
				}

				return nil
			}

			if err := editor.Open(conf.Editor, conf.Directory); err != nil {
				return err
			}

			return nil
		},
		Short: "open your notes directory in your specified editor",
		Use:   "edit",
	}

	cmd.Flags().BoolVar(&flags.Pick, "pick", false, "select the file you want to open before opening your editor")
	return cmd
}
