package main

import (
	"fmt"
)

// Type Switch with interface{}
func describeType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Integer: %d\n", v)
	case string:
		fmt.Printf("String: %q\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Unknown type: %T\n", v)
	}
}

// main function
func main() {
	// single value in a single case
	day := 1
	switch day {
	case 1:
		fmt.Println("Monday\n")
	case 2:
		fmt.Println("Tuesday\n")
	case 3:
		fmt.Println("Wednesday\n")
	case 4:
		fmt.Println("Thursday\n")
	case 5:
		fmt.Println("Friday\n")
	case 6:
		fmt.Println("Saturday\n")
	case 7:
		fmt.Println("Sunday\n")
	default:
		fmt.Println("Wrong day!\n")
	}

	// multiple values in a single case
	letter := 'c'
	switch letter {
	case 'a', 'e', 'i', 'o', 'u':
		fmt.Println("Vowel\n")
	default:
		fmt.Println("Consonant\n")
	}

	//  Switch Without an Expression
	num1 := 15
	num2 := 20
	switch { // there is no value provided
	case num1 < 10:
		fmt.Println("Less than 10\n")
	case num1 >= 10 && num2 < 20:
		fmt.Println("Between 10 and 20\n")
	case num1 == 15 && num2 == 20:
		fmt.Println("num1 = 15 and num2 = 20\n")
	default:
		fmt.Println("20 or more\n")
	}

	// Fallthrough Behavior: Go has a fallthrough keyword to explicitly continue to the next case.
	num := 15
	switch { // there is no value provided
	case num < 10:
		fmt.Println("Less than 10\n")
	case num >= 10 && num < 20:
		fmt.Println("Between 10 and 20\n")
		fallthrough // below case will always run
	case num == 0:
		fmt.Println("Fallthrough Behavior\n")
	default:
		fmt.Println("20 or more\n")
	}

	// Type Switch with interface{}
	describeType(42)      // Integer
	describeType("hello") // String
	describeType(true)    // Boolean
	describeType(3.14)    // eska type humne switch me define nhi kiya hai

}
