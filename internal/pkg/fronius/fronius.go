package fronius

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/config"
)

type FroniusController struct {
	Config   *config.Fronius
	meterUrl string
}

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

// NewController creates a new Controller
func NewController(config *config.Fronius) (*FroniusController, error) {

	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("Fronius Controller config invalid: %v", err)
	}

	meterUrl := fmt.Sprintf("%s/solar_api/v1/GetMeterRealtimeData.cgi?Scope=Device&DeviceId=0", config.BaseUrl)

	return &FroniusController{
		Config:   config,
		meterUrl: meterUrl,
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

// Main run loop
func (controller *FroniusController) Run(ctx context.Context) error {
	if err := controller.getMetrics(ctx); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			log.Debugf("Context Done. Shutting down")
			return nil
		case <-time.After(30 * time.Second):
			if err := controller.getMetrics(ctx); err != nil {
				return err
			}
		}
	}
}

// Retrieve all configured metrics
func (controller *FroniusController) getMetrics(ctx context.Context) error {
	if err := controller.getMeterData(ctx); err != nil {
		return err
	}

	return nil
}

// Retrive Fronius Meter data
func (controller *FroniusController) getMeterData(ctx context.Context) error {
	url := controller.meterUrl

	spaceClient := http.Client{
		Timeout: time.Second * 2,
	}

	log.WithFields(log.Fields{
		"url": url,
	}).Debug("Retrieving Fronius Meter data")

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "smarthome-metrics")

	res, err := spaceClient.Do(req)
	if err != nil {
		log.Errorf("Error retrieving Fronius Meter data: %v", err)
		return nil
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error reading Fronius Meter data: %v", err)
		return nil
	}

	log.WithFields(log.Fields{
		"body": body,
	}).Trace("Got body data")

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
