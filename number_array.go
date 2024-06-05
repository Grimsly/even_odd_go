package main

func createNumberSlice(length int) []int {
	var number_list []int
	for i := 0; i < length + 1; i++ {
		number_list = append(number_list, i)
	}

	return number_list
}