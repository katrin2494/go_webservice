package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"go_webservice/service"
	"log"
)

type Config struct {
	Host     string
	Port     string
	BotToken string
}

func main() {
	configPath := flag.String("config", "", "Path to config file")
	flag.Parse()

	cfg := &Config{}
	_, err := toml.DecodeFile(*configPath, cfg)

	if err != nil {
		log.Println("Ошибка декодирования конфигов")
	}

	go service.InitBot(cfg.BotToken)
	service.InitServer(cfg.Host, cfg.Port, cfg.BotToken)
}
