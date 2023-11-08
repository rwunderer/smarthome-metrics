package config

import (
	"fmt"
	"io/ioutil"

	yaml "gopkg.in/yaml.v3"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/ecotouch"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/fronius"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/graphite"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/watertemp"
)

type Config struct {
	ActiveAppliances  []string         `yaml:"appliances"`
	ActiveControllers []string         `yaml:"controllers"`
	Fronius           fronius.Config   `yaml:"fronius"`
	Ecotouch          ecotouch.Config  `yaml:"ecotouch"`
	Graphite          graphite.Config  `yaml:"graphite"`
	WaterTemperature  watertemp.Config `yaml:"watertemp"`
}

func (conf *Config) ReadFile(inFile string) error {
	content, err := ioutil.ReadFile(inFile)
	if err != nil {
		return fmt.Errorf("failed to read config file: %v", err)
	}

	conf.ActiveAppliances = []string{"fronius", "ecotouch"} // default value
	conf.ActiveControllers = []string{"ecotouch"}           // default value
	conf.Ecotouch = ecotouch.GetDefaultConfig()
	conf.Fronius = fronius.GetDefaultConfig()
	conf.Graphite = graphite.GetDefaultConfig()
	conf.WaterTemperature = watertemp.GetDefaultConfig()

	if err := yaml.Unmarshal(content, &conf); err != nil {
		return fmt.Errorf("failed to unmarshal config file: %v", err)
	}

	return nil
}
