// Given an integer N, return all divisors of N.
package strivers

func Divisors(num int) []int {

	divisors := []int{}

	for n:= 1; n<= num; n++ {
		if num % n == 0 {
			divisors = append(divisors, n)
		}
	}
	return divisors
}
