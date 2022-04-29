package ecotouch

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

// Retrieve Ecotouch data
func (controller *EcotouchController) retrieveHttpData(ctx context.Context, url string) ([]byte, error) {
	reqCtx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	log.WithFields(log.Fields{
		"url": url,
	}).Debug("Retrieving Ecotouch Meter data")

	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "smarthome-metrics")

	res, err := controller.client.Do(req)
	if err != nil {
		log.Errorf("Error retrieving Ecotouch Meter data: %v", err)
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error reading Ecotouch Meter data: %v", err)
		return nil, err
	}

	if strings.TrimSpace(string(body)) == "#E_NEED_LOGIN" {
		return nil, fmt.Errorf("Login required")
	}

	log.Tracef("Got body data: |%s|", string(body))

	return body, nil
}

// Login to Ecotouch
func (controller *EcotouchController) login(ctx context.Context, url string) error {
	reqCtx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "smarthome-metrics")

	res, err := controller.client.Do(req)
	if err != nil {
		log.Errorf("Error logging into Ecotouch: %v", err)
		return err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error reading Ecotouch login data: %v", err)
		return err
	}

	log.Tracef("Got body data: |%s|", body)

	return nil
}
