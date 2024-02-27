package main

import (
	"fmt"
	"time"
)

func main() {
	//Now, mendapatkan waktu saat ini
	//Date, membuat waktu
	//Parse, memparsing waktu dari string
		now := time.Now()
		fmt.Println(now)
		fmt.Println(now.Local())
		fmt.Println()

		utc := time.Date(2024, time.March, 17, 3, 0, 0, 0, time.UTC)
		fmt.Println(utc)
		fmt.Println(utc.Local())
		fmt.Println()

		//yyyy-MM-dd HH:mm:ss
		parse, _ := time.Parse(time.RFC3339, "2024-01-02 15:04:05")
		fmt.Println(parse)
		fmt.Println(parse.Local())
		fmt.Println()
		//atau
		formatter 	:= "2006-01-02 15:04:05"
		value 		:= "2020-10-10 20:10:10"
		valueTime, err := time.Parse(formatter, value)
		if err == nil {
			fmt.Println(valueTime)
		} else {
			fmt.Println("Error", err.Error())
		}
		fmt.Println(valueTime.Hour())
		fmt.Println()
}