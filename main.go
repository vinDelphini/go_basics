package main

import (
	"fmt"
	// "go_tutorials/main_concepts"
	// "go_tutorials/core_concepts"
	ch "go_tutorials/challenges/go_challenges"
)

func main(){
	fmt.Println("Hello World!")
	// fmt.Println(HideCardNumber(1234567812345678))
	fmt.Println(ch.HideCreditCardNumber(1234567812345678))
}
