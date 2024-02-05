package main

import "fmt"

func main() {
	person := map[string]string{ //string pertama untuk menentukan tipe key
		"name":    "David",
		"address": "Jogja",
		"hobby": "Swimming",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	//Function Map
	fmt.Println()
	fmt.Println(len(person)) //jumlah data di map
	
	delete(person, "hobby")
	fmt.Println(person)


}