package main

import "fmt"

func Ups() interface{} {
	//return 1
	//return true
	return "Ups"
} //semuanya bisa karena any

func main() {
	kosong := Ups()
	fmt.Println(kosong)
}