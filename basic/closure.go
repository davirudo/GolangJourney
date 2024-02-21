package main

import "fmt"

func main() {
	counter := 0
	increment := func() { //membuat anonymous func yang disimpan di variable increment
		fmt.Println("Increment")
		counter++
	}

	fmt.Println(counter)
	increment()
	increment() //memanggil ini dapat mengubah hasil counter di line 6 karena di line 9 menambahkan counter++
	fmt.Println(counter)
	//kemampuan function berinteraksi dengan data diatasnya.
	//digunakan dengan bijak agar tidak terjadi kesalahan
}