package nrgkick

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

type Config struct {
	BaseUrl string `yaml:"baseUrl"`
	Prefix  string `yaml:"prefix"`
}

type NrgkickController struct {
	Config     *Config
	infoUrl    string
	controlUrl string
	valuesUrl  string
}

// Get Default configuration
func GetDefaultConfig() Config {
	return Config{}
}

// NewController creates a new Controller
func NewController(config *Config) (*NrgkickController, error) {

	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("Nrgkick Controller config invalid: %v", err)
	}

	log.Infof("Nrgkick base url is %v", config.BaseUrl)

	infoUrl := fmt.Sprintf("%s/info?raw", config.BaseUrl)
	controlUrl := fmt.Sprintf("%s/control", config.BaseUrl)
	valuesUrl := fmt.Sprintf("%s/values?raw", config.BaseUrl)

	return &NrgkickController{
		Config:     config,
		infoUrl:    infoUrl,
		controlUrl: controlUrl,
		valuesUrl:  valuesUrl,
	}, nil
}

// Validate configuration
func validateConfig(conf *Config) error {
	var errs []string

	if conf.BaseUrl == "" {
		errs = append(errs, "Nrgkick BaseUrl not specified! Please set nrgkick.baseUrl in config file!")
	}

	if len(errs) > 0 {
		return fmt.Errorf("%s", strings.Join(errs, ", "))
	}

	return nil
}

// Empty function to satisfy interface
func (controller *NrgkickController) Close(ctx context.Context) {
}

// Empty function to satisfy writer interface
func (controller *NrgkickController) SetValue(ctx context.Context, fieldName string, desiredValue float64) error {
	return nil
}

// Main run loop
func (controller *NrgkickController) Run(ctx context.Context, metrics *metric.Metrics) error {
	var err error

	if err = controller.getMetrics(ctx, metrics); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			log.Debugf("Context Done. Shutting down")
			return nil
		case <-time.After(30 * time.Second):
			if err = controller.getMetrics(ctx, metrics); err != nil {
				return err
			}
		}
	}
}

// Retrieve all configured metrics
func (controller *NrgkickController) getMetrics(ctx context.Context, metrics *metric.Metrics) error {

	if err := controller.getInfoData(ctx, metrics); err != nil {
		return err
	}

	if err := controller.getControlData(ctx, metrics); err != nil {
		return err
	}

	if err := controller.getValuesData(ctx, metrics); err != nil {
		return err
	}

	return nil
}
