package main

import "fmt"

func findIntersect(a []int, b []int) []int {
	set := make(map[int]struct{}, len(a))

	for _, num := range a {
		set[num] = struct{}{}
	}

	intersect := make([]int, 0)
	for _, num := range b {
		if _, ok := set[num]; ok {
			intersect = append(intersect, num)
		}
	}

	return intersect
}

func main() {
	a := []int{1, 2, 3}
	b := []int{2, 3, 4}

	intersect := findIntersect(a, b)
	fmt.Println(intersect)
}
