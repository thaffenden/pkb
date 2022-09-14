package create

import (
	"io"
	"text/template"
	"time"
)

type templateVariables struct {
	Name string
	Date string
	Time string
}

// RenderTemplate reads the template content and renders out the default variables.
func RenderTemplate(templateContent string, fileName string, writer io.Writer) error {
	now := time.Now()

	config := templateVariables{
		Name: fileName,
		Date: now.Format("2006-01-02"),
		Time: now.Format("15:04"),
	}

	tpl, err := template.New("template").Parse(templateContent)
	if err != nil {
		return err
	}

	if err := tpl.Execute(writer, config); err != nil {
		return err
	}

	return nil
}
