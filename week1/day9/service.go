package main

import "fmt"

// startService depends on fetchConfig
func startService() error {
	_, err := fetchConfig("/wrong/path")
	if err != nil {
		// rap the error with context
		return fmt.Errorf("service startup failed: %w", err)
	}

	return nil
}