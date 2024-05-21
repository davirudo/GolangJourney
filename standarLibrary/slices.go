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
	fmt.Println(slices.Index(names, "George")) //2 karena index ke-2
	fmt.Println(slices.Min(values))
	fmt.Println(slices.Max(values))
	fmt.Println(slices.Min(names))
	fmt.Println(slices.Max(names))
}