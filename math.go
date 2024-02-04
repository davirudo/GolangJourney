package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	var c = a + b
	fmt.Println(c)

	//augmented assinments
	var d = 10
	d += 10
	fmt.Println(d)
	d *= 5
	fmt.Println(d) 

	//unary operator
	var e = 1
	e++
	fmt.Println(e)
	e--
	fmt.Println(e)

}