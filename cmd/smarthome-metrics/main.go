package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/rwunderer/smarthome-metrics/internal/pkg/config"
)

func init() {
	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load(".env")

		if err != nil {
			log.Fatal("Error loading .env file")
		}
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
	flag.StringVar(&config.Fronius.Hostname, "fronius-hostname", os.Getenv("FRONIUS_HOSTNAME"), "Fronius Hostname or IP")

	// read config from config file
	if _, err := os.Stat(configFile); err == nil {
		if err := config.ReadFile(configFile); err != nil {
			log.Print(fmt.Errorf("could not parse config file: %v", err))
			os.Exit(1)
		}
	}

	// parse command-line flags
	flag.Parse()

	log.Printf("Fronius hostname is %v", config.Fronius.Hostname)
}
