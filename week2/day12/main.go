package main	

import (
	"fmt"
	"os"
	"flag"
)

func main() {
	// 1. Read environment variables
	envPort := os.Getenv("APP_PORT")
	envEnv := os.Getenv("APP_ENV")

	// 2. Define flags (default empty)
	portFlag := flag.String("port", "", "port number")
	envFlag := flag.String("env", "", "Environement name")

	// 3. Parse flags
	flag.Parse()

	// 4. Decide final values (priority logic)
	port := "9090"	// default
	if envPort != "" {
		port = envPort	// env overrides default
	}
	if *portFlag != "" {
		port = *portFlag	// flag overrides env
	}


	env := "development" // default
	if envEnv != "" {
		env = envEnv
	}
	if *envFlag != "" {
		env = *envFlag
	}

	// 5. Print final configuration
	fmt.Println("Final Apllication Configuaration")
	fmt.Println("--------------------------------")
	fmt.Println("Port :", port)
	fmt.Println("Env  :", env)

}