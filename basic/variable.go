package main

import "fmt"

func main() {
	var name string
	name = "David"
	fmt.Println(name)

	//sistem bisa tau tanpa tipedata
	var nama = "Davirudo"
	fmt.Println(nama)

	//tanpa var
	namo := "Davi Rudo"
	fmt.Println(namo)

	//multiple variable
	var(
		first = "Davi"
		last = "Rudo"
	)
	fmt.Println(first)
	fmt.Println(last)

	//constant (cant change the value)
	const (
		firstName = "Car"
		lastName = "Rude"
	)
	//firstName = "Cat"
	//lastName = "Cute"

}