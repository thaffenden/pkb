package create_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thaffenden/pkb/internal/config"
	"github.com/thaffenden/pkb/internal/create"
)

func TestOutputPath(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		name      string
		rootDir   string
		templates []config.Template
		expected  string
	}{
		"returns path for single template": {
			name:    "simple.md",
			rootDir: "/home/username/notes",
			templates: []config.Template{
				{
					File:      "magic.tpl.md",
					OutputDir: "magic",
				},
			},
			expected: "/home/username/notes/magic/simple.md",
		},
		"creates full nested dir path when there are subtemplates": {
			name:    "nested-example.md",
			rootDir: "/home/username/notes",
			templates: []config.Template{
				{
					File:      "foo.tpl.md",
					OutputDir: "foo",
				},
				{
					File:      "bar.tpl.md",
					OutputDir: "bar",
				},
				{
					File:      "wow.tpl.md",
					OutputDir: "wow",
				},
			},
			expected: "/home/username/notes/foo/bar/wow/nested-example.md",
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := create.OutputPath(tc.rootDir, tc.name, tc.templates)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
