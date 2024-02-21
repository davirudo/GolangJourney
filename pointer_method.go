package main

import "fmt"

type Man struct {
	Name string
}

// func (man Man) Married() { //Man adalah value, artinya diduplikat bukan data asli
// 	man.Name = "Mr. " + man.Name
// }

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
}

func main() {
	budi := Man{"Budi"}
	budi.Married()
	fmt.Println(budi.Name)
}