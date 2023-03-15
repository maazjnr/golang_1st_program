package main

import (
	"fmt"
)


func beyondHelloWorld() {
	var x int
	x = 5
	var myName string
	myName = "Maazjnr"
	y := 10

	a, b := returnedValue()
	fmt.Println(
		"the return value is", a,
		"the return value is", b,
	)

	fmt.Println("the value is", x, "Y value is", y, "my name is", myName)
}

func returnedValue() (sum, production int) {
	w := 10
	z := 20
	return w + z, w * z
}

func arrayLoop() {
	var a [5]int
	a[1] = 200
	fmt.Println("the array is", a, "the length is", len(a))

	b := []int{20, 30, 40, 50, 60}
	b = append(b, 10)
	fmt.Println("the array array b is", b, "the length is", len(b))

	for i := 0; i < len(b); i++ {
		fmt.Println("the value os the index is", b[i])
	}

}


func main() {
	fmt.Println("my first go program")
	beyondHelloWorld()
	arrayLoop()
}
