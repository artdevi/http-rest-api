package main

import (
	"flag"
	"log"

	"github.com/BurntSushi/toml"
	"github.com/artdevi/http-rest-api/internal/app/apiserver"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "configs/apiserver.toml", "path to config file")
}

func main() {
	flag.Parse()

	config := apiserver.NewConfig()
	_, err := toml.DecodeFile("configs/apiserver.toml", config)
	if err != nil {
		log.Fatal(err)
	}

	s := apiserver.New(config)
	err = s.Start()
	if err != nil {
		log.Fatal(err)
	}
}
