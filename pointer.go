package main

import "fmt"

type Address struct {
	City, Prov, Country string
}

func main() {
	//Pass by Value
	//saat kirim variabel ke  func, method, var lain. data awal tidak berubah
	address1 := Address{"Jogja", "Jawa", "Indo"}
	address2 := address1 //duplikat value

	fmt.Println(address2) //sebelum
	address2.City = "Bandung"
	fmt.Println(address1) //data tidak berubah walau dirubah Bandung di address2 
	fmt.Println(address2)
	fmt.Println()


	//Pass by Reference (Pointer "&")
	//data akan berubah langsung ke variabel yang dikirim
	address3 := Address{"Jogja", "Jawa", "Indo"}
	address4 := &address3 //duplikat value

	fmt.Println(address4) //sebelum
	address4.City = "Bandung"
	fmt.Println(address3) //data berubah walau dirubah Bandung di address2 
	fmt.Println(address4)
	}

