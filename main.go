package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/tirlochanarora16/go_tmdb/requests"
)

type Response struct {
	page int
}

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("error loading .env file")
	}

	const URL string = "https://api.themoviedb.org/3/discover/movie?include_adult=true&include_video=false&language=en-US&page=1&sort_by=popularity.desc"

	var response Response
	data, err := requests.MakeHttpRequest(URL)
	err = json.Unmarshal(data, &response)

	fmt.Println(response.page)

}
