package main

import (
	"flag"
	"log"
	"os"

	"github.com/yaraloveyou/coffe-crafter.web-server/internal/app/webserver"
	"gopkg.in/yaml.v3"
)

var (
	configPath    string
	jwtConfigPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/dev_web.yaml", "path to config file")
	flag.StringVar(&jwtConfigPath, "jwt-config-path", "configs/jwt.yaml", "path to jwt config file")
}

func main() {
	flag.Parse()

	config := webserver.NewConfig()
	data, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
		return
	}

	if err = yaml.Unmarshal(data, &config); err != nil {
		log.Fatal(err)
		return
	}

	if err := webserver.Start(config, jwtConfigPath); err != nil {
		log.Fatal(err)
		return
	}
}
