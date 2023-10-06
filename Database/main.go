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

func loadEnvAndConfig() (string, string, string, error) {
	if err := godotenv.Load("../.env"); err != nil {
		return "", "", "", err
	}

	return os.Getenv("SEARCH_ENGINE_BULK_API_ENDPOINT"), os.Getenv("SEARCH_ENGINE_USER"), os.Getenv("SEARCH_ENGINE_PASSWORD"), nil
}

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

func readRecordsFromTarball(tarballPath string) (<-chan Record, error) {
	file, err := os.Open(tarballPath)
	if err != nil {
		return nil, err
	}

	out := make(chan Record)
	go func() {
		defer close(out)
		defer file.Close()

		tr := tar.NewReader(file)
		for {
			header, err := tr.Next()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Printf("Error reading tarball: %v", err)
				return
			}

			if header.Typeflag != tar.TypeReg {
				continue
			}

			content, err := io.ReadAll(tr)
			if err != nil {
				log.Printf("Error reading file content from tarball: %v", err)
				return
			}

			out <- Record{Content: string(content)}
		}
	}()

	return out, nil
}

func processRecordsAndSend(searchEngineAPI, searchEngineUser, searchEnginePassword string, records <-chan Record) {
	var batch []Record

	for record := range records {
		batch = append(batch, record)

		if len(batch) >= DefaultBatchSize {
			if err := sendBatchToSearchEngine(batch, searchEngineAPI, searchEngineUser, searchEnginePassword); err != nil {
				log.Printf("Error sending batch: %s\n", err)
			}
			batch = nil // Reset the batch
		}
	}

	if len(batch) > 0 {
		if err := sendBatchToSearchEngine(batch, searchEngineAPI, searchEngineUser, searchEnginePassword); err != nil {
			log.Printf("Error sending final batch: %s\n", err)
		}
	}
}

func main() {
	searchEngineAPI, searchEngineUser, searchEnginePassword, err := loadEnvAndConfig()
	if err != nil {
		log.Fatalf("Error during initialization: %v", err)
		return
	}

	if len(os.Args) < 2 {
		log.Fatal("Usage: ./indexer <path_to_tarball>")
		return
	}

	tarballPath := os.Args[1]
	records, err := readRecordsFromTarball(tarballPath)
	if err != nil {
		log.Fatalf("Error reading records from tarball: %v", err)
		return
	}

	processRecordsAndSend(searchEngineAPI, searchEngineUser, searchEnginePassword, records)
	log.Println("Finished indexing records!")
}
