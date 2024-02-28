package main

import (
	"fmt"
	"slices"
)

//Mengecek dan Manipulate slice
func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 90, 80, 70}

	fmt.Println(slices.Contains(names, "Paul"))
	fmt.Println(slices.Contains(names, "Dino"))
	fmt.Println(slices.Index(names, "George"))
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
}