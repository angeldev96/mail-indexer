package main

import (
	"archive/tar"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/joho/godotenv"
)

const defaultBatchSize = 1000 // Default number of emails per batch

type Email struct {
	Content string `json:"content"`
}

type BulkV2Payload struct {
	Index   string  `json:"index"`
	Records []Email `json:"records"`
}

func sendBatchToZincSearch(batch []Email, zincSearchURL, zincSearchUser, zincSearchPassword string) error {
	payload := BulkV2Payload{
		Index:   "enron_emails",
		Records: batch,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, zincSearchURL, bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth(zincSearchUser, zincSearchPassword)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error sending batch: %s", string(body))
	}

	return nil
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load("../.env") //
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	zincSearchAPI := os.Getenv("ZINC_SEARCH_BULKV2_API_ENDPOINT")
	zincSearchUser := os.Getenv("ZINC_SEARCH_USER")
	zincSearchPassword := os.Getenv("ZINC_SEARCH_PASSWORD")

	go func() {
		err := http.ListenAndServe(":6060", nil)
		if err != nil {
			log.Printf("Error starting profiling server: %v", err)
		}
	}()

	if len(os.Args) < 2 {
		fmt.Println("Usage: ./indexer <path_to_tarball>")
		return
	}

	tarball := os.Args[1]
	file, err := os.Open(tarball)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	tr := tar.NewReader(file)
	var batch []Email

	for {
		header, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		if header.Typeflag == tar.TypeReg {
			content, err := io.ReadAll(tr)
			if err != nil {
				panic(err)
			}

			email := Email{Content: string(content)}
			batch = append(batch, email)

			if len(batch) >= defaultBatchSize {
				err := sendBatchToZincSearch(batch, zincSearchAPI, zincSearchUser, zincSearchPassword)
				if err != nil {
					fmt.Printf("Error sending batch: %s\n", err)
				}
				batch = nil // Reset the batch
			}
		}
	}

	// Send any remaining emails in the last batch
	if len(batch) > 0 {
		err := sendBatchToZincSearch(batch, zincSearchAPI, zincSearchUser, zincSearchPassword)
		if err != nil {
			fmt.Printf("Error sending final batch: %s\n", err)
		}
	}

	fmt.Println("Finished indexing emails!")
}
