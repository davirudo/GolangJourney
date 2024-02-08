package main

import "fmt"

func main() {
	//break
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //hentikan total
		}
		fmt.Println("Nomor ke", i)
	}
	
	//comtinue
	for m := 0; m < 10; m++ {
		if m%2 == 0 {
			continue //untuk melanjutkan perulangan 
		}
		fmt.Println("Ini adalah", m)
		//mengexecute hasil yang bisa di %2 saja, sistem akan dilanjutkan yaitu mengeprint yang tidak bisa di %2
	}
}