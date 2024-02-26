package main

import (
	"fmt"
	"strconv"
)

func main() {
	//parseBool string to bool
	//FormatBool bool ke string

	boolean, err := strconv.ParseBool("true") //awalnya string "true", menjadi bool
	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println("Error", err.Error())
	}

	resultBinary := strconv.FormatInt(999, 2) //999 dalam binary itu berapa
	if err == nil {
		fmt.Println(resultBinary)
	} else {
		fmt.Println("Error", err.Error())
	}

	stringInt := strconv.Itoa(999) //dari int menjadi string
	if err == nil {
		fmt.Println(stringInt)
	} else {
		fmt.Println("Error", err.Error())
	}

}