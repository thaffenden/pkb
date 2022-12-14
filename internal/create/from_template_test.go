package create_test

import (
	"bytes"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thaffenden/pkb/internal/config"
	"github.com/thaffenden/pkb/internal/create"
)

func TestGetFileName(t *testing.T) {
	t.Parallel()

	testTime, _ := time.Parse(time.RFC3339, "2022-09-19T16:20:00Z")

	testCases := map[string]struct {
		renderer    create.TemplateRenderer
		expected    string
		assertError require.ErrorAssertionFunc
	}{
		"uses prompt when no value in config": {
			renderer: create.TemplateRenderer{
				NamePrompt: func() (string, error) {
					return "prompted for this string", nil
				},
				Templates: []config.Template{{}},
			},
			expected:    "prompted for this string",
			assertError: require.NoError,
		},
		"combines values when mutiple provided": {
			renderer: create.TemplateRenderer{
				NamePrompt: func() (string, error) {
					return "wow this is great", nil
				},
				SelectedTemplate: config.Template{
					NameFormat: "{{.Date}}-{{.Prompt}}-{{.Week}}-{{.Year}}-foo",
				},
				Time: testTime,
			},
			expected:    "2022-09-19-wow this is great-38-2022-foo",
			assertError: require.NoError,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.renderer.GetFileName()
			tc.assertError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestOutputPath(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		templateRenderer create.TemplateRenderer
		assertError      require.ErrorAssertionFunc
		expected         string
	}{
		"returns path for single template": {
			templateRenderer: create.TemplateRenderer{
				Config: config.Config{
					Directory: "/home/username/notes",
				},
				Name: "simple.md",
				Templates: []config.Template{
					{
						File:      "magic.tpl.md",
						OutputDir: "magic",
					},
				},
			},
			assertError: require.NoError,
			expected:    "/home/username/notes/magic/simple.md",
		},
		"creates full nested dir path when there are subtemplates": {
			templateRenderer: create.TemplateRenderer{
				Config: config.Config{
					Directory: "/home/username/notes",
				},
				Name: "nested-example.md",
				Templates: []config.Template{
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
			},
			assertError: require.NoError,
			expected:    "/home/username/notes/foo/bar/wow/nested-example.md",
		},
		"prompts user for directory input when specified in template config": {
			templateRenderer: create.TemplateRenderer{
				Config: config.Config{
					Directory: "/home/username/notes",
				},
				Name: "simple.md",
				DirectoryPrompt: func() (string, error) {
					return "foo/dir", nil
				},
				Templates: []config.Template{
					{
						File:      "magic.tpl.md",
						OutputDir: "{{.Prompt}}",
					},
				},
			},
			assertError: require.NoError,
			expected:    "/home/username/notes/foo/dir/simple.md",
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual, err := tc.templateRenderer.OutputPath()
			tc.assertError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestRender(t *testing.T) {
	t.Parallel()

	testTime, _ := time.Parse(time.RFC3339, "2022-09-19T16:20:00Z")

	testCases := map[string]struct {
		renderer        create.TemplateRenderer
		templateContent string
		expected        string
		assertError     require.ErrorAssertionFunc
	}{
		"expands expected variables": {
			renderer: create.TemplateRenderer{
				Config: config.Config{
					Templates: map[string]config.Template{},
				},
				Name: "example doc",
				SelectedTemplate: config.Template{
					CustomDateFormat: "Monday 2nd January",
				},
				Time: testTime,
			},
			templateContent: "{{.Date}}\n{{.Name}}\n{{.Time}}\n{{.CustomDateFormat}}\n{{.Week}}\n{{.Year}}",
			expected:        "2022-09-19\nexample doc\n16:20\nMonday 19th September\n38\n2022",
			assertError:     require.NoError,
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			var actual bytes.Buffer
			err := tc.renderer.Render(tc.templateContent, &actual)
			tc.assertError(t, err)

			assert.Equal(t, tc.expected, actual.String())
		})
	}
}
