package go_challenges

func FilterIntegers(input []interface{}) []int {
	var result []int
	for _, item:= range input {
		if val, ok:= item.(int); ok {
			result = append(result, val)
		}
	}
	return result
}
