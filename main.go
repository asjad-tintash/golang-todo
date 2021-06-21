package main

import (
	"./api/controllers"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main(){
	if err := godotenv.Load(); err != nil {
		log.Fatal("Unable to load the env file")
	}

	app := controllers.App{}
	app.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_PORT"),
		)

	app.RunServer()
}