package main

import (
	"fmt"
)

func returnedValue() (sum, production int) {
	w := 10
	z := 20
	return w + z, w * z
}

func beyondHelloWorld() {
	var x int
	x = 5
	var myName string
	myName = "Maazjnr"
	y := 10

	a, b := returnedValue()
	fmt.Println(
		"the return value is", a,
		"the ruturn value is", b,
	)

	fmt.Println("the value is", x, "Y value is", y, "my name is", myName)
}

func main() {
	fmt.Println("my first go program")
	beyondHelloWorld()
}
