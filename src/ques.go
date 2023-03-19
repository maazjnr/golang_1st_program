package main

import "fmt"

func main() {
	fmt.Println("Welcome to the quiz game!")

	fmt.Printf("Enter your name: ")

	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v Welcome to the quiz game!\n", name)
	var age uint
	fmt.Scan(&age)

	if age >= 10 {
		fmt.Println("Yahh, correct")
	} else {
		fmt.Println("Below required age")
		return
	}

	score := 0
	num_question := 2

	fmt.Printf("What is better between RTX 300 or RTX 390? ")

	var answer string
	var answer2 string

	fmt.Scan(&answer, &answer2)
	fmt.Println(answer, answer2)

	if answer+" "+answer2 == "RTX 390" || answer+" "+answer2 == "rtx 390" {
		fmt.Println("Correct!")
		score++
	} else {
		fmt.Println("Incorrect")
	}

	fmt.Println("How many cores does the Ryzen 9 3900x have? ")
	var cores uint
	fmt.Scan(&cores)

	if cores == 12 {
		fmt.Println("Correct answer")
		score++
	} else {
		fmt.Println("Incorrect answer")
	}

	fmt.Printf("You score %v out of %v \n", score, num_question)
	percent := (float64(score) / float64(num_question)) * 100
	fmt.Printf("You scored: %v%%.", percent)

}
