package main

import (
	"fmt"
	"maps"
	"sync"
)

type ConcurrentMap[K comparable, V any] struct {
	mutex sync.Mutex
	m     map[K]V
}

func NewConcurrentMap[K comparable, V any]() *ConcurrentMap[K, V] {
	return &ConcurrentMap[K, V]{
		mutex: sync.Mutex{},
		m:     make(map[K]V),
	}
}

func (cm *ConcurrentMap[K, V]) Set(key K, value V) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.m[key] = value
}

func (cm *ConcurrentMap[K, V]) Get(key K) (V, bool) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	value, ok := cm.m[key]
	return value, ok
}

func (cm *ConcurrentMap[K, V]) Del(key K) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	delete(cm.m, key)
}

func (cm *ConcurrentMap[K, V]) Len() int {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	return len(cm.m)
}

func (cm *ConcurrentMap[K, V]) All() map[K]V {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	return maps.Clone(cm.m)
}

func main() {
	cm := NewConcurrentMap[int, int]()
	sm := sync.Map{}
	wg := sync.WaitGroup{}

	for i := 0; i <= 10; i++ {
		wg.Go(func() {
			for i := 0; i <= 10; i++ {
				cm.Set(i, i)
			}
		})
	}

	for i := 0; i <= 10; i++ {
		wg.Go(func() {
			for i := 0; i <= 10; i++ {
				sm.Store(i, i)
			}
		})
	}

	wg.Wait()

	println("own Concurrent Map")
	for k, v := range cm.All() {
		fmt.Printf("%v:%v\n", k, v)
	}
	println()

	println("sync.Map")
	sm.Range(func(k, v any) bool {
		fmt.Printf("%v:%v\n", k, v)
		return true
	})
}
