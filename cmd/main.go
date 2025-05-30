package main

import (
	"github.com/GalahadKingsman/messenger_users/internal/app"
	"github.com/GalahadKingsman/messenger_users/internal/config"
	"github.com/GalahadKingsman/messenger_users/internal/database"
	"github.com/caarlos0/env"
	"log"
)

func main() {
	if err := database.Init(); err != nil {
		log.Fatalf("Ошибка инициализации базы: %v", err)
	}
	configg := new(config.Configs)
	if err := env.Parse(configg); err != nil {
		log.Fatal("can not parce config")
	}

	app.Run(configg)
}
