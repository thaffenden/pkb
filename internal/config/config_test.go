package config_test

import (
	"context"
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
				FilePath:  "testdata/xdg/valid/pkb/config.json",
				Templates: config.Templates{
					"foo": {
						File:      "bar.tpl.md",
						OutputDir: "bar",
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
				FilePath:  "testdata/home/.config/pkb/config.json",
				Templates: config.Templates{
					"foo": {
						File:      "bar.tpl.md",
						OutputDir: "",
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

func TestFromContext(t *testing.T) {
	t.Parallel()

	testCases := map[string]struct {
		ctxFunc       func() context.Context
		errorExpected require.ErrorAssertionFunc
		expected      config.Config
	}{
		"returns config struct when valid one is set in context": {
			ctxFunc: func() context.Context {
				return context.WithValue(
					context.Background(),
					"config",
					config.Config{Editor: "nvim"},
				)
			},
			errorExpected: require.NoError,
			expected: config.Config{
				Editor: "nvim",
			},
		},
		"returns error when no key called config exists on context": {
			ctxFunc: func() context.Context {
				return context.WithValue(
					context.Background(),
					"foo",
					[]string{},
				)
			},
			errorExpected: require.Error,
			expected:      config.Config{},
		},
		"returns error when config key is not of type config.Config": {
			ctxFunc: func() context.Context {
				return context.WithValue(
					context.Background(),
					"config",
					[]string{},
				)
			},
			errorExpected: require.Error,
			expected:      config.Config{},
		},
	}

	for name, testCase := range testCases {
		tc := testCase

		t.Run(name, func(t *testing.T) {
			t.Parallel()

			ctx := tc.ctxFunc()
			actual, err := config.FromContext(ctx)
			tc.errorExpected(t, err)

			assert.Equal(t, tc.expected, actual)
		})
	}
}
