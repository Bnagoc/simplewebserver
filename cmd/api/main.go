package main

import (
	"flag"
	"log"
	"simplewebserver/internal/app/api"

	"github.com/BurntSushi/toml"
)

var (
	configPath string
)

func init() {
	// Set paath before start app from cmd
	flag.StringVar(&configPath, "path", "configs/api.toml", "path to config file in .toml format")
}

func main() {
	// Initialization configPath
	flag.Parse()

	log.Println("Main func has been started.")

	// Server instance initialization (read config from .toml/.env files if new info there)
	config := api.NewConfig()
	// Deserialization toml file
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Println("Can not find configs file. Error: ", err)
	}
	server := api.New(config)

	// api server start
	if err := server.Start(); err != nil {
		log.Fatal(err)
	}
}
