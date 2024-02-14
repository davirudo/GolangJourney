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

//Function Value
func goodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	first, mid, last := namedReturnValues()
	fmt.Println(first,mid,last)

	total := sumAll(5, 15, 10, 20, 50)
	fmt.Println(total)
	//  fmt.Println(sumAll(5, 15, 10, 20, 50))

	//Slice parameter
	angka := []int{5, 15, 10, 20, 150} //terlanjur memiliki slice
	fmt.Println(sumAll(angka...)) //dikasih ... diakhir agar bisa dibaca oleh sumAll

	value := goodBye //bisa menjadikan function menjadi variable
	fmt.Println(value("Golang"))
}