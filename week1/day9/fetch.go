package main

import (
	"errors"
	"fmt"
)

// fetchConfig simulates reading a config file
func fetchConfig(path string) (string, error) {
	if path == "" {
		return "", errors.New("config path cannot be empty")
	}

	if path != "etc/app/config.yaml" {
		return "", fmt.Errorf("config file not found at %s", path)
	}

	return "config-loaded-successfully", nil
}