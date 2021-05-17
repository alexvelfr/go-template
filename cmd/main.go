package main

import (
	"log"

	"github.com/alexvelfr/go-template/pkg/config"
	"github.com/alexvelfr/go-template/pkg/logger"
	"github.com/alexvelfr/go-template/server"
	"github.com/spf13/viper"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}
	logger.InitLogger()
	app := server.NewApp()
	app.Run(viper.GetString("app.http_port"))
}
