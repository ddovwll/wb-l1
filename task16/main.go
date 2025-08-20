package main

import (
	"fmt"
	"math/rand"
)

func RecursionQuickSort(nums []int) {
	recursionQuickSort(nums, 0, len(nums)-1)
}

func recursionQuickSort(nums []int, l, r int) {
	if l < r {
		q := partition(nums, l, r)
		recursionQuickSort(nums, l, q)
		recursionQuickSort(nums, q+1, r)
	}
}

func partition(nums []int, l, r int) int {
	v := nums[(l+r)/2]
	i := l
	j := r
	for i <= j {
		for nums[i] < v {
			i++
		}
		for nums[j] > v {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}

	return j
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	rand.Shuffle(len(arr), func(i, j int) {
		arr[i], arr[j] = arr[j], arr[i]
	})
	fmt.Println("unsorted", arr)
	RecursionQuickSort(arr)
	fmt.Println("recursion sorted", arr)
}
