package main

import "fmt"

func toSet(strings []string) []string {
	set := make(map[string]struct{})

	for _, str := range strings {
		set[str] = struct{}{}
	}

	unique := make([]string, 0, len(set))
	for num := range set {
		unique = append(unique, num)
	}

	return unique
}

func main() {
	strings := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Printf("%v\n", toSet(strings))
}
