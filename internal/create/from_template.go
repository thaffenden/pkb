// Package create contains logic related to creating files.
package create

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/thaffenden/pkb/internal/config"
)

// FileFromTemplate creates the document as per the template config.
func FileFromTemplate(conf config.Config, name string, templates []config.Template) (string, error) {
	outputPath := OutputPath(conf.Directory, name, templates)

	parentDir := filepath.Dir(outputPath)

	// create parent directory if it does not already exist.
	if _, err := os.Stat(parentDir); os.IsNotExist(err) {
		if err := os.MkdirAll(parentDir, 0o750); err != nil {
			return "", fmt.Errorf("error creating file %s", outputPath)
		}
	}

	templateFile := filepath.Clean(filepath.Join(filepath.Dir(conf.FilePath), templates[len(templates)-1].File))
	contents, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return "", err
	}

	if err := ioutil.WriteFile(outputPath, contents, 0o600); err != nil {
		return "", err
	}

	fmt.Printf("output file created: %s\n", outputPath)

	return outputPath, nil
}

// OutputPath walks the sub template config to get build the full output path
// handling any nested sub templates.
func OutputPath(rootDir string, fileName string, templates []config.Template) string {
	output := []string{rootDir}

	for _, config := range templates {
		output = append(output, config.OutputDir)
	}

	output = append(output, fileName)

	return filepath.Join(output...)
}
