package main

import (
	_"basic/internal" // (_ Blank Identifier, menjalankan init tanpa mengeksekusi function)
	"basic/database"
	"fmt"
)

func main() {
	//Access Modifier
	//Function yang otomatis di eksekusi ketika sebuah package digunakan
	fmt.Println(database.GetDatabase())


}