package prompt

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/thaffenden/pkb/internal/config"
	"golang.org/x/exp/maps"
)

// SelectorFunc is the type def for the selector func used in the TemplateSelector struct.
type SelectorFunc func(config.Templates) (config.Template, error)

// TemplateSelector is a utility struct to enable mocking of calls to the
// survey prompt for easier testability.
type TemplateSelector struct {
	SelectFunc SelectorFunc
}

// NewTemplateSelector creates a new instance of the TemplateSelector struct.
func NewTemplateSelector() TemplateSelector {
	return TemplateSelector{
		SelectFunc: SelectTemplate,
	}
}

// Recursive function to select template with nested sub templates and return them in a slice
func (t TemplateSelector) SelectTemplateWithSubTemplates(templates config.Templates, selectedTemplates []config.Template) ([]config.Template, error) {
	selected, err := t.SelectFunc(templates)
	if err != nil {
		return []config.Template{}, err
	}

	selectedTemplates = append(selectedTemplates, selected)

	if !selected.HasSubTemplates() {
		return selectedTemplates, nil
	}

	return t.SelectTemplateWithSubTemplates(selected.SubTemplates, selectedTemplates)
}

// SelectTemplate prompts the user to select a template from the ones defined in config.
func SelectTemplate(templates config.Templates) (config.Template, error) {
	answer := struct {
		Selected string `survey:"template"`
	}{}

	err := survey.Ask([]*survey.Question{
		{
			Name: "template",
			Prompt: &survey.Select{
				Message: "select template:",
				Options: maps.Keys(templates),
			},
		},
	}, &answer)
	if err != nil {
		return config.Template{}, err
	}

	selected, ok := templates[answer.Selected]
	if !ok {
		return config.Template{}, fmt.Errorf("no template named '%s' exists in config file", answer.Selected)
	}

	return selected, nil
}
