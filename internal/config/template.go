package config

import "golang.org/x/exp/maps"

// Templates represents the structure of a collection of templates.
type Templates map[string]Template

// Template represents the config options for a custom template file used
// for template and sub templates.
type Template struct {
	File         string    `json:"file"`
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
