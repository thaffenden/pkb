package prompt

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
)

// SelectDirectory prompts the user to select a sub driectory in the provided
// parent. If the parent directory does not have any subdirectories this will
// error.
func SelectDirectory(parent string) (string, error) {
	// TODO: split out logic to get sub directories into separate func for tests.
	allFiles, err := os.ReadDir(parent)
	if err != nil {
		return "", err
	}

	subDirectories := []string{}

	for _, directory := range allFiles {
		if directory.IsDir() {
			subDirectories = append(subDirectories, directory.Name())
		}
	}

	if len(subDirectories) == 0 {
		return "", fmt.Errorf("no directories found in %s", parent)
	}

	fmt.Printf("found directories: %s", subDirectories)

	answer := struct {
		Selected string `survey:"directory"`
	}{}

	err = survey.Ask([]*survey.Question{
		{
			Name: "directory",
			Prompt: &survey.Select{
				Message: "select directory:",
				Options: subDirectories,
			},
		},
	}, &answer)
	if err != nil {
		return "", err
	}

	return answer.Selected, nil
}
