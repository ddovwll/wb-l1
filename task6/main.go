package main

import (
	"context"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

func conditionStop() {
	defer println("conditionStop stopped")
	println("conditionStop started")

	for {
		num := rand.Intn(5)
		if num == 3 {
			return
		}

		time.Sleep(200 * time.Millisecond)
	}
}

func notificationChanStop(ch <-chan struct{}) {
	defer println("notificationChanStop stopped")
	println("notificationChanStop started")

	for {
		select {
		case <-ch:
			return
		default:
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func contextStop(ctx context.Context) {
	defer println("contextStop stopped")
	println("contextStop started")

	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func goExitStop() {
	defer println("goExitStop stopped")
	println("goExitStop started")

	time.Sleep(400 * time.Millisecond)
	runtime.Goexit()
}

func closedChanStop(ch <-chan struct{}) {
	defer println("closedChanStop stopped")
	println("closedChanStop started")

	for range ch {
		time.Sleep(200 * time.Millisecond)
	}
}

func panicStop() {
	defer println("panicStop stopped")
	println("panicStop started")
	time.Sleep(200 * time.Millisecond)
	panic(nil)
}

func timeAfterStop() {
	defer println("timeAfterStop stopped")
	println("timeAfterStop started")
	select {
	case <-time.After(200 * time.Millisecond):
		return
	default:
		time.Sleep(200 * time.Millisecond)
	}
}

func main() {
	println("starting goroutines")
	defer println("goroutines stopped")

	wg := sync.WaitGroup{}

	wg.Go(conditionStop)

	ch := make(chan struct{})
	wg.Go(func() {
		notificationChanStop(ch)
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Go(func() {
		contextStop(ctx)
	})

	wg.Go(goExitStop)

	wg.Go(func() {
		closedChanStop(ch)
	})

	wg.Go(func() {
		defer func() {
			recover()
		}()
		panicStop()
	})

	wg.Go(timeAfterStop)

	time.Sleep(1 * time.Second)

	ch <- struct{}{}
	close(ch)

	cancel()

	wg.Wait()
}
