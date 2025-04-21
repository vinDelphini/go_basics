package strivers


func ReverseInt(num int) int {
	out := 0
	sign := 1
	if num < 0 {
		num = -num
		sign = -1
	}
	for num > 0 {
		digit := num % 10
		out = out * 10 + digit
		num /= 10
	}
	return out * sign
}
