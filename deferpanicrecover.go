package main

import (
	"fmt"
)

//Defer
//func yg bisa dijadwalkan untuk dieksekusi walau error
func logging() { 
	fmt.Println("Selesai")
}
func runApp() {
	defer logging()
	fmt.Println("Run App")
	//error
	//fmt.Println("Selesai")
	//menggunakan defer karena jika ditengah-tengah terjadi error, maka Logging tidak dijalankan
}

//Panic
//menghentikan program saat terjadi error misalnya, namun difer akan tetap dieksekusi
func endApp() {
	fmt.Println("End App")
}
func runApp2(error bool) {
	defer endApp()
	if error {
		panic("Ups, error")
	}
}

//Recover
//untuk menangkap data panic
func runApp3(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
	message := recover()
	fmt.Println("Terjadi Error", message)
}



func main() {
	//Defer
	runApp()

	//Panic
	runApp2(false)
	runApp2(true)

	//Recover
	runApp3(true)
	
}