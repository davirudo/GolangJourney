package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestAssert(t *testing.T) { //Fail()
	result := Hello("David")
	assert.Equal(t, "Hello David", result, "Result must be 'Hello David'")
	fmt.Println("TestAssert done")
}

func TestRequire(t *testing.T) { //FailNow()
	result := Hello("David")
	require.Equal(t, "Hello David", result, "Result must be 'Hello David'")
	fmt.Println("TestRequire done")
}

func TestSkip(t *testing.T) { //skil on some conditional
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on MacOS")
	}
	//if true, than u can run the test
	result := Hello("David")
	assert.Equal(t, "Hello David", result, "Result must be 'Hello David'")
	fmt.Println("TestAssert done")
}

func TestMain(m *testing.M) {
	//before
	fmt.Println("Before Unit Test")
	
	m.Run()

	//after
	fmt.Println("After Unit Test")
}

	//go test
	//go test -v (with detail execution information)
	//go test -v -run TestName (if u just wanna execute 1 test)
	//gotest ./... (run test on diffrent folder)
