package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "Guest", "Name of the user")
	flag.Parse()

	fmt.Printf("Hello %s, welcome to Go CLI tool!\n", *name)
}