package main

import "fmt"

func main() {
	age := 23
	fmt.Println(age > 18)
	fmt.Println(age < 18)
	fmt.Println()

	// if else and else if
	if age < 18 {
		fmt.Println("Age is greater than 18")
	} else if age < 22 {
		fmt.Println("Age is less than 18")
	} else {
		fmt.Println("Age is greater than 22")
	}
	fmt.Println()

	// continue: Break the current iteration, continues the loop without executing code below for that condition.
	names := []string{"Praveen", "Ankit", "Piyush", "Jeetu", "Mithun"}
	for index, value := range names {
		if index == 1 {
			fmt.Println("We are at index:", index)
			continue // what written below it will not be executed
		}
		fmt.Println("Name:", value)
	}
	fmt.Println()

	// break: Breaks the loop completely.
	for index, value := range names {
		if index == 3 {
			fmt.Println("We are at index:", index)
			break
		}
		fmt.Println("Name:", value)
	}
	fmt.Println()

}
