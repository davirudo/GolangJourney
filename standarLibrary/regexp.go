package main

import (
	"fmt"
	"regexp"
)

//Mencocokan berdasarkan format yg diinginkan
func main() {
	//MustCompile, membuat regexp
	//MatchString, is Regexp match with string?
	//FindAllString, match string with max result

	var regex = regexp.MustCompile(`d([a-z])v`) //tambahkan ([a-z]) lagi jika lebih dari 3 huruf

	fmt.Println(regex.MatchString("david")) //dav masuk
	fmt.Println(regex.MatchString("dov"))
	fmt.Println(regex.MatchString("David")) //D besar

	fmt.Println(regex.FindAllString("dav david David Dav dov davirudo", 10))
}