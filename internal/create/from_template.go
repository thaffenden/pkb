package create

import (
	"path/filepath"

	"github.com/thaffenden/pkb/internal/config"
)

// FileFromTemplate creates the document as per the template config.
func FileFromTemplate(name string, templates []config.Template) error {
	// check if parent dirs exist
	// if not create them
	// add id to name
	// create file

	// template to use will always be last in slice
	return nil
}

// OutputPath walks the sub template config to get build the full output path
// handling any nested sub templates.
func OutputPath(fileName string, templates []config.Template) (string, error) {
	output := []string{}

	for _, config := range templates {
		output = append(output, config.OutputDir)
	}

	output = append(output, fileName)

	return filepath.Join(output...), nil
}
