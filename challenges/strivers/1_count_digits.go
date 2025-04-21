package strivers

import "strconv"

// best
func CountDigits(num int) int {

	if num < 0 {
		num = -num
	}

	if num < 10 {
		return 1
	}

	count := 0
	for num > 0 {
		num /= 10
		count++
	}
	return count
}

// Option 2: Convert to String [not ideal]

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func CountDigits2(num int) int {
	return len(strconv.Itoa(abs(num)))
}

// Option 3: Recursive (Not preferred in GO due to lack of tail call optimization)

func CountDigits3(num int) int {
	num = num/10
	return 1 + CountDigits(num)
}
