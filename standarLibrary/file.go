package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

//Membuat file baru
func createNewFIle(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

//Membaca file
func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}
	return message, nil
}

//Membuat file baru
func addToFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	// createNewFIle("sample.log", "this is worked")

	// result, _ := readFile ("sample.log")
	// fmt.Println(result)

	addToFile("sample.log", "\nadd message")
	result, _ := readFile ("sample.log")
	fmt.Println(result)
}