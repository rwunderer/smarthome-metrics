package fronius

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

type FroniusBatteryController struct {
	CapacityMaximum float64 `json:"Capacity_Maximum"`
	ChargeRelative  float64 `json:"StateOfCharge_Relative"`
	CellTemperature float64 `json:"Temperature_Cell"`
}

type FroniusBatteryChannelsData struct {
	CapacityRemaining  float64 `json:"BAT_CAPACITY_ESTIMATION_REMAINING_F64"`
	LifetimeCharged    float64 `json:"BAT_ENERGYACTIVE_LIFETIME_CHARGED_F64"`
	LifetimeDischarged float64 `json:"BAT_ENERGYACTIVE_LIFETIME_DISCHARGED_F64"`
}

type FroniusBatteryAddData struct {
	Channels FroniusBatteryChannelsData `json:"channels"`
}

type FroniusBatteryData struct {
	Controller FroniusBatteryController `json:"Controller"`
	Battery    FroniusBatteryAddData    `json:"16580608"`
}

type FroniusBatteryHead struct {
	Timestamp string `json:"Timestamp"`
}

type FroniusBatteryBody struct {
	Data FroniusBatteryData `json:"Data"`
}

type FroniusBatteryDoc struct {
	Body FroniusBatteryBody `json:"Body"`
	Head FroniusBatteryHead `json:"Head"`
}

// Retrieve Fronius Battery data
func (controller *FroniusController) getBatteryData(ctx context.Context) error {
	err := controller.getBatteryBaseData(ctx)
	if err != nil {
		log.Errorf("Error retrieving Fronius Battery base data: %v", err)
		return nil
	}

	err = controller.getBatteryAdditionalData(ctx)
	if err != nil {
		log.Errorf("Error retrieving Fronius Battery additional data: %v", err)
		return nil
	}

	return nil
}

func (controller *FroniusController) getBatteryBaseData(ctx context.Context) error {
	body, err := controller.retrieveHttpData(controller.batteryUrl)
	if err != nil {
		log.Errorf("Error retrieving Fronius Battery data: %v", err)
		return nil
	}

	d := FroniusBatteryDoc{}

	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Errorf("Error parsing Fronius Battery data: %v", err)
		return nil
	}

	log.WithFields(log.Fields{
		"bat_max":     int(d.Body.Data.Controller.CapacityMaximum),
		"bat_remain":  int(d.Body.Data.Controller.CapacityMaximum * d.Body.Data.Controller.ChargeRelative / 100),
		"charge":      d.Body.Data.Controller.ChargeRelative,
		"temperature": d.Body.Data.Controller.CellTemperature,
		"timestamp":   d.Head.Timestamp,
	}).Debug("Successfully parsed Battery data")

	return nil
}

func (controller *FroniusController) getBatteryAdditionalData(ctx context.Context) error {
	body, err := controller.retrieveHttpData(controller.batteryAddUrl)
	if err != nil {
		log.Errorf("Error retrieving Fronius Battery additional data: %v", err)
		return nil
	}

	d := FroniusBatteryDoc{}

	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Errorf("Error parsing Fronius Battery additional data: %v", err)
		return nil
	}

	log.WithFields(log.Fields{
		"bat_remaining": int(d.Body.Data.Battery.Channels.CapacityRemaining),
		"charged":       d.Body.Data.Battery.Channels.LifetimeCharged,
		"discharged":    d.Body.Data.Battery.Channels.LifetimeDischarged,
		"timestamp":     d.Head.Timestamp,
	}).Debug("Successfully parsed Battery data")

	return nil
}
