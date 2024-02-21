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
//untuk menangkap data panic, dan panic akan dihentikan lalu data dijalankan kembali
//contoh recover yang salah
func runApp3salah(error bool) {
	defer endApp()
	if error {
		panic("ERROR") //akan berhenti disini
	}
	message := recover()
	fmt.Println("Terjadi Panic", message)
}
//contoh recover yang benar
func endAppRecover() {
	fmt.Println("End App")
	message := recover()
	fmt.Println("Terjadi Panic", message)
}

func runApp3(error bool) {
	defer endApp()
	if error {
		panic("ERROR")
	}
} //dimasukkan ke defer karena walau error code tetap di eksekusi



func main() {
	// Defer
	runApp()

	//Panic
	runApp2(false) //jika true, tidak bisa mengeksekusi code yang dibawah line ini karena apanic

	// Recover
	runApp3(false)
	fmt.Println("Halo")
	
}