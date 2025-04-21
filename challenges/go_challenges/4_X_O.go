package go_challenges

import "fmt"


// with if statement
func Count_x_o(word string) bool {
	x_counter := 0
	o_counter := 0

	for _, let:= range word {
		fmt.Println(let)
		if let == 'X' {
			x_counter++
		}
		if string(let) == "O" {
			o_counter++
		}
	}
	fmt.Println(x_counter)
	fmt.Println(o_counter)
	return x_counter == o_counter
}


// with switch case
func Count_x_o_2(word string) bool {
	x_counter := 0
	o_counter := 0

	for _, let:= range word {

		switch let {
		case 'X':
			x_counter++
		case 'O':
			o_counter++
		}
	}
	return x_counter == o_counter
}
