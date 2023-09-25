package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type ZincSearchRequest struct {
	SearchType string `json:"search_type"`
	Query      struct {
		Term  string `json:"term"`
		Field string `json:"field"`
	} `json:"query"`
	From       int `json:"from"`
	MaxResults int `json:"max_results"`
}

func SearchEmails(w http.ResponseWriter, r *http.Request) {
	// Load environment variables from .env file
	err := godotenv.Load("../.env") // Adjust the path based on the relative location of your .env file
	if err != nil {
		http.Error(w, "Error loading .env file", http.StatusInternalServerError)
		return
	}

	term := r.URL.Query().Get("term")

	// Prepare the request for ZincSearch
	zincReq := ZincSearchRequest{
		SearchType: "match",
		Query: struct {
			Term  string `json:"term"`
			Field string `json:"field"`
		}{
			Term:  term,
			Field: "_all",
		},
		From:       0,
		MaxResults: 10,
	}

	jsonData, err := json.Marshal(zincReq)
	if err != nil {
		http.Error(w, "Error preparing search request", http.StatusInternalServerError)
		return
	}

	// Fetch environment variables
	zincSearchURL := os.Getenv("ZINC_SEARCH_URL")
	zincSearchUser := os.Getenv("ZINC_SEARCH_USER")
	zincSearchPassword := os.Getenv("ZINC_SEARCH_PASSWORD")

	// Perform the call to ZincSearch
	req, err := http.NewRequest(http.MethodPost, zincSearchURL, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Error creating search request", http.StatusInternalServerError)
		return
	}

	req.SetBasicAuth(zincSearchUser, zincSearchPassword)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error executing search request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read and return the response from ZincSearch
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading search response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
