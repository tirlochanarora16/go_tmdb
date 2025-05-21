package main

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/tirlochanarora16/go_tmdb/constants"
	"github.com/tirlochanarora16/go_tmdb/models"
	"github.com/tirlochanarora16/go_tmdb/requests"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	start := time.Now()

	urls := []string{
		constants.DiscoverMoviesUrl,
		constants.PopularMoviesUrl,
		constants.NowPlayingMoviesUrl,
		constants.TopRatedMoviesUrl,
		constants.UpcomingMoviesUrl,
	}

	ch := make(chan models.MoviesResponse, len(urls))

	for _, url := range urls {
		go requests.FetchMoviesData(url, ch)
	}

	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Println(time.Since(start))

}
