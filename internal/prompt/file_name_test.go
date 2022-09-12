package prompt_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/thaffenden/pkb/internal/prompt"
)

func TestSanitiseFileName(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		input    string
		expected string
	}{
		"returns expected string for single word": {
			input:    "foo",
			expected: "foo.md",
		},
		"converts spaces in input to hyphens": {
			input:    "foo bar baz ",
			expected: "foo-bar-baz.md",
		},
		"removes special characters that aren't file path safe": {
			input:    "new&&test_who*dis+",
			expected: "new-test_who-dis.md",
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			actual := prompt.SanitiseFileName(tc.input)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
