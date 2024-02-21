package main

import "fmt"

func hello() {
	fmt.Println("hello")
}

//function with parameter
func helloParam(first string, last string) {
	fmt.Println("hello", first, last )
}

//function with return
func helloReturn(name string)  string {
	return "hello " + name
	//  hello := "hello : + name
	//  return hello
}

//returning multiple values
func  getFullName() (string, string) {
	return "David", "Go"
}

func main() {
	hello()

	helloParam("Davirudo", "Go")

	result := helloReturn("David")
	fmt.Println(result)
	//  fmt.Println(helloReturn("David"))

	one, two := getFullName()
	fmt.Println(one, two)
	//avoid some return value with _
	satu, _ := getFullName()
	fmt.Println(satu)
}