package main

import "fmt"

// ConsoleLogger logs message to terminal
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("[CONSOLE]" + message)
}