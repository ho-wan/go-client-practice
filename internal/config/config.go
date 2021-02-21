package config

import (
	"io/ioutil"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

// Config contains config for the project
type Config struct {
	Github Github `yaml:"github"`
}

// Github contains config for github
type Github struct {
	AccessToken string `yaml:"accessToken"`
}

// LoadConfig loads the config yaml file
func LoadConfig(filename string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read file")
	}

	var cfg Config
	if err = yaml.Unmarshal(yamlFile, &cfg); err != nil {
		return nil, errors.Wrap(err, "failed to unmarshal yaml")
	}
	return &cfg, nil
}
