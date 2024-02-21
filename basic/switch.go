package main

import "fmt"

func main() {
	name := "Daviddsadsda"

	switch name {
	case "Davirudo":
		fmt.Println("Salah")
	case "David":
		fmt.Println("Iya itu")
	default:
		fmt.Println("bukan semuanya")
	}

	//switch dengan short statement
	switch length := len(name); length < 8 {
	case true:
		fmt.Println("kurang panjang")
	case false:
		fmt.Println("sudah pas")
	}

	//switch tanpa kondisi
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Terlalu panjang")
	case length > 8:
		fmt.Println("Sudah pas")
	default:
		fmt.Println("Kurang panjang")
	}
}