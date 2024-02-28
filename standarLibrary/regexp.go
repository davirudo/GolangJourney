package main

import (
	"fmt"
	"regexp"
)

//Mencocokan sesuatu regexp yg di compile
func main() {
	//MustCompile, membuat regexp
	//MatchString, is Regexp match with string?
	//FindAllString, match string with max result

	var regex = regexp.MustCompile(`d([a-z])v`) //tambahkan ([a-z]) lagi jika lebih dari 3 huruf

	fmt.Println(regex.MatchString("david"))
	fmt.Println(regex.MatchString("dav"))
	fmt.Println(regex.MatchString("David"))

	fmt.Println(regex.FindAllString("dav david David Dav dov davirudo", 10))
}