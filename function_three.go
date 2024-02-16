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

//Anonymous Function
func registerUser(name string, blacklist func(string) bool) {
	if blacklist(name) {
		fmt.Println("You Are Blocked", name)
	} else {
	fmt.Println("Halo", name)	
	}
}

//

func main() {
	//Function as Paramter
	helloWithFilter("David", spamFilter)
	filter := spamFilter
	helloWithFilter("Budi", filter)

	//Function type declarations
	helloWithFiltertype("David", spamFiltertype)
	filtertype := spamFiltertype
	helloWithFiltertype("Budi", filtertype)

	//Anonymous function
	blacklist := func(name string) bool {
		return name == "Budi"
	}
	registerUser("David", blacklist)
	//atau bisa langsung
	registerUser("Budi", func(name string) bool {
		return name == "Budi"
	})



}