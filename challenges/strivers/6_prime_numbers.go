// Given an integer N, check whether it is prime or not.
// A prime number is a number that is only divisible by 1 and itself and the total number
// of divisors is 2

package strivers

func PrimeNumbers(num int) bool {
	if num < 2 {
		return false
	}

	divisors := []int{}

	for n:=1; n<=num; n++ {
		if num % n == 0 {
			divisors = append(divisors, n)
		}
	}
	return len(divisors) == 2
}