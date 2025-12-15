package main

import (
	"fmt"
	"os"
)

// FileLogger writes logs to a file
type FileLogger struct {
	FilePath string
}

func (f FileLogger) Log(message string) {
	file, err := os.OpenFile(f.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	defer file.Close()


file.WriteString(message + "\n")

}