package config

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"golang.org/x/exp/maps"
)

// Templates represents the structure of a collection of templates.
type Templates map[string]Template

// Template represents the config options for a custom template file used
// for template and sub templates.
type Template struct {
	File         string    `json:"file"`
	OutputDir    string    `json:"output_dir"`
	SubTemplates Templates `json:"sub_templates,omitempty"`
}

// HasSubTemplates checks if the Template struct has sub templates.
// SubTemplates are offered at the pointing of using a template to give a
// specific version of a file.
func (t Template) HasSubTemplates() bool {
	switch len(maps.Keys(t.SubTemplates)) {
	case 0:
		return false
	// handle an empty template inside sub templates
	case 1:
		return !maps.Values(t.SubTemplates)[0].isEmpty()
	default:
		return true
	}
}

// isEmpty checks if the config for the template object is empty.
func (t Template) isEmpty() bool {
	return t.File == ""
}

// Select prompts the user to select a template from the ones defined in config.
func (t Templates) Select() (Template, error) {
	answer := struct {
		Selected string `survey:"template"`
	}{}

	err := survey.Ask([]*survey.Question{
		{
			Name: "template",
			Prompt: &survey.Select{
				Message: "select a template:",
				Options: maps.Keys(t),
			},
		},
	}, &answer)
	if err != nil {
		return Template{}, err
	}

	selected, ok := t[answer.Selected]
	if !ok {
		return Template{}, fmt.Errorf("no template named '%s' exists in config file", answer.Selected)
	}

	return selected, nil
}
