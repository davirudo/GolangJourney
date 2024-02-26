package main

import (
"fmt"
"container/list"
)

//Double linked-list
//Seperti antrian atau tumpukan
func main() {
	data := list.New() //any
	data.PushBack("David")
	data.PushBack(1)
	data.PushBack(true)
	data.PushFront("Budi") //antri dari paling depan

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}