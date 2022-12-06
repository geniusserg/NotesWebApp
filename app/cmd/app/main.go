package main

import (
	"awesomeProject/app/internal/config"
	"awesomeProject/app/pkg/logging"
	"log"
)

func main() {
	log.Println("config initializing")
	config.GetConfig()
	log.Println("logger initialized")
	logging.GetLogger()

}
