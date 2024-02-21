package main

import "fmt"

//hanya bisa digunakan pada interface, function, map, slice, pointer, dan channel

func newMap(name string) map[string] string {
	if name == "" {
		return nil
	} else {
		return map[string]string {
			"name" : name,
		}
	}
}

func main() {
	data := newMap("Ada")
	if data == nil {
		fmt.Println("Data Kosong")
	} else {
		fmt.Println(data)
	}
}