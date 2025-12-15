package main

func main() {
	// Create different loggers
	console := ConsoleLogger{}
	file := FileLogger{FilePath: "logs.txt"}
	json := JSONLogger{}

	// A slice of Logger INTERFACE
	loggers := []Logger{console, file, json}

	// One function call = different behaviors depending on the struct
	for _, logger := range loggers {
		logger.Log("Day-8: Interface based logging!")
	}
}