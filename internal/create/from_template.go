// Package create contains logic related to creating files.
package create

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"text/template"
	"time"

	"github.com/thaffenden/pkb/internal/config"
)

// TemplateRenderer holds the config required to render and save the template.
type TemplateRenderer struct {
	Config           config.Config
	Name             string
	SelectedTemplate config.Template
	Time             time.Time
	Templates        []config.Template
}

// templateVariables are the variables that are expanded when rendering the template.
type templateVariables struct {
	CustomDateFormat string
	Date             string
	Name             string
	Time             string
}

// NewTemplateRenderer creates a new instance of the TemplateRenderer.
func NewTemplateRenderer(conf config.Config, name string, templates []config.Template) TemplateRenderer {
	return TemplateRenderer{
		Config:    conf,
		Name:      name,
		Time:      time.Now(),
		Templates: templates,
	}
}

// CreateAndSaveFile creates the required file from the provided template
// and saves it in the correct output directory.
func (t TemplateRenderer) CreateAndSaveFile() (string, error) {
	outputPath := OutputPath(t.Config.Directory, t.Name, t.Templates)

	if err := createParentDirectories(outputPath); err != nil {
		return "", err
	}

	t.SelectedTemplate = t.Templates[len(t.Templates)-1]

	templateFile := filepath.Clean(
		filepath.Join(
			filepath.Dir(t.Config.FilePath),
			t.SelectedTemplate.File,
		),
	)
	contents, err := ioutil.ReadFile(templateFile)
	if err != nil {
		return "", err
	}

	file, err := os.Create(filepath.Clean(outputPath))
	if err != nil {
		return "", err
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Fatalf("error closing created file: %s", err)
		}
	}()

	if err := t.Render(string(contents), file); err != nil {
		return "", err
	}

	fmt.Printf("output file created: %s\n", outputPath)

	return outputPath, nil
}

// Render reads the template content and expands any variables.
func (t TemplateRenderer) Render(content string, writer io.Writer) error {
	now := t.Time

	config := templateVariables{
		Name: t.Name,
		Date: now.Format("2006-01-02"),
		Time: now.Format("15:04"),
	}

	// if SelectedTemplate has custom date add to struct.
	if t.SelectedTemplate.CustomDateFormat != "" {
		config.CustomDateFormat = now.Format(t.SelectedTemplate.CustomDateFormat)
	}

	tpl, err := template.New("template").Parse(content)
	if err != nil {
		return err
	}

	if err := tpl.Execute(writer, config); err != nil {
		return err
	}

	return nil
}

// OutputPath walks the sub template config to get build the full output path
// handling any nested sub templates.
func OutputPath(rootDir string, fileName string, templates []config.Template) string {
	output := []string{rootDir}

	for _, config := range templates {
		output = append(output, config.OutputDir)
	}

	output = append(output, SanitiseFileName(fileName))

	return filepath.Join(output...)
}

// createParentDirectories creates the parent directories for the rendered file
// if they don't already exist.
func createParentDirectories(outputPath string) error {
	parentDir := filepath.Dir(outputPath)

	if _, err := os.Stat(parentDir); os.IsNotExist(err) {
		if err := os.MkdirAll(parentDir, 0o750); err != nil {
			return fmt.Errorf("error creating parent directory %s for file %s", parentDir, outputPath)
		}
	}

	return nil
}
