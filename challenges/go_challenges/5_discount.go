package go_challenges


func Discount(total int, discount int) int {
	return int(float64(total) * float64(100 - discount) / 100.0)
}
