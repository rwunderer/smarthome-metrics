package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/config"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/ecotouch"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/fronius"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/graphite"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/watertemp"
)

// Initialize log module
func init() {
	// load optional .env file
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file: %v", err)
		}
	}

	// set global log level
	lvl, ok := os.LookupEnv("LOG_LEVEL")
	if !ok {
		lvl = "info"
	}

	ll, err := log.ParseLevel(lvl)
	if err != nil {
		ll = log.DebugLevel
	}

	log.SetLevel(ll)

	// set log format
	lfmt, ok := os.LookupEnv("LOG_FORMAT")
	if !ok {
		lfmt = "text"
	}

	switch lfmt {
	case "text":
		log.SetFormatter(&log.TextFormatter{FullTimestamp: true})
	case "json":
		log.SetFormatter(&log.JSONFormatter{})
	default:
		log.Fatalf("Invalid log format: %v", lfmt)
	}
}

// Log cleanup
func cleanUp() {
	log.Infof("Clean up")
}

type SmarthomeController interface {
	Run(ctx context.Context, metrics *metric.Metrics) error
	Close(ctx context.Context)
}

// Read and parse config file
func readConfig() *config.Config {
	var configFile string

	// define config file
	configFile, ok := os.LookupEnv("CONFIG_FILE")
	if !ok {
		configFile = "./config.yaml"
	}

	config := &config.Config{}

	// define command-line flags
	flag.StringVar(&config.Fronius.BaseUrl, "fronius-baseurl", os.Getenv("FRONIUS_BASEURL"), "Fronius base url (eg http://fronius)")

	// read config from config file
	if _, err := os.Stat(configFile); err == nil {
		if err := config.ReadFile(configFile); err != nil {
			log.Errorf("could not parse config file: %v", err)
			os.Exit(1)
		}
	}

	// parse command-line flags
	flag.Parse()
	log.Infof("Fronius base url is %v", config.Fronius.BaseUrl)

	return config
}

// Create all controllers
func createControllers(config *config.Config) (map[string]SmarthomeController, *watertemp.WaterTemperatureController) {
	controllers := make(map[string]SmarthomeController)
	controllers["fronius"], _ = fronius.NewController(&config.Fronius)

	etController, _ := ecotouch.NewController(&config.Ecotouch)
	controllers["ecotouch"] = etController
	wtController := watertemp.NewController(&config.WaterTemperature, etController)

	return controllers, wtController
}

// Initialize metric storage
func initMetrics(config *config.Config) map[string]*metric.Metrics {
	metrics := make(map[string]*metric.Metrics)
	metrics["fronius"] = metric.NewMetrics(config.Fronius.Prefix)
	metrics["ecotouch"] = metric.NewMetrics(config.Ecotouch.Prefix)

	return metrics
}

// Main
func main() {

	// initialize
	config := readConfig()
	controllers, wtController := createControllers(config)
	graphiteClient := graphite.NewClient(&config.Graphite)
	metrics := initMetrics(config)

	// run main loop
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	ctx, cancel := context.WithCancel(context.Background())
	defer cleanUp()
	defer cancel()

	for _, c := range config.ActiveControllers {
		go controllers[c].Run(ctx, metrics[c])
	}

	time.Sleep(3 * time.Second)

	for c, m := range metrics {
		graphiteClient.Send(c, m)
	}
	wtController.Reconcile(ctx, metrics["fronius"], metrics["ecotouch"])

Loop:
	for {
		select {
		case <-time.After(30 * time.Second):
			for c, m := range metrics {
				graphiteClient.Send(c, m)
			}

			wtController.Reconcile(ctx, metrics["fronius"], metrics["ecotouch"])
		case <-s:
			for _, c := range config.ActiveControllers {
				controllers[c].Close(ctx)
			}
			break Loop
		}
	}
}
