package main

import (
	"fmt"
	"math"
)

func main() {
	//Round, membulatkan keatas atau ke bawah, sesuai dengan yang paling deket
	//Floor, kebawah
	//Ceil, keatas
	//Max, nilai paling besar
	//Min, nilai paling kecil

	fmt.Println(math.Round(1.40))
	fmt.Println(math.Floor(1.40))
	fmt.Println(math.Ceil(1.40))
	fmt.Println(math.Max(1, 2))
	fmt.Println(math.Min(1, 2))
}