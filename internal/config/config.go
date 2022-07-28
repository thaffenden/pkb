package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/thaffenden/notes/internal/sentinel"
)

type (
	// Config represents the options defined in the config file.
	Config struct {
		Directory string     `json:"directory"`
		Editor    string     `json:"editor"`
		Templates []Template `json:"templates"`
	}

	// Template represents the config options for a custom template file.
	Template struct {
		Cmd  string `json:"cmd"`
		File string `json:"file"`
	}
)

// Load reads the users config file and returns the config struct.
func Load() (Config, error) {
	root := os.Getenv("XDG_CONFIG_HOME")

	if root == "" {
		root = fmt.Sprintf("%s/.config", os.Getenv("HOME"))
	}

	configFilePath := fmt.Sprintf("%s/notes/config.json", root)
	if _, err := os.Stat(configFilePath); err != nil {
		return Config{}, sentinel.Wrap(nil, ErrConfigNotFound)
	}

	contents, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, sentinel.WithMessagef(err, ErrReadingConfigFile, "error reading config file at %s", configFilePath)
	}

	var configContents Config
	if err := json.Unmarshal(contents, &configContents); err != nil {
		return Config{}, sentinel.Wrap(err, ErrUnmashallingJSON)
	}

	return configContents, nil
}

// FromContext returns the Config struct from the provided context with the
// correct type asserted from the default context interface{} return value.
func FromContext(ctx context.Context) (Config, error) {
	conf, ok := ctx.Value("config").(Config)
	if ok == false {
		return Config{}, errors.New("error getting config from context")
	}

	return conf, nil
}
