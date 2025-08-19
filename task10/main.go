package main

import (
	"fmt"
)

func groupTemp(temps []float32) map[int][]float32 {
	groups := make(map[int][]float32)

	for _, temp := range temps {
		key := int(temp) / 10 * 10
		groups[key] = append(groups[key], temp)
	}

	return groups
}

func main() {
	temp := []float32{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	groups := groupTemp(temp)

	for key, group := range groups {
		fmt.Printf("%v\t%v\n", key, group)
	}
}
