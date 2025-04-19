package main

import "fmt"

func square(thing *[5]float64) [5]float64{
	fmt.Printf("%p\n", thing)
	for i:= range thing{
		thing[i] = thing[i] * thing[i]
	}
	return *thing
}
