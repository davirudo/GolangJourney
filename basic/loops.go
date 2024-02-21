package main

import "fmt"

func main() {
	//cara 1
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}
	fmt.Println()

	//cara 2
	for counter2:= 1; counter2 <= 10; counter2++ {
		fmt.Println("Ini nomor", counter2)
	}
	fmt.Println()

	//
	//cara manual
	names := []string{"1", "2", "3"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	fmt.Println()

	//cara For Range
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
	fmt.Println()

	//jika tidak menggunakan index
	for _, name2 := range names {
		fmt.Println(name2)
	}
}