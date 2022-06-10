package config_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thaffenden/notes/pkg/config"
	"github.com/thaffenden/notes/pkg/test"
)

func TestLoad(t *testing.T) {
	testCases := map[string]struct {
		conf          config.Config
		configDir     string
		errorExpected require.ErrorAssertionFunc
	}{
		"errors when config file is not found": {
			conf:          config.Config{},
			configDir:     "",
			errorExpected: test.IsSentinelError(config.ErrConfigNotFound),
		},
		"errors when config file is not valid json": {
			conf:          config.Config{},
			configDir:     "invalid",
			errorExpected: test.IsSentinelError(config.ErrUnmashallingJSON),
		},
		"returns config struct when valid file exists": {
			conf: config.Config{
				Directory: "/home/username/notes",
				Editor:    "nvim",
				Templates: []config.Template{
					{
						Cmd:  "foo",
						File: "bar.tpl.md",
					},
				},
			},
			configDir:     "valid",
			errorExpected: require.NoError,
		},
	}

	for description, testCase := range testCases {
		tc := testCase

		t.Run(description, func(t *testing.T) {
			if tc.configDir != "" {
				os.Setenv("XDG_CONFIG_HOME", filepath.FromSlash(fmt.Sprintf("testdata/%s", tc.configDir)))
			}

			conf, err := config.Load()
			tc.errorExpected(t, err)
			assert.Equal(t, tc.conf, conf)
		})
	}
}
