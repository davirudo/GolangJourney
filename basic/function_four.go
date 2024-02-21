package main

import "fmt"

//Recursive Function
//coba menggunakan factorial for loop biasa
func factorialLoop(value int) int {
	result := 1 
	for i := value; i > 0; i-- {
		result *= i 
	} 
	return result
}
//sekarang menggunakan Recursive Function
func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive(value-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}