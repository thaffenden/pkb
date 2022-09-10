package new

import (
	"fmt"

	"github.com/spf13/cobra"
)

// NewCmdNew creates the new command "new" used to create new notes.
func NewCmd() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			fmt.Println("new command")
			// check config for templates
			// if number of templates == 1, use that
			// else present template types as option
			// open new doc with selected template type
			return nil
		},
		Short: "create a new note",
		Use:   "new",
	}

	return cmd
}
