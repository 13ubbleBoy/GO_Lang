/*
--> It is not necessary to declare a function above the main function in Go. The Go compiler
	allows functions to be declared in any order within the same package. This is because Go
	uses a single-pass compiler, which analyzes the entire file during compilation and resolves
	function declarations and calls regardless of their order.
*/

package main

import (
	"fmt"
	"math"
)

// function declaration
func greeting(name string) {
	fmt.Println("Welcome!", name)
	fmt.Println()
}

func leaving(name string) {
	fmt.Println("Good Bye!", name)
	fmt.Println()
}

func printNames(names []string) {
	for _, value := range names {
		fmt.Println(value)
	}
	fmt.Println()
}

// we can also pass function to a function as an parameter.
func printNamesAndGreet(names []string, fun func(string)) {
	for _, name := range names {
		fun(name) // calling greeting function from line no. 13
	}
	fmt.Println()
}

// returning a value from a function
func add(a int, b int) int { // 'int' before { is the return type
	return a + b
}

func perimeterOfCircle(radius float32) float64 {
	return 2 * math.Pi * float64(radius)
}

// returning multiple values
func addAndDivide(num1 int, num2 int) (int, float32) {
	sum := num1 + num2
	division := float32(num1) / float32(num2)
	return sum, division
}

func main() {
	// calling a function
	greeting("Praveen")
	leaving("Kumar")

	names := []string{"Praveen", "Ankit", "Piyush", "Jeetu", "Mithun"}
	printNames(names)

	// calling a function and passing another function as an argument
	printNamesAndGreet(names, greeting) // we only passed the reference for function greeting
	// because it we use 'greeting()' then the function will be called automatically.

	// add to numbers using a function
	fmt.Println(add(2, 5))
	fmt.Println()

	fmt.Printf("Perimeter of a circle of radius 7.14: %0.2f\n", perimeterOfCircle(7.14))
	fmt.Println()

	// accepting 2 return values from a function
	sum, division := addAndDivide(7, 2)
	fmt.Printf("Sum: %v\nDivision: %0.2f\n\n", sum, division)
}
