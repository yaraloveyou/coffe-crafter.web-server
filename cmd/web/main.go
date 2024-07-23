package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/webserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/web.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := webserver.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}

	s := webserver.New(config)
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
