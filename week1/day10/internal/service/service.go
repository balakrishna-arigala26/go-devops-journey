package service

import "github.com/balakrishna-arigala26/go-devops-journey/week1/day10/internal/logger"

// Run executes business logic
func Run(log logger.Logger) {
	log.Log("Servise is starting")
	log.Log("Service is running")
	log.Log("Service finished successfully")
}