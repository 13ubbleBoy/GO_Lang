package main

import "fmt"

type bill struct { // custom data type
	name  string
	items map[string]int16
	tip   int16
}

// make new bills
func newbill(name string) bill { // the function will return custom bill type
	b := bill{ // creating a new instance of type bill variable / object
		name:  name,
		items: map[string]int16{}, // declaration of an empty map
		tip:   0,
	}

	// another way of creating instances of structure
	var b1 bill
	b1.name = "Kumar"
	fmt.Println(b1, "\n")
	// by default another values will be empty, if no value is provided like not even name (here),
	// then every value inside b1 will be empty. For name it will a empty string.

	return b
}
