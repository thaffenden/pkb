package cmd

import (
	"github.com/spf13/cobra"
	"github.com/thaffenden/pkb/internal/config"
	"github.com/thaffenden/pkb/internal/create"
	"github.com/thaffenden/pkb/internal/editor"
	"github.com/thaffenden/pkb/internal/flags"
	"github.com/thaffenden/pkb/internal/prompt"
)

// CreateNew creates the new command "new" used to create new notes.
func CreateNew() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := config.Get()
			if err != nil {
				return err
			}

			selected := []config.Template{}
			selector := prompt.NewTemplateSelector()

			selected, err = selector.SelectTemplateWithSubTemplates(conf.Templates, selected)
			if err != nil {
				return err
			}

			renderer := create.NewTemplateRenderer(conf, selected)
			createdFile, err := renderer.CreateAndSaveFile()
			if err != nil {
				return err
			}

			if !flags.NoEdit {
				if err := editor.OpenFile(conf.Editor, conf.Directory, createdFile); err != nil {
					return err
				}
			}

			return nil
		},
		Short: "Create a new note. Select from templates defined in your config file.",
		Use:   "new",
	}

	cmd.Flags().BoolVar(&flags.NoEdit, "no-edit", false, "don't open the file in your editor after creating")
	return cmd
}
