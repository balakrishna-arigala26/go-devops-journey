package main

import (
	"errors"
	"fmt"
)

func main() {
	err := startService()
	if err != nil {
		fmt.Println("ERROR:", err)

		// check root cause
		if errors.Is(err, errors.New("config path canot be empty")) {
			fmt.Println("Root cause: missing config path")
		}

		fmt.Println("Shutting down application safely")
		return
	}

	fmt.Println("service started successfully")
}