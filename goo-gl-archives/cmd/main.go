package main

import (
	"goo-gl-archives/internal/database"
	"goo-gl-archives/internal/url_processor"
	"goo-gl-archives/pkg/utils"
	"log"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
)

func main() {
	const numParallelExecutions = 1000 // Number of parallel executions
	var wg sync.WaitGroup
	resultsChan := make(chan url_processor.Link, numParallelExecutions) // Channel to collect results

	logger := utils.SetupLogger("logfile.log")

	// Initialize the database
	db, err := database.InitializeDatabase("../db/archives.db")
	if err != nil {
		logger.Fatal(err)
	}

	// Create and start the progress bar
	bar := pb.StartNew(numParallelExecutions)
	bar.SetRefreshRate(100 * time.Millisecond) // Adjust as needed

	// Start parallel processing
	for i := 0; i < numParallelExecutions; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			processRequest(resultsChan, logger) // Assuming processRequest does not need a WaitGroup pointer
			bar.Increment()
		}()
	}

	// Wait for all goroutines to complete and close the results channel
	go func() {
		wg.Wait()
		close(resultsChan)
		bar.Finish()
	}()

	// Collect results
	results := collectResults(resultsChan)

	// Insert or update results in the database
	if err := database.StoreLinks(db, results, logger); err != nil {
		logger.Fatal(err)
	}
}

// processRequest performs a URL processing request and sends the result to the results channel.
func processRequest(resultsChan chan<- url_processor.Link, logger *log.Logger) {
	result, err := url_processor.ProcessRequest()
	if err != nil {
		logger.Println(err)
		return
	}

	// Send result to the channel if it is not an empty Link
	if result != (url_processor.Link{}) {
		resultsChan <- result
	}
}

// collectResults collects all results from the channel into a slice.
func collectResults(resultsChan <-chan url_processor.Link) []url_processor.Link {
	var results []url_processor.Link
	for result := range resultsChan {
		results = append(results, result)
	}
	return results
}
