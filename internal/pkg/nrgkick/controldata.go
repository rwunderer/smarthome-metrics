package nrgkick

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

type NrgkickControlDoc struct {
	CurrentSet  float64 `json:"current_set"`
	ChargePause uint8   `json:"charge_pause"`
	EnergyLimit uint32  `json:"energy_limit"`
	PhaseCount  uint8   `json:"phase_count"`
}

// Retrieve Nrgkick Control data
func (controller *NrgkickController) getControlData(ctx context.Context, metrics *metric.Metrics) error {
	body, err := controller.retrieveHttpData(ctx, controller.controlUrl)
	if err != nil {
		log.Errorf("Error retrieving Nrgkick Control data: %v", err)
		return nil
	}

	d := NrgkickControlDoc{}

	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Errorf("Error parsing Nrgkick Control data: %v", err)
		return nil
	}

	metrics.Set("control.current_set", d.CurrentSet)
	metrics.Set("control.charge_pause", float64(d.ChargePause))
	metrics.Set("control.energy_limit", float64(d.EnergyLimit))
	metrics.Set("control.phase_count", float64(d.PhaseCount))

	return nil
}
