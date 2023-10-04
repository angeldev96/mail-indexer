package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// SearchEngineRequest represents the structure of a search request for a generic search engine.
type SearchEngineRequest struct {
	SearchType string `json:"search_type"`
	Query      struct {
		Term  string `json:"term"`
		Field string `json:"field"`
	} `json:"query"`
	Aggregations map[string]interface{} `json:"aggregations,omitempty"`
	From         int                    `json:"from"`
	MaxResults   int                    `json:"max_results"`
}

// SearchRecords handles search queries and returns search results.
func SearchRecords(w http.ResponseWriter, r *http.Request) {
	if err := godotenv.Load("../.env"); err != nil {
		http.Error(w, "Error loading .env file", http.StatusInternalServerError)
		return
	}

	term := r.URL.Query().Get("term")

	searchReq := SearchEngineRequest{
		SearchType: "match",
		Query: struct {
			Term  string `json:"term"`
			Field string `json:"field"`
		}{
			Term:  term,
			Field: "_all",
		},
		Aggregations: map[string]interface{}{
			"count_by_term": map[string]interface{}{
				"terms": map[string]interface{}{
					"field":   "_all",
					"include": term,
				},
			},
		},
		From:       0,
		MaxResults: 10,
	}

	jsonData, err := json.Marshal(searchReq)
	if err != nil {
		http.Error(w, "Error preparing search request", http.StatusInternalServerError)
		return
	}

	searchEngineURL := os.Getenv("SEARCH_ENGINE_URL")
	searchEngineUser := os.Getenv("SEARCH_ENGINE_USER")
	searchEnginePassword := os.Getenv("SEARCH_ENGINE_PASSWORD")

	req, err := http.NewRequest(http.MethodPost, searchEngineURL, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Error creating search request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(searchEngineUser, searchEnginePassword)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error executing search request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading search response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(body); err != nil {
		log.Printf("Could not write response: %v", err)
	}
}
