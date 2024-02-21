package main

import "fmt"

type Tempat struct {
	City, Prov, Country string
}

func main() {
	//New - operator yang mengembalikan pointer ke data kosong
	//tempat1 := &Address{}
	tempat1 := new(Tempat)
	tempat2 := tempat1
	tempat2.Country = "Indo"
	fmt.Println(tempat1) //tempat1 ikut berubah
	fmt.Println(tempat2)
	fmt.Println()


	//Function Operator

}