package watertemp

import (
	"context"
	"math"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/ecotouch"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

const produceThreshold = -2000
const produceMinTime = 10
const producingTemp = 55
const normalTemp = 48
const tempChangeMinTime = 30

type gridState int

const (
	nonproducing gridState = iota
	producing
)

type WaterTemperatureController struct {
	ecotouch *ecotouch.EcotouchController

	gridState      gridState
	gridStateSince time.Time
	tempWanted     float64
	lastTempChange time.Time
}

// Create a new Controller
func NewController(ecotouch *ecotouch.EcotouchController) *WaterTemperatureController {
	now := time.Now()

	return &WaterTemperatureController{
		ecotouch:       ecotouch,
		gridStateSince: now,
		lastTempChange: now,
	}
}

// Determine state and set water temperature if necessary
func (wt *WaterTemperatureController) Reconcile(ctx context.Context, froniusMetrics *metric.Metrics, ecotouchMetrics *metric.Metrics) {
	now := time.Now()

	wt.determineProduceState(froniusMetrics, now)

	// Abort if produce state changed recently
	if now.Sub(wt.gridStateSince).Minutes() < produceMinTime {
		log.Debugf("Current grid state too young (%0.0f min): aborting", now.Sub(wt.gridStateSince).Minutes())
		return
	}

	wt.determineTemperatureWanted()

	if wt.temperatureChangeWanted(ecotouchMetrics, now) {
		log.Infof("Setting water temperature to %v", wt.tempWanted)
		if wt.ecotouch.SetWaterTemp(ctx, wt.tempWanted) == nil {
			wt.lastTempChange = now
		}
	}
}

// Determine if we are producing (enough) power
func (wt *WaterTemperatureController) determineProduceState(froniusMetrics *metric.Metrics, now time.Time) {
	var gridStateWanted gridState
	if froniusMetrics.Get("inverter.p_grid").Value < produceThreshold {
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
		wt.tempWanted = producingTemp
	} else {
		wt.tempWanted = normalTemp
	}

	log.Debugf("Current temperature wanted: %v", wt.tempWanted)
}

// Determine if need and should adjust the temprature
func (wt *WaterTemperatureController) temperatureChangeWanted(ecotouchMetrics *metric.Metrics, now time.Time) bool {
	tempDiff := (math.Abs(ecotouchMetrics.Get("water.temp_set").Value-wt.tempWanted) > 1)
	if !tempDiff {
		log.Debugf("Temperature changed not needed: %v == %v", ecotouchMetrics.Get("water.temp_set").Value, wt.tempWanted)
	}

	timeDiff := (now.Sub(wt.lastTempChange).Minutes() >= tempChangeMinTime)
	if !timeDiff {
		log.Debugf("Last temperature change less than %v min ago", tempChangeMinTime)
	}

	return tempDiff && timeDiff
}
