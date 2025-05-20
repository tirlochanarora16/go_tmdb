package requests

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/tirlochanarora16/go_tmdb/constants"
	"github.com/tirlochanarora16/go_tmdb/models"
)

func makeHttpRequest(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating a new request")
		return nil, err
	}

	token := fmt.Sprintf("Bearer %s", os.Getenv("API_READ_ACCESS_TOKEN"))

	req.Header.Add("Authorization", token)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("error hitting the api")
		return nil, err
	}

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("error decoding the response body")
		return nil, err
	}

	resp.Body.Close()

	return body, nil
}

func FetchMoviesData() {
	var response models.MoviesResponse
	data, err := makeHttpRequest(constants.TopRatedMoviesUrl)

	if err != nil {
		fmt.Println("failed hitting the api")
		return
	}

	err = json.Unmarshal(data, &response)

	if err != nil {
		fmt.Println("failed hitting the api")
		return
	}

	fmt.Println(response.Results[0])
}
