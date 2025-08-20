package main

import (
	"fmt"
	"sync"
)

type MutexCounter struct {
	mutex sync.Mutex
	count uint
}

func NewMutexCounter() *MutexCounter {
	return &MutexCounter{
		mutex: sync.Mutex{},
		count: 0,
	}
}

func (m *MutexCounter) Inc() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.count++
}

func (m *MutexCounter) Count() uint {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.count
}

func main() {
	counter := NewMutexCounter()
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Go(func() {
			for i := 0; i < 100; i++ {
				counter.Inc()
			}
		})
	}

	wg.Wait()
	fmt.Println(counter.Count())
}
