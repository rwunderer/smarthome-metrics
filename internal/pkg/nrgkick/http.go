package nrgkick

import (
	"context"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Retrieve Nrgkick data
func (controller *NrgkickController) retrieveHttpData(ctx context.Context, url string) ([]byte, error) {
	reqCtx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	log.WithFields(log.Fields{
		"url": url,
	}).Debug("Retrieving Nrgkick data")

	req, err := http.NewRequestWithContext(reqCtx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "smarthome-metrics")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("Error retrieving Nrgkick Meter data: %v", err)
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error reading Nrgkick Meter data: %v", err)
		return nil, err
	}

	log.Tracef("Got body data: %s", body)

	return body, nil
}
