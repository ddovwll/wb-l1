package main

import (
	"errors"
	"fmt"
	"slices"
)

func remove[T any](sl []T, index int) ([]T, error) {
	if index >= len(sl) || index < 0 {
		return sl, errors.New("index out of range")
	}

	copy(sl[index:], sl[index+1:])
	sl = slices.Clip(sl[:len(sl)-1])

	return sl, nil
}

func main() {
	sl := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Printf("%v len %v, cap %v\n", sl, len(sl), cap(sl))
	sl, _ = remove(sl, 6)
	fmt.Printf("%v len %v, cap %v\n", sl, len(sl), cap(sl))
}
