package main

import (
	"fmt"
	// "go_tutorials/main_concepts"
	// "go_tutorials/core_concepts"
	ch "go_tutorials/challenges/go_challenges"
	// st "go_tutorials/challenges/strivers"
)

func main(){
	fmt.Println("Hello World!")
	// fmt.Println(HideCardNumber(1234567812345678))
	// nums := []int{5, 3, 1, 4, 2}
	// fmt.Println(ch.SortList(nums))
	// fmt.Println(st.CountDigits2(123))
	// fmt.Println(st.PrimeNumbers(23))
	fmt.Println(ch.FilterIntegers([]interface{}{"one", 2, 3, "four", 5, 0, true}))
	fmt.Println(ch.RepeatChar("123a!"))
	fmt.Println(ch.ConvertToUpper("goodbye"))
	fmt.Println(ch.AddDots3("skills"))

}
