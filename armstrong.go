//Given an integer N, return the number of digits in N.
package main

import (
	"strconv"
	"math"
)

func armstrong(x int) bool {
	str_x := strconv.Itoa(x)
	len_x := len(str_x)
	sum_ := 0
	for _, val:= range str_x{
		digit, _ := strconv.Atoi(string(val))
		sum_ +=  int(math.Pow(float64(digit), float64(len_x)))
	}
	return sum_ == x
}
