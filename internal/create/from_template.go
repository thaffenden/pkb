// Package create contains logic related to creating files.
package create

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"github.com/thaffenden/pkb/internal/config"
	"github.com/thaffenden/pkb/internal/date"
	"github.com/thaffenden/pkb/internal/prompt"
)

// TemplateRenderer holds the config required to render and save the template.
type TemplateRenderer struct {
	Config           config.Config
	DirectoryPrompt  func() (string, error)
	Name             string
	NamePrompt       func() (string, error)
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
func NewTemplateRenderer(conf config.Config, templates []config.Template) TemplateRenderer {
	return TemplateRenderer{
		Config:          conf,
		DirectoryPrompt: prompt.EnterDirectory,
		NamePrompt:      prompt.EnterFileName,
		Time:            time.Now(),
		Templates:       templates,
	}
}

// CreateAndSaveFile creates the required file from the provided template
// and saves it in the correct output directory.
func (t TemplateRenderer) CreateAndSaveFile() (string, error) {
	t.SelectedTemplate = t.Templates[len(t.Templates)-1]

	fileName, err := t.GetFileName()
	if err != nil {
		return "", err
	}

	t.Name = fileName

	outputPath := t.OutputPath()

	if err := createParentDirectories(outputPath); err != nil {
		return "", err
	}

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

// GetFileName either prompts the user for input or uses one of the supported
// name specifiers to automatically set the date.
func (t TemplateRenderer) GetFileName() (string, error) {
	if t.SelectedTemplate.NameFormat == "" {
		return t.NamePrompt()
	}

	outputString := t.SelectedTemplate.NameFormat

	if strings.Contains(outputString, "DATE") {
		outputString = strings.ReplaceAll(outputString, "DATE", t.Time.Format("2006-01-02"))
	}

	if strings.Contains(outputString, "PROMPT") {
		promptString, err := t.NamePrompt()
		if err != nil {
			return "", err
		}

		outputString = strings.ReplaceAll(outputString, "PROMPT", promptString)
	}

	return outputString, nil
}

// Render reads the template content and expands any variables.
func (t TemplateRenderer) Render(content string, writer io.Writer) error {
	now := t.Time

	config := templateVariables{
		Name: t.Name,
		Date: now.Format("2006-01-02"),
		Time: now.Format("15:04"),
	}

	// If a custom date format is specified on the template config run it through
	// the date utils to better support human friendly output.
	if t.SelectedTemplate.CustomDateFormat != "" {
		config.CustomDateFormat = now.Format(t.SelectedTemplate.CustomDateFormat)

		if date.IncludesSuffixFormat(config.CustomDateFormat) {
			fixedDate, err := date.ReplaceSuffixFormatter(config.CustomDateFormat)
			if err != nil {
				return err
			}

			config.CustomDateFormat = fixedDate
		}
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
// handling any nested sub templates and prompts or selections for output
// directories.
func (t TemplateRenderer) OutputPath() string {
	output := []string{t.Config.Directory}

	for _, config := range t.Templates {
		output = append(output, config.OutputDir)
	}

	output = append(output, SanitiseFileName(t.Name))

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
