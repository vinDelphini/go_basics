package go_challenges

import (
	"strings"
	"strconv"
)

func HideCreditCardNumber(num int) string {
	ast := strings.Repeat("*", 12)
	last_4_digits := strconv.Itoa(num % 10000)
	return ast + last_4_digits
}
