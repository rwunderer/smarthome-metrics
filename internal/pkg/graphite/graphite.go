package graphite

import (
	graphiteClient "github.com/jtaczanowski/go-graphite-client"
	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
)

type Config struct {
	Hostname string `yaml:"hostname"`
	Port     int    `yaml:"port"`
	Prefix   string `yaml:"prefix"`
	Protocol string `yaml:"protocol"`
}

type Client struct {
	config *Config
	client *graphiteClient.Client
}

// Get Default configuration
func GetDefaultConfig() Config {
	return Config{}
}

// Create new client object
func NewClient(config *Config) *Client {
	return &Client{
		config: config,
		client: graphiteClient.NewClient(
			config.Hostname,
			config.Port,
			config.Prefix,
			config.Protocol,
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
