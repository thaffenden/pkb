package config

// Template represents the config options for a custom template file used
// for template and sub templates.
type Template struct {
	Type         string     `json:"type"`
	File         string     `json:"file"`
	SubTemplates []Template `json:"sub_templates,omitempty"`
}

// HasSubTemplates checks if the Template struct has sub templates.
// SubTemplates are offered at the pointing of using a template to give a
// specific version of a file.
func (t Template) HasSubTemplates() bool {
	switch len(t.SubTemplates) {
	case 0:
		return false
	// handle an empty template inside sub templates
	case 1:
		return !t.SubTemplates[0].IsEmpty()
	default:
		return true
	}
}

// IsEmpty checks if the config for the template object is empty.
func (t Template) IsEmpty() bool {
	return t.Type == ""
}
