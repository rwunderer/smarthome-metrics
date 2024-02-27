package watertemp

import (
	"context"
	"math"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/controllers"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

type gridState int

const (
	nonproducing gridState = iota
	producing
)

type Config struct {
	ProduceThreshold  float64 `yaml:"produceThreshold"`
	ProduceMinTime    float64 `yaml:"produceMinTime"`
	ProducingTemp     float64 `yaml:"producingTemp"`
	NormalTemp        float64 `yaml:"normalTemp"`
	TempChangeMinTime float64 `yaml:"tempChangeMinTime"`
}

type WaterTemperatureController struct {
	ecotouch controllers.SmarthomeAppliance
	config   *Config

	gridState      gridState
	gridStateSince time.Time
	tempWanted     float64
	lastTempChange time.Time
}

// Get Default configuration
func GetDefaultConfig() Config {
	return Config{
		ProduceThreshold:  -2000,
		ProduceMinTime:    10,
		ProducingTemp:     55,
		NormalTemp:        48,
		TempChangeMinTime: 30,
	}
}

// Create a new Controller
func NewController(config *Config, ecotouch controllers.SmarthomeAppliance) *WaterTemperatureController {
	now := time.Now()

	log.Infof("Initialized water temperature controller with producing temp %0.1f, normal temp %0.1f", config.ProducingTemp, config.NormalTemp)

	return &WaterTemperatureController{
		ecotouch:       ecotouch,
		config:         config,
		gridStateSince: now,
		lastTempChange: now,
	}
}

// Determine state and set water temperature if necessary
func (wt *WaterTemperatureController) Reconcile(ctx context.Context, metrics metric.MetricsMap) {
	froniusMetrics := metrics["fronius"]
	ecotouchMetrics := metrics["ecotouch"]
	now := time.Now()

	wt.determineProduceState(froniusMetrics, now)

	// Abort if produce state changed recently
	if now.Sub(wt.gridStateSince).Minutes() < wt.config.ProduceMinTime {
		log.Debugf("Current grid state too young (%0.1f min): aborting", now.Sub(wt.gridStateSince).Minutes())
		return
	}

	wt.determineTemperatureWanted()

	if wt.temperatureChangeWanted(ecotouchMetrics, now) {

		if wt.ecotouch == nil {
			log.Warnf("Not actually setting water temperature: no writer appliance!")
			wt.lastTempChange = now

		} else {
			log.Infof("Setting water temperature to %v", wt.tempWanted)
			if wt.ecotouch.SetValue(ctx, "water.temp_set2", wt.tempWanted) == nil {
				wt.lastTempChange = now
			}
		}
	}
}

// Determine if we are producing (enough) power
func (wt *WaterTemperatureController) determineProduceState(froniusMetrics *metric.Metrics, now time.Time) {
	var gridStateWanted gridState
	if froniusMetrics.Get("inverter.p_grid").Value < wt.config.ProduceThreshold {
		gridStateWanted = producing
	} else {
		gridStateWanted = nonproducing
	}

	if wt.gridState != gridStateWanted {
		wt.gridState = gridStateWanted
		wt.gridStateSince = now
	}

	log.Debugf("Current grid state: %v (%v W)", wt.gridState, froniusMetrics.Get("inverter.p_grid").Value)
}

// Determine desired water temperature
func (wt *WaterTemperatureController) determineTemperatureWanted() {
	if wt.gridState == producing {
		wt.tempWanted = wt.config.ProducingTemp
	} else {
		wt.tempWanted = wt.config.NormalTemp
	}

	log.Debugf("Current temperature wanted: %v", wt.tempWanted)
}

// Determine if need and should adjust the temprature
func (wt *WaterTemperatureController) temperatureChangeWanted(ecotouchMetrics *metric.Metrics, now time.Time) bool {
	tempDiff := (math.Abs(ecotouchMetrics.Get("water.temp_set").Value-wt.tempWanted) > 1)
	if !tempDiff {
		log.Debugf("Temperature changed not needed: %v == %v", ecotouchMetrics.Get("water.temp_set").Value, wt.tempWanted)
	}

	timeDiff := (now.Sub(wt.lastTempChange).Minutes() >= wt.config.TempChangeMinTime)
	if !timeDiff {
		log.Debugf("Last temperature change less than %v min ago", wt.config.TempChangeMinTime)
	}

	return tempDiff && timeDiff
}
