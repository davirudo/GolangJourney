package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	for _, arg := range args {
	fmt.Println(arg)
	}
	//go run os.go on Terminal with name after 

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Println(hostname)
	} else {
		fmt.Println("error", err.Error())
	}
}