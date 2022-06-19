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
			return nil
		},
		Short: "create a new note",
		Use:   "new",
	}

	return cmd
}
