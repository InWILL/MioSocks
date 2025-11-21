package main

import (
	"encoding/json"
	"flag"
	"log"
	"os"

	"github.com/InWILL/MioSocks/config"
	"github.com/InWILL/MioSocks/service"
)

func ParseConfig(file string) config.Options {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	config := config.Options{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func main() {
	config := flag.String("c", "config.json", "Path to the configuration file")

	flag.Parse()
	globalConfig := ParseConfig(*config)

	service, err := service.NewService(globalConfig)
	if err != nil {
		log.Fatalf("Failed to parse proxy: %v", err)
	}

	service.Start()
}
