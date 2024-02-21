package main

import "fmt"

func main() {
	type NoKTP string
	var KTP NoKTP = "3404045658945551"

	var contoh string = "22222222222222"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(KTP)
	fmt.Println(contohKtp)
}