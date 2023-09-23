package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
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
	term := r.URL.Query().Get("term")

	// Preparar la solicitud para ZincSearch
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
		MaxResults: 100,
	}

	jsonData, err := json.Marshal(zincReq)
	if err != nil {
		http.Error(w, "Error preparing search request", http.StatusInternalServerError)
		return
	}

	// Realizar la llamada a ZincSearch
	req, err := http.NewRequest("POST", "http://localhost:4080/api/enron_emails/_search", bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Error creating search request", http.StatusInternalServerError)
		return
	}

	req.SetBasicAuth("admin", "maiden")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Error executing search request", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Leer y devolver la respuesta de ZincSearch
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Error reading search response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(body)
}
