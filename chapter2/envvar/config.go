package envvar

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
// 	"github.com/pkg/errors"
)

// LoadConfig will load files optionally from the json file stored
// at path, then will override those values based on the envconfig
// struct tags. The envPrefix is how we prefix our environment
// variables.
func LoadConfig(path, envPrefix string, config interface{}) error {
	if path != "" {
		err := LoadFile(path, config)
		if err != nil {
			return fmt.Errorf("error loading config from file: %w", err) 
		}
	}
	if err := envconfig.Process(envPrefix, config); err != nil {
		return fmt.Errorf("error loading config from env: %w", err)
	}
	return nil
}

// LoadFile unmarshalls a json file into a config struct
func LoadFile(path string, config interface{}) error {
	configFile, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)
	if err = decoder.Decode(config); err != nil {
		return fmt.Errorf("failed to decode config file: %w", err)
	}
	return nil
}
