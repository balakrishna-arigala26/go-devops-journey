package main

import (
	"encoding/json" // Decode JSON from API
	"fmt"           // Print output nicely
	"log"           // Logging retries, errors, info
	"net/http"      // HTTP client
	"time"          // Timeout + retry delay
)

// fetchPostWithRetry fetches a post from an API.
// It retries multiple times if something  goes wrong.
// id               → Post ID (1..5)
// maxRetries       → how many attempts before giving up
// delay            → how long to sleep between retries
func fetchPostWithRetry(id, maxRetries int, delay time.Duration) {

	// Create an HTTP client with timeout so the program NEVER hangs forever.
	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)

	// Attempt the request multiple times.
	for attempt := 1; attempt <= maxRetries; attempt++ {

		log.Printf("Post %d: attempt %d to fetch %s", id, attempt, url)

		resp, err := client.Get(url)
		if err != nil {
			// Network error → retry later
			log.Printf("Post %d: request error: %v", id, err)
		} else {
			// Decode JSON directly into struct
			var post Post
			decodeErr := json.NewDecoder(resp.Body).Decode(&post)
			resp.Body.Close()

			if decodeErr == nil {
				// SUCCESS → print and return
				fmt.Printf("Post %d → %s\n", post.ID, post.Title)
				return
			}

			log.Printf("Post %d: JSON decode error: %v", id, decodeErr)
		}

		// If not the last attempt → sleep before retrying
		if attempt < maxRetries {
			log.Printf("Pst %d: retrying in %v ...", id, delay)
			time.Sleep(delay)
		}
	}

	// If we reach here → all  retries failed
	log.Printf("Post %d: failed after %d attempts", id, maxRetries)
}