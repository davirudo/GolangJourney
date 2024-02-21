package main

import "fmt"

func main() {
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)
	//nilai 16 menjadi minus karena nilai max nya adalah 32767

	var name = "Davirudo"
	var e = name[0]
	var eString = string(e)
	//konversi tipedata

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)
}