/*
--> b *bill, it means passing a pointer of data type bill.
*/

package main

import "fmt"

type bill struct { // custom data type
	name  string
	items map[string]float32
	tip   float32
}

// make new bills
func newbill(name string) bill { // the function will return custom bill type
	b := bill{ // creating a new instance of type bill variable / object
		name:  name,
		items: map[string]float32{}, // declaration of an empty map
		tip:   0,
	}

	return b
}

// reciever function, which is associated with only bill object. No other unit can use or call this function.
func (b *bill) format() string { // this can only be called by bill objects.
	formatedString := "Bill breakdown: \n"
	var total float32 = 0

	for key, value := range b.items {
		formatedString += fmt.Sprintf("%-20v ...₹%v\n", key+":", value) // fmt.Sprintf returns a formated string
		total += value
	}

	// adding tip in the bill
	formatedString += fmt.Sprintf("%-20v ...₹%0.2f\n", "tip:", b.tip)

	// total
	formatedString += fmt.Sprintf("%-20v ...₹%0.2f", "Total:", total+b.tip) // "%-10v" makes the value 10 character long
	// -ve means on space on right side and 10 means space on left side

	return formatedString
}

// update tip
func (b *bill) addTip(t float32) { // t int16 should match the type
	b.tip = t // pointer to the bill so that value at the reference gets updated as go lang uses call by value.
}

// add an item to the bill
func (b *bill) addItem(itemName string, price float32) { // "b *bill" it is passing a pointer of type bill.
	b.items[itemName] = price // adding a new value to the items map.
}
