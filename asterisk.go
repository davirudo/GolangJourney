package main

import "fmt"

type Alamat struct {
	City, Prov, Country string
}

func main() {
	//Pointer
	address3 := Alamat{"Jogja", "Jawa", "Indo"}
	address4 := &address3
	address4.City = "Bandung"
	fmt.Println(address3) //ikut berubah
	fmt.Println(address4)

	address4 = &Alamat{"Kuala", "Lumpur", "Malaysia"}
	fmt.Println(address3)
	fmt.Println(address4) //sudah menjadi alamat baru
	fmt.Println()


	//Asterisk Operator - menggunakan bintang (*)
	address1 := Alamat{"Jogja", "Jawa", "Indo"}
	address2 := &address1
	address4.City = "Bandung"
	fmt.Println(address1) //ikut berubah
	fmt.Println(address2)

	*address2 = Alamat{"Kuala", "Lumpur", "Malaysia"}
	fmt.Println(address1) //tetap ikut berubah walau yang dirubah Address2
	fmt.Println(address2) 
}