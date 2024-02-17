package main

import (
	"fmt"
)

func random() interface{} {
	return "Random"
}

func main() {
	result := random()              //masih berbentuk any
	resultString := result.(string) //dirubah menjadi string
	fmt.Println(resultString)

	// resultInt := result.(int) //panic karena paksa string menjadi int
	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}