package main

import (
	"fmt"
	"squeeze-auction/docs"
	"squeeze-auction/internal/config"
	"squeeze-auction/internal/logger"
	"squeeze-auction/server"
	"squeeze-auction/server/routes"
)

// @title           Squeeze Auction API
// @version         1.0
// @description     Squeeze Auction API
// @termsOfService  http://swagger.io/terms/

// @BasePath /
func main() {
	docs.SwaggerInfo.Host = ""

	cfg, err := config.NewConfig()
	if err != nil {
		fmt.Println("Error loading config: ", err)
		return
	}

	log := logger.New(cfg)
	log.Info("Starting the application...")

	router := routes.New(log)

	serv := server.New(cfg, router)
	err = serv.Start()
	if err != nil {
		log.Fatalln("Error starting the server: ", err)
	}
}
