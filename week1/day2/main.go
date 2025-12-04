package main

import (
	"fmt"
	"time"
)

func printCount(prefix string) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("%s : %d\n", prefix, i)
        time.Sleep(200 * time.Millisecond)
    }
}

func main() {
    go printCount("Goroutine-1")
    go printCount("Goroutine-2")

    fmt.Println("Starting two goroutines...")

    // wait for goroutines to finish
    time.Sleep(2 * time.Second)

    fmt.Println("Done with Day 2!")
}