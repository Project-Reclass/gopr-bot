package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const defaultFileName = "config.yml"

// Config holds all the configurations for gopr
type Config struct {
	Slack  SlackConfig  `yaml:"slack"`
	Github GithubConfig `yaml:"github"`
}

// NewFromDefaultFile creates a new config structure from the default config `config.yml`
func NewFromDefaultFile() Config {
	return New(defaultFileName)
}

// New creates a new config structure from the passed filename
func New(filename string) Config {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	configBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	config := Config{}
	if err = yaml.Unmarshal(configBytes, &config); err != nil {
		panic(err)
	}
	return config
}
