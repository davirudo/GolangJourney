package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

//struktur data yang diakhir e akan kembali ke e awal
func main() {
	data := ring.New(5)
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i) //merubah i menjadi int
		data = data.Next()
	}

	data.Do(func(value interface {}) {
		fmt.Println(value)
	})
}