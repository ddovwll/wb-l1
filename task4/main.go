package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"runtime"
	"syscall"
	"time"
)

func writer(ctx context.Context, ch chan int) {
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
	workersCount := flag.Int("workers", runtime.GOMAXPROCS(0), "Number of workers")

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())

	ch := make(chan int)

	for i := 0; i < *workersCount; i++ {
		go reader(ch)
	}

	go func() {
		<-signals
		cancel()
	}()

	writer(ctx, ch)
}
