package main

import "fmt"

func main() {

	// Print, line will not break
	fmt.Print("Hello World!")

	// Print then break line
	fmt.Println("Hello")

	// Print multiple values with comma seperated
	name := "Praveen"
	age := 23
	fmt.Println("My name is", name, "and my age is", age) // it also adds spaces before and after if used between strings.

	// Printf like in c langauge
	fmt.Printf("My name is %v and my age is %v\n", name, age) // $v normally puts the variable
	fmt.Printf("My name is %q and my age is %q\n", name, age) // %q puts double quote on the variable but it does not apply on numbers.

	// find data type
	fmt.Printf("Data type of name: %T\nData type of age: %T\n", name, age)

	// print decimal numbers
	dec := 12.469678
	fmt.Printf("Decimal no. is: %v\n", dec)
	// or
	fmt.Printf("Decimal no. is: %f\n", dec)

	// set precision on decimal numbers
	fmt.Printf("Decimal no. is: %0.2f\n", dec) // it also rounds off the value

	// Sprintf: Returns a formatted string, according to what we provide.
	str := fmt.Sprintf("My name is %q and my age is %v", name, age)
	fmt.Println(str)

}
