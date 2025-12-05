package main

import (
	"fmt"
	"sync"
)

func square(num int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	result := num * num 
	ch <- result // send result to channel
}

func main() {
	numbers := []int{2, 4, 6, 8, 10}
	ch := make(chan int)
	var wg sync.WaitGroup

	for _, num := range numbers {
		wg.Add(1)
		go square(num, &wg, ch)
	}

	// goroutine to close channel after all work is done
	go func() {
		wg.Wait()
		close(ch)
	}()

	fmt.Println("Squares:")
	for val := range ch { // recieve from channel
		fmt.Println(val)

	}

	fmt.Println("Done with Day 3!")
}