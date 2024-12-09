package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := []int{23, 43, 33, 21, 31, 27, 23, 36}
	sort.Ints(ages) // it will sort the existing slice, but only sort slice not array.
	fmt.Println(ages, "\n")

	// get index of an element
	index := sort.SearchInts(ages, 33)
	fmt.Println(index, "\n")
	/*
		--> what happen when we pass an element that does not exist
			It return the length of the slice here 8, but indexing starts from 0, so 0 to 7 elements
			and we get 8, it indicates that the element does not exits in the slice.
	*/
	index2 := sort.SearchInts(ages, 50)
	fmt.Println(index2, "\n")

	// sort an slice of strings
	names := []string{"Praveen", "Ankit", "Piyush", "Jeetu"}
	sort.Strings(names)
	fmt.Println(names, "\n")

	index3 := sort.SearchStrings(names, "Piyush")
	fmt.Println(index3, "\n")

}
