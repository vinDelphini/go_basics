package go_challenges

import "strings"

// easy approach
func RepeatChar(word string) string {
	out := ""
	for _, letter:= range word {
		out += string(letter) + string(letter)
	}
	return out
}

// optimized approach
func RepeatChar2(word string) string {
	var builder strings.Builder
	builder.Grow(len(word) * 2)

	for _, letter:= range word {
		builder.WriteRune(letter)
		builder.WriteRune(letter)
	}
	return builder.String()
}
