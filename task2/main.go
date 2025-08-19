package main

import (
	"fmt"
	"sync"
)

func square(num int) {
	fmt.Println(num * num)
}

func main() {
	nums := [5]int{2, 4, 6, 8, 10}
	wg := sync.WaitGroup{}
	for _, num := range nums {
		wg.Go(func() {
			square(num)
		})
	}
	wg.Wait()
}
