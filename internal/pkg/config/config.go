package config

import (
	"fmt"
	"io/ioutil"
	"strings"

	yaml "gopkg.in/yaml.v2"
)

type Fronius struct {
	BaseUrl        string `yaml:"baseUrl"`
	MeterOffsetIn  int    `yaml:"meterOffsetIn"`
	MeterOffsetOut int    `yaml:"meterOffsetOut"`
}

type Config struct {
	Fronius Fronius `yaml:"fronius"`
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

func (conf *Config) Validate() error {
	var errs []string

	if conf.Fronius.BaseUrl == "" {
		errs = append(errs, "Fronius BaseUrl not specified! Please set fronius.baseUrl in config file!")
	}

	if len(errs) > 0 {
		return fmt.Errorf("%s", strings.Join(errs, ", "))
	}

	return nil
}
