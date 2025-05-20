package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tirlochanarora16/go_tmdb/requests"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	requests.FetchMoviesData()

}
