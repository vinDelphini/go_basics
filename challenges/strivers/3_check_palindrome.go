package strivers

func CheckPalindrome(num int) bool {

	if num < 0 {
		return false
	}

	temp_num := num
	out := 0

	for num > 0 {
		digit := num % 10
		out = out * 10 + digit
		num /= 10
	}
	return out == temp_num
}
