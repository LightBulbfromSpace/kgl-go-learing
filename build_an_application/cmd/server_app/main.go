package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	server "github.com/LightBulbfromSpace/build_an_application/internal/app/server_app"
	"log"
	"net/http"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config-path", "configs/server_app.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := server.NewConfig()
	_, err := toml.DecodeFile(configPath, config)
	if err != nil {
		log.Fatal(err)
	}
	server := server.NewServer(config)

	if err := server.Start(); err != nil {
		log.Fatal(err)
	}

	log.Fatal(http.ListenAndServe(":8080", server))
}
