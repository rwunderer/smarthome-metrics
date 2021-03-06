package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type Fronius struct {
	BaseUrl        string `yaml:"baseUrl"`
	Prefix         string `yaml:"prefix"`
	MeterOffsetIn  int    `yaml:"meterOffsetIn"`
	MeterOffsetOut int    `yaml:"meterOffsetOut"`
}

type Ecotouch struct {
	BaseUrl  string `yaml:"baseUrl"`
	Prefix   string `yaml:"prefix"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Graphite struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Prefix   string `yaml:"prefix"`
	Protocol string `yaml:"protocol"`
}

type Config struct {
	Fronius  Fronius  `yaml:"fronius"`
	Ecotouch Ecotouch `yaml:"ecotouch"`
	Graphite Graphite `yaml:"graphite"`
}

func (conf *Config) ReadFile(inFile string) error {
	content, err := ioutil.ReadFile(inFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	if err := yaml.Unmarshal(content, &conf); err != nil {
		return fmt.Errorf("failed to unmarshal config file: %v", err)
	}

	return nil
}
