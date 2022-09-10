package editor

import (
	"os"
	"os/exec"

	"github.com/pkg/errors"
)

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
