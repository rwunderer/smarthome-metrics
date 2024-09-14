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
	"github.com/rwunderer/smarthome-metrics/internal/pkg/controllers"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/ecotouch"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/fronius"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/graphite"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/nrgkick"
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

	return config
}

// Create all controllers
func createControllers(config *config.Config) (controllers.ApplianceMap, controllers.ControllersMap) {
	appliances := make(controllers.ApplianceMap)

	for _, name := range config.ActiveAppliances {
		switch name {
		case "fronius":
			appliances[name], _ = fronius.NewController(&config.Fronius)
		case "nrgkick":
			appliances[name], _ = nrgkick.NewController(&config.Nrgkick)
		case "ecotouch":
			appliances[name], _ = ecotouch.NewController(&config.Ecotouch)
		default:
			log.Warnf("Ignoring unkown appliance type %v", name)
		}
	}

	controllers := make(controllers.ControllersMap)

	for _, name := range config.ActiveControllers {
		switch name {
		case "ecotouch":
			controllers[name] = watertemp.NewController(&config.WaterTemperature, appliances["ecotouch"])
		default:
			log.Warnf("Ignoring unkown controller type %v", name)
		}
	}

	return appliances, controllers
}

// Initialize metric storage
func initMetrics(config *config.Config) map[string]*metric.Metrics {
	metrics := make(metric.MetricsMap)

	for _, name := range config.ActiveAppliances {
		switch name {
		case "fronius":
			metrics[name] = metric.NewMetrics(config.Fronius.Prefix)
		case "nrgkick":
			metrics[name] = metric.NewMetrics(config.Nrgkick.Prefix)
		case "ecotouch":
			metrics[name] = metric.NewMetrics(config.Ecotouch.Prefix)
		default:
			log.Warnf("Ignoring unkown metric type %v", name)
		}
	}

	return metrics
}

// Main
func main() {

	// initialize
	config := readConfig()
	appliances, controllers := createControllers(config)
	graphiteClient := graphite.NewClient(&config.Graphite)
	metrics := initMetrics(config)

	// run main loop
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	ctx, cancel := context.WithCancel(context.Background())
	defer cleanUp()
	defer cancel()

	for name, appliance := range appliances {
		go appliance.Run(ctx, metrics[name])
	}

	time.Sleep(3 * time.Second)

	for c, m := range metrics {
		graphiteClient.Send(c, m)
	}

	for _, controller := range controllers {
		controller.Reconcile(ctx, metrics)
	}

Loop:
	for {
		select {
		case <-time.After(30 * time.Second):
			for c, m := range metrics {
				graphiteClient.Send(c, m)
			}

			for _, controller := range controllers {
				controller.Reconcile(ctx, metrics)
			}
		case <-s:
			for _, appliance := range appliances {
				appliance.Close(ctx)
			}
			break Loop
		}
	}
}
