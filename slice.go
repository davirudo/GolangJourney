package main

import (
	"fmt"

)

func main() {
	names := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "Minggu"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[:3]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	//function slice
	fmt.Println()

	namesSlice1 := names[5:] // Sabtu, Minggu
	fmt.Println(namesSlice1)
	
	namesSlice1[0] = "Sabtu Baru"
	namesSlice1[1] = "Minggu Baru"
	fmt.Println(namesSlice1)
	fmt.Println(names) //array berubah menjadi baru
	
	//append = membuat array baru (namesSlice2) dari nameSlice1
	namesSlice2 := append(namesSlice1, "Libur Baru")
	fmt.Println()
	fmt.Println(namesSlice1)
	fmt.Println(namesSlice2)
	fmt.Println(names)

	//data dari namesSlice2 tidak merubah array inti
	namesSlice2[0] = "Sabtu Lama"
	fmt.Println()
	fmt.Println(namesSlice2)

	//
	//Make Slice, membuat secara manual
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Senin" 
	newSlice[1] = "Senin"
	// akan error untuk [2] karena panjang hanya 2

	fmt.Println()
	fmt.Println(newSlice)
	fmt.Println(len(newSlice)) //panjang data saat ini
	fmt.Println(cap(newSlice)) //kapasitas max data

	newSlice2 := append(newSlice, "Senin") //menambah data untuk [2]
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Selasa" //menambah data baru tanpa merusak newSlice
	fmt.Println(newSlice)
	fmt.Println(newSlice2)
	fmt.Println()

	//copy = mengcopy data dari array lain
	fromSlice := names[:] //mengambil semua data names
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	//
	//perbedaan slice dengan array
	iniArray := [...]int{1,2,3,4,5}
	iniSlice := []int{1,2,3,4,5}
	fmt.Println(iniArray)
	fmt.Println(iniSlice)

	









}