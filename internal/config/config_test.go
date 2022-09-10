package config_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thaffenden/pkb/internal/config"
	"github.com/thaffenden/pkb/internal/test"
)

func TestLoad(t *testing.T) {
	testCases := map[string]struct {
		conf          config.Config
		xdgConfigDir  string
		errorExpected require.ErrorAssertionFunc
	}{
		"errors when config file is not found": {
			conf:          config.Config{},
			xdgConfigDir:  "foo",
			errorExpected: test.IsSentinelError(config.ErrConfigNotFound),
		},
		"errors when config file is not valid json": {
			conf:          config.Config{},
			xdgConfigDir:  "invalid",
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
			xdgConfigDir:  "valid",
			errorExpected: require.NoError,
		},
		"tries home directory if XDG_CONFIG_HOME is not set": {
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
			xdgConfigDir:  "",
			errorExpected: require.NoError,
		},
	}

	for description, testCase := range testCases {
		tc := testCase

		t.Run(description, func(t *testing.T) {
			if tc.xdgConfigDir == "" {
				os.Setenv("HOME", filepath.FromSlash("testdata/home"))
				defer os.Unsetenv("HOME")
			}

			if tc.xdgConfigDir != "" {
				os.Setenv("XDG_CONFIG_HOME", filepath.FromSlash(fmt.Sprintf("testdata/xdg/%s", tc.xdgConfigDir)))
				defer os.Unsetenv("XDG_CONFIG_HOME")
			}

			conf, err := config.Load()
			tc.errorExpected(t, err)
			assert.Equal(t, tc.conf, conf)
		})
	}
}
