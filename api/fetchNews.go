package api

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mmcdole/gofeed"
)

// Configuration represents the configuration for the news API.
type Configuration struct {
	MacrumorsAPIURL string `json:"MacrumorsAPIURL"` // Change to a single string
}

// FetchNews fetches news articles from the news API.
func FetchNews() ([]*gofeed.Item, error) {
	fp := gofeed.NewParser()
	// First, try to get the directory of the executable
	executablePath, err := os.Executable()
	if err != nil {
		return nil, fmt.Errorf("failed to get executable path: %w", err)
	}
	executableDir := filepath.Dir(executablePath)

	// Try to open the configuration file in the executable directory
	configPath := filepath.Join(executableDir, "conf.json")
	file, err := os.Open(configPath)
	if err != nil {
		// If it fails, fall back to the current working directory
		currentDir, err := os.Getwd()
		if err != nil {
			return nil, fmt.Errorf("failed to get current working directory: %w", err)
		}
		configPath = filepath.Join(currentDir, "conf.json")
		file, err = os.Open(configPath)
		if err != nil {
			return nil, fmt.Errorf("failed to open configuration file: %w", err)
		}
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		return nil, fmt.Errorf("failed to decode configuration: %w", err)
	}

	if len(configuration.MacrumorsAPIURL) == 0 {
		return nil, fmt.Errorf("no Macrumors API URL found in configuration")
	}

	url := configuration.MacrumorsAPIURL

	feed, _ := fp.ParseURL(url)
	return feed.Items, nil
}
