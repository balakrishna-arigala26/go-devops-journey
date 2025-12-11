package main

import (
	"log"  // logging output
	"sync" // WaitGroup to sync goroutines
	"time" // retry delay
)

func main() {
	log.Println("Starting Day-7: Reliable API Fetcher")

	// List of post IDs to fetch in parallel
	postIDs := []int{1, 2, 3, 4, 5}

	maxRetries := 3                  // retry attempts
	retryDelay := 1 * time.Second    // wait between retries

	var wg sync.WaitGroup

	// Launch one goroutine per post ID
	for _, id := range postIDs {
		wg.Add(1)

		go func(pid int) {
			defer wg.Done()
			fetchPostWithRetry(pid, maxRetries, retryDelay)
		}(id)
	}

	// Wait for all goroutines to complete
	wg.Wait()

	log.Println("Done with Day-7!")
}