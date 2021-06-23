package main

import (
	"flag"
	"github.com/anatolethien/forum/internal/app/server"
	"log"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/config.json", "path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	if err := server.ReadConfig(configPath, config); err != nil {
		log.Fatal(err)
	}

	s := server.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
