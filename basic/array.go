package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Github"
	names[1] = "David"
	names[2] = "Davirudo"
	fmt.Println(names[1])
	
	var values = [3]int{
		90, 30, 95,
	}
	fmt.Println(values)
	fmt.Println(len(values)) //hitung isi
	values[0] = 100 //mengubah data
	fmt.Println(values)
}