package main

import (
	"fmt"
	"time"
)

func chanSleep(duration time.Duration) {
	<-time.After(duration)
}

func loopSleep(duration time.Duration) {
	end := time.Now().Add(duration)
	for time.Now().Before(end) {

	}
}

func main() {
	t := time.Now()
	chanSleep(5 * time.Second)
	fmt.Printf("chanSleep: %v\n", time.Now().Sub(t))
	t = time.Now()
	loopSleep(3 * time.Second)
	fmt.Printf("loopSleep: %v\n", time.Now().Sub(t))
}
