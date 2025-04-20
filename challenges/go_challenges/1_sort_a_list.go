package go_challenges

import (
	"fmt"
	"sort"
)

func SortByInput(nums []int, order string) []int {

	result := make([]int, len(nums))
	copy(result, nums)

	switch order{
	case "asc":
		sort.Ints(result)
	case "desc":
		sort.Sort(sort.Reverse(sort.IntSlice(result)))
	case "none":
		// Return as-is
	default:
		fmt.Println("Invalid order type!")
		return nums
	}
	return result
}
