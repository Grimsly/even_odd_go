package main

import "fmt"

func main () {
	number_list := createNumberSlice(10)

	for _, number := range number_list {
		if (number % 2 == 0) {
			// You don't need to add a space to separate the variable from the string!?
			fmt.Println(number, "is an even number")
		} else {
			fmt.Println(number, "is an odd number")
		}
	}
}