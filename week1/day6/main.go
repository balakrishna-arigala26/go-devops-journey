package main

import (
	"encoding/json" // Helps decode JSON data into Go structs
	"fmt"           // For printing messages
	"net/http"      // For making HTTP GET requests
	"sync"          // for WaitGroup (syncing goroutines)
)

// Post represents the structure of a post returned by the API.
// The JSON tags (json:"id") tell Go which JSON fields map to which struct fields.
type Post struct {
	ID      int  `json:"id"`
	Title string `json:"title"`

}

// fetchPost fetches a post with a given ID from the API.
// It runs as a goroutine, so multiple posts are fetched in parallel.
func fetchPost(id int, wg *sync.WaitGroup) {

	// Tell the WaitGroup that this goroutine is done when the function exists.
	defer wg.Done()
	
	// Create the API URL dynamically using the post ID.
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)

	// Send an HTTP GET request ti the API.
	resp, err := http.Get(url)
	if err != nil {
		// If there's an error (network issue, invalid RL), log the issue and exit this goroutine.
		fmt.Printf("Error fetching post %d: %v\n", id, err)
		return
	}

	// Ensure the response body is closed after we're done reading it.
	defer resp.Body.Close()

	// Create a variable to store the decoded JSON response.
	var post Post

	// Decode JSON from the response body into the 'post' variable.
	if err := json.NewDecoder(resp.Body).Decode(&post); err != nil {
		fmt.Printf("Error decoding post %d: %v\n", id, err)
		return
	}

	// Print the result nicely formatted.
	fmt.Printf("Post %d → %s\n", post.ID, post.Title)
}

func main() {

	// Create a WaitGroup to sync all goroutines.
	var wg sync.WaitGroup

	// List of post IDs we want to fetch from the API.
	// Each of these will be processed in separate goroutines.
	postIDs := []int{1, 2, 3, 4, 5}

	// Loop through all IDs.
	for _, id := range postIDs {

		// Increase WaitGroup counter by 1 (one goroutine starting).
		wg.Add(1)

		// Launch fetchPost() as a goroutine.
		// Passing &wg  so that fetchPost can call wg.Done().
		go fetchPost(id, &wg)
	}

	// Wait until ALL gorputines have finished executing.
	wg.Wait()

	// After everything is done, print completion message.
	fmt.Println("Done with Day 6!")



}
