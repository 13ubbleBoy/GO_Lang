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
		items: map[string]int16{"tea": 20, "coffee": 40}, // declaration of an empty map
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

// reciever function, which is associated with only bill object. No other unit can use or call this function.
func (b bill) format() string { // this can only be called by bill objects.
	formatedString := "Bill breakdown: \n"
	var total int16 = 0

	for key, value := range b.items {
		formatedString += fmt.Sprintf("%-10v ...₹%v\n", key+":", value) // fmt.Sprintf returns a formated string
		total += value
	}

	// total
	formatedString += fmt.Sprintf("%-10v ...₹%v", "Total:", total) // "%-10v" makes the value 10 character long
	// -ve means on space on right side and 10 means space on left side

	return formatedString
}
