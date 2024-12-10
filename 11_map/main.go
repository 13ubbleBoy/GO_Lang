/*
--> These are like dictonaries from python.

--> Value get stored in incresing order of the 'key'.

--> Data type in a map cannot be changed.
*/

package main

import "fmt"

func main() {
	// map has key value pairs, key type is defined in [] square brackerts.
	// value type is declared after [] the brackets.
	menu := map[string]int16{ // map declaration
		"soup":    50,
		"noodles": 60,
		"tea":     20,
		"coffee":  30, // last me bhi comma "," dena jaruri hai warna error aayega
	}

	// accessing all the items
	fmt.Println(menu)
	fmt.Println()

	// accessing single item
	fmt.Println(menu["tea"])
	fmt.Println()

	// looping in map
	for key, value := range menu {
		fmt.Printf("%v - ₹%v\n", key, value)
	}
	fmt.Println()

	phoneBook := map[int]string{
		836153: "Ankit",
		234151: "Praveen",
		654871: "Jeetu",
	}
	fmt.Println(phoneBook[654871])
	fmt.Println()

	fmt.Println(phoneBook)
	fmt.Println()

	// update an item in a map
	phoneBook[836153] = "Piyush"
	fmt.Println(phoneBook)
	fmt.Println()

	// No, you cannot directly change a key in a map in Go. Maps in Go are designed such that
	// keys are immutable once added. But we can delete a key and add a new one
	phoneBook[347193] = phoneBook[654871] // adding a new entry just before the old entry
	delete(phoneBook, 654871)
	fmt.Println(phoneBook, "\n")

	// adding a new key value pair
	phoneBook[746254] = "Mithun"
	fmt.Println(phoneBook, "\n")

}
