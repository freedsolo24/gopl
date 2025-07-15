package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type Movie struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Poster   string `json:"Poster"`
	Response string `json:"Response"`
	Error    string `json:"Error"`
}

const (
	apiBaseURL = "http://www.omdbapi.com/"
	dataDir2   = "posters"
	apiKeyEnv  = "OMDB_API_KEY"
)

func searchMovie(title, year, apiKey string) error {
	movie, err := fetchMovie(title, year, apiKey)
	if err != nil {
		return err
	}
	fmt.Printf("Title: %s\nYear: %s\nPoster: %s\n", movie.Title, movie.Year, movie.Poster)
	return nil
}

func fetchMovie(title, year, apiKey string) (*Movie, error) {
	query := url.Values{}
	query.Set("t", title)
	if year != "" {
		query.Set("y", year)
	}
	query.Set("apikey", apiKey)

	u := apiBaseURL + "?" + query.Encode()
	resp, err := http.Get(u)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch movie data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch movie data: HTTP %d", resp.StatusCode)
	}

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return nil, fmt.Errorf("failed to parse movie data: %w", err)
	}

	if movie.Response != "True" {
		return nil, fmt.Errorf("movie not found: %s", movie.Error)
	}

	return &movie, nil
}

func downloadPoster(title, year, apiKey string) error {
	movie, err := fetchMovie(title, year, apiKey)
	if err != nil {
		return err
	}
	if movie.Poster == "" || movie.Poster == "N/A" {
		return fmt.Errorf("no poster available for %s (%s)", movie.Title, movie.Year)
	}

	if err := os.MkdirAll(dataDir2, 0755); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", dataDir2, err)
	}

	fileName := fmt.Sprintf("%s_%s%s", sanitizeFileName(movie.Title), movie.Year, filepath.Ext(movie.Poster))
	filePath := filepath.Join(dataDir2, fileName)

	resp, err := http.Get(movie.Poster)
	if err != nil {
		return fmt.Errorf("failed to download poster from %s: %w", movie.Poster, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download poster: HTTP %d", resp.StatusCode)
	}

	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %s: %w", filePath, err)
	}
	defer file.Close()

	// Writes the binary image data (e.g., JPEG or PNG) to the file at filePath
	if _, err := io.Copy(file, resp.Body); err != nil {
		return fmt.Errorf("failed to save poster to %s: %w", filePath, err)
	}

	fmt.Printf("Downloaded poster for %s (%s) to %s\n", movie.Title, movie.Year, filePath)
	return nil
}

func sanitizeFileName(name string) string {
	name = strings.ReplaceAll(name, "/", "_")
	name = strings.ReplaceAll(name, "\\", "_")
	name = strings.ReplaceAll(name, ":", "_")
	name = strings.ReplaceAll(name, " ", "_")
	return name
}
