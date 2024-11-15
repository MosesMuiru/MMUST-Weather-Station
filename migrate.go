package main

import (
	"log"
	"mmust-weather-station/db"
	"mmust-weather-station/db/models"
)

func main(){
	err := db.DB().AutoMigrate(
		&models.Rainfall{},
		&models.Temperature{},
	)
	if err != nil{
		log.Fatal(err)
	}
}
