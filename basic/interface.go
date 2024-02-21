package main

import "fmt"

type HasName interface {
	GetName() string
} //kontrak
func Sayhello(hasname HasName) {
	fmt.Println("Hello", hasname.GetName())
}

//
type Person struct {
	Name string
}
func (person Person) GetName() string { //sesuai kontrak
	return person.Name
}
//

type Animal struct {
	Name string 
}
func (animal Animal) GetName() string { //sesuai kontrak
	return animal.Name
}

func main() {
	person := Person{Name: "David"}
	Sayhello(person)
	animal := Animal{Name: "Woof"}
	Sayhello(animal)
}