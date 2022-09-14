package fronius

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

type FroniusPowerFlowSite struct {
	ETotal             float64 `json:"E_Total"`
	PAkku              float64 `json:"P_Akku"`
	PGrid              float64 `json:"P_Grid"`
	PLoad              float64 `json:"P_Load"`
	PPV                float64 `json:"P_PV"`
	RelAutonomy        float64 `json:"rel_Autonomy"`
	RelSelfConsumption float64 `json:"rel_SelfConsumption"`
}

type FroniusPowerFlowData struct {
	Site FroniusPowerFlowSite `json:"Site"`
}

type FroniusPowerFlowHead struct {
	Timestamp string `json:"Timestamp"`
}

type FroniusPowerFlowBody struct {
	Data FroniusPowerFlowData `json:"Data"`
}

type FroniusPowerFlowDoc struct {
	Body FroniusPowerFlowBody `json:"Body"`
	Head FroniusPowerFlowHead `json:"Head"`
}

// Retrieve Fronius PowerFlow data
func (controller *FroniusController) getPowerFlow(ctx context.Context, metrics *metric.Metrics) error {
	body, err := controller.retrieveHttpData(ctx, controller.flowUrl)
	if err != nil {
		log.Errorf("Error retrieving Fronius PowerFlow data: %v", err)
		return nil
	}

	d := FroniusPowerFlowDoc{}

	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Errorf("Error parsing Fronius PowerFlow data: %v", err)
		return nil
	}

	log.WithFields(log.Fields{
		"energy":          int(d.Body.Data.Site.ETotal),
		"akku":            int(d.Body.Data.Site.PAkku),
		"grid":            int(d.Body.Data.Site.PGrid),
		"load":            int(d.Body.Data.Site.PLoad),
		"pv":              int(d.Body.Data.Site.PPV),
		"autonomy":        d.Body.Data.Site.RelAutonomy,
		"selfconsumption": d.Body.Data.Site.RelSelfConsumption,
		"timestamp":       d.Head.Timestamp,
	}).Debug("Successfully parsed PowerFlow data")

	metrics.Set("inverter.e_total", d.Body.Data.Site.ETotal)
	metrics.Set("inverter.p_akku", d.Body.Data.Site.PAkku)
	metrics.Set("inverter.p_grid", d.Body.Data.Site.PGrid)
	metrics.Set("inverter.p_load", d.Body.Data.Site.PLoad)
	metrics.Set("inverter.p_pv", d.Body.Data.Site.PPV)
	metrics.Set("inverter.rel_autonomy", d.Body.Data.Site.RelAutonomy)
	metrics.Set("inverter.rel_selfconsumption", d.Body.Data.Site.RelSelfConsumption)

	return nil
}
