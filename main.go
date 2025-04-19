package main

import (
	"fmt"
)

func main(){
	fmt.Println(armstrong(153))

	var myEngine gasEngine = gasEngine{25, 15}
	canMakeIt(myEngine, 50)

	var myEngine2 electricEngine = electricEngine{25, 15}
	canMakeIt(myEngine2, 100)

	nums := [5]float64{2, 3, 4, 5, 6}
	fmt.Printf("%p\n", &nums)
	result := square(&nums)
	fmt.Println(result)
}
