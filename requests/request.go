package requests

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func MakeHttpRequest(url string) ([]byte, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("Error creating a new request")
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("API_READ_ACCESS_TOKEN")))
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
