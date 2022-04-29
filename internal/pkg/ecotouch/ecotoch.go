package ecotouch

import (
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strings"
	"strconv"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/config"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

type tagDescription struct {
	name	  string
	module	string
	fact    float64
	usage		int
}

var tags = map[string]tagDescription {
	"A1": {
		name:   "temp",
		module: "outside",
		fact:   0.1,
    usage:  1,
  },
	"A2": {
		name:   "temp_1h",
		module: "outside",
		fact:   0.1,
    usage:  2,
	},
	"A3": {
		name:   "temp_24h",
		module: "outside",
		fact:   0.1,
    usage:  2,
	},
}

type EcotouchController struct {
	Config    *config.Ecotouch
	loginUrl  string
	logoutUrl string
	readUrl   string
	client    http.Client
}

// NewController creates a new Controller
func NewController(config *config.Ecotouch) (*EcotouchController, error) {

	if err := validateConfig(config); err != nil {
		return nil, fmt.Errorf("Ecotouch Controller config invalid: %v", err)
	}

	loginUrl := fmt.Sprintf("%s/cgi/login?username=%s&password=%s", config.BaseUrl, config.Username, config.Password)
	logoutUrl := fmt.Sprintf("%s/cgi/logout", config.BaseUrl)

	var tagPar string
	i := 0
	for t := range tags {
		i += 1
		tagPar += fmt.Sprintf("&t%v=%v", i, t)
	}
	readUrl := fmt.Sprintf("%s/cgi/readTags?n=%v%v", config.BaseUrl, len(tags), tagPar)

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, fmt.Errorf("Couldn't create cookie jar: %s", err)
	}

	client := http.Client{
		Jar: jar,
	}

	return &EcotouchController{
		Config:    config,
		loginUrl:  loginUrl,
		logoutUrl: logoutUrl,
		readUrl:   readUrl,
		client:    client,
	}, nil
}

// Validate configuration
func validateConfig(conf *config.Ecotouch) error {
	var errs []string

	if conf.BaseUrl == "" {
		errs = append(errs, "Ecotouch BaseUrl not specified! Please set ecotouch.baseUrl in config file!")
	}

	if len(errs) > 0 {
		return fmt.Errorf("%s", strings.Join(errs, ", "))
	}

	return nil
}

// Main run loop
func (controller *EcotouchController) Run(ctx context.Context, metrics *metric.Metrics) error {
	var err error

	if err = controller.getMetrics(ctx, metrics); err != nil {
		return err
	}

	for {
		select {
		case <-ctx.Done():
			log.Debugf("Context Done. Shutting down")
			controller.Close(ctx)
			return nil
		case <-time.After(30 * time.Second):
			if err = controller.getMetrics(ctx, metrics); err != nil {
				controller.Close(ctx)
				return err
			}
		}
	}
}

// Log out of ecotouch
func (controller *EcotouchController) Close(ctx context.Context) {
	if err := controller.login(ctx, controller.logoutUrl); err != nil {
		log.Errorf("Error logging out of Ecotouch: %v", err)
	} else {
		log.Infof("Successfully logged out of Ecotouch")
	}
}

// Retrieve all configured metrics
func (controller *EcotouchController) getMetrics(ctx context.Context, metrics *metric.Metrics) error {
	var err error
	var body []byte
	retry := 2

	for {
		retry--
		body, err = controller.retrieveHttpData(ctx, controller.readUrl)

		if err != nil {
			if err.Error() == "Login required" {
				err = controller.login(ctx, controller.loginUrl)
			}

			if err != nil {
				log.Errorf("Error retrieving Ecotouch Meter data: %v", err)
				return nil
			}

		} else {
			retry = 0
		}

		if retry < 1 {
			break
		}
	}

	for _, v := range strings.Split(string(body), "#") {
		m := strings.Fields(v)
		if len(m) >= 4 {
			if val, err := strconv.ParseFloat(m[3], 64); err == nil {
				log.Infof("metric: %v = %v", m[0], val * tags[m[0]].fact)
			}
		}
	}

	return nil
}
