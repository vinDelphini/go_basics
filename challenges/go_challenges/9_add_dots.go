package go_challenges

import (
	"fmt"
	"strings"
)

// brute force
func AddDots(word string) string {
	new_output := ""
	for i, val:= range word {
		fmt.Println(i)
		if i == len(word) - 1 {
			new_output += string(val)
		}else{
			new_output += string(val) + "."
		}
	}
	return new_output
}

// split & join
func AddDots2(word string) string {
	letters := strings.Split(word, "")
	return strings.Join(letters, ".")
}


// strings builder
func AddDots3(word string) string {
	var builder strings.Builder
	builder.Grow(len(word) * 2 - 1)

	for i, val:= range word {
		if i == len(word) - 1 {
			builder.WriteRune(val)
		}else{
			builder.WriteRune(val)
			builder.WriteRune('.')
		}
	}
	return builder.String()
}
