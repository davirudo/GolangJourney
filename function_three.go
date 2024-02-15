package main

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

func main() {
	//Function as Paramter
	helloWithFilter("David", spamFilter)
	filter := spamFilter
	helloWithFilter("Budi", filter)

}