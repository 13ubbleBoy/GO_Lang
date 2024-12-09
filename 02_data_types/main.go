/*

--> Once variable's type is declared, we cannot change its type. If its a string we cannot
	make it an integer.

--> String
	Strings are immutable because they are represented as a read-only slice of bytes. When you
	create a string, it is stored in memory as a sequence of bytes, and you cannot modify these
	bytes directly. Instead, any operation that appears to "modify" a string (e.g., appending,
	replacing, etc.) actually creates a new string.

	Ex: s := "Hello"
    	s[0] = 'h' // This will result in a compile-time error

--> Efficiency: String immutability allows safe sharing of string values without worrying about
	concurrent modifications. This is especially useful in multithreaded programs.
--> Memory Optimization: Since strings are immutable, Go can safely reuse string data and optimize
	memory usage.
--> Predictability: Immutability avoids unexpected side effects when strings are passed between
	functions.

--> Channel type
	In Go, channels are a powerful type used for communication between goroutines, enabling
	synchronization and data exchange. Channels provide a way to send and receive values of a
	specific type between goroutines.

--> rune
	a rune is a data type used to represent a single Unicode code point. It is an alias for the
	int32 type, meaning it can store any Unicode character, including those outside the ASCII range.
	This makes it suitable for handling multilingual text and special symbols.

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
	var age1 int = 20 // Explicit Declaration
	var age2 = 30     // Type Inference
	age3 := 400       // Short Declaration
	fmt.Println(age1, age2, age3, "\n")

	// we can specify the int size according to our need
	number := 130
	fmt.Println(number, "\n")

	var age4 int8 = 23  // // this is int8, 8 bits integer ranging from -128 to 127
	var age5 int16 = 23 // 16 bit integer
	fmt.Println(age4, age5, "\n")

	// uint, it is also a 8 bit but its range is 0 - 255
	var age6 uint = 240
	fmt.Println(age6, "\n")

	// decimal numbers
	var price1 float32 = 12.5 // here we have to specify if its float32 or float64
	var price2 = 13.28747
	price3 := 31.389634
	fmt.Println(price1, price2, price3, "\n")

	// boolean
	flag := true
	fmt.Println(flag)

	// pointer
	var ptr *int = &age1
	println("Value inside ptr:", ptr, "\nValue at address", ptr, "is:", *ptr, "\n")

	// custom type variable
	type Celsius = float64
	var celsius Celsius = 42.32354
	fmt.Printf("Temperature: %0.2f\n", celsius)

	// constants
	const PI = 3.14
	fmt.Printf("Value of pie: %0.2f\n", PI)

	// channel type

	// character type
	ch := 65
	fmt.Printf("Character: %c\n", ch)

	// rune
	var r rune = 'A'           // Single quotes for a rune
	fmt.Println(r)             // Output: 65 (Unicode code point for 'A')
	fmt.Printf("rune %c\n", r) // Output: A
}
