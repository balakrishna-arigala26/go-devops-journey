package main

import (
  "os"
  "fmt"
)

func main() {
  // Read environment variables
  appName := os.Getenv("APP_NAME")
  port := os.Getenv("APP_PORT")
  env := os.Getenv("APP_ENV")

// Provide defaults if not set
if appName == "" {
  appName = "go-devops-app"
}

if port == "" {
  port = "9090"
}

if env == "" {
  env = "development"
}

// Print configuration
fmt.Println("Application Configuration")
fmt.Println("-------------------------")
fmt.Println("App Name :", appName)
fmt.Println("Port     :", port)
fmt.Println("Env      :", env)

}
 
