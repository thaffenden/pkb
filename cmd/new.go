package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thaffenden/pkb/internal/config"
	"github.com/thaffenden/pkb/internal/create"
	"github.com/thaffenden/pkb/internal/prompt"
)

// CmdNew creates the new command "new" used to create new notes.
func CmdNew() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := config.FromContext(ccmd.Context())
			if err != nil {
				return err
			}

			selected := []config.Template{}
			selector := prompt.NewTemplateSelector()

			selected, err = selector.SelectTemplateWithSubTemplates(conf.Templates, selected)
			if err != nil {
				return err
			}

			fileName, err := prompt.EnterFileName()
			if err != nil {
				return err
			}

			err = create.FileFromTemplate(conf, fileName, selected)
			if err != nil {
				return err
			}

			return nil
		},
		Short: "create a new note",
		Use:   "new",
	}

	return cmd
}
