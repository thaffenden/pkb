package cmd

import (
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/thaffenden/pkb/internal/config"
)

// NewCmdEdit creates the new command "edit" used to open your editor to edit existing notes.
func NewCmdEdit() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(ccmd *cobra.Command, args []string) error {
			conf, err := config.FromContext(ccmd.Context())
			if err != nil {
				return err
			}

			cmd := exec.Command(conf.Editor)
			cmd.Dir = conf.Directory
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr

			if err := cmd.Run(); err != nil {
				return errors.Wrapf(err, "error launching editor %s in directory %s", conf.Editor, conf.Directory)
			}

			return nil
		},
		Short: "open your notes directory in your specified editor",
		Use:   "edit",
	}

	return cmd
}
