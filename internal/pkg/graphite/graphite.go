package graphite

import (
	graphiteClient "github.com/jtaczanowski/go-graphite-client"
	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/config"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

type Client struct {
	config config.Graphite
	client *graphiteClient.Client
}

// Create new client object
func NewClient(config *config.Config) *Client {
	return &Client{
		config: config.Graphite,
		client: graphiteClient.NewClient(
			config.Graphite.Hostname,
			config.Graphite.Port,
			config.Graphite.Prefix,
			config.Graphite.Protocol,
		),
	}
}

// Send one metric set to graphite
func (graphite *Client) Send(controllerName string, metric *metric.Metrics) {

	if err := graphite.client.SendData(metric.GetGraphiteMap()); err != nil {
		log.Errorf("Error sending %v metrics to %v:%v: %v",
			controllerName,
			graphite.config.Hostname,
			graphite.config.Port,
			err,
		)

	} else {
		log.Infof("Sent %v metrics to %v:%v",
			controllerName,
			graphite.config.Hostname,
			graphite.config.Port,
		)
	}
}
