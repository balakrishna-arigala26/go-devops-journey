package logger

import "fmt"

// Logger defines a logging contract
type Logger interface {
	Log(message string)
}

// ConsoleLogger logs to stdout
type ConsoleLogger struct{}

func (c ConsoleLogger) Log(message string) {
	fmt.Println("[LOG]", message)
}