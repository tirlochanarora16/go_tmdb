package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/tirlochanarora16/go_tmdb/requests"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	start := time.Now()

	requests.FetchMoviesData()
	requests.FetchMoviesData()
	requests.FetchMoviesData()
	requests.FetchMoviesData()

	fmt.Println(time.Since(start))

}
