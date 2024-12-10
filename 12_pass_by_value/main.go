package main

import "fmt"

func updateName1(name string) { // this 'name' is a copy of actual 'name' from 'main' function.
	name = "Kumar"
}

func updateName2(name string) string { // this 'name' is a copy of actual 'name' from 'main' function.
	name = "Kumar"
	return name
}

func main() {
	name := "Praveen"
	fmt.Println(name, "\n")

	updateName1(name) // when we pass an argument go creates a copy of the variable and then sent it.
	fmt.Println(name, "\n")
	/*
		--> To fix this we can update the value and return it from the function and store it in the
			same vaiable name
	*/
	name = updateName2(name)
	fmt.Println(name, "\n")

}
