package main

import "fmt"

type Customer struct {
	Name, Anddress string
	Age            int
}

func (customer Customer) hello(name string) {
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func main() {
	//Cara 1
	var budi Customer
	budi.Name = "Budi"
	budi.Anddress = "Jakarta Timur"
	budi.Age = 35
	fmt.Println(budi)
	fmt.Println(budi.Name)

	//Cara 2
	david := Customer {
		Name:  "David",
		Anddress: "Jogja",
		Age : 0,
	}
	fmt.Println(david)

	//Cara 3
	agus := Customer { "Agus", "Malang", 40}
	fmt.Println(agus)

	//cara memanggil method
	budi.hello("Dino")


}