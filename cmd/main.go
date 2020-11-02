package main

import (
	"log"

	"github.com/alexvelfr/go-template/pkg/config"
	"github.com/alexvelfr/go-template/server"
)

func main() {
	if err := config.InitConfig(); err != nil {
		log.Fatal(err)
	}
	app := server.NewApp()
	app.Run()
}
