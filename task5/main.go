package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

func writer(ctx context.Context, ch chan<- int) {
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		case ch <- rand.Intn(100):
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func reader(ch <-chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func main() {
	duration := flag.Int("duration", 5, "program duration in seconds")
	flag.Parse()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(*duration)*time.Second)
	defer cancel()
	ch := make(chan int)

	go reader(ch)
	go writer(ctx, ch)

	<-ctx.Done()
	println("time is up, Honey)")
}
