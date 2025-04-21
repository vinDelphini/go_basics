// Given an integer N, return true it is an Armstrong number otherwise return false.

package strivers

import (
	"strconv"
	"math"
)

func Armstrong(num int) bool {

	temp_num := num
	out := 0
	power := len(strconv.Itoa(num))

	for num > 0 {
		digit := num % 10
		out += int(math.Pow(float64(digit), float64(power)))
		num /= 10
	}
	return temp_num == out
}
