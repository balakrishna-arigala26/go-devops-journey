package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

type User struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

func main() {
	file := flag.String("file", "data.json", "JSON foile to read")
	flag.Parse()

	jsonData, err := os.ReadFile(*file)
	if err != nil {
		fmt.Println("Error reaing file:", err)
		os.Exit(1)
	}

	var users []User
	err = json.Unmarshal(jsonData, &users)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}

	fmt.Println("Users Information:")
	for _, user := range users {
		fmt.Printf("Name: %s | Role: %s\n", user.Name, user.Role)
	}

}


