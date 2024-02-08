package main

import "fmt"

//Named Return Values
func namedReturnValues() (first, mid, last string) {
	first = "David"
	mid = "Go"
	last = "Lang"

	return first, mid, last
}

//Variadic Function
func sumAll(numbers ...int) int { //..int adalah slice integer
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	first, mid, last := namedReturnValues()
	fmt.Println(first,mid,last)

	total := sumAll(5, 15, 10, 20, 50)
	fmt.Println(total)
	//  fmt.Println(sumAll(5, 15, 10, 20, 50))

	//Slice parameter
	numbers := []int{5, 15, 10, 20, 50} //terlanjur memiliki slice
	fmt.Println(sumAll(numbers...)) //dikasih ... diakhir agar bisa dibaca oleh sumAll
}