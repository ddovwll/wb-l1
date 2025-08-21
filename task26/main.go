package main

import "unicode"

func checkUniqueness(str string) bool {
	m := make(map[rune]struct{})
	for _, v := range str {
		if _, ok := m[unicode.ToLower(v)]; ok {
			return false
		}
		m[v] = struct{}{}
	}

	return true
}

func main() {
	s1 := "abcd"
	s2 := "abCdefAaf"
	s3 := "aabcd"

	println(checkUniqueness(s1))
	println(checkUniqueness(s2))
	println(checkUniqueness(s3))
}
