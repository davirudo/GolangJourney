package main

import (
	"fmt"
	"path"
	"path/filepath"
)

func main() {
	fmt.Println(path.Dir("hello/david.go"))
	fmt.Println(path.Base("hello/david.go"))
	fmt.Println(path.Ext("hello/david.go"))
	fmt.Println(path.Join("hello", "david", "main.go"))
	fmt.Println()

	//bisa tau sistemnya windows atau bukan (hasil menjadi backslash)
	fmt.Println(filepath.Dir("hello/david.go"))
	fmt.Println(filepath.Base("hello/david.go"))
	fmt.Println(filepath.Ext("hello/david.go"))
	fmt.Println(filepath.IsAbs("hello/david.go"))
	fmt.Println(filepath.IsLocal("hello/david.go"))
	fmt.Println(filepath.Join("hello", "david", "main.go"))
}