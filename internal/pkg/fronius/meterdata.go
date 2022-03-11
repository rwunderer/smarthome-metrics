package fronius

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type FroniusMeterData struct {
	Consumed float64 `json:"EnergyReal_WAC_Sum_Consumed"`
	Produced float64 `json:"EnergyReal_WAC_Sum_Produced"`
}

type FroniusMeterHead struct {
	Timestamp string `json:"Timestamp"`
}

type FroniusMeterBody struct {
	Data FroniusMeterData `json:"Data"`
}

type FroniusMeterDoc struct {
	Body FroniusMeterBody `json:"Body"`
	Head FroniusMeterHead `json:"Head"`
}

// Retrieve Fronius Meter data
func (controller *FroniusController) getMeterData(ctx context.Context) error {
	body, err := controller.retrieveHttpData(controller.meterUrl)
	if err != nil {
		log.Errorf("Error retrieving Fronius Meter data: %v", err)
		return nil
	}

	d := FroniusMeterDoc{}

	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Errorf("Error parsing Fronius Meter data: %v", err)
		return nil
	}

	log.WithFields(log.Fields{
		"consumed":  int(d.Body.Data.Consumed),
		"produced":  int(d.Body.Data.Produced),
		"timestamp": d.Head.Timestamp,
	}).Debug("Successfully parsed Meter data")

	return nil
}
