package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

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

func loadEnv() error {
	return godotenv.Load("../.env")
}

func prepareSearchRequest(r *http.Request) (*SearchEngineRequest, error) {
	term := r.URL.Query().Get("term")
	pageStr := r.URL.Query().Get("page")
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	const resultsPerPage = 10
	from := (page - 1) * resultsPerPage

	return &SearchEngineRequest{
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
		From:       from,
		MaxResults: resultsPerPage,
	}, nil
}

func executeSearchRequest(searchReq *SearchEngineRequest) (*http.Response, error) {
	jsonData, err := json.Marshal(searchReq)
	if err != nil {
		return nil, err
	}

	searchEngineURL := os.Getenv("SEARCH_ENGINE_URL")
	searchEngineUser := os.Getenv("SEARCH_ENGINE_USER")
	searchEnginePassword := os.Getenv("SEARCH_ENGINE_PASSWORD")

	req, err := http.NewRequest(http.MethodPost, searchEngineURL, bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(searchEngineUser, searchEnginePassword)

	client := &http.Client{}
	return client.Do(req)
}

func SearchRecords(w http.ResponseWriter, r *http.Request) {
	if err := loadEnv(); err != nil {
		http.Error(w, "Error loading .env file", http.StatusInternalServerError)
		return
	}

	searchReq, err := prepareSearchRequest(r)
	if err != nil {
		http.Error(w, "Error preparing search request", http.StatusInternalServerError)
		return
	}

	resp, err := executeSearchRequest(searchReq)
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
