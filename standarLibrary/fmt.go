package main

import "fmt"

func main() {
	firstName := "David"
	lastName := "Github"

	fmt.Println("Hello '", firstName, lastName, "!'") //ribet dan terbatas
	fmt.Printf("Hello '%s %s!' \n", firstName, lastName)
}