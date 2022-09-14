// Package prompt contains logic for prompts and user interactions with the CLI.
package prompt

import (
	"github.com/AlecAivazis/survey/v2"
)

// EnterFileName prompts the user to enter the name of the file they are going
// to save a template as, and returns a sanitised.
func EnterFileName() (string, error) {
	name := ""
	prompt := &survey.Input{
		Message: "enter file name:",
	}
	err := survey.AskOne(prompt, &name)
	if err != nil {
		return "", err
	}

	return name, nil
}
