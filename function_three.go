package main

import "fmt"

//Function as Paramter
//ada function didalam function
func helloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Budi" {
		return "tak dikenal"
	} else {
		return name
	}
}

//Function type Declarations
//membuat alias agar tidak panjang
type Filter func(string) string

func helloWithFiltertype(name string, filter Filter) {
	fmt.Println("Hello", filter(name))
}

func spamFiltertype(name string) string {
	if name == "Budi" {
		return "tak dikenal"
	} else {
		return name
	}
}

func main() {
	//Function as Paramter
	helloWithFilter("David", spamFilter)
	filter := spamFilter
	helloWithFilter("Budi", filter)

	//Function type declarations
	helloWithFiltertype("David", spamFiltertype)
	filtertype := spamFiltertype
	helloWithFiltertype("Budi", filtertype)

}