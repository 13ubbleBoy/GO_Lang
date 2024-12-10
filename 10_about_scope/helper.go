package main

/*
--> 'package main': It mean that this file is also a part of the main package and if we dont do then
	we cannot share variables and functions between files.

--> In Go, every file begins with a package declaration. If multiple files have the same package
	declaration and are in the same directory, they automatically share code.
*/

/* this will only run when the files are not connected.
func main() {
	fmt.Println("hello2")
}
*/

import "fmt"

// points := []int{26,13,47,34,50} // this cannot be used outside of a function
var points = []int{26, 13, 47, 34, 50}

func sayHello(name string) {
	fmt.Println("Hello!", name)
	fmt.Println()
}

func showScore() {
	fmt.Println("Score:", score)
	fmt.Println()
}
