package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
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
	configFilePath := ""
	xdg := os.Getenv("XDG_CONFIG_HOME")

	if xdg != "" {
		configFilePath = fmt.Sprintf("%s/notes/config.json", xdg)
	}

	contents, err := os.ReadFile(configFilePath)
	if err != nil {
		return Config{}, errors.New(fmt.Sprintf("error reading config file at %s", configFilePath))
	}

	var configContents Config
	err = json.Unmarshal(contents, &configContents)
	if err != nil {
		return Config{}, errors.New("error unmarshalling config file")
	}

	return configContents, nil
}
