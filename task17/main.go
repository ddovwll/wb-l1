package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func binSearch[T constraints.Ordered](arr []T, target T) int {
	l := 0
	r := len(arr) - 1
	for l <= r {
		m := l + (r-l)/2
		if arr[m] < target {
			l = m + 1
		} else if arr[m] > target {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}

func main() {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(binSearch(arr, 7))
}
