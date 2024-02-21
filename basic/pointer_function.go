package main

import "fmt"

type Lokasi struct {
	City, Prov, Country string
}

func ChangeAddressToIndo(address *Lokasi) {
	address.Country = "Indo"
}

func main() {
	//Function Pointer
	//function yg bisa mengubah data asli parameter
	address := Lokasi{"Subang", "JaBar", ""}
	
	ChangeAddressToIndo(&address)
	fmt.Println(address)
}