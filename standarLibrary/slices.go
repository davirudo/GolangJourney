package main

import (
	"fmt"
	"slices"
)

//Manipulate slice
func main() {
	names := []string{"John", "Paul", "George", "Ringo"}

	fmt.Println(slices.Contains(names, "Paul"))
	fmt.Println(slices.Contains(names, "Dino"))
	fmt.Println(slices.Index(names, "George"))
}