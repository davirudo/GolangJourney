package main

func main() {
	//Reader Membaca Input
	type Reader interface {
		Read(p []byte) (n int, err error)
	}

	//Writer Untuk menulis output didalam Package IO
	type Writer interface {
		Write(p []byte) (n int, err error)
	}
}