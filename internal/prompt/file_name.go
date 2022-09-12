package prompt

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

// EnterFileName prompts the user to enter the name of the file they are going
// to save a template as, and returns a sanitised
func EnterFileName() (string, error) {
	name := ""
	prompt := &survey.Input{
		Message: "enter file name:",
	}
	err := survey.AskOne(prompt, &name)
	if err != nil {
		return "", err
	}

	return SanitiseFileName(name), nil
}

// SanitiseFileName removes any spaces or special characters so the format
// is valid to use as a file name
func SanitiseFileName(name string) string {
	baseName := strings.Trim(name, " ")

	separators := regexp.MustCompile(`[ &=+:*]`)
	baseName = separators.ReplaceAllString(baseName, "-")

	baseNameSeparators := regexp.MustCompile(`[./]`)
	baseName = baseNameSeparators.ReplaceAllString(baseName, "-")

	doubleSeparators := regexp.MustCompile(`--`)
	baseName = doubleSeparators.ReplaceAllString(baseName, "-")

	baseName = strings.Trim(baseName, "-")

	if strings.HasSuffix(baseName, ".md") {
		return baseName
	}

	return fmt.Sprintf("%s.md", baseName)
}
