package main

import (
	"fmt"
	"strings" // "Contains" comes from this strings package
)

func main() {
	// check for a substring
	str := "praveen kumar is a good boy. He is from Uttarakhand."
	fmt.Println(strings.Contains(str, "kumar"), "\n") // returns true or false

	// replace a word appearing multiple times with another
	fmt.Println(strings.ReplaceAll(str, "is", "REPLACED")) // returns a new string, does not modify the existing one.
	fmt.Println(str, "\n")                                 // original string is unchanged

	// Upper case
	name := "Praveen"
	fmt.Println(strings.ToUpper(name), "\n")

	// get starting index of a word / index of a character
	fmt.Println(strings.Index(str, "kumar"))    // starting index of a word
	fmt.Println(strings.Index(name, "e"), "\n") // index of a single character

	// split string based on a charater
	newStr := strings.Split(str, " ") // return a string array
	fmt.Println(newStr)
	fmt.Printf("Type of newStr: %T\n\n", newStr)

}
