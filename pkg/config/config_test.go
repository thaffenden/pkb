package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/thaffenden/notes/pkg/config"
	"github.com/thaffenden/notes/pkg/test"
)

func TestLoad(t *testing.T) {
	testCases := map[string]struct {
		conf          config.Config
		errorExpected require.ErrorAssertionFunc
		useXDG        bool
	}{
		"ErrorsWhenJSONUnmarshalFails": {
			conf:          config.Config{},
			errorExpected: test.IsSentinelError(config.ErrUnmashallingJSON),
			useXDG:        true,
		},
	}

	for description, testCase := range testCases {
		tc := testCase

		t.Run(description, func(t *testing.T) {
			t.Parallel()

			conf, err := config.Load()
			tc.errorExpected(t, err)
			assert.Equal(t, tc.conf, conf)
		})
	}
}
