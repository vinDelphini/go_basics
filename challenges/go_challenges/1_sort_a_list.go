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


// just sort a list

func SortList(nums []int) ([]int, []int) {

	asc := make([]int, len(nums))
	copy(asc, nums)

	desc:= make([]int, len(nums))
	copy(desc, nums)

	sort.Ints(asc)
	sort.Sort(sort.Reverse(sort.IntSlice(desc)))
	return asc, desc
}
