package main

import "fmt"

var score = 100

// if 'score' we define this inside of main function then it will bo gone from its scope making it
// unaccessable

func main() { // this main function only appear once in the program.
	sayHello("Praveen")

	for _, item := range points {
		fmt.Println(item)
	}
	fmt.Println()

	showScore()
}

/*
func main() { // redecalration error
	fmt.Println("hello")
}
*/
