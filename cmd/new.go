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

			selected, err := conf.Templates.Select()
			if err != nil {
				return err
			}

			var subTemplate config.Template
			if selected.HasSubTemplates() {
				subTemplate, err = selected.SubTemplates.Select()
				if err != nil {
					return err
				}
			}

			fmt.Printf("%+v\n", selected)
			fmt.Printf("%+v\n", subTemplate)
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
