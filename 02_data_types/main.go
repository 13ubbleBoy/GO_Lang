/*

--> Once variable's type is declared, we cannot change its type. If its a string we cannot
	make it an integer.

*/

package main

import "fmt"

func main() {
	// variable of type string with a 'var' keyword, declaring a variable but not using it, gives an error.
	var name string = "Praveen" // string are in "" double quotes, not in single quotes.

	var name2 = "Kumar" // this is same as the above but here data type is provided automatically.

	var name3 string // here 'name3' is an empty string
	fmt.Println(name, name2, name3)

	name2 = "Piyush" // values can be changed
	name3 = "Jeetu"
	fmt.Println(name, name2, name3)

	// without the 'var' keyword declaring a variable
	name4 := "Ankit" // the ':=' is used only for variable declaration.
	// ':=' cannot be used outside of a function.
	fmt.Println(name4)

	name4 = "Mithun\n" // no need for ':=' this to modify variables.
	fmt.Println(name4)

	// Integers
	var age1 int = 20 // this is int8, 8 bits integer ranging from -128 to 127
	var age2 = 30
	age3 := 40
	fmt.Println(age1, age2, age3, "\n")

	// we can specify the int size according to our need
	number := 130
	fmt.Println(number, "\n")

	var age4 int8 = 23  // 8 bit integer
	var age5 int16 = 23 // 16 bit integer
	fmt.Println(age4, age5, "\n")
}
