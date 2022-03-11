package fronius

import (
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

// Retrieve Fronius data
func (controller *FroniusController) retrieveHttpData(url string) ([]byte, error) {
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
		return nil, err
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Errorf("Error reading Fronius Meter data: %v", err)
		return nil, err
	}

	log.Tracef("Got body data: %s", body)

	return body, nil
}
