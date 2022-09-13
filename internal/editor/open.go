// Package editor contains logic for sending commands to or interacting with
// the editor the user defined in config.
package editor

import (
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

// Open opens the provided editor in the specified directory.
func Open(editorCmd string, directory string) error {
	cmd := exec.Command(editorCmd)
	cmd.Dir = directory
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return errors.Wrapf(err, "error launching editor %s in directory %s", editorCmd, directory)
	}

	return nil
}
