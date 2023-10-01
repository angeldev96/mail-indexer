package main

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// DefaultBatchSize is the number of records sent in each batch.
const DefaultBatchSize = 1000

// Record represents the structure of a record to be indexed.
type Record struct {
	Content string `json:"content"`
}

// BulkPayload is the payload structure for the search engine.
type BulkPayload struct {
	Index   string   `json:"index"`
	Records []Record `json:"records"`
}

// sendBatchToSearchEngine sends a batch of records to the search engine.
func sendBatchToSearchEngine(batch []Record, searchEngineURL, searchEngineUser, searchEnginePassword string) error {
	payload := BulkPayload{
		Index:   "generic_index",
		Records: batch,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, searchEngineURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(searchEngineUser, searchEnginePassword)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error sending batch: %s", body)
	}

	return nil
}

func main() {
	// Load environment variables.
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("Error loading .env file.")
		return
	}

	searchEngineAPI := os.Getenv("SEARCH_ENGINE_BULK_API_ENDPOINT")
	searchEngineUser := os.Getenv("SEARCH_ENGINE_USER")
	searchEnginePassword := os.Getenv("SEARCH_ENGINE_PASSWORD")

	go func() {
		if err := http.ListenAndServe(":6060", nil); err != nil {
			log.Printf("Error starting profiling server: %v", err)
		}
	}()

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./indexer <path_to_tarball>")
		return
	}

	tarballPath := os.Args[1]
	file, err := os.Open(tarballPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tr := tar.NewReader(file)
	var batch []Record

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if header.Typeflag != tar.TypeReg {
			continue
		}

		content, err := io.ReadAll(tr)
		if err != nil {
			panic(err)
		}

		batch = append(batch, Record{Content: string(content)})

		if len(batch) >= DefaultBatchSize {
			if err := sendBatchToSearchEngine(batch, searchEngineAPI, searchEngineUser, searchEnginePassword); err != nil {
				fmt.Printf("Error sending batch: %s\n", err)
			}
			batch = nil // Reset the batch
		}
	}

	// Send any remaining records.
	if len(batch) > 0 {
		if err := sendBatchToSearchEngine(batch, searchEngineAPI, searchEngineUser, searchEnginePassword); err != nil {
			fmt.Printf("Error sending final batch: %s\n", err)
		}
	}

	fmt.Println("Finished indexing records!")
}
