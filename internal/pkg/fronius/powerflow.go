package fronius

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type FroniusPowerFlowSite struct {
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
func (controller *FroniusController) getPowerFlow(ctx context.Context) error {
	body, err := controller.retrieveHttpData(controller.flowUrl)
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
		"akku":            int(d.Body.Data.Site.PAkku),
		"grid":            int(d.Body.Data.Site.PGrid),
		"load":            int(d.Body.Data.Site.PLoad),
		"pv":              int(d.Body.Data.Site.PPV),
		"autonomy":        d.Body.Data.Site.RelAutonomy,
		"selfconsumption": d.Body.Data.Site.RelSelfConsumption,
		"timestamp":       d.Head.Timestamp,
	}).Debug("Successfully parsed PowerFlow data")

	return nil
}
