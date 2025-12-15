package main

// Logger is an interface that defines behavior.
// Any struct with a Log(mesasage string) method automatically IMPLEMENTS this interface.
type Logger interface {
	Log(message string)
}

