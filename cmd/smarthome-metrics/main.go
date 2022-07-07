package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	graphite "github.com/jtaczanowski/go-graphite-client"
	log "github.com/sirupsen/logrus"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/config"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/ecotouch"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/fronius"
	"github.com/rwunderer/smarthome-metrics/internal/pkg/metric"
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

func cleanUp() {
	log.Infof("Clean up")
}

type SmarthomeController interface {
	Run(ctx context.Context, metrics *metric.Metrics) error
	Close(ctx context.Context)
}

func main() {
	var configFile string
	activeControllers := []string{"fronius", "ecotouch"}

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

	// create controllers
	controllers := make(map[string]SmarthomeController)
	controllers["fronius"], _ = fronius.NewController(&config.Fronius)
	controllers["ecotouch"], _ = ecotouch.NewController(&config.Ecotouch)

	// create graphite output
	graphiteClient := graphite.NewClient(
		config.Graphite.Hostname,
		config.Graphite.Port,
		config.Graphite.Prefix,
		config.Graphite.Protocol,
	)

	// run main loop
	s := make(chan os.Signal, 1)
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	ctx, cancel := context.WithCancel(context.Background())
	defer cleanUp()
	defer cancel()

	metrics := make(map[string]*metric.Metrics)
	metrics["fronius"] = metric.NewMetrics(config.Fronius.Prefix)
	metrics["ecotouch"] = metric.NewMetrics(config.Ecotouch.Prefix)

	for _, c := range activeControllers {
		go controllers[c].Run(ctx, metrics[c])
	}

	time.Sleep(3 * time.Second)

	for c, m := range metrics {
		if err := graphiteClient.SendData(m.GetGraphiteMap()); err != nil {
			log.Errorf("Error sending metrics to %v:%v: %v",
				config.Graphite.Hostname,
				config.Graphite.Port,
				err,
			)
		} else {
			log.Infof("Sent %v metrics to %v:%v",
				c,
				config.Graphite.Hostname,
				config.Graphite.Port,
			)
		}
	}

Loop:
	for {
		select {
		case <-time.After(30 * time.Second):
			for c, m := range metrics {
				if err := graphiteClient.SendData(m.GetGraphiteMap()); err != nil {
					log.Errorf("Error sending metrics to %v:%v: %v",
						config.Graphite.Hostname,
						config.Graphite.Port,
						err,
					)
				} else {
					log.Infof("Sent %v metrics to %v:%v",
						c,
						config.Graphite.Hostname,
						config.Graphite.Port,
					)
				}
			}
		case <-s:
			for _, c := range activeControllers {
				controllers[c].Close(ctx)
			}
			cancel()
			break Loop
		}
	}
}
