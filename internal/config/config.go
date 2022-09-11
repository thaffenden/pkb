package config

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/thaffenden/pkb/internal/sentinel"
)

type (
	// Config represents the options defined in the config file.
	Config struct {
		Directory string     `json:"directory"`
		Editor    string     `json:"editor"`
		Templates []Template `json:"templates"`
	}
)

// Load reads the users config file and returns the config struct.
func Load() (Config, error) {
	root := os.Getenv("XDG_CONFIG_HOME")

	if root == "" {
		root = fmt.Sprintf("%s/.config", os.Getenv("HOME"))
	}

	configFilePath := fmt.Sprintf("%s/pkb/config.json", root)
	if _, err := os.Stat(configFilePath); err != nil {
		return Config{}, sentinel.Wrap(nil, ErrConfigNotFound)
	}

	contents, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, sentinel.Wrap(err, ErrReadingConfigFile)
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
