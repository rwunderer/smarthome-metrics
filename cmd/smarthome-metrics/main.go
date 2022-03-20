package main

import (
	"context"
	"flag"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/config"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/fronius"
)

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

func main() {
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

	// create fronius controller
	fronius, err := fronius.NewController(&config.Fronius)
	if err != nil {
		log.Errorf("could not initialize fronius controller: %v", err)
		os.Exit(1)
	}

	// run main loop

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()
	metricsCh := make(chan metric.Metric)

	go fronius.Run(ctx, metricsCh)

	for {
		select {
		case metrics := <-metricsCh:
			for k, v := range metrics {
				log.Infof("metric %v = %v", k, v)
			}
		}
	}
}
