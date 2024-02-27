package main

import (
	"fmt"
	"reflect"
)


type Sample struct {
	Name string 
}

type Person struct {
	Name string `required:"true" max:"10"`
	Age int `required:"true" max:"10"`
	Address string	`required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Println("Type Name", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
		valueField.Tag.Get("required")
		valueField.Tag.Get("max")
		
	}
}

//Validation Library
func isValid(value any) (result bool) {
	result = true
	t := reflect.TypeOf(value)
	for i := 0; i< t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			data := reflect.ValueOf(value).Field(i).Interface()
			result = data != ""
			if result == false {
				return result
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"David"}
	sampleType := reflect.TypeOf(sample)
	structField := sampleType.Field(0)
	fmt.Println(structField.Name)
	fmt.Println()
	//melihat isi yang ada di dalam Sample struct

	readField(Sample{"David"})
	fmt.Println()
	readField(Person{"David", 40, "Indo"})
	fmt.Println()
	//membaca type name dan isi keseluruhan di Person struct

	person := Person {
		Name : "David",
		Age : 40,
		Address: "Indo",
	}
	fmt.Println(isValid(person))
	//mengecek jika semua Datanya valid dan semua terisi
}