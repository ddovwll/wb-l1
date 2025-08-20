package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func binSearch[T constraints.Ordered](arr []T, target T) int {
	l := -1
	r := len(arr)
	for l < r-1 {
		m := (l + r) / 2
		if arr[m] < target {
			l = m
		} else {
			r = m
		}
	}

	return r
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binSearch(arr, 8))
}
