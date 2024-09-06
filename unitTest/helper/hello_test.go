package helper

import (
	"fmt"
	"testing"
)
// func TestPanic(t *testing.T) {
// 	result := Hello("David")

// 	if result != "Hello David" {
// 		panic("The result is not Hello David")
// 	}
// }
func TestHello(t *testing.T) { //error still run other test
	result := Hello("David")

	if result != "Hello David" {
		t.Fail()
	}
	fmt.Println("TestHello done")
}
func TestNow(t *testing.T) { //error and stop there
	result := Hello("David")

	if result != "Hello David" {
		t.FailNow()
	}
	fmt.Println("TestNow done")
}
func TestError(t *testing.T) { //Fail() with argumen
	result := Hello("David")

	if result != "Hello David" {
		t.Error("The result must be 'Hello David'")
	}
	fmt.Println("TestError done")
}

func TestFatal(t *testing.T) { //FailNow() with argumen
	result := Hello("David")

	if result != "Hello David" {
		t.Error("The result must be 'Hello David'")
	}
	fmt.Println("TestFatal done")
}

	//go test
	//go test -v (with detail execution information)
	//go test -v -run TestName (if u just wanna execute 1 test)
	//gotest ./... (run test on diffrent folder)