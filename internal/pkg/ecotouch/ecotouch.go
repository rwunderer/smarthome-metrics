package ecotouch

import (
	"context"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/config"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

var modules = map[string]struct{}{
	"main":    {},
	"geo":     {},
	"water":   {},
	"heating": {},
	"cooling": {},
	"comp1":   {},
	"comp2":   {},
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
	tagCount := 0
	for t, v := range tags {
		if _, ok := modules[v.module]; ok {
			tagCount += 1
			tagPar += fmt.Sprintf("&t%v=%v", tagCount, t)
		}
	}
	readUrl := fmt.Sprintf("%s/cgi/readTags?n=%v%v", config.BaseUrl, tagCount, tagPar)

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

// Log out of ecotouch
func (controller *EcotouchController) Close(ctx context.Context) {
	if err := controller.login(ctx, controller.logoutUrl); err != nil {
		log.Errorf("Error logging out of Ecotouch: %v", err)
	} else {
		log.Infof("Successfully logged out of Ecotouch")
	}
}

// Main run loop
func (controller *EcotouchController) Run(ctx context.Context, metrics *metric.Metrics) error {
	var err error

	if err = controller.getMetrics(ctx, metrics); err != nil {
		controller.Close(ctx)
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

	var day, month, year, hour, minute int
	for _, v := range strings.Split(string(body), "#") {
		m := strings.Fields(v)
		if len(m) >= 4 {
			if val, err := strconv.ParseFloat(m[3], 64); err == nil {
				tag := tags[m[0]]

				// Deal with special tags
				switch m[0] {
				case "I51":
					for _, w := range stateWord {
						metrics.Set(fmt.Sprintf("%s.%s", w.module, w.name), float64(int(val)&w.flag))
					}
				case "I5":
					day = int(val)
				case "I6":
					month = int(val)
				case "I7":
					year = 2000 + int(val)
				case "I8":
					hour = int(val)
				case "I9":
					minute = int(val)
				default:
					metrics.Set(fmt.Sprintf("%s.%s", tag.module, tag.name), val*tag.fact)
				}
			}
		}
	}

	if zone, err := time.LoadLocation("Europe/Vienna"); err == nil {
		t := time.Date(year, time.Month(month), day, hour, minute, 0, 0, zone)
		metrics.Set("main.datetime", float64(t.Unix()))
		metrics.Set("main.time", float64(hour*3600+minute*60))
		metrics.Set("main.timediff", time.Since(t).Seconds())
		log.Debugf("datetime=%v", t.UTC())
	}

	return nil
}
