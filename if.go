package main

import "fmt"

func main() {
	name := "David"
	name2 := "David"

	if name == name2 {
		fmt.Println("iya itu saya")
	} else if name == "Davirudo" {
		fmt.Println("nama lain saya")
	} else {
		fmt.Println("salah orang")
	}
	
	//if dengan short statement
	if length := len(name); length < 8 {
		fmt.Println("terlalu pendek")
	} else {
		fmt.Println("bisa dipakai")
	}
}