package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	//base64
	var encoded = base64.StdEncoding.EncodeToString([]byte("David Github"))
	fmt.Println(encoded)

	var decoded, err = base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(decoded))
		fmt.Println()
	}

	//CSV Reader
	csvString := "david,github,davirudo\n" +
	"budi,afi,bima\n" +
	"asep,irusan,mikel"

	reader := csv.NewReader(strings.NewReader(csvString))
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}
	fmt.Println()

	//CSV Writer 
	writer := csv.NewWriter(os.Stdout)
	_ = writer.Write([]string{"david", "github", "davirudo"})
	_ = writer.Write([]string{"budi", "afi", "bima"})
	_ = writer.Write([]string{"asep", "irusan", "mikel"})
	writer.Flush()
}