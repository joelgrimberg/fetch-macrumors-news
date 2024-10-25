package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Article represents a news article.
type Article struct {
	Title string `json:"title"`
}

// Response represents a response from the news API.
type Response struct {
	Articles []Article `json:"articles"`
}

// Configuration represents the configuration for the news API.
type Configuration struct {
	NewsAPIURL string `json:"newsAPIURL"` // Change to a single string
}

// FetchNews fetches news articles from the news API.
func FetchNews() ([]Article, error) {
	file, err := os.Open("conf.json")
	if err != nil {
		return nil, fmt.Errorf("failed to open configuration file: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		return nil, fmt.Errorf("failed to decode configuration: %w", err)
	}

	if len(configuration.NewsAPIURL) == 0 {
		return nil, fmt.Errorf("no URLs found in configuration")
	}

	url := configuration.NewsAPIURL

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error executing request: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Unmarshal the JSON response
	var data Response
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return data.Articles, nil
}
