package main

import (
	"sync"
)

func sendToNumsChan(numsChan chan<- int, nums []int) {
	defer close(numsChan)

	for _, num := range nums {
		numsChan <- num
	}
}

func readFromNumsChan(numsChan <-chan int, squareChan chan<- int) {
	defer close(squareChan)

	for num := range numsChan {
		squareChan <- num * num
	}
}

func readFromSquareChan(squareChan <-chan int) {
	for square := range squareChan {
		println(square)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numsChan := make(chan int)
	squareChan := make(chan int)

	wg := sync.WaitGroup{}

	wg.Go(func() {
		sendToNumsChan(numsChan, nums)
	})

	wg.Go(func() {
		readFromNumsChan(numsChan, squareChan)
	})

	wg.Go(func() {
		readFromSquareChan(squareChan)
	})

	wg.Wait()
}
