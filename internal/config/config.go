package config

import (
	"os"

	"gopkg.in/yaml.v3"
)

const configFile = "config.yaml"

type Config struct {
	Token  string `yaml:"token"`
	DBName string `yaml:"db_name"`
}

func New() (*Config, error) {
	config := &Config{}
	rawYaml, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(rawYaml, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
