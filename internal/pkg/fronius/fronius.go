package fronius

import (
	"context"
	"fmt"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/config"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

type FroniusController struct {
	Config        *config.Fronius
	meterUrl      string
	flowUrl       string
	batteryUrl    string
	batteryAddUrl string
}

// NewController creates a new Controller
func NewController(config *config.Fronius) (*FroniusController, error) {

	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("Fronius Controller config invalid: %v", err)
	}

	meterUrl := fmt.Sprintf("%s/solar_api/v1/GetMeterRealtimeData.cgi?Scope=Device&DeviceId=0", config.BaseUrl)
	flowUrl := fmt.Sprintf("%s/solar_api/v1/GetPowerFlowRealtimeData.fcgi?Scope=Device&DeviceId=0", config.BaseUrl)
	batteryUrl := fmt.Sprintf("%s/solar_api/v1/GetStorageRealtimeData.cgi?Scope=Device&DeviceId=0", config.BaseUrl)
	batteryAddUrl := fmt.Sprintf("%s/components/BatteryManagementSystem/readable", config.BaseUrl)

	return &FroniusController{
		Config:        config,
		meterUrl:      meterUrl,
		batteryAddUrl: batteryAddUrl,
		flowUrl:       flowUrl,
		batteryUrl:    batteryUrl,
	}, nil
}

// Validate configuration
func validateConfig(conf *config.Fronius) error {
	var errs []string

	if conf.BaseUrl == "" {
		errs = append(errs, "Fronius BaseUrl not specified! Please set fronius.baseUrl in config file!")
	}

	if len(errs) > 0 {
		return fmt.Errorf("%s", strings.Join(errs, ", "))
	}

	return nil
}

// Empty function to satisfy interface
func (controller *FroniusController) Close(ctx context.Context) {
}

// Main run loop
func (controller *FroniusController) Run(ctx context.Context, metrics *metric.Metrics) error {
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
func (controller *FroniusController) getMetrics(ctx context.Context, metrics *metric.Metrics) error {

	if err := controller.getMeterData(ctx, metrics); err != nil {
		return err
	}

	if err := controller.getPowerFlow(ctx, metrics); err != nil {
		return err
	}

	if err := controller.getBatteryData(ctx, metrics); err != nil {
		return err
	}

	return nil
}
