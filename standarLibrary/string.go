package main

import (
	"fmt"
	"strings"
)

func main() {
	//Trim (memotong cutset awalakhir)
	//ToLower (lower casee)
	//ToUpper (upper space)
	//Split (memotong berdasarkan separator)
	//Contains (mengecek apakah mengandung string lain)
	//ReplaceAll (mengubah string)

	fmt.Println(strings.Trim("   David   ", " "))
	fmt.Println(strings.ToLower("David Github"))
	fmt.Println(strings.ToUpper("David Github"))
	fmt.Println(strings.Split("David Github", " "))
	fmt.Println(strings.Contains("David Github", "vid"))
	fmt.Println(strings.ReplaceAll("David Budi David David", "David", "Github"))

}