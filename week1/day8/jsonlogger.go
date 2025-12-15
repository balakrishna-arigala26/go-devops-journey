package main

import (
	"encoding/json"
	"fmt"
)

type JSONLogger struct{}

func (j JSONLogger) Log(message string) {
	jsonData, _ := json.Marshal(map[string]string{
		"log": message,
	})
	fmt.Println(string(jsonData))
}