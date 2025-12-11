package main

// Post represents the structure of JSON returned by the API.
//
// Example API response:
// {
//	"id":1,
//	"title": "some title"
// }
// Go will automatically map the JSON keys to these fields.
type Post struct {
	ID	int      `json:"id"`
	Title string `json:"title"`
}