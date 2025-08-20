package main

func reverse(runes []rune) {
	n := len(runes)
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
}

func reverseWords(s string) string {
	runes := []rune(s)
	reverse(runes)

	wordStart := 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == ' ' {
			reverse(runes[wordStart:i])
			wordStart = i + 1
		} else if i == len(runes)-1 {
			reverse(runes[wordStart : i+1])
		}
	}

	return string(runes)
}

func main() {
	words := "snow dog sun"
	println(reverseWords(words))
}
