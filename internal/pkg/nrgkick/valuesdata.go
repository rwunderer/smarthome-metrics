package nrgkick

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

type NrgkickValuesEnergy struct {
	TotalChargedEnergy uint64 `json:"total_charged_energy"`
	LastChargedEnergy  uint32 `json:"charged_energy"`
}

type NrgkickValuesPowerflowPhase struct {
	Voltage       float64 `json:"voltage"`
	Current       float64 `json:"current"`
	ActivePower   float64 `json:"active_power"`
	ReactivePower float64 `json:"reactive_power"`
	ApparentPower float64 `json:"apparent_power"`
	PowerFactor   float64 `json:"power_factor"`
}

type NrgkickValuesPowerflowNeutral struct {
	Current float64 `json:"current"`
}

type NrgkickValuesPowerflow struct {
	ChargingVoltage    float64                       `json:"charging_voltage"`
	ChargingCurrent    float64                       `json:"charging_current"`
	GridFrequency      float64                       `json:"grid_frequency"`
	PeakPower          float64                       `json:"peak_power"`
	TotalActivePower   float64                       `json:"total_active_power"`
	TotalReactivePower float64                       `json:"total_reactive_power"`
	TotalApparentPower float64                       `json:"total_apparent_power"`
	TotalPowerFactor   float64                       `json:"total_power_factor"`
	L1                 NrgkickValuesPowerflowPhase   `json:"l1"`
	L2                 NrgkickValuesPowerflowPhase   `json:"l2"`
	L3                 NrgkickValuesPowerflowPhase   `json:"l3"`
	N                  NrgkickValuesPowerflowNeutral `json:"n"`
}

type NrgkickValuesGeneral struct {
	ChargingRate         float64 `json:"charging_rate"`
	Vehicle_connectTime  uint32  `json:"vehicle_connect_time"`
	Vehicle_chargingTime uint32  `json:"vehicle_charging_time"`
	Status               uint8   `json:"status"`
	ChargePermitted      uint8   `json:"charge_permitted"`
	RelayState           uint8   `json:"relay_state"`
	ChargeCount          uint16  `json:"charge_count"`
	RcdTrigger           uint8   `json:"rcd_trigger"`
	WarningCode          uint8   `json:"warning_code"`
	ErrorCode            uint8   `json:"error_code"`
}

type NrgkickValuesTemperatures struct {
	Housing       float64 `json:"housing"`
	ConnectorL1   float64 `json:"connector_l1"`
	ConnectorL2   float64 `json:"connector_l2"`
	ConnectorL3   float64 `json:"connector_l3"`
	DomesticPlug1 float64 `json:"domestic_plug_1"`
	DomesticPlug2 float64 `json:"domestic_plug_2"`
}

type NrgkickValuesDoc struct {
	Energy       NrgkickValuesEnergy       `json:"energy"`
	Powerflow    NrgkickValuesPowerflow    `json:"powerflow"`
	General      NrgkickValuesGeneral      `json:"general"`
	Temperatures NrgkickValuesTemperatures `json:"temperatures"`
}

// Retrieve Nrgkick Values data
func (controller *NrgkickController) getValuesData(ctx context.Context, metrics *metric.Metrics) error {
	body, err := controller.retrieveHttpData(ctx, controller.valuesUrl)
	if err != nil {
		log.Errorf("Error retrieving Nrgkick Values data: %v", err)
		return nil
	}

	d := NrgkickValuesDoc{}

	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Errorf("Error parsing Nrgkick Values data: %v", err)
		return nil
	}

	metrics.Set("values.energy.total_charged_energy", float64(d.Energy.TotalChargedEnergy))
	metrics.Set("values.energy.last_charged_energy", float64(d.Energy.LastChargedEnergy))

	metrics.Set("values.powerflow.charging_voltage", d.Powerflow.ChargingVoltage)
	metrics.Set("values.powerflow.charging_current", d.Powerflow.ChargingCurrent)
	metrics.Set("values.powerflow.grid_frequency", d.Powerflow.GridFrequency)
	metrics.Set("values.powerflow.peak_power", d.Powerflow.PeakPower)
	metrics.Set("values.powerflow.total_active_power", d.Powerflow.TotalActivePower)
	metrics.Set("values.powerflow.total_reactive_power", d.Powerflow.TotalReactivePower)
	metrics.Set("values.powerflow.total_apparent_power", d.Powerflow.TotalApparentPower)
	metrics.Set("values.powerflow.total_power_factor", d.Powerflow.TotalPowerFactor)

	metrics.Set("values.powerflow.l1.voltage", d.Powerflow.L1.Voltage)
	metrics.Set("values.powerflow.l1.current", d.Powerflow.L1.Current)
	metrics.Set("values.powerflow.l1.active_power", d.Powerflow.L1.ActivePower)
	metrics.Set("values.powerflow.l1.reactive_power", d.Powerflow.L1.ReactivePower)
	metrics.Set("values.powerflow.l1.apparent_power", d.Powerflow.L1.ApparentPower)
	metrics.Set("values.powerflow.l1.power_factor", d.Powerflow.L1.PowerFactor)

	metrics.Set("values.powerflow.l2.voltage", d.Powerflow.L2.Voltage)
	metrics.Set("values.powerflow.l2.current", d.Powerflow.L2.Current)
	metrics.Set("values.powerflow.l2.active_power", d.Powerflow.L2.ActivePower)
	metrics.Set("values.powerflow.l2.reactive_power", d.Powerflow.L2.ReactivePower)
	metrics.Set("values.powerflow.l2.apparent_power", d.Powerflow.L2.ApparentPower)
	metrics.Set("values.powerflow.l2.power_factor", d.Powerflow.L2.PowerFactor)

	metrics.Set("values.powerflow.l3.voltage", d.Powerflow.L3.Voltage)
	metrics.Set("values.powerflow.l3.current", d.Powerflow.L3.Current)
	metrics.Set("values.powerflow.l3.active_power", d.Powerflow.L3.ActivePower)
	metrics.Set("values.powerflow.l3.reactive_power", d.Powerflow.L3.ReactivePower)
	metrics.Set("values.powerflow.l3.apparent_power", d.Powerflow.L3.ApparentPower)
	metrics.Set("values.powerflow.l3.power_factor", d.Powerflow.L3.PowerFactor)

	metrics.Set("values.powerflow.n.current", d.Powerflow.N.Current)

	metrics.Set("values.general.charging_rate", d.General.ChargingRate)
	metrics.Set("values.general.vehicle_connect_time", float64(d.General.Vehicle_connectTime))
	metrics.Set("values.general.vehicle_charging_time", float64(d.General.Vehicle_chargingTime))
	metrics.Set("values.general.status", float64(d.General.Status))
	metrics.Set("values.general.charge_permitted", float64(d.General.ChargePermitted))
	metrics.Set("values.general.relay_state", float64(d.General.RelayState))
	metrics.Set("values.general.charge_count", float64(d.General.ChargeCount))
	metrics.Set("values.general.rcd_trigger", float64(d.General.RcdTrigger))
	metrics.Set("values.general.warning_code", float64(d.General.WarningCode))
	metrics.Set("values.general.error_code", float64(d.General.ErrorCode))

	metrics.Set("values.temperatures.housing", d.Temperatures.Housing)
	metrics.Set("values.temperatures.connector_l1", d.Temperatures.ConnectorL1)
	metrics.Set("values.temperatures.connector_l2", d.Temperatures.ConnectorL2)
	metrics.Set("values.temperatures.connector_l3", d.Temperatures.ConnectorL3)
	metrics.Set("values.temperatures.domestic_plug_1", d.Temperatures.DomesticPlug1)
	metrics.Set("values.temperatures.domestic_plug_2", d.Temperatures.DomesticPlug2)

	return nil
}
