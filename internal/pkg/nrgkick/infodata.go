package nrgkick

import (
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

type NrgkickInfoGeneral struct {
	SerialNumber string `json:"serial_number"`
	ModelType    string `json:"model_type"`
	DeviceName   string `json:"device_name"`
	RatedCurrent uint8  `json:"rated_current"`
}

type NrgkickInfoConnector struct {
	PhaseCount uint8   `json:"phase_count"`
	MaxCurrent float64 `json:"phase_count"`
	Type       uint8   `json:"type"`
	Serial     string  `json:"serial"`
}

type NrgkickInfoGrid struct {
	Voltage   uint8 `json:"voltage"`
	Frequency uint8 `json:"frequency"`
	Phases    uint8 `json:"phases"`
}

type NrgkickInfoNetwork struct {
	IpAddress  string `json:"ip_address"`
	MacAddress string `json:"mac_address"`
	ssid       string `json:"ssid"`
	rssi       int8   `json:"rssi"`
}

type NrgkickInfoVersions struct {
	SwSm string `json:"sw_sm"`
	HwSm string `json:"hw_sm"`
	SwMa string `json:"sw_ma"`
	HwMa string `json:"hw_ma"`
	SwTo string `json:"sw_to"`
	HwTo string `json:"hw_to"`
	SwSt string `json:"sw_st"`
	HwSt string `json:"hw_st"`
	SwCm string `json:"sw_cm"`
}

type NrgkickInfoDoc struct {
	General   NrgkickInfoGeneral   `json:"general"`
	Connector NrgkickInfoConnector `json:"connector"`
	Grid      NrgkickInfoGrid      `json:"grid"`
	Network   NrgkickInfoNetwork   `json:"network"`
	Versions  NrgkickInfoNetwork   `json:"versions"`
}

// Retrieve Nrgkick Info data
func (controller *NrgkickController) getInfoData(ctx context.Context, metrics *metric.Metrics) error {
	body, err := controller.retrieveHttpData(ctx, controller.infoUrl)
	if err != nil {
		log.Errorf("Error retrieving Nrgkick Info data: %v", err)
		return nil
	}

	d := NrgkickInfoDoc{}

	err = json.Unmarshal(body, &d)
	if err != nil {
		log.Errorf("Error parsing Nrgkick Info data: %v", err)
		return nil
	}

	// metrics.Set("info.general.serial_number", d.General.SerialNumber)
	// metrics.Set("info.general.model_type", d.General.ModelType)
	// metrics.Set("info.general.device_name", d.General.DeviceName)
	metrics.Set("info.general.rated_current", float64(d.General.RatedCurrent))

	metrics.Set("info.connector.phase_count", float64(d.Connector.PhaseCount))
	metrics.Set("info.connector.max_current", d.Connector.MaxCurrent)
	metrics.Set("info.connector.type", float64(d.Connector.Type))
	// metrics.Set("info.connector.serial", d.Connector.Serial)

	metrics.Set("info.grid.voltage", float64(d.Grid.Voltage))
	metrics.Set("info.grid.frequency", float64(d.Grid.Frequency))
	metrics.Set("info.grid.phases", float64(d.Grid.Phases))

	// metrics.Set("info.network.ip_address", d.Network.IpAddress)
	// metrics.Set("info.network.mac_address", d.Network.MacAddress)
	// metrics.Set("info.network.ssid", d.Network.Ssid)
	// metrics.Set("info.network.rssi", d.Network.Rssi)

	// metrics.Set("info.versions.sw_sm", d.Versions.SwSm)
	// metrics.Set("info.versions.hw_sm", d.Versions.HwSm)
	// metrics.Set("info.versions.sw_ma", d.Versions.SwMa)
	// metrics.Set("info.versions.hw_ma", d.Versions.HwMa)
	// metrics.Set("info.versions.sw_to", d.Versions.SwTo)
	// metrics.Set("info.versions.hw_to", d.Versions.HwTo)
	// metrics.Set("info.versions.sw_st", d.versions.swSt)
	// metrics.Set("info.versions.hw_st", d.versions.hwSt)

	return nil
}
