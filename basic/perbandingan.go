package main

import "fmt"

func main() {
	var name1 = "Cat"
	var name2 = "Car"
	var result bool = name1 == name2
	fmt.Println(result)

	var number1 = 2
	var number2 = 3
	var resultNumber bool = number1 < number2
	fmt.Println(resultNumber)

	//AND &&, semua harus true agar true
	//OR ||, semua harus false agar false
	//REVERSE !, true dibalik menjadi false
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi
	fmt.Println(lulus)
}